package mapper

import (
	"encoding/json"
)

func Map(source interface{}, target interface{}) {
	data, _ := json.Marshal(source)
	_ = json.Unmarshal(data, &target)
}
