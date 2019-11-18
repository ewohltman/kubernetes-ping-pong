.PHONY: lint test build

lint:
	golangci-lint run --enable-all --deadline=5m ./...

test:
	go test -v -race -coverprofile=coverage.out ./...

buildPing:
	CGO_ENABLED=0 go build -o build/package/ping/ping cmd/ping/ping.go

buildPong:
	CGO_ENABLED=0 go build -o build/package/pong/pong cmd/pong/pong.go

build: buildPing buildPong

docker: build
	docker image build -t ewohltman/ping:latest build/package/ping
	docker image build -t ewohltman/pong:latest build/package/pong

push: docker
	docker push ewohltman/ping:latest
	docker push ewohltman/pong:latest
