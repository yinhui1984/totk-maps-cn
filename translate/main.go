package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func loadJsonData() *MapData {
	file, err := os.Open("../data_cn.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	var mapData MapData
	json.Unmarshal(data, &mapData)
	return &mapData
}

func saveJsonData(mapData *MapData) {
	file, err := os.Create("../data_cn.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := json.Marshal(mapData)
	if err != nil {
		fmt.Println(err)
	}

	file.Write(data)
}

func main() {
	mapData := loadJsonData()

	for index, feature := range mapData.Features {
		// trans := traslate(feature.Properties.Title)
		// fmt.Println(feature.Properties.Title, " => ", trans)
		// mapData.Features[index].Properties.Title = trans
		trans := traslate(feature.Properties.Subcat)
		fmt.Println(feature.Properties.Subcat, " => ", trans)
		mapData.Features[index].Properties.Subcat = trans
	}

	saveJsonData(mapData)

}
