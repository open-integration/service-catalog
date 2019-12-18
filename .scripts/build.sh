build () {
    NAME=$1
    echo "Generating temp file for $NAME"
    cd $NAME
    go generate scripts/generate.go
    VERSION=$(cat ../VERSION)

    echo "Building $NAME-$VERSION service fro darwin and amd64"
    GOOS=darwin GOARCH=amd64 packr build -o ./dist/$NAME-$VERSION-darwin-amd64 .

    echo "Building $NAME-$VERSION service fro linux and amd64"
    GOOS=linux GOARCH=amd64 packr build -o ./dist/$NAME-$VERSION-linux-amd64 .

    echo "Building $NAME-$VERSION service fro linux and 386"
    GOOS=linux GOARCH=386 packr build -o ./dist/$NAME-$VERSION-linux-386 .
}

build $NAME