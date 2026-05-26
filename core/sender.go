package core

import (
	"io"

	"github.com/chanify/chanify/model"
	"github.com/gin-gonic/gin"
)

type sendContext interface {
	JSON(code int, obj interface{})
	DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string)
}

func (c *Core) handleSender(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func (c *Core) handlePostSender(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func (c *Core) sendDirect(ctx sendContext, token *model.Token, msg *model.Message) {
	_ = "STUB: not implemented"
	return
}

func (c *Core) sendForward(ctx sendContext, token *model.Token, msg *model.Message) {
	_ = "STUB: not implemented"
	return
}

func (c *Core) sendMsg(ctx sendContext, token *model.Token, msg *model.Message) {
	_ = "STUB: not implemented"
	return
}

func (c *Core) saveUploadImage(ctx *gin.Context, token *model.Token, data []byte) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Core) saveUploadAudio(ctx *gin.Context, token *model.Token, fname string, title string, data []byte) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Core) saveUploadFile(ctx *gin.Context, token *model.Token, data []byte, filename string, desc string, actions []string) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Core) makeTextContent(msg *model.Message, text string, title string, copytext string, autocopy string, actions []string) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Core) makeActionContent(msg *model.Message, text string, title string, actions []string) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
