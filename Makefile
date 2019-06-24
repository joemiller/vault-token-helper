deps:
	@go get

lint:
	@golangci-lint run -v

test:
	@go test -coverprofile=cover.out -v ./...

test-update-golden-files:
	@go test -v ./cmd/... -update

cov:
	@go tool cover -html=cover.out

build:
	@go build .

build-linux:
	@GOOS=linux go build .

snapshot:
	@goreleaser --snapshot --rm-dist --debug

build-dev-docker-image:
	@docker build -t joemiller/vault-token-helper-dev -f ./dev/Dockerfile.dev ./dev

run-dev-docker-image:
	#docker run --rm -it -v$$(PWD):/src -w /src joemiller/vault-token-helper-dev /bin/bash
	docker run --rm -it -v$$(PWD):/src --privileged -w /src joemiller/vault-token-helper-dev /bin/bash

todo:
	@grep \
		--exclude-dir=vendor \
		--exclude-dir=dist \
		--exclude-dir=Attic \
		--text \
		--color \
		-nRo -E 'TODO:.*' .

.PHONY: build build-linux test snapshot todo