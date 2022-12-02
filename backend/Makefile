build: bin/
	go build -o bin/trackr src/main.go

clean: bin/
	go mod tidy
	rm -rf bin/*

run: bin/
	go run src/main.go

test: bin/
	go test ./tests/...

bin/:
	mkdir bin/
