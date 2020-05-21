deps:
	@go get

lint:
	@golangci-lint run -v --timeout=3m
	@if command -v goreleaser >/dev/null; then \
		goreleaser check; \
	else \
		echo "goreleaser not installed, skiping goreleaser linting"; \
	fi

test:
	@go test -coverprofile=cover.out -v ./...

cov:
	@go tool cover -html=cover.out

build:
	@go build .

release:
	@docker run \
		--rm \
		-e "GITHUB_TOKEN=$$GITHUB_TOKEN" \
		-e "GPG_KEY=$$GPG_KEY" \
		-v `pwd`:/src \
		-w /src \
		dockercore/golang-cross \
			/src/scripts/release.sh $(GORELEASER_ARGS)

snapshot: GORELEASER_ARGS= --rm-dist --snapshot
snapshot: release

sign-and-promote-release:
	bash ./scripts/sign-and-promote-release.sh

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