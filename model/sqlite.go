package model

import (
	"database/sql"
	"log"
	"strings"

	_ "modernc.org/sqlite" // sqlite driver
)

type sqlite struct {
	db *sql.DB
}

func init() {
	drivers["sqlite"] = func(dsn string) (DB, error) {
		items := strings.Split(dsn, "://")
		path := items[1]
		db, _ := sql.Open(items[0], "file:"+path)
		if err := db.Ping(); err != nil {
			return nil, err
		}
		log.Println("Open sqlite database:", path)
		s := &sqlite{db: db}
		if err := s.fixDB(); err != nil {
			return nil, err
		}
		return s, nil
	}
}

func (s *sqlite) Close() { _ = "STUB: not implemented"; return }

func (s *sqlite) SetOption(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlite) GetOption(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlite) GetUser(uid string) (*User, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *sqlite) UpsertUser(u *User) error { _ = "STUB: not implemented"; return nil }

func (s *sqlite) BindDevice(uid string, uuid string, key []byte, devType int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlite) UnbindDevice(uid string, uuid string) error { _ = "STUB: not implemented"; return nil }

func (s *sqlite) UpdatePushToken(uid string, uuid string, token []byte, sandbox bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sqlite) GetDeviceKey(uuid string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sqlite) GetDevices(uid string) ([]*Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nolint: errcheck

func (s *sqlite) fixDB() error { _ = "STUB: not implemented"; return nil }

// nolint: errcheck

// nolint: errcheck
