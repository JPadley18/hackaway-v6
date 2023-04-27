.PHONY: dev

dev:
	@bash scripts/dev.sh


.PHONY: test

test:
	go mod tidy
	go test -cover ./test/...