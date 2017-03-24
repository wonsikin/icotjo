package parser

import (
	"encoding/json"
	"testing"
)

func TestMapJson(t *testing.T) {
	mapList := map[string]string{"用户名": "User,name", "密码": "Password"}

	data, err := json.Marshal(mapList)
	if err != nil {
		t.Errorf("fail %s", err)
	}

	t.Logf("result is %s", string(data))
}
