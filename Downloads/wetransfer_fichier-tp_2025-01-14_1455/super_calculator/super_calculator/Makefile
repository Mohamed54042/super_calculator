# Variables
GO_FILES=$(shell find . -name '*.go')

# Commandes
build:
	go build -o supercalculator main.go

test:
	go test ./...

lint:
	semgrep --config=auto
	docker run --rm -i hadolint/hadolint < Dockerfile

release:
	goreleaser release --clean

docker-build:
	docker build -t ${DOCKERHUB_USERNAME}/supercalculator .

docker-push:
	docker push ${DOCKERHUB_USERNAME}/supercalculator
