package Watcher

import (
	"log"
	"github.com/fsnotify/fsnotify"
)

type fn func()

func Start(doMethod fn) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case _ , ok := <-watcher.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				doMethod()
				/*if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}*/
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/home/david/.taskbook/storage")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

