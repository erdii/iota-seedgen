release: clean all zip

.PHONY: all
all: windows linux mac

.PHONY: clean
clean:
	rm -rf build/
	rm -rf release/


.PHONY: windows
windows: win-32 win-64

win-32: build/iota-seedgen_win-32bit.exe
win-64: build/iota-seedgen_win-64bit.exe

build/iota-seedgen_win-32bit.exe: main.go
	GOOS=windows GOARCH=386 go build -o build/iota-seedgen_win-32bit.exe main.go

build/iota-seedgen_win-64bit.exe: main.go
	GOOS=windows GOARCH=amd64 go build -o build/iota-seedgen_win-64bit.exe main.go


.PHONY: linux
linux: linux-32 linux-64

linux-32: build/iota-seedgen_linux-32bit
linux-64: build/iota-seedgen_linux-64bit

build/iota-seedgen_linux-32bit: main.go
	GOOS=linux GOARCH=386 go build -o build/iota-seedgen_linux-32bit main.go

build/iota-seedgen_linux-64bit: main.go
	GOOS=linux GOARCH=amd64 go build -o build/iota-seedgen_linux-64bit main.go


.PHONY: mac
mac: darwin-32 darwin-64

darwin-32: build/iota-seedgen_mac-32bit
darwin-64: build/iota-seedgen_mac-64bit

build/iota-seedgen_mac-32bit: main.go
	GOOS=darwin GOARCH=386 go build -o build/iota-seedgen_mac-32bit main.go

build/iota-seedgen_mac-64bit: main.go
	GOOS=darwin GOARCH=amd64 go build -o build/iota-seedgen_mac-64bit main.go


.PHONY: zip
zip: release/iota-seedgen.zip all

release/iota-seedgen.zip:
	mkdir release
	zip -j release/iota-seedgen.zip build/*
	keybase pgp sign -d -i release/iota-seedgen.zip -o release/iota-seedgen.zip.asc