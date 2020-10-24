#!/bin/bash

set -e

docker build -t lc3compiler compiler/
docker run --rm -it -v $(pwd):/data lc3compiler /data/$1.asm

rm -f $1.sym
mv -f $1.obj programs/bin
