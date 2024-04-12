package run

import (
	"fmt"
	"os/exec"
)

func Run(command string, port int, device string) {
	adbCommand := fmt.Sprintf("adb -s %s:%d shell %s", device, port, command)
	cmd := exec.Command("sh", "-c", adbCommand)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command on device %s: %v\n", device, err)
	}

	result := string(output)
	fmt.Printf("\n[+] %s : %s\n", device, result)
}
