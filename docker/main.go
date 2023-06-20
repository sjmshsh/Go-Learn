package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 创建一个命令对象, 封装了Linux clone这个系统调用接口, 并且执行的命令是sh
	cmd := exec.Command("sh")
	// 用于设置新进程的属性
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	// 分别讲新进程的标准输入, 标准输出和标准错误连接到当前进行的标准输入, 标准输出和标准错误输出
	// 这样可以确保新进程的输入输出与当前进程保持一致
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
