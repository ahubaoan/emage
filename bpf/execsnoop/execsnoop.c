//go:build ignore
// +build ignore
#include "com.h"

const volatile u64  filter_cg = 0;

struct event {
    u32     pid;
    u32     ppid;
    char    comm[16];
    char    filename[256];
};

const struct event *unused __attribute__((unused));

struct {
	__uint(type, BPF_MAP_TYPE_RINGBUF);
	__uint(max_entries, 1 << 24);
} events SEC(".maps");


struct {
	__uint(type, BPF_MAP_TYPE_CGROUP_ARRAY);
	__type(key, u32);
	__type(value, u32);
	__uint(max_entries, 1);
} cgroup_map SEC(".maps");


struct execve_args {
    unsigned long long unused;
    int __syscall_nr;
    char *filename;
    const char *const *argv;
    const char *const *envp;
};

// Userspace pathname: /sys/kernel/debug/tracing/events/syscalls/sys_enter_execve/
SEC("tracepoint/syscalls/sys_enter_execve")
int sys_enter_execve(struct execve_args *args)
{

    struct event *e;
    struct task_struct *task;
    u32 ret;
    u32 pid;

    if (filter_cg && !bpf_current_task_under_cgroup(&cgroup_map, 0))
        return 0;

	e = bpf_ringbuf_reserve(&events, sizeof(struct event), 0);
	if (!e) {
		return 0;
	}

    pid = bpf_get_current_pid_tgid();
    e->pid = pid;
	bpf_get_current_comm(&e->comm, sizeof(e->comm));
    task = (struct task_struct*)bpf_get_current_task();
    e->ppid = BPF_CORE_READ(task, real_parent, tgid);

    bpf_probe_read_user_str(e->filename, sizeof(e->filename), args->filename); // 用于读取数据

    bpf_printk("filter_cg=%d\n", filter_cg);

    bpf_ringbuf_submit(e, 0);

	return 0;
}
