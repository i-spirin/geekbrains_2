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
	ch := make(chan string, 2)

	tmpDirectory, tmpFile := createTemp()

	wg := sync.WaitGroup{}
	wg.Add(1)

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

	go file_walk.Search(tmpDirectory, &wg, ch)
	wg.Wait()

	foundFile := <-ch

	if tmpFile != foundFile {
		t.Errorf("Unexpected file found %s", foundFile)
	}

	if _, ok := <-ch; ok {
		t.Errorf("Channel is not closed")
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

	go file_walk.CheckFiles(ch, &wg, &info, errorChannel)
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

	ch := make(chan string, 1)
	ch <- "/tmp/guess-this-is-unexisting-file"
	close(ch)

	go file_walk.CheckFiles(ch, &wg, &info, errorChannel)
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

	go file_walk.CheckFiles(ch, &wg, &info, errorChannel)
	wg.Wait()

	err = <-errorChannel
	if err == nil {
		t.Errorf("CheckFiles does not returned error when expected")
	}
}
