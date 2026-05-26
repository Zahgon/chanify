package core

import (
	"errors"
	"net/http"

	"github.com/chanify/chanify/logic"
	"github.com/gin-gonic/gin"
)

// error define
var (
	ErrNoContent       = errors.New("NoContent")
	ErrTooLargeContent = errors.New("TooLargeContent")
	ErrInvalidContent  = errors.New("InvalidContent")
)

// Core instance
type Core struct {
	logic *logic.Logic
}

// New core instance
func New() *Core { _ = "STUB: not implemented"; return nil }

// Init core with option
func (c *Core) Init(opts *logic.Options) error { _ = "STUB: not implemented"; return nil }

// Close & cleaup for core
func (c *Core) Close() { _ = "STUB: not implemented"; return }

// APIHandler return handler for http
func (c *Core) APIHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// nolint: errcheck

func (c *Core) handleHome(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func (c *Core) handleInfo(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func (c *Core) handleQRCode(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func loggerMiddleware(c *gin.Context) { _ = "STUB: not implemented"; return }
