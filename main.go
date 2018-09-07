package main


import (
	"os/user"
	"os"
	"path/filepath"
	"fmt"
	//Sync "./Sync"
	//Restore "./Restore"
	Utils "./Utils"
	Remitter "./Remitter"
	Watcher "./Watcher"
	//Reader "./Reader"
)

func main(){
	usr, _ := user.Current()
	//fmt.Println(usr.HomeDir)
	//Restore.Start()

  if (len(os.Args) > 1){
		Utils.ServerUrl = fmt.Sprintf(Utils.ServerUrl, os.Args[1])
  }
	if (len(os.Args) > 2){
		Utils.Tag = os.Args[2]
  }

	go Watcher.Start(SendDataServer, usr.HomeDir + "/.taskbook/storage")
	Watcher.Start(SendDataServer, usr.HomeDir + "/.taskbook/archive")
}

func SendDataServer(Resources map[string]Utils.Resource, filePath string){
  var targetName string = filepath.Base(filepath.Dir(filePath))
  Remitter.SendDataToServer(Resources, targetName + "_" + Utils.Tag, Utils.ServerUrl)
}

