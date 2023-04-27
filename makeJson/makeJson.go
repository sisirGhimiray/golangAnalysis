package makejson

import (
	"encoding/json"
	"log"
)

func MakeJsonData(d interface{}) []byte {
	jsonD, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	return jsonD
}
