package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var host string
var start_port int
var end_port int

var resChan = make(chan string)
var errChan = make(chan string)

func main() {
	user_input()
}

func check_port(host string, start_port, end_port int) {

	for portNum := start_port; portNum <= end_port; portNum++ {
		go doConnectionTest(portNum)
        }
	for res := start_port; res <= end_port; res++ {
		select {
		case aResult := <- resChan:
			fmt.Println("\n\u2620 ", aResult)
		case aBlockedPort := <- errChan:
			fmt.Print("\u26D4 :", aBlockedPort)
		}
	}
	fmt.Println("\n\nAll tested!")
}

func doConnectionTest(port int) {
		qualified_host := fmt.Sprintf("%s%s%d", host, ":", port)
		conn, err := net.DialTimeout("tcp", qualified_host, 1*time.Second)  // Code from: https://stackoverflow.com/questions/37294052/golang-why-net-dialtimeout-get-timeout-half-of-the-time
		if (err != nil) { 
			errChan <-fmt.Sprintf("%d, ",port) 
		} else {
			fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
			status, _ := bufio.NewReader(conn).ReadString('\n')
		        resChan <- fmt.Sprintf("%d:%s",port,status)	
		} 
}

func user_input() {
	fmt.Println("Host> ")
	fmt.Scan(&host)
	fmt.Println("Starting Port (i.e. 80)> ")
	fmt.Scan(&start_port)
	fmt.Println("End Port (i.e. 8080)> ")
	fmt.Scan(&end_port)
	fmt.Println("Running scan... ")
	check_port(host, start_port, end_port)
}
