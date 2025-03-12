# Music Collection

This project represents a **music collection application**.

## Prerequisites

Before you begin, ensure you have the following tools installed:

1. **Go** (version 1.23 or higher)
2. **Taskfile:**  
   A task runner for automated workflows and commands.
   ```bash
   go install github.com/go-task/task/v3/cmd/task@latest
   ```
3. **Templ:**  
   A Go package for safe and efficient HTML/templating.
   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```

## Setup and Build Instructions

This project uses [Taskfile](https://taskfile.dev) to manage build tasks and dependencies. Here's how to use it:

### Viewing Available Tasks

To see all available tasks in the project:

```bash
task
# or
task --list-all
```

### Building the Application

To build the application:

```bash
task build
```

This will:
1. Generate Go code from Templ templates with `templ generate`
2. Compile the application to `./out/music-collection`

### Running the Application

To build and run the application:

```bash
task run
```

This command:
1. First executes the `build` task as a dependency
2. Then runs the compiled binary located at `./out/music-collection`

## Project Structure

The application server is built using **[Echo](https://echo.labstack.com)**, a high-performance, extensible, and minimalist Go web framework.

## Additional Notes

- The build output is placed in the `./out/` directory
- Templ is used for HTML templating - ensure any template changes are regenerated with `templ generate` before building
- For development purposes, you can use `task run` which handles both building and running in one command