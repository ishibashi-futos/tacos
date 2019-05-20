MAKEFLAGS += --no-print-directory
GOOS=linux
GOARCH=amd64

build:
	GOOS=$(GOOS) go build -o ./plugin.exe

package:
	tar -zcvf ./plugin.tar.gz ./plugin.exe ./plugin.json

clean:
	rm -f ./plugin.exe ./plugin.tar.gz

repackage:
	@make clean
	@make build
	@make package