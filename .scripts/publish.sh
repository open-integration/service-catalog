echo "Authentication with GCS project $GCS_PROJECT"
gcloud auth activate-service-account --key-file $KEY_FILE_PATH
echo "Authenticated!"

FILES=$(ls -ls dist | awk 'NR>1' | awk '{print $10}')
VERSION=$(cat VERSION)
echo "Files to be uploaded:"
echo "$FILES"

gsutil mv dist/* $GCS_BUCKET
