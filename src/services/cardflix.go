package service

import (
	"encoding/json"
	"log"
	database "serverless/src/database/mongo"
	model "serverless/src/models"
)

func AddCardflix(cardflix model.Cardflix) (string, error) {
	var object_id, err = database.AddData("cardflix", cardflix)
	if err != nil {
		return "", err
	}
	return object_id, nil
}

func GetAllCardflix() []model.Cardflix {
	var data = database.GetAllData("cardflix")

	var result = make([]model.Cardflix, len(data))

	for i, d := range data {
		var dataJson, err = json.Marshal(d)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(dataJson, &result[i])
	}

	return result
}
