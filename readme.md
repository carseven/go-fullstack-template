# GO server side rendring with Templ, HTMX and TailwindCSS

## Install

Install templ and air locally on your machine:

```shell
go install github.com/a-h/templ/cmd/templ@latest
```

```shell
go install github.com/cosmtrek/air@latest
```

Make sure go binary are exported to the path:

```shell
export PATH="$HOME/go/bin:$PATH"
```

Install nodejs and npm to be able to use the tailwind compiler:
Personally I install nodejs using nvm or fvm.

Install tailwind and tailwind plugins:

```shell
npm i
```

Generate tailwind css from templates:

```shell
npm run tailwind:build
```

For more info about tailwind check https://tailwindcss.com/docs/installation

## Build

Use make to build or run the application:

```shell
make build
```

```shell
make run
```

Generate templ and tailwind:

```shell
make generate
```

Use air to launch the hot reloading mode:

```shell
air
```

## TODO

Basic features (MVP)

- [x] Implement HTMX to have dynamic content and state management (https://templ.guide/server-side-rendering/htmx/ and https://htmx.org)
- [ ] Format DEV logs properly for easy DEBUG
- [ ] Load markdown from external files and render component (Example: CMS to load md files)
- [ ] Use Notion as a CMS to retrieve markdown
- [ ] Add database integration layer
- [x] Configure tailwind user color-preference media query with Dark mode
- [ ] Implement basic structure https://templ.guide Nice example to follow https://github.com/a-h/templ/tree/main/examples/counter (separate API from frontend (/api and /app))
- [ ] Implement basic testing
- [ ] Implement CI/CD with github actions and AWS lambda function?
- [ ] Add metadata parametrization to improve SEO by default
- [x] Add language translation
- [ ] Load language translation from JSON file instead of struct
- [ ] Swagger for documenting the APIs
- [ ] Documentation page
- [ ] Default error handling, 404 not found and other methods
- [x] Upload videos and transform to HLS to stream
- [ ] Use case to connect webcam thro RTMP and transform to HLS
- [ ] Observability
- [x] Generate the same template but for static files (Example: Blog)
- [x] Dev and Prod mode
- [x] Client reloading web server has changed (Websocket client communication to reload HTML pages when view has changed) Replicate Vite implementation (WS to send event of reload to the client. Also, client PING websocket to check connection alive and then polling to relaunch when the server connect again?)
- [x] Add markdown support to render to HTML and be able to inject components inside
- [x] Add cache mechanism for static files
- [x] Add vscode settings and extensions recommendation HTMX, templ syntax, tailwind, etc.
- [x] CSS design system (Tailwind) (https://tailwindcss.com/docs/installation)

Optional features

- [ ] Inject JS or Typescript using vite https://github.com/pilcrowOnPaper/go-vite (As a alternative, because sometimes is need it to done client side outside HTMX, for example, integrate with legacy code or use complex things already done in JS!)
- [ ] Find more inspiration with Astro templates, really great ideas

## Bugs

Sometimes when killing the dev server does not kill the port 3000. For the moment kill it manually

```shell
lsof -i:3000

pkill pid
or
kill -9 pid
```
