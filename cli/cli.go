package main

import (
	"syscall"

	"github.com/kevwan/adhoc"
)

func main() {
	adhoc.SayHello("0.0.7-1")
	_ = syscall.SIGTERM
}
