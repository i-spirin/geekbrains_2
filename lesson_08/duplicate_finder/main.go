package main

import (
	"flag"
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
	defer wg.Wait()

	go file_walk.Search(*path, &wg, toCheck)

	info := files_info.New()

	for i := 0; i < *threads; i++ {
		go file_walk.CheckFiles(toCheck, &wg, &info)
	}
}
