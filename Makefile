default: golangci-lint test

test:
	@sh -c "'$(CURDIR)/scripts/test.sh'"

golangci-lint:
	@sh -c "'$(CURDIR)/scripts/golangci-lint.sh'"

.PHONY: test golangci-lint
