package main

import (
	"flag"
	"log"
	"sync"

	"github.com/i-spirin/geekbrains_2/lesson_08/duplicate_finder/file_walk"
	files_info "github.com/i-spirin/geekbrains_2/lesson_08/duplicate_finder/files_hash_info"
)

func main() {

	path := flag.String("path", "./", "Root directory to start searching")
	threads := flag.Int("threads", 1, "Workers to start")
	flag.Parse()

	toCheck := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(1 + *threads)

	go file_walk.Search(*path, &wg, toCheck)

	info := files_info.New()

	errorChannel := make(chan error, 10)

	for i := 0; i < *threads; i++ {
		go file_walk.CheckFiles(toCheck, &wg, &info, errorChannel)
	}
	wg.Wait()
	close(errorChannel)

	for {
		err, ok := <-errorChannel
		if !ok {
			break
		} else {
			log.Println(err)
		}
	}
}
