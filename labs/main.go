package main

import (
  "fmt"
  "os/user"
)


func main(){
  //usr, _ := user.Current()
  fmt.Println((*user.Current()).HomeDir)
}
