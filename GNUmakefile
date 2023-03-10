default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: generate
generate:
	rm -rf ./tools/grok/terraform
	rm -rf ./internal/provider/zzz_*.go
	cd ./tools/grok && MIN_MATURITY=merged GRAFANA_VERSION=v9.4.0 go generate ./
	cp ./tools/grok/terraform/v9.4.0/data_sources/* ./internal/provider/
	go generate ./...
