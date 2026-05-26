package model

import (
	"crypto/cipher"

	"github.com/chanify/chanify/crypto"
)

// NewAESGCM for aes-gcm chiper
func NewAESGCM(key []byte) (cipher.AEAD, error) {
	_ = "STUB: not implemented"
	return *new(cipher.AEAD), nil
}

// DecodePushToken for APNS
func DecodePushToken(token string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// CalcDeviceKey from device public key
func CalcDeviceKey(uuid string, key string) (*crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CalcUserKey from user public key
func CalcUserKey(uid string, key string) (*crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calc user id
