package scan

import (
	"fmt"
	"net/http"
	"time"

	"github.com/iamritikbhardwaj/gorecon/utils"
)

func ScanHTTP(target string) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(target)
	if err != nil {
		utils.Error(err.Error())
		return
	}
	defer resp.Body.Close()

	utils.Info(resp.Status)
	fmt.Printf("%s ->  %s, %s\n", target, resp.StatusCode, resp.Status)

}
