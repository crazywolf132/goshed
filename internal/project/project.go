package project

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/crazywolf132/goshed/internal/config"
	"github.com/crazywolf132/goshed/internal/model"
	"github.com/crazywolf132/goshed/internal/styles"
	"github.com/crazywolf132/goshed/internal/template"
)

// Create creates a new project with the given configuration
func Create(p *model.Project) error {
	projectDir := filepath.Join(config.ProjectsDir, p.Name)

	// Check if project already exists
	if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
		return fmt.Errorf("project %s already exists", p.Name)
	}

	// Create project directory
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Set project path
	p.Path = projectDir

	// Create metadata file
	metadataPath := filepath.Join(projectDir, ".goshed.json")
	metadata, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal project metadata: %w", err)
	}

	if err := os.WriteFile(metadataPath, metadata, 0644); err != nil {
		return fmt.Errorf("failed to write project metadata: %w", err)
	}

	// Initialize Go module
	if err := initGoModule(projectDir, p.Name); err != nil {
		return fmt.Errorf("failed to initialize Go module: %w", err)
	}

	// Create initial files based on template
	if err := createTemplateFiles(p); err != nil {
		return fmt.Errorf("failed to create template files: %w", err)
	}

	// Initialize Git repository
	if err := InitGit(p); err != nil {
		fmt.Printf("%s: %v\n", styles.Warning("Warning: Failed to initialize Git"), err)
	}

	return nil
}

func initGoModule(dir, name string) error {
	modFile := filepath.Join(dir, "go.mod")
	content := fmt.Sprintf("module %s\n\ngo 1.23\n", name)

	return os.WriteFile(modFile, []byte(content), 0644)
}

func createTemplateFiles(p *model.Project) error {
	tmpl, err := template.Get(p.Template)
	if err != nil {
		// Fallback to basic template if specified template not found
		tmpl, err = template.Get("basic")
		if err != nil {
			return fmt.Errorf("failed to get template: %w", err)
		}
	}

	// Create each file from the template
	for filename, content := range tmpl.Files {
		filePath := filepath.Join(p.Path, filename)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", filename, err)
		}
	}

	return nil
}

// Get retrieves a project by name
func Get(name string) (*model.Project, error) {
	projectDir := filepath.Join(config.ProjectsDir, name)

	// Check if project exists
	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("project %s does not exist", name)
	}

	// Read metadata file
	metadataPath := filepath.Join(projectDir, ".goshed.json")
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read project metadata: %w", err)
	}

	var p model.Project
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, fmt.Errorf("failed to parse project metadata: %w", err)
	}

	// Set path
	p.Path = projectDir

	return &p, nil
}

// Update updates a project's metadata
func Update(p *model.Project) error {
	if p.Path == "" {
		p.Path = filepath.Join(config.ProjectsDir, p.Name)
	}

	metadataPath := filepath.Join(p.Path, ".goshed.json")
	metadata, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal project metadata: %w", err)
	}

	if err := os.WriteFile(metadataPath, metadata, 0644); err != nil {
		return fmt.Errorf("failed to write project metadata: %w", err)
	}

	return nil
}

// CopyTo copies a project to a new location
func CopyTo(p *model.Project, dest string) error {
	// Copy all files except .goshed.json
	return filepath.Walk(p.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip .goshed.json
		if info.Name() == ".goshed.json" {
			return nil
		}

		// Get relative path
		relPath, err := filepath.Rel(p.Path, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// Create destination path
		destPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		// Copy file
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		if err := os.WriteFile(destPath, data, info.Mode()); err != nil {
			return fmt.Errorf("failed to write file %s: %w", destPath, err)
		}

		return nil
	})
}

// Remove removes a project from GoShed
func Remove(name string) error {
	projectDir := filepath.Join(config.ProjectsDir, name)
	return os.RemoveAll(projectDir)
}

// List returns all projects
func List() ([]*model.Project, error) {
	entries, err := os.ReadDir(config.ProjectsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read projects directory: %w", err)
	}

	var projects []*model.Project
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		p, err := Get(entry.Name())
		if err != nil {
			// Log error but continue with other projects
			fmt.Printf("Warning: failed to read project %s: %v\n", entry.Name(), err)
			continue
		}

		projects = append(projects, p)
	}

	return projects, nil
}
