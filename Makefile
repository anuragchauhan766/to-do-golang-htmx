BINARY_NAME=main.out

run: 
	air

build:
	go build -o dist/${BINARY_NAME} cmd/main.go

clean:
	go clean
	rm -rf dist

## css: build tailwindcss
.PHONY: css
css:
	./tailwindcss -i static/css/index.css -o static/css/output.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	./tailwindcss -i static/css/index.css -o static/css/output.css --watch