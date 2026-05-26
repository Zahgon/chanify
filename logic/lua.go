package logic

import (
	lua "github.com/yuin/gopher-lua"
)

var luaMods = map[string]map[string]lua.LGFunction{
	"hex": {
		"decode": luaHEXDecode,
		"encode": luaHEXEncode,
	},
	"json": {
		"decode": luaJsonDecode,
		"encode": luaJsonEncode,
	},
	"crypto": {
		"equal": luaCryptoEqual,
		"hmac":  luaCryptoHmac,
	},
}

func luaHEXDecode(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

func luaHEXEncode(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

func luaJsonDecode(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

func luaJsonEncode(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

// nolint:errcheck

func luaInterface2LValue(v interface{}) lua.LValue {
	_ = "STUB: not implemented"
	return *new(lua.LValue)
}

func luaLValue2Interface(v lua.LValue) interface{} { _ = "STUB: not implemented"; return nil }

func luaCryptoEqual(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

func luaCryptoHmac(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

func initLua(l *lua.LState) { _ = "STUB: not implemented"; return }
