package main

import (
	"fmt"
	"net"
)

func macTest1() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}

	fmt.Println(macAddrs)
	return macAddrs
}

func macTest2() {
	addr, err := net.ParseMAC("de:53:60:8f:44:bd")
	if err != nil {
		fmt.Println("Invalid MAC")
	} else {
		fmt.Println(addr)
	}

	addr, err = net.ParseMAC("de:53:60:8f:44:b")
	if err != nil {
		fmt.Println("Invalid MAC")
	} else {
		fmt.Println(addr)
	}
}

func main() {
	//获取本机MAc
	//macTest1()
	//检查MAC格式
	macTest2()
}
