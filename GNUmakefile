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

build:
	go build
	mkdir -p "C:/Users/agnes/AppData/Roaming/terraform.d/plugins/terraform.local/grafana/schemas/1.0.0/windows_amd64"
	cp terraform-provider-schemas.exe "C:/Users/agnes/AppData/Roaming/terraform.d/plugins/terraform.local/grafana/schemas/1.0.0/windows_amd64/terraform-provider-schemas.exe