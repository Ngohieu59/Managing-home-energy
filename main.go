package main

import (
	"Managing-home-energy/cmd"
	"Managing-home-energy/log"
	"context"
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Println("pid:", pid)
	log.Infow(context.Background(), fmt.Sprintf("Process ID: %v", pid))
	cmd.Execute()
}
