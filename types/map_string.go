package types

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type MapStrings map[string]string

// Set 添加值 force=true 覆盖原值
func (maps MapStrings) Set(k, v string) {
	maps.SetForce(k, v, true)
}

// SetForce Set 添加值 force=true 覆盖原值
func (maps MapStrings) SetForce(k, v string, force bool) {
	if maps.Exist(k) {
		if force {
			maps[k] = v
		}
		return
	}
	maps[k] = v
}

// Exist 键值是否存在
func (maps MapStrings) Exist(k string) bool {
	_, ok := maps[k]
	return ok
}

// Delete 删除指定到key值
func (maps MapStrings) Delete(k string) {
	delete(maps, k)
}

// GetDefault 获取值并携带默认值
func (maps MapStrings) GetDefault(k, d string) string {
	if v, ok := maps[k]; ok {
		return v
	}
	return d
}

// Keys keys
func (maps MapStrings) Keys() []string {
	keys := make([]string, len(maps))
	for k := range maps {
		keys = append(keys, k)
	}
	return keys
}

// Values values
func (maps MapStrings) Values() []string {
	values := make([]string, len(maps))
	for _, v := range maps {
		values = append(values, v)
	}
	return values
}

// Implode values 转字符串
func (maps MapStrings) Implode(sep string) string {
	if len(maps) <= 0 {
		return ""
	}
	buf := new(bytes.Buffer)
	for _, s := range maps {
		buf.Write([]byte(s))
		buf.Write([]byte(sep))
	}
	str := buf.String()
	return str[:len(str)-1]
}

// JsonEncode 转json字符串
func (maps MapStrings) JsonEncode() string {
	marshal, err := json.Marshal(maps)
	if err != nil {
		panic(fmt.Sprintf("MapStrings ToJson err=%v", err))
	}
	return string(marshal)
}

// JsonDecode 解析json到实体上
func (maps MapStrings) JsonDecode(v any) error {
	return json.Unmarshal([]byte(maps.JsonEncode()), v)
}
