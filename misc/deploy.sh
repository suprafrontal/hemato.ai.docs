#!/bin/sh

set -e

echo this should be run from the root of the project


echo Building the static website

# if running for the first time you need to first fetch the image
# docker build . -t slatedocs/slate
docker run --rm --name slate -v $(pwd)/build:/srv/slate/build -v $(pwd)/source:/srv/slate/source slatedocs/slate build
mkdir -p static
cp -r build/* static/


echo CLOUD BUILDING
gcloud builds submit --tag gcr.io/stimulator/docs.hemato.ai --timeout=3m
echo success
echo -------------------

echo DEPLOYING
gcloud beta run deploy docs-hemato-ai \
  --labels service=docs-hemato-ai,version="${VERSION}",commit=${COMMIT} \
  --region us-east1 \
  --timeout=1m \
  --use-http2 \
  --cpu 2 \
  --memory 2Gi \
  --image gcr.io/stimulator/docs.hemato.ai \
  --platform managed

rm -rf static