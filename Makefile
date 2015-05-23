all: deps
	@echo "Building..."
	@mkdir -vp bin/
	@go build -v -o bin/tissue
	@echo "Installing into GOPATH/bin..."
	@cp -v bin/* $(shell echo $(GOPATH) | sed -e 's/:.*//g')/bin
	@echo "...done!"

deps:
	@echo "Fetching deps..."
	@go get -d -v ./...
