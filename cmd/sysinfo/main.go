package main

import (
	"os"
	"github.com/spf13/cobra"
	"fmt"
	"net"
)

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}

func init() {
	basename := os.Args[0][2:]

	rootCmd = &cobra.Command{
		Use:   basename,
		Short: basename + ": 系统信息查看工具",
	}

	hostnameCmd := &cobra.Command{
		Use:   "hostname",
		Short: "查看主机名",
		Run: func(cmd *cobra.Command, args []string) {
			printHostname()
		},
	}

	var v4, v6 bool
	netCmd := &cobra.Command{
		Use:   "ip",
		Short: "查看主机IP",
		Run: func(cmd *cobra.Command, args []string) {
			if v4 {
				printIpv4Addrs()
			}
			if v6 {
				printIpv6Addrs()
			}
			if !v4 && !v6 {
				cmd.Usage()
			}
		},
	}
	netCmd.PersistentFlags().BoolVarP(
		&v4,
		"ipv4",
		"i",
		false,
		"仅显示IPv4地址",
	)
	netCmd.PersistentFlags().BoolVarP(
		&v6,
		"ipv6",
		"x",
		false,
		"显示IPv6地址",
	)

	rootCmd.AddCommand(hostnameCmd, netCmd)
}

func printHostname() {
	hostname, _ := os.Hostname()
	fmt.Println(hostname)
}

func printIpv4Addrs() {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}

func printIpv6Addrs() {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To16() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
