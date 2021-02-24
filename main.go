package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/kardianos/service"
	"github.com/sevlyar/go-daemon"
)

const name = "tink-reboot"

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	doReboot()
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

const customSysVinit = `#!/sbin/openrc-run

name="$SVCNAME"
command="/usr/sbin/$SVCNAME"
command_args="-daemon"
pidfile="/var/run/$SVCNAME.pid"
`

func main() {
	installFlag := flag.Bool("install", false, "install the system service.")
	uninstallFlag := flag.Bool("uninstall", false, "uninstall the system service.")
	startFlag := flag.Bool("start", false, "start the system service.")
	daemonFlag := flag.Bool("daemon", false, "run the daemon.")

	flag.Parse()
	svcConfig := &service.Config{
		Name:        name,
		DisplayName: "Tinkerbell reboot action",
		Description: "Tinkerbell reboot action",
		Option: service.KeyValue{
			"SysvScript": customSysVinit,
			"Restart":    "on-failure",
		},
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if *installFlag {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if *uninstallFlag {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if *startFlag {
		err = s.Start()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if *daemonFlag {
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
	}
	doReboot()

}

func doReboot() {
	log.Println("waiting 10 seconds before rebooting...")
	time.Sleep(10 * time.Second)
	log.Println("rebooting...")
	syscall.Sync()
	if err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART); err != nil {
		log.Fatal(err)
	}
	log.Println("reboot called")
}
