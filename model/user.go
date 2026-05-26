package model

// User information
type User struct {
	UID       string
	PublicKey []byte
	SecretKey []byte
	Flags     uint
}

// IsServerless for user configuration
func (u *User) IsServerless() bool { _ = "STUB: not implemented"; return false }

// SetServerless for user configuration
func (u *User) SetServerless(s bool) { _ = "STUB: not implemented"; return }

// GetPublicKeyString return the user public key
func (u *User) GetPublicKeyString() string { _ = "STUB: not implemented"; return "" }

// PublicKeyEncrypt return encrypted public key
func (u *User) PublicKeyEncrypt(data []byte) []byte { _ = "STUB: not implemented"; return nil }

// nolint: errcheck
