package core

import (
	"github.com/chanify/chanify/model"
	"github.com/gin-gonic/gin"
)

const (
	pngHeader  = "\x89PNG\r\n\x1a\n"
	gifHeader  = "GIF"
	riffHeader = "RIFF"
	webpHeader = "WEBP"
)

func (c *Core) bindBodyJSON(ctx *gin.Context, obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyUser(ctx *gin.Context, key string) bool { _ = "STUB: not implemented"; return false }

func verifyDevice(ctx *gin.Context, key string) bool { _ = "STUB: not implemented"; return false }

func verifySign(key string, sign []byte, data []byte) bool { _ = "STUB: not implemented"; return false }

func (c *Core) parseToken(token string) (*model.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getToken(ctx *gin.Context) string { _ = "STUB: not implemented"; return "" }

func parsePriority(priority string) int { _ = "STUB: not implemented"; return 0 }

func parseImageContentType(data []byte) string { _ = "STUB: not implemented"; return "" }

func createThumbnail(data []byte) *model.Thumbnail { _ = "STUB: not implemented"; return nil }

func fileBaseName(path string) string { _ = "STUB: not implemented"; return "" }

func fixLog(s string) string { _ = "STUB: not implemented"; return "" }

// JSONString define boolean string
type JSONString string

// UnmarshalJSON for boolean string
func (s *JSONString) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
