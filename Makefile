cover:
	# go test -covermode=count -coverprofile=coverage.out ./...
	go test -covermode=set -coverprofile=coverage.out ./...
.PHONY: cover

cover-show:
	go tool cover -html=coverage.out
.PHONY: cover-show

# Tags thing
# go test -tags=integration -covermode=count -coverprofile=count.out ./...