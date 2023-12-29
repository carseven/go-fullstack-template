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

Install nodejs and npm to be able to use the tailwind compiler:
Personally I install nodejs using nvm or fvm.

Check tailwind cli works properly.

```bash
npx tailwind --help
```

If a specific version is need instead of the latest, just change the make file with the proper version:
Be aware of the tailwind.config.js file might change. Edited as you like it

```bash
npx tailwind@3.4.0 -i ./assets/css/tailwindcss/tailwind.css -o ./assets/css/tailwindcss/dist/style.css
```

For more info about tailwind check https://tailwindcss.com/docs/installation

## Build

Use make to build or run the application:

```bash
make build
```

```bash
make run
```

Generate templ and tailwind:

```bash
make generate
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
- [x] Generate the same template but for static files (Example: Blog)
- [] Add metadata parametrization to improve SEO by default
- [] Add language translation
- [x] Add markdown support to render to HTML and be able to inject components inside
- [] Load markdown from external files and render component (Example: CMS to load md files)
- [] Add cache mechanism for static files
- [] Find more inspiration with Astro templates, really great ideas
- [] Render components from other JS frameworks, sometimes is need to avoid massive refactors and for maintenance of legacy code
- [x] Add vscode settings and extensions recommendation HTMX, templ syntax, tailwind, etc.
- [] CLI tool to create the boilerplate?
- [] Add database integration layer
  Implement basic structure https://templ.guide Nice example to follow https://github.com/a-h/templ/tree/main/examples/counter

- [x] Components
- [x] Layouts
- [x] View
- [x] CSS design system (Tailwind) (https://tailwindcss.com/docs/installation)
- [] Dark mode
- [] Final refactor of the structure
