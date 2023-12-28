# GO server with templates

## Install

Install templ and air locally on your machine:

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

```bash
go install github.com/cosmtrek/air@latest
```

Make sure go binary are exported to the path:

```bash
export PATH="$HOME/go/bin:$PATH"
```

## Build

Use make to build or run the application:

```bash
make run
```

```bash
make run
```

Use air to launch the hot reloading mode:

```bash
air
```

## TODO

[] Implement HTMX to have dynamic content and state management
[] Increase basic structure to have a design system, components, layouts, etc. https://templ.guide
[] Generate the same template but for static files (Example: Blog)
