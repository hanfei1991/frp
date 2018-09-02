package client

import (
	"os"
	"time"
	"fmt"
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
		return false
	}
	if pro != nil {
		fmt.Printf("find pid success: %d\n", pro.Pid)
		state, err := pro.Wait();
		if err != nil {
			return false;
		}
		return !state.Exited()
	}
	return false
}

