GOPKGNAME = github.com/whosthatknocking/jocular-octo-moo

BUILD_SHORTCOMMIT=$(shell git rev-parse --short HEAD 2>/dev/null)
ifeq ($(OS),Darwin)
BUILD_VERSION=macos
else
BUILD_VERSION=0.1
endif
BUILD_RELEASE=1
BUILD_DATE=$(shell date -u)


.PHONY: local clean get build vet unit test 

local: build test

clean:
	go clean -i -x $(GOPKGNAME)/...
	rm -fr $(RPM_DIR)

get:
	go get -t -d -tags testing $(GOPKGNAME)/...

build: get
	go install -ldflags "-X \"main.Version=$(BUILD_VERSION)-$(BUILD_RELEASE)\" -X \"main.BuildCommit=$(BUILD_SHORTCOMMIT)\" -X \"main.BuildDate=$(BUILD_DATE)\"" $(GOPKGNAME)/...

vet:
	go vet $(GOPKGNAME)/...

test:
	go test -v $(GOPKGNAME)/...
