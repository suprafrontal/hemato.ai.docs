#!/bin/sh

set -e

echo this should be run from the root of the project


echo Building the static website

# if running for the first time you need to first fetch the image
# docker build . -t slatedocs/slate
docker run --rm --name slate -v $(pwd)/build:/srv/slate/build -v $(pwd)/source:/srv/slate/source slatedocs/slate build


echo Building the docker image to serve the static website

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 460080181115.dkr.ecr.us-east-1.amazonaws.com

export TS=`date +%s`

echo Current Tag is $TS

docker build -t docs.hemato.ai:$TS -t docs.hemato.ai:latest .

docker tag docs.hemato.ai:$TS 460080181115.dkr.ecr.us-east-1.amazonaws.com/docs.hemato.ai:$TS
docker tag docs.hemato.ai:latest 460080181115.dkr.ecr.us-east-1.amazonaws.com/docs.hemato.ai:latest

docker push 460080181115.dkr.ecr.us-east-1.amazonaws.com/docs.hemato.ai:$TS
docker push 460080181115.dkr.ecr.us-east-1.amazonaws.com/docs.hemato.ai:latest
