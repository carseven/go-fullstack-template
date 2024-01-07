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

Install ffmpeg for transcoding mp4 video to HLS format

```shell
brew install ffmpeg
```

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
- [x] Configure tailwind user color-preference media query with Dark mode
- [x] Add language translation
- [ ] Load language translation from JSON file instead of struct
- [x] Upload videos and transform to HLS to stream
- [x] Generate the same template but for static files (Example: Blog)
- [x] Dev and Prod mode
- [x] Client reloading web server has changed (Websocket client communication to reload HTML pages when view has changed) Replicate Vite implementation (WS to send event of reload to the client. Also, client PING websocket to check connection alive and then polling to relaunch when the server connect again?)
- [x] Add markdown support to render to HTML and be able to inject components inside
- [x] Add cache mechanism for static files
- [x] Add vscode settings and extensions recommendation HTMX, templ syntax, tailwind, etc.
- [x] CSS design system (Tailwind) (https://tailwindcss.com/docs/installation)

## Bugs

Sometimes when killing the dev server does not kill the port 3000. For the moment kill it manually

```shell
lsof -i:3000

pkill pid
or
kill -9 pid
```
