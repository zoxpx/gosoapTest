.PHONY: test clean

all: myservice/myservice.go

$(GOPATH)/bin/gowsdl:
	@echo ":: Building gowsdl ..."
	go get github.com/hooklift/gowsdl/...

myservice.wsdl:
	@echo ":: Downloading WSDL ..."
	curl -fsSL "http://www.dneonline.com/calculator.asmx?wsdl" -o $@

myservice/myservice.go: $(GOPATH)/bin/gowsdl myservice.wsdl
	@echo ":: Composing service from WSDL ..."
	$(GOPATH)/bin/gowsdl myservice.wsdl

test: myservice/myservice.go
	@echo ":: Running tests ..."
	go test -v ./...

clean:
	rm -rf myservice.wsdl myservice

