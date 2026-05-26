package model

import (
	"time"

	"github.com/chanify/chanify/pb"
)

// MsgTimeItem define data for timeline
type MsgTimeItem struct {
	Name  string
	Value interface{}
}

// Message for notification
type Message struct {
	pb.Message
	isTimeline bool
	ilValue    string
}

// NewMessage with sender token
func NewMessage(tk *Token) *Message { _ = "STUB: not implemented"; return nil }

// DisableToken clear token
func (m *Message) DisableToken() *Message { _ = "STUB: not implemented"; return nil }

// LinkContent set link notification
func (m *Message) LinkContent(link string) *Message { _ = "STUB: not implemented"; return nil }

// TimelineContent set timeline notification
func (m *Message) TimelineContent(code string, title string, ts *time.Time, items []*MsgTimeItem) *Message {
	_ = "STUB: not implemented"
	return nil
}

// TextContent set text notification
func (m *Message) TextContent(text string, title string, copytext string, autocopy string) *Message {
	_ = "STUB: not implemented"
	return nil
}

// ActionContent set custom action notification
func (m *Message) ActionContent(text string, title string, actions []string) *Message {
	_ = "STUB: not implemented"
	return nil
}

// FileContent set file notification
func (m *Message) FileContent(path string, filename string, desc string, size int, actions []string) *Message {
	_ = "STUB: not implemented"
	return nil
}

// ImageContent set image notification
func (m *Message) ImageContent(path string, t *Thumbnail, size int) *Message {
	_ = "STUB: not implemented"
	return nil
}

// AudioContent set audio notification
func (m *Message) AudioContent(path string, fname string, title string, duration uint64, size int) *Message {
	_ = "STUB: not implemented"
	return nil
}

// TextFileContent set text file notification
func (m *Message) TextFileContent(path string, filename string, title string, desc string, size int, actions []string) *Message {
	_ = "STUB: not implemented"
	return nil
}

// IsTimeline return is timeline notification
func (m *Message) IsTimeline() bool { _ = "STUB: not implemented"; return false }

// SetTimeline set timeline notification
func (m *Message) SetTimeline(timeline bool) *Message { _ = "STUB: not implemented"; return nil }

// SoundName set notification sound
func (m *Message) SoundName(sound string) *Message { _ = "STUB: not implemented"; return nil }

// SetPriority set notification priority
func (m *Message) SetPriority(priority int) *Message { _ = "STUB: not implemented"; return nil }

// SetInterruptionLevel set time sensitive notification
func (m *Message) SetInterruptionLevel(interruptionLevel string) *Message {
	_ = "STUB: not implemented"
	return nil
}

// EncryptContent return encrypted content with key
func (m *Message) EncryptContent(key []byte) { _ = "STUB: not implemented"; return }

// nolint: errcheck

// EncryptData return encrypted body with key & timestamp
func (m *Message) EncryptData(key []byte, ts uint64) []byte { _ = "STUB: not implemented"; return nil }

// Marshal return binary data
func (m *Message) Marshal() []byte { _ = "STUB: not implemented"; return nil }

func (m *Message) fixChannel() { _ = "STUB: not implemented"; return }

func parseActions(actions []string) []*pb.ActionItem { _ = "STUB: not implemented"; return nil }
