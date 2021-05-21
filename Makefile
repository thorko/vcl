.PHONY: vcl

vcl:
	rm -f /tmp/vcl
	go get github.com/hashicorp/vault/api
	go get gopkg.in/alecthomas/kingpin.v2
	go build -o vcl vcl.go

clean:
	rm -f vcl

test:
	./vcl --help
