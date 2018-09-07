package Watcher

import (
	"log"
	//"fmt"
	Utils "../Utils"
	Reader "../Reader"
	"github.com/fsnotify/fsnotify"
)

type fn func(Resources map[string]Utils.Resource, filePath string)

func Start(doMethod fn, dir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//fmt.Println(event.Name)
				data := Reader.GetResourceData(event.Name)
				doMethod(data, event.Name)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

