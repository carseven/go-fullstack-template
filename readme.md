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

[] Implement HTMX to have dynamic content and state management (https://templ.guide/server-side-rendering/htmx/ and https://htmx.org)
[] Implement basic testing
[] Increase basic structure to have a design system, components, layouts, etc. https://templ.guide
[] Implement CI/CD with github actions and AWS lambda function?
[] Generate the same template but for static files (Example: Blog)
