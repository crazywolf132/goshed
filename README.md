# GoShed 🛠️

GoShed is your personal Go workshop - a powerful, interactive tool for creating and managing Go project playgrounds. It's designed to make experimentation and prototyping in Go as frictionless as possible.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D%201.16-blue)

## Why GoShed?

Every developer knows the dance: you have an idea, want to try a new library, or need to prototype a concept. You create a new directory, initialize a Go module, set up the basic files... and often delete everything later because it was just an experiment.

GoShed streamlines this entire process by providing:

- **Instant Playgrounds**: Create ready-to-use Go environments in seconds
- **Smart Management**: Keep track of all your experiments without cluttering your workspace
- **Automatic Cleanup**: No more stale test projects taking up space
- **Seamless Promotion**: Turn successful experiments into real projects with a single command
- **Interactive UI**: Beautiful terminal UI for easy project management

## Features

### 🚀 Quick Start
```bash
# Start interactive mode
goshed -i

# Or use CLI commands
goshed new -n myapi -t web
```

### 🎯 Project Templates
- **Basic**: Simple Go program structure
- **Web**: HTTP server with routing setup
- **CLI**: Command-line application boilerplate
- *More templates coming soon!*

### 💡 Smart Features
- **Auto-cleanup**: Automatically remove unused playgrounds
- **Project Promotion**: Convert experiments to full projects
- **Project Tags**: Organize and categorize your work
- **Notes System**: Keep track of your experiments
- **VS Code Integration**: Open projects directly in your editor

### 🎨 Interactive Mode
- Beautiful terminal UI powered by Bubble Tea
- Intuitive keyboard controls
- Project insights at a glance
- Quick actions for common tasks

## Installation

```bash
# Install with go install
go install github.com/crazywolf132/goshed@latest

# Or build from source
git clone https://github.com/crazywolf132/goshed.git
cd goshed
go build
```

## Usage

### Interactive Mode
```bash
goshed -i
```
Navigate through projects, create new ones, and manage everything with an intuitive interface.

### CLI Commands
```bash
# Create a new playground
goshed new -n myproject -t web

# List all playgrounds
goshed list

# Open in VS Code
goshed open -n myproject

# Promote to full project
goshed promote -n myproject

# Clean old playgrounds
goshed clean --older-than 720h
```

## Why You'll Love GoShed

- **Zero Configuration**: Start coding immediately with sensible defaults
- **Non-Intrusive**: Experiments stay in their own space until you decide to keep them
- **Productive**: Focus on coding, not project setup
- **Organized**: Never lose track of your experiments again
- **Flexible**: Works with your existing Go tools and workflows

## Perfect For

- 🧪 Experimenting with new libraries
- 🎓 Learning Go concepts
- 💡 Prototyping ideas
- 🧪 Testing different approaches
- 📚 Managing code examples
- 🎯 Interview preparation

## Contributing

We love contributions! Whether it's:
- 🐛 Bug reports
- 💡 Feature suggestions
- 🔨 Code contributions
- 📚 Documentation improvements

Open an issue or submit a pull request at [github.com/crazywolf132/goshed](https://github.com/crazywolf132/goshed).

## Roadmap

- [ ] Custom template system
- [ ] Git integration status
- [ ] Dependency insights
- [ ] Project sharing
- [ ] Cloud backup integration
- [ ] Multi-workspace support

## License

MIT License - feel free to use in your own projects!

---

## Get Started

Don't let project setup slow down your experimentation. Get GoShed now and focus on what matters - writing great Go code.

```bash
go install github.com/crazywolf132/goshed@latest
```

Built with ❤️ by the Go community, for the Go community.