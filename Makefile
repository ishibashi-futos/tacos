GOOS=linux
GOARCH=amd64
GOPATH=<@yourGOPATH>

build:
	go build -o ./plugin.exe

package:
	tar -zcvf ./plugin.tar.gz ./plugin.exe ./plugin.json

clean:
	rm -f ./plugin.exe ./plugin.tar.gz
