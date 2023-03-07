package types

import (
	"encoding/json"
	"fmt"
)

type MapStringAny map[string]any

// Set 添加值 force=true 覆盖原值
func (maps MapStringAny) Set(k string, v any) {
	maps.SetForce(k, v, true)
}

// SetForce Set 添加值 force=true 覆盖原值
func (maps MapStringAny) SetForce(k string, v any, force bool) {
	if maps.Exist(k) {
		if force {
			maps[k] = v
		}
		return
	}
	maps[k] = v
}

// Exist 键值是否存在
func (maps MapStringAny) Exist(k string) bool {
	_, ok := maps[k]
	return ok
}

func (maps MapStringAny) Delete(k string) {
	delete(maps, k)
}

func (maps MapStringAny) Keys() []string {
	keys := make([]string, len(maps))
	for key, _ := range maps {
		keys = append(keys, key)
	}
	return keys
}
func (maps MapStringAny) Values() []any {
	values := make([]any, len(maps))
	for _, value := range maps {
		values = append(values, value)
	}
	return values
}

// GetDefault 获取value 要设置默认值
func (maps MapStringAny) GetDefault(k string, d any) any {
	if v, ok := maps[k]; ok {
		return v
	}
	return d
}

// JsonEncode  转json 字节流
func (maps MapStringAny) JsonEncode() string {
	marshal, err := json.Marshal(maps)
	if err != nil {
		panic(fmt.Sprintf("MapStringAny ToJson err=%v", err))
	}
	return string(marshal)
}

// JsonDecode 解析json到实体上
func (maps MapStringAny) JsonDecode(v any) error {
	return json.Unmarshal([]byte(maps.JsonEncode()), v)
}
