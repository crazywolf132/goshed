package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/crazywolf132/goshed/internal/model"
)

// GitStatus represents the Git status of a project
type GitStatus struct {
	Initialized bool
	Clean       bool
	Branch      string
	Remote      string
}

// GetGitStatus returns the Git status of a project
func GetGitStatus(p *model.Project) (*GitStatus, error) {
	status := &GitStatus{}

	// Check if git is initialized
	gitDir := filepath.Join(p.Path, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return status, nil
	}
	status.Initialized = true

	// Get current branch
	cmd := exec.Command("git", "-C", p.Path, "branch", "--show-current")
	output, err := cmd.Output()
	if err == nil {
		status.Branch = string(output)
	}

	// Check if working directory is clean
	cmd = exec.Command("git", "-C", p.Path, "status", "--porcelain")
	output, err = cmd.Output()
	if err == nil {
		status.Clean = len(output) == 0
	}

	// Get remote URL
	cmd = exec.Command("git", "-C", p.Path, "config", "--get", "remote.origin.url")
	output, err = cmd.Output()
	if err == nil {
		status.Remote = string(output)
	}

	return status, nil
}

// InitGit initializes a Git repository for the project
func InitGit(p *model.Project) error {
	cmd := exec.Command("git", "-C", p.Path, "init")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	// Create .gitignore
	gitignore := filepath.Join(p.Path, ".gitignore")
	content := `# GoShed metadata
.goshed.json

# Go build
/bin/
/pkg/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Go testing
*.test
*.out

# IDE specific files
.idea/
.vscode/
*.swp
*.swo
`
	if err := os.WriteFile(gitignore, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}

	return nil
}
