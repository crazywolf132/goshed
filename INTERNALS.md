# GoShed: Under the Hood 🔧

This document explains the internal workings of GoShed, how it manages projects, and its core architecture.

## Directory Structure

```
$HOME/
└── .goshed/
    ├── project-1/
    │   ├── .goshed.json    # Project metadata
    │   ├── main.go         # Project files
    │   ├── go.mod
    │   └── .git/           # Git repository
    └── project-2/
        ├── .goshed.json
        └── ...
```

## Core Components

### 1. Project Management

#### Metadata System
Each project is tracked through a `.goshed.json` file containing:
```json
{
    "created": "2024-12-03T10:00:00Z",
    "lastAccessed": "2024-12-03T11:00:00Z",
    "name": "project-name",
    "template": "web",
    "tags": ["api", "experiment"],
    "notes": "Testing JWT implementation"
}
```

This metadata is used for:
- Project tracking
- Cleanup decisions
- Template identification
- Organization
- Search and filtering

### 2. Template Engine

Templates are defined in-memory as Go structs:
```go
type Template struct {
    Name         string
    Description  string
    Files        map[string]string
    Dependencies []string
}
```

When creating a new project:
1. Template is selected
2. Files are created from the template map
3. Dependencies are installed (if any)
4. Git repository is initialized
5. Metadata file is created

### 3. User Interface

#### CLI Mode
Uses `urfave/cli` for command parsing:
1. Commands are defined in the main app structure
2. Each command maps to a specific action function
3. Flags provide additional configuration

#### TUI Mode (Interactive)
Built with Bubble Tea framework:

1. **State Management**
```go
type model struct {
    screen          screen
    mainList        list.Model
    templateList    list.Model
    projectList     list.Model
    nameInput       textinput.Model
    spinner         spinner.Model
    help           help.Model
    // ...
}
```

2. **Screen Flow**
```
Main Menu ─────┬─── New Project ─── Template Selection
               │
               ├─── List Projects ─── Project Details
               │
               └─── Clean Projects
```

3. **Update Loop**
- Handles user input
- Updates model state
- Manages screen transitions
- Executes commands

### 4. Project Lifecycle

#### Creation
1. Generate project directory
2. Apply template
3. Initialize Git repository
4. Create Go module
5. Install dependencies
6. Create metadata file

#### Access
1. Update last accessed timestamp
2. Load project in editor
3. Update metadata

#### Promotion
1. Verify target location
2. Move project files
3. Update Go module path
4. Remove from GoShed tracking

#### Cleanup
1. Scan all projects
2. Check last accessed time
3. Remove projects older than threshold
4. Update project list

### 5. File Operations

All file operations are handled through Go's standard library:
```go
os.MkdirAll()     // Create directories
os.WriteFile()    // Write files
os.Rename()       // Move files
os.RemoveAll()    // Delete directories
```

Safety measures:
- Operations confined to GoShed directory
- Metadata validation before operations
- Error handling for all file operations

### 6. External Integrations

#### Editor Integration
1. Try VS Code first:
```go
cmd := exec.Command("code", projectDir)
```

2. Fall back to system editor:
```go
defaultEditor := os.Getenv("EDITOR")
cmd = exec.Command(defaultEditor, projectDir)
```

#### Git Integration
Basic Git operations:
```go
exec.Command("git", "init").Run()
```

## Performance Considerations

1. **File System Operations**
   - Metadata operations are lightweight
   - Template application is sequential
   - Cleanup runs as a single pass

2. **Memory Usage**
   - Templates stored in memory
   - Project list loaded on demand
   - TUI components use minimal state

3. **Concurrency**
   - File operations are synchronous
   - UI updates are non-blocking
   - Background cleanup possible

## Error Handling

1. **Graceful Degradation**
   - Editor fallback
   - Template validation
   - File system checks

2. **User Feedback**
   - Clear error messages
   - Status updates
   - Operation confirmation

## Future Extensibility

The architecture supports future additions:

1. **Template System**
   - Custom template loading
   - Template sharing
   - Version control

2. **Project Management**
   - Remote backups
   - Project sharing
   - Dependency tracking

3. **UI Enhancements**
   - Custom themes
   - Additional views
   - Keyboard shortcuts

## Security Considerations

1. **File System**
   - Operations confined to GoShed directory
   - No execution of external code
   - Safe file paths handling

2. **External Commands**
   - Limited to Git and editor
   - No arbitrary command execution
   - Environment validation

3. **Project Data**
   - Local storage only
   - No sensitive data collection
   - Metadata validation

## Debugging

To enable debug logging:
```bash
export GOSHED_DEBUG=1
```

Debug information includes:
- File operations
- Command execution
- State transitions
- Error details

This internal documentation should help contributors understand the codebase and make informed decisions when adding features or fixing bugs.