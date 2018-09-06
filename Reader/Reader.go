package Reader


import (
	"os"
	//"fmt"
	"encoding/json"
	"io/ioutil"
)

type Resource struct {
	Id uint64 `json:"_id"`
	Date string `json:"_date"`
	Tmestamp int `json:"_timestamp"`
	Description string `json:"description"`
	IsStarred bool `json:"isStarred"`
	_isTask bool `json:"_isTask"`
	IsComplete bool `json:"isComplete"`
	Priority int `json:"priority"`
}

func GetResourceData() []Resource{
	jsonFile, _ := os.Open("/home/david/.taskbook/storage/storage.json")
	/*if err != nil {
		fmt.Println(err)
	}*/
	//fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	//fmt.Println(jsonFile)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//resources : = map[string]Resource{}
	var resources map[string]Resource
	//fmt.Println(reflect.TypeOf(byteValue))
	json.Unmarshal([]byte(byteValue), &resources)
	//fmt.Println(string(byteValue))
	//fmt.Printf("total items: %d \n", len(resources))
	//fmt.Println(resources)

	defer jsonFile.Close()
	var results []Resource
	for _, row := range resources {
	    //keys = append(keys, key)
	    //fmt.Println(key)
	    //fmt.Println(row.Description)
	    //fmt.Println(row)
	    results = append(results, row)
	}
	return results


	//for i := 0; i < len(resources); i++ {
		//fmt.Println(resources.2.name)
		//fmt.Println("xxxxx")
		//fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		//fmt.Println("User Name: " + users.Users[i].Name)
		//fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	//}
}
