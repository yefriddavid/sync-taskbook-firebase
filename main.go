package main


import (
	"os/user"
	"path/filepath"
	//"fmt"
	Utils "./Utils"
	Remitter "./Remitter"
	Watcher "./Watcher"
	//Reader "./Reader"
)

func main(){
	usr, _ := user.Current()
	go Watcher.Start(onStorageDataChange, usr.HomeDir + "/.taskbook/storage")
	Watcher.Start(onArchiveDataChange, usr.HomeDir + "/.taskbook/archive")
}

func onStorageDataChange(Resources []Utils.Resource, filePath string){
	var targetName string = filepath.Base(filepath.Dir(filePath))
	Remitter.SendDataToServer(Resources, targetName + "_" + Utils.Tag)
}

func onArchiveDataChange(Resources []Utils.Resource, filePath string){
	var targetName string = filepath.Base(filepath.Dir(filePath))
	Remitter.SendDataToServer(Resources, targetName + "_" + Utils.Tag)
}
