SERVICES="trello google-spreadsheet git docker"

echo ""

for s in $SERVICES
do  
    echo "Building service $s"
    cd $s
    make build
    cd ../
done
