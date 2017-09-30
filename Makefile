VERSION := $(shell cat VERSION.txt)
LDFLAGS := -ldflags="-X main.version=$(VERSION)"

release: | clean all zip checksums


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
	GOOS=windows GOARCH=386 go build -o build/iota-seedgen_win-32bit.exe $(LDFLAGS) main.go

build/iota-seedgen_win-64bit.exe: main.go
	GOOS=windows GOARCH=amd64 go build -o build/iota-seedgen_win-64bit.exe $(LDFLAGS) main.go


.PHONY: linux
linux: linux-32 linux-64

linux-32: build/iota-seedgen_linux-32bit
linux-64: build/iota-seedgen_linux-64bit

build/iota-seedgen_linux-32bit: main.go
	GOOS=linux GOARCH=386 go build -o build/iota-seedgen_linux-32bit $(LDFLAGS) main.go

build/iota-seedgen_linux-64bit: main.go
	GOOS=linux GOARCH=amd64 go build -o build/iota-seedgen_linux-64bit $(LDFLAGS) main.go


.PHONY: mac
mac: darwin-32 darwin-64

darwin-32: build/iota-seedgen_mac-32bit
darwin-64: build/iota-seedgen_mac-64bit

build/iota-seedgen_mac-32bit: main.go
	GOOS=darwin GOARCH=386 go build -o build/iota-seedgen_mac-32bit $(LDFLAGS) main.go

build/iota-seedgen_mac-64bit: main.go
	GOOS=darwin GOARCH=amd64 go build -o build/iota-seedgen_mac-64bit $(LDFLAGS) main.go


.PHONY: zip
zip: | all release/iota-seedgen.zip 

release/iota-seedgen.zip:
	mkdir release || true
	zip -j release/iota-seedgen_v$(VERSION).zip build/*
	cp release/iota-seedgen_v$(VERSION).zip release/iota-seedgen_latest.zip
	keybase pgp sign -d -i release/iota-seedgen_v$(VERSION).zip -o release/iota-seedgen_v$(VERSION).zip.asc
	cp release/iota-seedgen_v$(VERSION).zip.asc release/iota-seedgen_latest.zip.asc


MD5 = $(shell md5sum release/iota-seedgen_v$(VERSION).zip | cut -d ' ' -f 1)
SHA1 = $(shell sha1sum release/iota-seedgen_v$(VERSION).zip | cut -d ' ' -f 1)
SHA256 = $(shell sha256sum release/iota-seedgen_v$(VERSION).zip | cut -d ' ' -f 1)
SHA512 = $(shell sha512sum release/iota-seedgen_v$(VERSION).zip | cut -d ' ' -f 1)

.PHONY: checksums
checksums:
	echo "\nrelease checksum: for release/iota-seedgen_v$(VERSION).zip\nMD5: $(MD5)\nSHA1: $(SHA1)\nSHA256: $(SHA256)\nSHA512: $(SHA512)"

.PHONY: test
test:
	go test
