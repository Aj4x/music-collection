# Music Collection

This project is a **music collection application** that allows users to manage and interact with their music library.

## Features

- **Album Management:** Browse and organize your music collection by albums
- **Settings Panel:** Customize application appearance and behavior
- **Responsive Design:** Enhanced UI for various screen sizes


## Prerequisites

Before you begin, ensure you have the following tools installed:

1. **Go** (version 1.23 or higher)
2. **Taskfile:**  
   A task runner for automated workflows and commands.
   ```bash
   go install github.com/go-task/task/v3/cmd/task@latest
   ```
3.**Air**
  Live reload for Go apps
  ```bash
  go install github.com/air-verse/air@latest
  ```
4.**Templ:**  
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

For live reload using Air:

```bash
task dev
```

This will run `templ generate` as well as build and execute
the Go build output executable. This will then re-run each time
project files are changed.

This command:
1. First executes the `build` task as a dependency
2. Then runs the compiled binary located at `./out/music-collection`

## Project Structure

- **`/views`**: Contains Templ template files for the UI components
- **`/models`**: Data models including settings definitions
- **`/data`**: Music collection data management
- **`/static`**: CSS styles organized by component
   - `/styles/styles.css`: Main application styles
   - `/styles/album.css`: Album view styles
   - `/styles/settings.css`: Settings panel styles
   - `/styles/colour.css`: Color scheme definitions

## Technology Stack

- **Backend**: Go with [Echo](https://echo.labstack.com) framework
- **Frontend**: HTML/CSS with Templ templating
- **Build System**: Taskfile for task automation

## Development Notes

- Modify template files in `/views` directory and run `templ generate` before building
- CSS is organized by component for easier maintenance
- The application uses a component-based architecture for better code organization