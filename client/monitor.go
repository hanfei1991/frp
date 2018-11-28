package client

import (
	"os"
	"time"
	"fmt"
	"syscall"
)

func (svr *Service) monitor() {
	for {
		ticker := time.NewTicker(time.Second)
		select {
			case <- ticker.C:
		}
		if !findProcess(svr.hostPID) {
			close(svr.closedCh)
		}
		ticker.Stop()
	}
}

func findProcess(pid int) bool {
	pro, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("find process error: %v", err)
		return false
	}
	if pro != nil {
		//fmt.Printf("find pid success: %d\n", pro.Pid)
		err :=pro.Signal(syscall.Signal(0))
		if (err != nil) {
			fmt.Printf("ping process error: %v", err)
		}
		//fmt.Printf("ping pid success: %d\n", pro.Pid)
		return true;
	}
	return false
}

