#!/bin/bash

set -e

name=${1::-4}

docker build -t lc3compiler compiler/
docker run --rm -it -v $(pwd):/data lc3compiler /data/$name.asm

rm -f $name.sym
mv -f $name.obj programs/bin
