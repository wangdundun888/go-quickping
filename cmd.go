package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	Ip        string
	Row       int
	FileModel bool
	Intervel  int //second
	Args      []string
}

func NewCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = Usage
	flag.StringVar(&cmd.Ip, "ip", "192.168.0.0", "-ip 192.168.0.0")
	flag.IntVar(&cmd.Row, "r", 20, "")
	flag.BoolVar(&cmd.FileModel, "m", false, "")
	flag.IntVar(&cmd.Intervel, "t", 1, "")
	flag.Parse()
	args := os.Args
	if len(args) > 0 {
		cmd.Args = args[1:]
	}
	return cmd
}

func Usage() {
	fmt.Printf("Usage: %s -ip subIp\n\t eg : %s -ip 192.168.0 \n", os.Args[0],os.Args[0])
}
