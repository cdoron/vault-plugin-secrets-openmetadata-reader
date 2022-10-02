module github.com/cdoron/vault-plugin-secrets-openmetadata-reader

go 1.13

require (
	fybrik.io/openmetadata-connector/datacatalog-go-client v0.0.0-00010101000000-000000000000 // indirect
	github.com/hashicorp/go-hclog v0.15.0
	github.com/hashicorp/vault/api v1.0.4
	github.com/hashicorp/vault/sdk v0.1.13
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime v0.8.2
)

replace fybrik.io/openmetadata-connector/datacatalog-go-client => github.com/fybrik/openmetadata-connector/auto-generated/client v0.0.0-20220928091421-bc7556a9adbb
