package main

import "learningGo/daemon"

func main() {
	// Use Daemon
	learn_daemon.NewDaemon(daemonExample)
}

func daemonExample() {
	select{}
}