#!/bin/bash
cd $1
BPF_CLANG=clang go generate
