package dbkit

import (
	"encoding/json"
	"fmt"
	"log"
)

func ParseToStruct[T any](data any) (dest []T) {
	buff, err := json.Marshal(&data)
	if err != nil {
		log.Println(err)
		return []T{}
	}

	fmt.Println(string(buff))
	var out []T
	err = json.Unmarshal(buff, &out)
	if err != nil {
		log.Println(err)
		return []T{}
	}

	dest = out
	return
}

func responseQueryErr(err error) ResponseQuery {
	return ResponseQuery{
		Success:   false,
		Message:   err.Error(),
		ItemFound: false,
		Items:     []map[string]interface{}{},
	}
}

func responseQuerySuccess(msg string, items []map[string]interface{}) ResponseQuery {
	return ResponseQuery{
		Success:   true,
		Message:   msg,
		ItemFound: len(items) != 0,
		Items:     items,
	}
}
