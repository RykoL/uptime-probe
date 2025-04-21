build:
    go build cmd/main.go

generate:
    templ generate

clean:
    rm ./main

run: build generate
    ./main

test: generate
    go test ./...
