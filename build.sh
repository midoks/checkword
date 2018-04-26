#!/bin/sh

bee pack -be GOOS=darwin
mv checkword.tar.gz checkword.macosx.tar.gz

bee pack -be GOOS=windows
mv checkword.tar.gz checkword.win64.tar.gz

bee pack -be GOOS=linux
mv checkword.tar.gz checkword.linux.tar.gz
