package main


import (
	"fmt"
	. "./Remitter"
	. "./Watcher"
	. "./Reader"
)

func main(){
	fmt.Println("I going to start!!")
	Start(onDataChange)
}

func onDataChange(){
	fmt.Println("Anything it was changed!!")
	data := GetResourceData()
	fmt.Println(data[0].Id)
	SendDataToServer(data)
}
