package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"net"
	"time"
)

func main() {
	argsWithProg := os.Args
	var hostIndex int;
	portsIndex := -1
	for i := range argsWithProg {
		if (argsWithProg[i] == "--help") {
			fmt.Println("Enter the host IP address and the port range. You may use domain name instead of IP address and you don't have to enter the port range(it will be set by default).")
			os.Exit(0)
		}
	}

	for i := range argsWithProg {
		if (argsWithProg[i] == "--host") {
			hostIndex = i;
		} else if (argsWithProg[i] == "--ports") {
			portsIndex = i;
		}
	}
	hostIp := argsWithProg[hostIndex+ 1]
	port_string := "0-65535"
	if (portsIndex != -1) {
		port_string = argsWithProg[portsIndex + 1];
	}
	ports := strings.Split(port_string, "-")
	low_port, _ := strconv.Atoi(ports[0])
	high_port, _ := strconv.Atoi(ports[1])
	for i := low_port; i <= high_port; i++ {
		current_port := strconv.Itoa(i);
		_, err := net.DialTimeout("tcp", hostIp + ":" + current_port, 300 * time.Millisecond)
		if (err == nil) {
			fmt.Print(".")
			fmt.Print(current_port)
		}
	}
}