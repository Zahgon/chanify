package logic

import (
	lua "github.com/yuin/gopher-lua"
)

func fixPath(path string) error { _ = "STUB: not implemented"; return nil }

func saveFile(path string, data []byte) error { _ = "STUB: not implemented"; return nil }

func readOptString(opts map[string]interface{}, key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func readOptTable(opts map[string]interface{}, key string) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func compileLua(filePath string) (*lua.FunctionProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
