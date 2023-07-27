
.PHONY: gen
gen:
	@echo "Generating SDK"
	curl -sSL https://insulator.conductor.one/api/v1/openapi.yaml -o openapi.yaml
	speakeasy generate sdk -s openapi.yaml -o . -d
	rm openapi.yaml
	@echo "Fixing Permissions"
	# We do this because speakeasy's platform generates files as 0755 but we don't want all of these files to be executable by default
	# Once this is fixed on their end we can remove this
	find * -type f -perm '+rwx' | xargs chmod -x
