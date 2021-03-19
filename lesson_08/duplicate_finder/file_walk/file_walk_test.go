package file_walk_test

import (
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/i-spirin/geekbrains_2/lesson_08/duplicate_finder/file_walk"
	files_info "github.com/i-spirin/geekbrains_2/lesson_08/duplicate_finder/files_hash_info"
)

func createTemp() (string, string) {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	tmpDirectory := "/tmp/" + string(b)
	tmpFile := tmpDirectory + "/" + string(b)

	return tmpDirectory, tmpFile

}

func TestSearch(t *testing.T) {
	tmpDirectory, tmpFile := createTemp()

	err := os.Mkdir(tmpDirectory, 0755)
	if err != nil {
		t.Errorf("Cannot create directory")
	}
	defer os.RemoveAll(tmpDirectory)
	err = os.Mkdir(tmpDirectory+"/1", 0755)
	if err != nil {
		t.Errorf("Cannot create directory")
	}
	file, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Errorf("Cannot create file")
	}
	defer file.Close()

	found := file_walk.Search(tmpDirectory)

	if len(found) != 1 {
		t.Error("Found more than 1 file")
	}
	if found[0] != tmpFile {
		t.Errorf("Unexpected found")
	}

}

func TestFileMD5(t *testing.T) {
	tmpDirectory, tmpFile := createTemp()
	err := os.Mkdir(tmpDirectory, 0755)
	if err != nil {
		t.Errorf("Error creating tmpDirectory: %v", err)
	}
	defer os.RemoveAll(tmpDirectory)

	file, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Errorf("Error creating tmpFile: %v", err)
	}

	file.Write([]byte("123"))
	file.Close()

	sum, err := file_walk.FileMD5(tmpFile)
	if err != nil {
		t.Errorf("Error counting hash %v", err)
	}
	if sum != "202cb962ac59075b964b07152d234b70" {
		t.Errorf("Unexpected sum got")
	}

	_, err = file_walk.FileMD5(tmpFile + "1")
	if err == nil {
		t.Error("file_walk.FileMD5 does not returned error when expected")
	}

}

func TestCheckFiles1(t *testing.T) {
	tmpDirectory, tmpFile := createTemp()

	err := os.Mkdir(tmpDirectory, 0755)
	if err != nil {
		t.Errorf("Error creating tmpDirectory: %v", err)
	}
	defer os.RemoveAll(tmpDirectory)

	file, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Errorf("Error creating tmpFile: %v", err)
	}

	file.Write([]byte("123"))
	file.Close()

	ch := make(chan string, 1)
	ch <- tmpFile
	close(ch)

	wg := sync.WaitGroup{}
	wg.Add(1)
	info := files_info.New()
	errorChannel := make(chan error, 10)
	duplicates := make(chan string)

	go file_walk.CheckFiles(ch, &wg, &info, errorChannel, duplicates)
	wg.Wait()
	close(errorChannel)

	err = <-errorChannel
	if err != nil {
		t.Errorf("CheckFiles returned error: %v", err)
	}
}

func TestCheckFiles2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	info := files_info.New()
	errorChannel := make(chan error, 10)
	duplicates := make(chan string)

	ch := make(chan string, 1)
	ch <- "/tmp/guess-this-is-unexisting-file"
	close(ch)

	go file_walk.CheckFiles(ch, &wg, &info, errorChannel, duplicates)
	wg.Wait()

	err := <-errorChannel
	if err == nil {
		t.Errorf("CheckFiles does not returned error when expected")
	}
}

func TestCheckFiles3(t *testing.T) {
	tmpDirectory, tmpFile := createTemp()

	err := os.Mkdir(tmpDirectory, 0755)
	if err != nil {
		t.Errorf("Error creating tmpDirectory: %v", err)
	}
	defer os.RemoveAll(tmpDirectory)

	file, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Errorf("Error creating tmpFile: %v", err)
	}

	file.Write([]byte("123"))
	file.Close()

	file, err = os.OpenFile(tmpFile+"1", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Errorf("Error creating tmpFile: %v", err)
	}

	file.Write([]byte("123"))
	file.Close()

	ch := make(chan string, 2)
	ch <- tmpFile
	ch <- tmpFile + "1"
	close(ch)

	wg := sync.WaitGroup{}
	wg.Add(1)
	info := files_info.New()
	errorChannel := make(chan error, 10)
	duplicates := make(chan string, 10)

	go file_walk.CheckFiles(ch, &wg, &info, errorChannel, duplicates)
	wg.Wait()
	close(duplicates)

	err = <-errorChannel
	if err == nil {
		t.Errorf("CheckFiles does not returned error when expected")
	}

	dup, ok := <-duplicates
	if !ok {
		t.Errorf("CheckFiles does not found any duplicate")
	}
	if dup != tmpFile && dup != tmpFile+"1" {
		t.Errorf("CheckFiles found wrong duplicate: expected `%s' or `%s', got `%s'", tmpFile, tmpFile+"1", dup)
	}
}
