package main


import (
	Rm "./Remitter"
	W "./Watcher"
	R "./Reader"
)

func main(){
//	fmt.Println("I going to start!!")
	W.Start(onDataChange)
}

func onDataChange(){
//	fmt.Println("Anything it was changed!!")
	data := R.GetResourceData()
//	fmt.Println(data[0].Id)
	Rm.SendDataToServer(data)
}
