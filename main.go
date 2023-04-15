package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	Cmd("GOPROXY", "https://proxy.golang.com.cn,direct")
	Cmd("DOTNET_CLI_TELEMETRY_OPTOUT", "1")
	time.Sleep(500 * time.Millisecond)
	os.Exit(0)
}
func Cmd(valname string, val string) {
	cmd := exec.Command("cmd", "/C", "setx", "/M", valname, val)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running setx command: ", err)
	}
	fmt.Println("Success : " + valname + " " + val)
}
