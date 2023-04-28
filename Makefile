.PHONY: dev

dev:
	@bash scripts/dev.sh


.PHONY: test

test:
	go clean -testcache 
	go mod tidy
	go test -cover ./test/...