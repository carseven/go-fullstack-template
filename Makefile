run:
	@templ generate
	@go run cmd/main.go
build:
	@templ generate
	@go build -o ./tmp/main cmd/main.go
htmx-install:
	@curl https://unpkg.com/htmx.org@1.9.10/dist/htmx.min.js -o 3p/htmx/htmx@1.9.10.min.js