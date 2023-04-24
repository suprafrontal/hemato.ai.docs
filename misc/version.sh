#!/bin/sh

#read more at http://alimoeenysbrain.blogspot.com/2013/10/automatic-versioning-with-git-commit.html
#replace all hemato.docs with your project name
#for languages other than matlab replace all the .go extensions with your language extensions and modify BODY to reflect you language syntax
#for go (golang) see this gits: https://gist.github.com/alimoeeny/7005725

VERBASE=$(git rev-parse --verify HEAD | cut -c 1-7)
echo $VERBASE
NUMVER=$(awk '{printf("%s", $0); next}' hematoDocsVersion.go | sed 's/.*hemato.docs\.//' | sed 's/\..*//')
echo "old version: $NUMVER"
NEWVER=$(expr $NUMVER + 1)
BRANCH=$(git rev-parse --abbrev-ref HEAD)

echo "new version: hemato.docs.$NEWVER.$VERBASE"
#BODY="package main\n\nfunc hematoDocsVersion() string {\n\treturn \"hemato.docs.$NEWVER.$BRANCH.$VERBASE\"\n}"
#echo $BODY > hematoDocsVersion.go
printf "package main\n\nfunc hematoDocsVersion() string {\n\treturn \"hemato.docs.%s.%s.%s\"\n}" $NEWVER $BRANCH $VERBASE > hematoDocsVersion.go
git add hematoDocsVersion.go

echo "new version: hemato.docs.$NEWVER.$VERBASE"
BODY="export class hematoDocsVersion {\n\tString() {\n\t\treturn \"hemato.docs.$NEWVER.$BRANCH.$VERBASE\"\n\t}\n}\nexport default hematoDocsVersion;\n\n//"
echo $BODY > source/javascripts/app/hematoDocsWWWVersion.js
git add source/javascripts/app/hematoDocsWWWVersion.js
