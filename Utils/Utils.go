package Utils


import (
  "os/user"
)

var Tag string = "Sync"
var FirebaseProject string = "sync-00000"
var ServerUrl string = "https://%s.firebaseio.com"


func GetAccountPathFileKey() string{
  usr, _ := user.Current()
  return usr.HomeDir + "/.ssh/serviceAccountKey-taskbook.json"
}


type Resource struct {
	Id uint64 `json:"_id"`
	Date string `json:"_date"`
	Tmestamp int `json:"_timestamp"`
	Description string `json:"description"`
	IsStarred bool `json:"isStarred"`
	IsTask bool `json:"_isTask"`
	IsComplete bool `json:"isComplete"`
	Boards []string `json:"boards"`
	Priority int `json:"priority"`
}

