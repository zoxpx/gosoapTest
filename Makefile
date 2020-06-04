.PHONY: test clean

all: myservice/myservice.go

$(GOPATH)/bin/gowsdl:
	@echo ":: Getting gowsdl ..."
	go get -u github.com/hooklift/gowsdl/...

$(GOPATH)/bin/wsdl2go:
	@echo ":: Getting wsdl2go ..."
	go get -u github.com/fiorix/wsdl2go

myservice.wsdl:
	@echo ":: Downloading WSDL ..."
	curl -fsSL "http://www.dneonline.com/calculator.asmx?wsdl" -o $@

myservice/myservice.go: $(GOPATH)/bin/gowsdl myservice.wsdl
	@echo ":: Composing service from WSDL via gowsdl ..."
	$(GOPATH)/bin/gowsdl myservice.wsdl

calculatorsoap12/myservice.go: $(GOPATH)/bin/wsdl2go myservice.wsdl
	@echo ":: Composing service from WSDL via wsdl2go ..."
	mkdir -p $(dir $@)
	$(GOPATH)/bin/wsdl2go < myservice.wsdl > $@

test: myservice/myservice.go calculatorsoap12/myservice.go
	@echo ":: Running tests ..."
	go test -v ./...

clean:
	rm -rf myservice.wsdl myservice calculatorsoap12

