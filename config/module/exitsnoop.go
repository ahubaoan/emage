package module

type ExitSnoopKern struct {
	FilterCroup     bool
	TargetPid       uint32
	TraceFailedOnly bool
	//TraceByProcess  bool
}

type ExitSnoop struct {
	ExitSnoopKern
}
