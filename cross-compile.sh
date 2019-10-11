#!/bin/sh
set -ex

cd build

for os in windows linux darwin; do
  for arch in 386 amd64; do
    tag=$os-$arch
    if [ ! -d "$tag/" ]; then
      mkdir $tag
    fi

    ext=""
    if [ "$os" == "windows" ]; then
      ext=.exe
    fi
    filename=slack-recommend-channel-bot$ext

    cd ..
    env GOOS=$os GOARCH=$arch go build -o build/$tag/$filename
    cd build
    gzipped=$tag/$filename.gz
    gzip -f $tag/$filename
    mv $gzipped slack-recommend-channel-bot-$tag.gz
  done
done
