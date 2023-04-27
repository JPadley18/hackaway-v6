.PHONY: dev

dev:
	@bash scripts/dev.sh


.PHONY: test

test:
	@go test -cover ./test/...