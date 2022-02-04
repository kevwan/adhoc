package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevwan/adhoc"
)

func cancelOnSignals() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTSTP, syscall.SIGQUIT)
		select {
		case <-c:
			fmt.Println(`migrate failed, reason: "User Canceled"`)
			os.Exit(0)
		}
	}()
}

func main() {
	cancelOnSignals()
	adhoc.SayHello("0.0.7-1")
}
