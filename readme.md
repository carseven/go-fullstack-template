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

- [x] Implement HTMX to have dynamic content and state management (https://templ.guide/server-side-rendering/htmx/ and https://htmx.org)

- [] Implement basic testing
- [] Implement structure to separate API from frontend (/api and /app)
- [] Implement CI/CD with github actions and AWS lambda function?
- [] Generate the same template but for static files (Example: Blog)
- [] Add metadata parametrization to improve SEO by default
- [] Add language translation
- [] Add markdown support to render to HTML and be able to inject components inside
- [] Find more inspiration with Astro templates, really great ideas
- [] Render components from other JS frameworks, sometimes is need to avoid massive refactors and for maintenance of legacy code

Implement basic structure https://templ.guide Nice example to follow https://github.com/a-h/templ/tree/main/examples/counter

- [x] Components
- [x] Layouts
- [x] View
- [] CSS design system (Probably be compatible with tailwind?) Probably better idea ti implement tailwind for better adoption, but nice to have ways to still as templ defines
- [] Final refactor of the structure
