build: generate
	go build cmd/main.go

generate:
	templ generate

clean:
	rm ./main

run-dev:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run cmd/main.go config/testdata/example.yaml" --open-browser=false

test: generate
	go test ./...
