package Utils

var Tag string = "Personal"
var ServerUrl string = "Personal"


type Resource struct {
	Id uint64 `json:"_id"`
	Date string `json:"_date"`
	Tmestamp int `json:"_timestamp"`
	Description string `json:"description"`
	IsStarred bool `json:"isStarred"`
	_isTask bool `json:"_isTask"`
	IsComplete bool `json:"isComplete"`
	Boards []string `json:"boards"`
	Priority int `json:"priority"`
}

