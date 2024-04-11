package run

import (
	"fmt"
	"os/exec"
	"sync"
)

func Run(command string, port int) {
	devices := []string{
		"20.15.200.180",
	}

	var wg sync.WaitGroup

	for _, device := range devices {
		wg.Add(1)
		go func(device string) {
			defer wg.Done()
			adbCommand := fmt.Sprintf("adb -s %s:%d shell %s", device, port, command)
			cmd := exec.Command("sh", "-c", adbCommand)
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Error running command on device %s: %v\n", device, err)
				return
			}

			result := string(output)
			fmt.Printf("[+] %s : %s\n", device, result)
		}(device)
	}

	wg.Wait()
}
