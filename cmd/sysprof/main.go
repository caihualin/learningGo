package main

import (
	"flag"
	"fmt"
	"time"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var (
	loadF, cpuF, memF bool
)

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
	}
	cmdRun()
}

func init() {
	flag.BoolVar(&loadF, "l", false, "查看CPU负载")
	flag.BoolVar(&cpuF, "u", false, "查看CPU占用")
	flag.BoolVar(&memF, "m", false, "查看内存占用")
}

func cmdRun() {
	if loadF {
		printLoad()
	}
	if cpuF {
		printCPU()
	}
	if memF {
		printMem()
	}
}

func printLoad() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Printf("load 1: %f\n",loadInfo())
		}
	}
}

func printCPU() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Printf("cpu usage(%%): %f\n",cpuUsage())
		}
	}
}

func printMem() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Printf("mem usage(%%): %f\n",memUsage())
		}
	}
}

func loadInfo() float64 {
	l, _ := load.Avg()
	return l.Load1
}

func cpuUsage() float64 {
	u, _ := cpu.Percent(1 * time.Second,false)
	return u[0]
}

func memUsage() float64 {
	m, _ := mem.VirtualMemory()
	return m.UsedPercent
}
