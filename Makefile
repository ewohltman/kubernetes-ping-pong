.PHONY: lint
lint:
	golangci-lint run --enable-all --deadline=5m ./...

.PHONY: test
test:
	go test -v -race -coverprofile=coverage.out ./...

.PHONY: buildPing
buildPing:
	CGO_ENABLED=0 go build -o build/package/ping/ping cmd/ping/ping.go

.PHONY: buildPong
buildPong:
	CGO_ENABLED=0 go build -o build/package/pong/pong cmd/pong/pong.go

.PHONY: build
build: buildPing buildPong

.PHONY: image
image: build
	docker image build -t ewohltman/ping:latest build/package/ping
	docker image build -t ewohltman/pong:latest build/package/pong

.PHONY: push
push: image
	docker push ewohltman/ping:latest
	docker push ewohltman/pong:latest
