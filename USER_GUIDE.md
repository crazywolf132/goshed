# GoShed User Guide 📖

## Getting Started

### Installation
```bash
go install github.com/crazywolf132/goshed@latest
```

### First Steps
1. Create your first playground:
   ```bash
   goshed create -n hello -t basic
   ```

2. List your playgrounds:
   ```bash
   goshed list
   ```

3. Open in your editor:
   ```bash
   goshed open -n hello
   ```

## Interactive Mode

Launch with:
```bash
goshed i
```

### Navigation
- `↑/↓`: Navigate items
- `Enter`: Select
- `Esc`: Go back
- `?`: Toggle help
- `q`: Quit

### Project Management
- `n`: New project
- `o`: Open project
- `d`: Delete project
- `e`: Edit project
- `/`: Search projects

## Templates

### Available Templates
- `basic`: Simple Go program
- `web`: HTTP server
- `cli`: Command-line app
- `api`: RESTful API
- `graphql`: GraphQL API

### Using Templates
```bash
goshed create -n myproject -t web
```

## Project Features

### Tags
Add tags when creating:
```bash
goshed create -n myproject -t web --tags="api,demo"
```

### Notes
Add notes to projects:
```bash
goshed notes -n myproject -t "This is a test API"
```

### Git Integration
Git commands are available in project view:
- Initialize repository
- View status
- Commit changes
- Push/pull

### Project Promotion
Convert playground to full project:
```bash
goshed promote -n myproject -d ~/projects/myapi
```

## Configuration

### Default Settings
Configuration file: `~/.goshed/config.yaml`
```yaml
editor: code
cleanup:
  older_than: 720h
```

### Environment Variables
- `GOSHED_EDITOR`: Preferred editor
- `GOSHED_CONFIG`: Config file location

## Tips & Tricks

1. Quick Search
   - Use `/` in interactive mode
   - Filter by tags: `goshed list --filter-tag=api`

2. Workspace Organization
   - Use tags consistently
   - Add descriptive notes
   - Clean up regularly

3. Git Integration
   - Initialize early
   - Commit frequently
   - Push important work

## Troubleshooting

### Common Issues
1. Editor not opening
   - Check `EDITOR` environment variable
   - Verify editor installation

2. Git errors
   - Ensure Git is installed
   - Check repository status

3. Template errors
   - Verify template name
   - Check template requirements

### Getting Help
- Use `goshed help [command]`
- Check documentation
- Open GitHub issues