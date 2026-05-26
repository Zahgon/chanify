package core

import (
	"mime/multipart"
	"time"

	"github.com/chanify/chanify/model"
	"github.com/gin-gonic/gin"
)

// TimeContent define timeline content
type TimeContent struct {
	Code      string
	Timestamp *time.Time
	Items     []*model.MsgTimeItem
}

// MsgParam parse message parameters
type MsgParam struct {
	Token             *model.Token
	Text              string
	Link              string
	Title             string
	Sound             string
	AutoCopy          string
	CopyText          string
	Filename          string
	Priority          int
	InterruptionLevel string
	Actions           []string
	TimeContent       TimeContent
}

// ParsePlainText process text/plain
func (m *MsgParam) ParsePlainText(ctx *gin.Context) { _ = "STUB: not implemented"; return }

// ParseJSON process application/json
func (m *MsgParam) ParseJSON(c *Core, ctx *gin.Context) { _ = "STUB: not implemented"; return }

// ParseForm process form
func (m *MsgParam) ParseForm(c *Core, ctx *gin.Context) { _ = "STUB: not implemented"; return }

// ParseFormData process multipart/form-data
func (m *MsgParam) ParseFormData(c *Core, ctx *gin.Context) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseImage process image
func (m *MsgParam) ParseImage(c *Core, ctx *gin.Context) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseAudio process audio
func (m *MsgParam) ParseAudio(c *Core, ctx *gin.Context) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MsgParam) parsePriorityFromForm(form *multipart.Form) { _ = "STUB: not implemented"; return }

func parseTimeContentItems(items map[string]interface{}) []*model.MsgTimeItem {
	_ = "STUB: not implemented"
	return nil
}

func parseTimeContentStringItems(items map[string]string) []*model.MsgTimeItem {
	_ = "STUB: not implemented"
	return nil
}

func parseTimestamp(t interface{}) *time.Time { _ = "STUB: not implemented"; return nil }

func readFileFromForm(form *multipart.Form, name string) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func tryStringValue(value string, newValue string) string { _ = "STUB: not implemented"; return "" }

func tryFormValue(form *multipart.Form, name string, value string) string {
	_ = "STUB: not implemented"
	return ""
}

func tryFormValues(form *multipart.Form, name string, value []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func tryFormMap(form *multipart.Form, name string, items []*model.MsgTimeItem) []*model.MsgTimeItem {
	_ = "STUB: not implemented"
	return nil
}

func tryFormTimestamp(form *multipart.Form, name string, ts *time.Time) *time.Time {
	_ = "STUB: not implemented"
	return nil
}
