package module

type ExitSnoopKern struct {
	FilterCroup     bool   `bpf:"filter_cg" default:"false"`
	TargetPid       uint64 `bpf:"target_pid" default:"0"`
	TraceFailedOnly bool   `bpf:"trace_failed_only" default:"false"`
	//TraceByProcess  bool
}

type ExitSnoop struct {
	ExitSnoopKern
}
