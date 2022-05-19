#!/bin/bash
BPF_CLANG=clang go generate
sed -i 's/return spec.LoadAndAssign(obj, opts)/if err := RewriteConstants(obj.(*bpfObjects), spec); err != nil {\n        return err\n    }\n\n    return spec.LoadAndAssign(obj, opts)/g' bpf_bpfe*.go
sed -i '/type bpfObjects struct {/ a\\tConsts map[string]interface{}'  bpf_bpfe*.go
