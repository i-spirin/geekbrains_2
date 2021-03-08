package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
)

var (
	wg      sync.WaitGroup
	ch      = make(chan struct{}, *threads)
	threads = flag.Int("threads", 1, "Number of threads to start")
	mm      = make([]int, 0)
	mu      sync.RWMutex
)

func main() {

	minPort := flag.Int("min_port", 21, "Port to start from")
	maxPort := flag.Int("max_port", 65535, "End port")
	host := flag.String("host", "127.0.0.1", "Host to scan")
	flag.Parse()

	wg.Add(*maxPort - *minPort + 1)

	for port := *minPort; port <= *maxPort; port++ {
		ch <- struct{}{}
		go checkPort(*host, port)
	}

	wg.Wait()
	log.Println("Opened ports:", mm)
}

func checkPort(host string, port int) bool {
	conn, err := net.Dial("tcp", host+":"+fmt.Sprint(port))
	defer func() {
		<-ch
	}()
	defer wg.Done()
	if err != nil {
		// log.Println(err)
		// log.Printf("Port %d is closed\n", port)
		return false
	}
	defer conn.Close()
	// log.Printf("Port %d is open", port)
	appendOpenedPort(port)
	return true
}

func appendOpenedPort(port int) {
	mu.RLock()
	defer mu.RUnlock()

	mm = append(mm, port)

}
