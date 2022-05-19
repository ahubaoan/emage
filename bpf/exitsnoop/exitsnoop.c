#include "com.h"

struct event {
    u64     start_time;
    u64     exit_time;
    u32     pid;
    u32     tid;
    u32     ppid;
    u32     sig;
    u32     exit_code;
    char    comm[16];
};

const volatile u64 target_pid = 0

struct {
	__uint(type, BPF_MAP_TYPE_CGROUP_ARRAY);
	__type(key, u32);
	__type(value, u32);
	__uint(max_entries, 1);
} cgroup_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
	__uint(key_size, sizeof(__u32));
	__uint(value_size, sizeof(__u32));
} events SEC(".maps");

SEC("tracepoint/sched/sched_process_exit")
int sched_process_exit(void *ctx)
{
	u64 id = bpf_get_current_pid_tgid();
	u32 pid = id >> 32;
	u32 tid = (__u32)pid_tgid;
	int exit_code;
	struct task_struct *task;
	struct event event = {};

	if (filter_cg && !bpf_current_task_under_cgroup(&cgroup_map, 0))
		return 0;

	if (target_pid && target_pid != pid)
		return 0;


	task = (struct task_struct *)bpf_get_current_task();
	exit_code = BPF_CORE_READ(task, exit_code);
	if (trace_failed_only && exit_code == 0)
		return 0;

	event.start_time = BPF_CORE_READ(task, start_time);
	event.exit_time = bpf_ktime_get_ns();
	event.pid = pid;
	event.tid = tid;
	event.ppid = BPF_CORE_READ(task, real_parent, tgid);
	event.sig = exit_code & 0xff;
	event.exit_code = exit_code >> 8;
	bpf_get_current_comm(event.comm, sizeof(event.comm));
	bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &event, sizeof(event));
	return 0;
}