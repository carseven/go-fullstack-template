generate:
	@npx tailwindcss -i ./assets/css/tailwindcss/tailwind.css -o ./assets/css/tailwindcss/dist/style.css --minify
	@templ generate
run:
	make generate
	@go run cmd/main.go
build:
	make generate
	@go build -o ./tmp/main cmd/main.go
htmx-install:
	@curl https://unpkg.com/htmx.org@1.9.10/dist/htmx.min.js -o 3p/htmx/htmx@1.9.10.min.js