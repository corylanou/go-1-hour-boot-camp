//START IMPORT-OMIT
package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

//END IMPORT-OMIT

//START MAIN-OMIT
func main() {
	// Parse all arguments
	flag.Parse()

	// flag.Args contains all non-flag arguments
	sites := flag.Args()

	// Lets keep a reference to when we started
	start := time.Now()

	processSites(sites)

	fmt.Printf("Entire process took %s\n", time.Since(start))
}

//END MAIN-OMIT

//START PROCESS-OMIT
func processSites(sites []string) {
	for i, site := range sites {
		// start a timer for this request
		begin := time.Now()

		// Retrieve the site
		resp, err := http.Get(site)
		if err != nil {
			fmt.Println(site, err)
			continue
		}

		fmt.Printf("%d) Site %q took %s to retrieve with status code of %d.\n", i, site, time.Since(begin), resp.StatusCode)
	}
}

//END PROCESS-OMIT
