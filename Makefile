
.PHONY: cover
cover:
	@go test -coverpkg=./... -coverprofile=coverage.txt ./...

.PHONY: test
test:
	@go test ./...
