package main

import (
	"fmt"
	"net"
)

func ipTest1() {
	address := net.ParseIP("1.2.3.4")
	if address == nil {
		fmt.Printf("invalid IP Address: %s\n", address)
	}
	address = net.ParseIP("aaa")
	if address == nil {
		fmt.Printf("invalid IP Address: %s\n", address)
	}
}

func ipTest2() (ips []string) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}

	fmt.Println(ips)
	return ips
}

func main() {
	//IP合法性检查
	//ipTest1()
	//获取本机IP
	ipTest2()
}
