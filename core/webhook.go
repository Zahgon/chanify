package core

import (
	"io"

	"github.com/gin-gonic/gin"
	lua "github.com/yuin/gopher-lua"
)

const coreKey = "_chanify/http/core"

func (c *Core) handlePostWebhook(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func initHttpLua(l *lua.LState, ctx *gin.Context) { _ = "STUB: not implemented"; return }

func getHttpLuaReturn(l *lua.LState) (int, string, string) {
	_ = "STUB: not implemented"
	return 0, "", ""
}

var luaContextMethods = map[string]lua.LGFunction{
	"request": luaContextGetRequest,
	"send":    luaContextSend,
}

var luaRequestMethods = map[string]lua.LGFunction{
	"token":  luaContextGetToken,
	"url":    luaContextGetUrl,
	"body":   luaContextGetBody,
	"query":  luaContextGetQuery,
	"header": luaContextGetHeader,
}

func luaCheckContext(l *lua.LState) *gin.Context { _ = "STUB: not implemented"; return nil }

func luaContextGetRequest(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

type luaSendContext struct {
	msg string
}

func luaContextSend(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

func (l *luaSendContext) String() string { _ = "STUB: not implemented"; return "" }

func (l *luaSendContext) JSON(code int, obj interface{}) { _ = "STUB: not implemented"; return }

func (l *luaSendContext) DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string) {
	_ = "STUB: not implemented"
	return
}

func luaGetOptsString(opts *lua.LTable, key string) string { _ = "STUB: not implemented"; return "" }

func luaGetOptsArray(opts *lua.LTable, key string) []string { _ = "STUB: not implemented"; return nil }

func luaContextGetToken(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

func luaContextGetUrl(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

func luaContextGetBody(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

func luaContextGetQuery(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }

func luaContextGetHeader(l *lua.LState) int { _ = "STUB: not implemented"; return 0 }
