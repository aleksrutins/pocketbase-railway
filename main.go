package main

import (
	"log"
	"os"
	"sync"

	"github.com/aleksrutins/pocketbase-railway/watchman"
	"github.com/pocketbase/pocketbase"
)

func main() {
	var wg sync.WaitGroup

	app := pocketbase.New()

	watcher := watchman.NewWatcher("pb_data/data.db")

	wg.Add(2)
	go func() {
		if err := app.Start(); err != nil {
			log.Fatal(err)
			wg.Done()
			os.Exit(1)
		}
		wg.Done()
	}()
	go func() {
		watcher.Start()
		wg.Done()
	}()

	wg.Wait()
}
