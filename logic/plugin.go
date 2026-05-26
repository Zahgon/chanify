package logic

import (
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	lua "github.com/yuin/gopher-lua"
)

type luaFunc struct {
	modTime time.Time
	lfunc   *lua.FunctionProto
}

// Webhook item
type Webhook struct {
	name  string
	env   map[string]interface{}
	lfunc *luaFunc
}

type pluginManager struct {
	watcher  *fsnotify.Watcher
	luaMutex sync.Mutex
	luaFiles map[string]*luaFunc
	webHooks map[string]*Webhook
}

func loadWebhookPlugin(path string, wbOpts []map[string]interface{}) *pluginManager {
	_ = "STUB: not implemented"
	return nil
}

func (p *pluginManager) Close() { _ = "STUB: not implemented"; return }

func (p *pluginManager) GetWebhook(name string) (*Webhook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *pluginManager) loadLuaFile(file string, path string) *luaFunc {
	_ = "STUB: not implemented"
	return nil
}

// nolint: errcheck

func (p *pluginManager) luaWatch() { _ = "STUB: not implemented"; return }

func (p *pluginManager) ReloadWebhook(file string) { _ = "STUB: not implemented"; return }

func (lf *luaFunc) Reload(file string) error { _ = "STUB: not implemented"; return nil }

func (w *Webhook) DoCall(l *lua.LState) error { _ = "STUB: not implemented"; return nil }
