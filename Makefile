#################
###  Build  ###
#################
bd:
	@mkdir -p build
	@echo "--> ensure dependencies have not been modified"
	@go mod verify
	@echo "--> building void"
	@go build -mod=readonly -o build/void ./cli/

install: bd
	@export GOBIN=$(go env GOPATH)/bin
	@echo "--> installing void to $(GOBIN)"
	@cp build/void $(GOBIN)/void

.PHONY: all install

###################
### Development ###
###################

govet:
	@echo Running go vet...
	@go vet ./...

.PHONY: govet