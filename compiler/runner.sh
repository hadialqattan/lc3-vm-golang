#!/bin/bash

set -e

compiler/assembler.sh programs/source/$1.asm && \
    go run main.go programs/bin/$1.obj
