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
	for _, site := range sites {
		// start a timer for this request
		begin := time.Now()

		// Retrieve the site
		if _, err := http.Get(site); err != nil {
			fmt.Println(site, err)
			continue
		}

		fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
	}
}

//END PROCESS-OMIT
