package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/iamritikbhardwaj/gorecon/scan"
)

func main() {
	targets := flag.String("targets", "", "Path to file with target domains")
	flag.Parse()

	if *targets == "" {
		fmt.Println("Usage: gorecon -targets <file>")
		os.Exit(1)
	}

	file, err := os.Open(*targets)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		scan.LookupCNAME(domain)
		scan.ScanPorts(domain) // 👈 Add this line
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading targets:", err)
	}
}
