#!/bin/bash

set -x
export BPF_CLANG=clang

cd $1
go generate
