test: _ensure
	go test -race -v ./...

build: test
	docker build -t jmervine/arp:latest .

_ensure:
	go mod vendor
	go mod tidy
	go vet
