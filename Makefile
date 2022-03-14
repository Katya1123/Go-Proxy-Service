.PHONY: lint
default:

.PHONY: lint
lint:
	golangci-lint run --fix ./...

lint-ci:
	golangci-lint run ./...

# make new-migration name=some_name
.PHONY: new-migration
new-migration:
	migrate create -ext sql -dir migrations $(name)

.PHONY: docs
docs:
	swag i -g cmd/main/main.go --parseInternal

.PHONY: test
test:
	@go test --race --vet= ./... -v
