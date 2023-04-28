VERSION ?= 1.0.0

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/terraform-provider-sams-test-tf-provider_v$(VERSION)-darwin_amd64
	GOOS=linux GOARCH=amd64 go build -o bin/terraform-provider-sams-test-tf-provider_v$(VERSION)-linux_amd64
