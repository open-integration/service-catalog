SERVICES="trello google-spreadsheet git docker jira slack kubernetes"

echo ""

for s in $SERVICES
do  
    echo "Building service $s"
    cd $s
    VERSION=$(cat ./VERSION)
    GO_MOD_CANDIDATE=$s/v$VERSION
    LIST=$(git tag --list)
    ALREADY_EXISTS="false"
    for t in $LIST
    do
        if [ $GO_MOD_CANDIDATE == $t ]
        then
            echo "$GO_MOD_CANDIDATE == $t skipping new release for $s"
            ALREADY_EXISTS="true"
        fi
    done
    
    make build
    if [ $ALREADY_EXISTS == "false" ]
    then
        echo "Creating new git tag $GO_MOD_CANDIDATE"
        git tag $GO_MOD_CANDIDATE
        git push --tags
    fi
    cd ../
done
