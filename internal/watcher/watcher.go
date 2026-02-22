package watcher

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// FileWatcher watches a file for changes
type FileWatcher struct {
	filename string
	lastMod  time.Time
	callback func()
}

// NewFileWatcher creates a new file watcher
func NewFileWatcher(filename string, callback func()) *FileWatcher {
	info, err := os.Stat(filename)
	if err != nil {
		return nil
	}

	return &FileWatcher{
		filename: filename,
		lastMod:  info.ModTime(),
		callback: callback,
	}
}

// Watch starts watching the file
func (fw *FileWatcher) Watch(interval time.Duration) error {
	fmt.Printf("👀 Watching %s for changes... (Press Ctrl+C to stop)\n\n", fw.filename)

	for {
		time.Sleep(interval)

		info, err := os.Stat(fw.filename)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("⚠️  File deleted: %s\n", fw.filename)
				continue
			}
			return err
		}

		if info.ModTime().After(fw.lastMod) {
			fw.lastMod = info.ModTime()
			fmt.Printf("\n🔄 File changed: %s\n", filepath.Base(fw.filename))
			fmt.Println(string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}))
			fw.callback()
			fmt.Println(string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}) + string([]rune{0x2500}))
			fmt.Println()
		}
	}
}
