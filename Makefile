build: generate
	go build cmd/main.go

generate:
	templ generate

clean:
	rm ./main

run:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run cmd/main.go" --open-browser=false
	./main

test: generate
	go test ./...
