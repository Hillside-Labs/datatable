
build:
	templ fmt .
	templ generate
	go mod tidy
	go build .
