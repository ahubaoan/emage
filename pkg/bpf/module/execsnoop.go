package module

type ExecSnoopKern struct {
	FilterCroup bool `bpf:"filter_cg" default:"false"` // 过滤cgroup应用
}

type ExecSnoop struct {
	ExecSnoopKern
}
