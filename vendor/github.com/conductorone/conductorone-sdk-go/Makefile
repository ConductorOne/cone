
.PHONY: gen
gen:
	@echo "Generating SDK"
	curl -sSL https://insulator.conductor.one/api/v1/openapi.yaml -o openapi.yaml
	speakeasy generate sdk -s openapi.yaml -o . -d --lang go
	rm openapi.yaml


.PHONY: testacc
testacc:
	GO_ACC=1 go test -v -cover -timeout=5m ./...