package files_info

import (
	"fmt"
	"sync"
)

type FilesHashInfo struct {
	Files map[string]string
	mu    sync.RWMutex
}

func New() FilesHashInfo {
	return FilesHashInfo{
		Files: make(map[string]string),
		mu:    sync.RWMutex{},
	}
}

func (f *FilesHashInfo) Add(path string, hash string) error {
	f.mu.RLock()
	defer f.mu.RUnlock()
	if existPath, isExists := f.Files[hash]; isExists {
		return fmt.Errorf("hash %s already found by path %s", hash, existPath)
	}
	f.Files[hash] = path
	return nil
}

func (f *FilesHashInfo) Get(hash string) (string, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	if path, isExists := f.Files[hash]; isExists {
		return path, nil
	}
	return "", fmt.Errorf("%s not found", hash)
}
