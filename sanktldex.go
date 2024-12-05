package main

import (
        "bufio"
        "flag"
        "fmt"
        "os"
        "strings"
        "sync"

        "github.com/jpillora/go-tld"
)

func main() {
        // Flags for command line arguments
        concurrencyPtr := flag.Int("t", 8, "Number of threads to utilise. Default is 8.")
        subdomainsPtr := flag.Bool("s", false, "Dump subdomains instead of base domains")
        filterDomainPtr := flag.String("f", "", "Extract subdomains only for this domain (single domain).")
        filterDomainsListPtr := flag.String("fL", "", "Extract subdomains only for these domains from a file.")
        flag.Parse()

        var domainsList []string
        filterDomain := strings.ToLower(*filterDomainPtr)
        filterDomainsList := strings.ToLower(*filterDomainsListPtr)

        // If both -f and -fL are set, prioritize -fL (list of domains)
        if filterDomainsList != "" {
                // Read the domains from the file provided with -fL
                domainsList = readDomainsFromFile(filterDomainsList)
        } else if filterDomain != "" {
                domainsList = append(domainsList, filterDomain)
        }

        numWorkers := *concurrencyPtr
        work := make(chan string)

        go func() {
                s := bufio.NewScanner(os.Stdin)
                for s.Scan() {
                        work <- s.Text()
                }
                close(work)
        }()

        wg := &sync.WaitGroup{}

        for i := 0; i < numWorkers; i++ {
                wg.Add(1)
                go doWork(work, wg, *subdomainsPtr, domainsList)
        }
        wg.Wait()
}

func readDomainsFromFile(filename string) []string {
        file, err := os.Open(filename)
        if err != nil {
                fmt.Println("Error opening file:", err)
                return nil
        }
        defer file.Close()

        var domainsList []string
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                domainsList = append(domainsList, strings.ToLower(scanner.Text()))
        }

        if err := scanner.Err(); err != nil {
                fmt.Println("Error reading file:", err)
        }

        return domainsList
}

func doWork(work chan string, wg *sync.WaitGroup, subdomainsPtr bool, domainsList []string) {
        defer wg.Done()
        for url := range work {
                result, err := tld.Parse(url)
                if err != nil {
                        continue
                }

                domain := strings.ToLower(result.Domain + "." + result.TLD)
                // Check if the domain matches any of the provided domains
                if len(domainsList) == 0 || contains(domainsList, domain) {
                        if subdomainsPtr && result.Subdomain != "" {
                                fmt.Println(result.Subdomain + "." + domain)
                        } else if !subdomainsPtr {
                                fmt.Println(domain)
                        }
                }
        }
}

// contains checks if a domain exists in the list of domains
func contains(domainsList []string, domain string) bool {
        for _, d := range domainsList {
                if d == domain {
                        return true
                }
        }
        return false
}
