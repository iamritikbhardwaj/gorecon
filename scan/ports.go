package scan

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// CommonPorts holds the web-related ports to scan
var CommonPorts = []int{80, 443, 8080, 8443, 8000, 3000, 5000, 8888}

// ScanPorts scans a domain/IP for open common web ports
func ScanPorts(target string) {
	fmt.Printf("[PORT SCAN] Scanning %s...\n", target)

	for _, port := range CommonPorts {
		address := net.JoinHostPort(target, strconv.Itoa(port))

		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err != nil {
			continue // Port closed or unreachable
		}

		conn.Close()
		fmt.Printf("[OPEN] %s:%d\n", target, port)
	}
}
