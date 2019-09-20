package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	vmStat, err := mem.VirtualMemory()
	errorCheck(err)
	fmt.Println(vmStat)

	diskStat, err := disk.Usage("/")
	errorCheck(err)
	fmt.Println(diskStat)

	cpuStat, err := cpu.Info()
	errorCheck(err)
	fmt.Println(cpuStat)

	percentage, err := cpu.Percent(0, true)
	errorCheck(err)
	fmt.Println(percentage)

	hostStat, err := host.Info()
	errorCheck(err)
	fmt.Println(hostStat)

	interfStat, err := net.Interfaces()
	errorCheck(err)
	fmt.Println(interfStat)

}
