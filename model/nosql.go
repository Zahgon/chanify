package model

import (
	"crypto/sha256"
	"net/url"

	"github.com/chanify/chanify/crypto"
)

type nosql struct {
	secret []byte
	seckey []byte
}

func init() {
	drivers["nosql"] = func(dsn string) (DB, error) {
		u, _ := url.Parse(dsn)
		secret := []byte(u.Query().Get("secret"))
		if len(secret) <= 0 {
			return nil, ErrInvalidDSN
		}
		return &nosql{
			secret: sha256.New().Sum(secret),
			seckey: crypto.GenerateSecretKey(secret).MarshalSecretKey(),
		}, nil
	}
}

func (s *nosql) Close() { _ = "STUB: not implemented"; return }

func (s *nosql) GetOption(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *nosql) SetOption(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *nosql) GetUser(uid string) (*User, error) { _ = "STUB: not implemented"; return nil, nil }

// nolint: errcheck
// nolint: errcheck

func (s *nosql) UpsertUser(u *User) error { _ = "STUB: not implemented"; return nil }

func (s *nosql) BindDevice(uid string, uuid string, key []byte, devType int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *nosql) UnbindDevice(uid string, uuid string) error { _ = "STUB: not implemented"; return nil }

func (s *nosql) UpdatePushToken(uid string, uuid string, token []byte, sandbox bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *nosql) GetDeviceKey(uuid string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *nosql) GetDevices(uid string) ([]*Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
