package model

import (
	"github.com/chanify/chanify/pb"
)

// Token for sender
type Token struct {
	data     pb.Token
	signSys  []byte
	signNode []byte
	rawData  []byte
	raw      string
}

// ParseToken create token from base64 string
func ParseToken(token string) (*Token, error) { _ = "STUB: not implemented"; return nil, nil }

// GetUserID return user id string
func (tk *Token) GetUserID() string { _ = "STUB: not implemented"; return "" }

// GetNodeID return node id
func (tk *Token) GetNodeID() []byte { _ = "STUB: not implemented"; return nil }

// GetChannel return channel code
func (tk *Token) GetChannel() []byte { _ = "STUB: not implemented"; return nil }

// IsExpires check token expires timestamp(UTC)
func (tk *Token) IsExpires() bool { _ = "STUB: not implemented"; return false }

// VerifySign check token sign
func (tk *Token) VerifySign(key []byte) bool { _ = "STUB: not implemented"; return false }

// nolint: errcheck

// VerifyDataHash check the hash of uri limit
func (tk *Token) VerifyDataHash(data []byte) bool { _ = "STUB: not implemented"; return false }

// RawToken return raw value
func (tk *Token) RawToken() string {
	_ = "STUB: not implemented"

	// HashValue return sha1 with token raw value
	return ""
}

func (tk *Token) HashValue() []byte { _ = "STUB: not implemented"; return nil }
