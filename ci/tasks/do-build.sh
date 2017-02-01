#!/bin/sh
basedir=`pwd`/gopath-tested/src/github.com/jghiloni/cf-apptree
build_dir=`pwd`/build-output/build
version_file=`pwd`/version/number

mkdir ${build_dir} > /dev/null 2>&1

set -e
set -x

export GOPATH=`pwd`/gopath-tested

# Run tests
cd ${basedir}

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X github.com/jghiloni/cf-apptree/main.version=`cat ${version_file}`" -o ${build_dir}/cf-apptree-linux-amd64
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-X github.com/jghiloni/cf-apptree/main.version=`cat ${version_file}`" -o ${build_dir}/cf-apptree-linux-386
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-X github.com/jghiloni/cf-apptree/main.version=`cat ${version_file}`" -o ${build_dir}/cf-apptree-windows-amd64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-X github.com/jghiloni/cf-apptree/main.version=`cat ${version_file}`" -o ${build_dir}/cf-apptree-windows-386.exe
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-X github.com/jghiloni/cf-apptree/main.version=`cat ${version_file}`" -o ${build_dir}/cf-apptree-macosx
