package core

import (
	"github.com/chanify/chanify/model"
	"github.com/gin-gonic/gin"
)

func (c *Core) handleImageDownload(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func (c *Core) handleAudioDownload(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func (c *Core) handleFileDownload(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func (c *Core) downloadImageFile(ctx *gin.Context, token *model.Token) {
	_ = "STUB: not implemented"
	return
}

func (c *Core) downloadAudioFile(ctx *gin.Context, token *model.Token) {
	_ = "STUB: not implemented"
	return
}

func (c *Core) downloadFile(ctx *gin.Context, token *model.Token) {
	_ = "STUB: not implemented"
	return
}
