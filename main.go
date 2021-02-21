package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/sevlyar/go-daemon"
)

const name = "tink-reboot"

// To terminate the daemon use:
//  kill `cat tink-reboot`
func main() {
	cntxt := &daemon.Context{
		PidFileName: fmt.Sprintf("%v.pid", name),
		PidFilePerm: 0644,
		LogFileName: fmt.Sprintf("%v.log", name),
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{name},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release() // nolint

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("tink-reboot started")

	doReboot()
}

func doReboot() {
	time.Sleep(10 * time.Second)
	syscall.Sync()
	if err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART); err != nil {
		log.Fatal(err)
	}
}
