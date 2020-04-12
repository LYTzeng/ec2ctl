.PHONY: clean

GO_CROSSCOMPILE :=
	ifeq ($(OS),Windows_NT)
		GO_CROSSCOMPILE = CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on
	else
		UNAME_S := $(shell uname -s)
		ifeq ($(UNAME_S),Darwin)
			GO_CROSSCOMPILE = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
		endif
	endif

all: clean main main.zip update_lambda

main: main.go
	$(GO_CROSSCOMPILE) go build -ldflags="-s -w" -o $@ *.go
	upx -f --brute -o $@.upx $@
	rm $@

main.zip: main
	zip $@ $<.upx
	# zip $@ $<

clean:
	rm -f main main.zip main.upx

update_lambda: main.zip
	aws lambda update-function-code --function-name ec2ctl --zip-file fileb://./$<