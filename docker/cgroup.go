package example

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

// 挂载了memory subsystem 的 hierarchy的根目录的位置
const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {
	// os.Args[0]表示程序运行时的第一个命令行参数, 即表示可执行文件的路径
	// "/proc/self/exe"是一个特殊的文件路径，在Linux系统中用于表示当前正在执行的可执行文件的路径
	if os.Args[0] == "/proc/self/exe" {
		// 判断程序是否在容器内运行, 如果条件成立，说明程序正在以容器的方式运行
		// 因为容器内部执行的可执行文件路径会被替换成"/proc/self/exe"
		// 容器进程
		fmt.Printf("current pid %d", os.Getpid())
		fmt.Println()
		// 创建一个新的进程用来模拟容器对内存的消耗
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	} else {
		// 得到fork出来进程映射在外部命名空间的pid
		fmt.Printf("%v", cmd.Process.Pid)

		// 在系统默认创建挂在了memory subsystem的Hierarchy上创建cgroup
		os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit"), 0755)
		// 讲容器进程放到cgroup中
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
		// 限制cgroup进程使用
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "memory.limit_in_bytes"), []byte("100m"), 0644)
	}
	cmd.Process.Wait()
}
