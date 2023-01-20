test-coverage:
	@go test -v ./... -covermode=atomic -race -coverpkg=./... -coverprofile coverage/coverage.out
	@go tool cover -html coverage/coverage.out -o coverage/coverage.html
	@open coverage/coverage.html

test:
	@go test ./... -race -covermode=atomic

.PHONY: test-coverage test