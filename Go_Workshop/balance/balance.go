// TCP load balancer
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// Channel of backends
var (
	backends = make(chan string)
)

func init() {
	// Supply next backend, no locks
	go func() {
		hosts := []string{
			"www.google.com:80",
			"www.ibm.com:80",
			"ww.intel.com:80",
			// "www.ge.com:80", // GE socket blocks
		}
		i := 0
		for {
			backends <- hosts[i]
			i = (i + 1) % len(hosts)
		}
	}()
}

// forward proxies traffic from conn to next backend
func forward(conn net.Conn) {
	be := <-backends // Get next backend
	fmt.Printf("forwarding to %s\n", be)
	remote, err := net.Dial("tcp", be)
	if err != nil {
		fmt.Printf("can't dial to %s - %s", be, err)
	}
	// proxy traffic on both directions
	go io.Copy(conn, remote)
	go io.Copy(remote, conn)
}

func main() {
	local, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("can't listen - %s", err)
	}

	for {
		conn, err := local.Accept()
		if err != nil {
			log.Fatalf("can't accept - %s", err)
		}
		go forward(conn)
	}
}