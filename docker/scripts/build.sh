#!/bin/bash
set -e

build () {
    NAME=$1
    VERSION=$(cat ./VERSION)

    echo "Building $NAME-$VERSION service for darwin and amd64"
    GOOS=darwin GOARCH=amd64 packr build -o ./dist/$NAME-$VERSION-darwin-amd64 .

    echo "Building $NAME-$VERSION service for linux and amd64"
    GOOS=linux GOARCH=amd64 packr build -o ./dist/$NAME-$VERSION-linux-amd64 .

    echo "Building $NAME-$VERSION service for linux and 386"
    GOOS=linux GOARCH=386 packr build -o ./dist/$NAME-$VERSION-linux-386 .
}

build $1
