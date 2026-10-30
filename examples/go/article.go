// PR Newswire Article API – Go example
// Usage: PILOTERR_API_KEY=your-key go run article.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	params := url.Values{}
	params.Set("query", "https://www.prnewswire.com/news-releases/pumpkin-awarded-11-000-in-donations-to-veterinary-clinics-for-cost-of-care-education-302924942.html")

	req, _ := http.NewRequest("GET", "https://api.piloterr.com/v2/prnewswire/article?"+params.Encode(), nil)
	req.Header.Set("x-api-key", os.Getenv("PILOTERR_API_KEY"))

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Piloterr error %d: %s\n", resp.StatusCode, body)
		os.Exit(1)
	}

	var pretty bytes.Buffer
	json.Indent(&pretty, body, "", "  ")
	fmt.Println(pretty.String())
}
