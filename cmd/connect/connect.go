package connect

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func Connect(devices []string) {
	var wg sync.WaitGroup

	for _, device := range devices {
		wg.Add(1)
		go func(device string) {
			defer wg.Done()
			command := fmt.Sprintf("adb connect %s > /dev/null", device)
			cmd := exec.Command("sh", "-c", command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error connecting to device %s: %v\n", device, err)
			}
		}(device)
	}

	wg.Wait()
}
