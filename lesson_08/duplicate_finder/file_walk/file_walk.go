package file_walk

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	files_info "github.com/i-spirin/geekbrains_2/lesson_08/duplicate_finder/files_hash_info"
)

func Search(rootPath string, wg *sync.WaitGroup, toCheck chan string) {
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			log.Println("Found", path)
			toCheck <- path
		}
		return nil
	})
	close(toCheck)
	wg.Done()
}

func CheckFiles(toCheck <-chan string, wg *sync.WaitGroup, info *files_info.FilesHashInfo) {
	log.Println("Starting checkFiles")
	defer wg.Done()

	for {
		path := <-toCheck
		if path == "" {
			log.Println("checkFiles end")
			return
		}

		sum, err := FileMD5(path)
		if err != nil {
			log.Println("An error occured when reading file:", err)
		}

		err = info.Add(path, sum)
		if err != nil {
			log.Printf("Error checking %s: %v", path, err)
		}
	}
}

func FileMD5(path string) (string, error) {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
