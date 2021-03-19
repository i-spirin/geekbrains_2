package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/i-spirin/geekbrains_2/lesson_08/duplicate_finder/file_walk"
	files_info "github.com/i-spirin/geekbrains_2/lesson_08/duplicate_finder/files_hash_info"
)

func main() {

	path := flag.String("path", "./", "Root directory to start searching")
	threads := flag.Int("threads", 1, "Workers to start")
	remove := flag.Bool("remove", false, "Should application remove all duplicates")
	flag.Parse()

	wg := sync.WaitGroup{}
	wg.Add(*threads)

	found := file_walk.Search(*path)
	toCheck := make(chan string, len(found))
	for _, item := range found {
		toCheck <- item
		log.Println("Put", item)
	}
	close(toCheck)
	info := files_info.New()

	errorChannel := make(chan error, len(found))
	duplicates := make(chan string, len(found)*2)

	for i := 0; i < *threads; i++ {
		go file_walk.CheckFiles(toCheck, &wg, &info, errorChannel, duplicates)
	}
	wg.Wait()
	close(errorChannel)
	close(duplicates)

	for {
		err, ok := <-errorChannel
		if !ok {
			break
		} else {
			log.Println(err)
		}
	}
	if *remove {
		for {
			file, ok := <-duplicates
			if !ok {
				break
			}
			yes := ""
			fmt.Printf("Remove %s? (yes/NO):", file)
			if _, err := fmt.Scan(&yes); err != nil {
				log.Println(err)
			}
			if yes == "yes" {
				log.Println("Removing", file)
				err := os.Remove(file)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
