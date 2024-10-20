# 🪐🌌 Nova 🌌🪐
Template repository to create Go projects

## Prerequisites
Before you can build and run this project, ensure you have the following installed on your machine:
- Go (version 1.23.1+)
- golangci-lint for linting

## Installation
### Clone the repository:
```bash
git clone git@github.com:senonide/nova.git
cd nova
```

### Install dependencies (if any):
```bash
go mod tidy
```

### Install golangci-lint (if not already installed):
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Makefile Commands
The project includes a Makefile that simplifies common tasks. Here are the available commands:

1. Lint the Code
Run static analysis and check the code for potential issues:
```bash
make lint
```
This command uses golangci-lint to perform static analysis and report any issues.

2. Run Tests
Execute all unit tests with the race detector enabled:
```bash
make test
```

3. Build the Project (Debug Mode)
Compile the project with debugging symbols included:
```bash
make build-debug
```

4. Build the Project (Release Mode)
Compile the project for production:
```bash
make build
```

5. Run the Project
Run the project directly:
```bash
make run
```
