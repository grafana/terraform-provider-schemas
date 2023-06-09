default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: generate
generate:
	rm -rf ./internal/provider/*_gen.go
	cd ./gen && MIN_MATURITY=merged go generate ./
	go generate ./
