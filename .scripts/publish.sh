echo "Authentication with GCS project $GCS_PROJECT"
gcloud auth activate-service-account --key-file $KEY_FILE_PATH
echo "Authenticated!"

SERVICES="trello google-spreadsheet git jira slack kubernetes exec"

echo ""

for s in $SERVICES
do  
    cd $s
    VERSION=$(cat VERSION)
    FILES=$(ls -ls dist | awk 'NR>1' | awk '{print $10}')
    echo "Files to be uploaded:"
    echo "$FILES"
    gsutil mv dist/* $GCS_BUCKET
    cd ..
done

