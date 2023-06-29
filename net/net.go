package main

import (
	"fmt"
	"net"
)

func main() {
	// 定义一个IPv4网络
	ip := net.ParseIP("192.168.0.1")
	subnet := net.IPNet{
		IP:   ip,
		Mask: net.CIDRMask(24, 32), // 子网掩码为24位, 表示前24是网络部分, 后8位是主机部分
	}

	// 判断一个IP地址是否属于该网络
	ipToCheck := net.ParseIP("192.168.0.100")
	if subnet.Contains(ipToCheck) {
		fmt.Println("IP belongs to the network")
	} else {
		fmt.Println("IP does not belong to the network")
	}

	// 获取网络的起始地址和结束地址
	networkStart := subnet.IP.Mask(subnet.Mask)
	networkEnd := make(net.IP, len(networkStart))
	copy(networkEnd, networkStart)
	for i := range networkEnd {
		networkEnd[i] |= ^subnet.Mask[i]
	}

	fmt.Println("Netword start: ", networkStart.String())
	fmt.Println("Network end: ", networkEnd.String())
}
