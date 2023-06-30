#
# You can override these e.g. as
#     make test TEST_PKG=./packages/vm/core/testcore/ TEST_ARG="-v --run TestAccessNodes"
#
TEST_PKG=./...
TEST_ARG=

test:
	go test ./...  --count 1 -failfast

lint:
	golangci-lint run --timeout 5m

gofumpt-list:
	gofumpt -l ./

.PHONY: test lint gofumpt-list