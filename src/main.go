// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var version = "version"
var url = "https://github.com/yuki-bara/glc/blob/main"

type License struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: lses <license-name>")
	}
	var wg sync.WaitGroup
	for i := 1; i < len(os.Args); i++ {
		wg.Add(1)
		go func(licenseName string) {
			defer wg.Done()
			fmt.Printf("GLC-%s  %s\n", version, url)

			url := "https://api.github.com/licenses/" + licenseName

			resp, err := http.Get(url)
			if err != nil {
				log.Fatalf("Error fetching license: %v", err)
			}

			if resp.StatusCode != http.StatusOK {
				log.Fatalf("Error: License '%s' not found (HTTP %d)", licenseName, resp.StatusCode)
			}
			defer resp.Body.Close()
			jsonData, _ := io.ReadAll(resp.Body)
			var lic License
			err = json.Unmarshal(jsonData, &lic)
			if err != nil {
				fmt.Printf("Error parsing JSON: %v\n", err)
				return
			}
			err = os.WriteFile(lic.Name, []byte(lic.Body), 0644)
			if err != nil {
				log.Fatalf("Error writing file: %v", err)
				return
			}

			fmt.Printf("Successfully generated %s LICENSE file!\n", licenseName)
		}(os.Args[i])
	}
	wg.Wait()

}
