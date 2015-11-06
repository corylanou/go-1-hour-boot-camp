//START IMPORT-OMIT
package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

//END IMPORT-OMIT

//START MAIN-OMIT
func main() {
	flag.Parse()
	sites := flag.Args()
	// Args might be a single string with spaces to delimit sites, break them apart
	if len(sites) == 1 {
		sites = strings.Split(sites[0], " ")
	}

	// Use a waitgroup to make sure all go routines finish
	var wg sync.WaitGroup

	// Lets keep a reference to when we started
	start := time.Now()
	// Set the value for the waitgroup

	wg.Add(len(sites))
	processSites(sites, &wg)
	// Block until all routines finish
	wg.Wait()

	fmt.Printf("Entire process took %s\n", time.Since(start))
}

//END MAIN-OMIT

// START PROCESS-OMIT
func processSites(sites []string, wg *sync.WaitGroup) {
	for _, site := range sites {
		// Launch each retrieval in a go routine.  This makes each request concurrent
		go func(site string) {
			defer wg.Done()

			// start a timer for this request
			begin := time.Now()

			// Retrieve the site
			if _, err := http.Get(site); err != nil {
				fmt.Println(site, err)
				return
			}

			fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
		}(site)
	}
}

// END PROCESS-OMIT
