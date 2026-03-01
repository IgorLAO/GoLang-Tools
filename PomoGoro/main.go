package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"os/exec"
)

func main () {
	args, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing argument:", err)
		return
	}
	
	for i := 0; i < args; i++ {
		ClearTerminal()
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func ClearTerminal(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}