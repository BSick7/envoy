SHELL := /bin/bash

.PHONY: deps vet build test

deps:
	go get -u github.com/Masterminds/glide
	go get github.com/jstemmer/go-junit-report
	glide install

vet:
	glide nv | xargs go vet

build:
	go build

install:
	go install

test:
	set -o pipefail;glide nv \
		| xargs go test -v \
		| tee /dev/tty \
		| go-junit-report > unit-tests.xml

release:
	go get github.com/mitchellh/gox
	go get github.com/tcnksm/ghr
	gox -os "linux darwin windows" -arch "amd64 386" -ldflags "-X main.Version=`cat VERSION`" -output="dist/envoy_{{.OS}}_{{.Arch}}"
	ghr -t $$GITHUB_TOKEN -u BSick7 -r envoy --replace `cat VERSION` dist/
