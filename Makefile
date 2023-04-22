build:
	go build -o bin/go-grep main.go
run:
	./bin/go-grep
clean:
	rm ./bin/go-grep
