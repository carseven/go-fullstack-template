generate:
	@npm run tailwind:build
	@templ generate
run:
	@make generate
	@go run cmd/main.go
build:
	@make generate
	@go build -o ./tmp/server cmd/main.go	
htmx-install:
	@curl https://unpkg.com/htmx.org@1.9.10/dist/htmx.min.js -o assets/js/htmx@1.9.10.min.js
hls-install:
	@curl https://cdn.jsdelivr.net/npm/hls.js@1 -o assets/js/hls@1.js