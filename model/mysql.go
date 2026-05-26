package model

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

type mysql struct {
	db *sql.DB
}

func init() {
	drivers["mysql"] = func(dsn string) (DB, error) {
		items := strings.Split(dsn, "://")
		db, _ := sql.Open(items[0], items[1])
		if db == nil {
			return nil, ErrInvalidDSN
		}
		log.Println("Open mysql database:", dsn)
		s := &mysql{db: db}
		if err := s.fixDB(); err != nil {
			return nil, err
		}
		return s, nil
	}
}

func (s *mysql) Close() { _ = "STUB: not implemented"; return }

func (s *mysql) GetOption(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *mysql) SetOption(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *mysql) GetUser(uid string) (*User, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *mysql) UpsertUser(u *User) error { _ = "STUB: not implemented"; return nil }

func (s *mysql) BindDevice(uid string, uuid string, key []byte, devType int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *mysql) UnbindDevice(uid string, uuid string) error { _ = "STUB: not implemented"; return nil }

func (s *mysql) UpdatePushToken(uid string, uuid string, token []byte, sandbox bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *mysql) GetDeviceKey(uuid string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *mysql) GetDevices(uid string) ([]*Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nolint: errcheck

func (s *mysql) fixDB() error { _ = "STUB: not implemented"; return nil }

// nolint: errcheck

// nolint: errcheck
