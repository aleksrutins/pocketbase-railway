package watchman

import "os"

type Watcher struct {
	DBPath  string
	fileRef *os.File
}

func NewWatcher(dbPath string) Watcher {
	return Watcher{
		DBPath: dbPath,
	}
}

func (w *Watcher) Start() error {
	//fileRef, err := os.OpenFile(w.DBPath, os.O_RDONLY, fs.Mode)
	return nil
}
