package Reader


import (
	"os"
	Utils "../Utils"
	//"os/user"
	"encoding/json"
	"io/ioutil"
)

func GetResourceData(filePath string) []Utils.Resource{
	jsonFile, _ := os.Open(filePath)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var resources map[string]Utils.Resource
	json.Unmarshal([]byte(byteValue), &resources)
	defer jsonFile.Close()
	var results []Utils.Resource
	for _, row := range resources {
	    results = append(results, row)
	}
	return results
}
