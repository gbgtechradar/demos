package main

import (
	"encoding/json"
	"fmt"
	"github.com/anvie/port-scanner"
	"os"
	"time"
)

type Scan struct {
	Address   string
	Timeout   time.Duration
	StartPort int
	EndPort   int
}

func main() {

	s := &Scan{Address: "127.0.0.1", Timeout: 500000, StartPort: 3999, EndPort: 4001} // Some defaults
	json.NewDecoder(os.Stdin).Decode(s)

	fmt.Printf("Scanning %s for %s", s.Address, s.Timeout)
	ps := portscanner.NewPortScanner(s.Address, s.Timeout)

	// get opened ports...
	fmt.Printf("Scanning port %d-%d...\n", s.StartPort, s.EndPort)

	openedPorts := ps.GetOpenedPort(s.StartPort, s.EndPort)

	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		fmt.Print(" ", port, " [open]")
		fmt.Println("  -->  ", ps.DescribePort(port))
	}
}
