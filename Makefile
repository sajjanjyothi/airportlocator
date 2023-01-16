BUILDNAME?=airport_detector

build:
	go build -o bin/$(BUILDNAME) ./cmd
test:
	go test -cover ./...

run:
	go run ./cmd

clean:
	rm -rf bin/
	
build_docker:
	docker build . -t airport_detector