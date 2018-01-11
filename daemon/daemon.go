package learn_daemon

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"runtime"
)

type F func()

func newLogger() *log.Logger {
	l := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)
	return l
}

func NewDaemon(f F) {
	name := filepath.Base(os.Args[0])
	args := os.Args[:]
	logger := newLogger()
	plat := runtime.GOOS

	runLock := "/tmp/" + name + ".pid"
	lockFile, err := os.OpenFile(runLock, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}

	stat, _ := lockFile.Stat()
	if stat.Size() != 0 {
		logger.Println("Already in running (PID file is existed)")
		os.Exit(2)
	}

	if os.Getppid() != 1 {
		syscall.Umask(0)
		if plat != "darwin" {
			os.Chdir("/")
		}

		os.StartProcess(
			name,
			args,
			&os.ProcAttr{Files: []*os.File{nil, nil, nil}},
		)
		return
	}

	lockFile.WriteString(fmt.Sprintf("%d\n", os.Getpid()))

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGUSR2, syscall.SIGQUIT)

	go f()

	for {
		s := <-c
		switch s {
		case syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT:
			exit(runLock)
		case syscall.SIGUSR2:
			continue
		}
	}
}

func exit(f string) {
	_ = os.Remove(f)
	os.Exit(0)
}
