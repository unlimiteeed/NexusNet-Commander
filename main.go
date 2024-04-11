package main

import (
	"fmt"

	connect "lbot.init/project/cmd/connect"
	run "lbot.init/project/cmd/run"
	clear "lbot.init/project/pkg/clear"
)

func main() {
	clear.ClearTerminal()
	fmt.Printf("\n[+]Running Command\n")
	excutec := "whoami"
	customPort := 5555
	connect.Connect()
	run.Run(excutec, customPort)
}
