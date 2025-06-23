package scan

import (
	"fmt"
	"net"
)

// LookupCNAME resolves the canonical name (CNAME) of a domain.
func LookupCNAME(domain string) {
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		fmt.Printf("[DNS] %s -> Error: %v\n", domain, err)
		return
	}

	fmt.Printf("[DNS] %s -> CNAME: %s\n", domain, cname)
}
