package crypto

import (
	"crypto/aes"
	"crypto/ecdsa"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"hash"
	"io"
)

// variable define
var (
	eciesKeyLen       = aes.BlockSize
	Base64Encode      = base64.RawURLEncoding
	Base32Encode      = base32.StdEncoding.WithPadding(base32.NoPadding)
	ErrInvalidKey     = errors.New("InvalidKey")
	ErrInvalidMessage = errors.New("InvalidMessage")
)

// PublicKey of ECDSA
type PublicKey struct {
	ecdsa.PublicKey
}

// SecretKey of ECDSA
type SecretKey struct {
	ecdsa.PrivateKey
}

// LoadPublicKey from binary data
func LoadPublicKey(key []byte) (*PublicKey, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadSecretKey from binary data
func LoadSecretKey(key []byte) (*SecretKey, error) { _ = "STUB: not implemented"; return nil, nil }

// GenerateSecretKey with secret key
func GenerateSecretKey(secret []byte) *SecretKey { _ = "STUB: not implemented"; return nil }

// MarshalPublicKey return binary public key
func (k *PublicKey) MarshalPublicKey() []byte { _ = "STUB: not implemented"; return nil }

// ToID calc public key id with code
func (k *PublicKey) ToID(code byte) string { _ = "STUB: not implemented"; return "" }

// Verify message with sign
func (k *PublicKey) Verify(msg []byte, sig []byte) bool { _ = "STUB: not implemented"; return false }

// Encrypt data with public key
func (k *PublicKey) Encrypt(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPublicKey from secret key
func (k *SecretKey) GetPublicKey() *PublicKey { _ = "STUB: not implemented"; return nil }

// MarshalSecretKey return binary secret key
func (k *SecretKey) MarshalSecretKey() []byte { _ = "STUB: not implemented"; return nil }

// MarshalPublicKey return binary public key
func (k *SecretKey) MarshalPublicKey() []byte { _ = "STUB: not implemented"; return nil }

// ToID calc secret key id with code
func (k *SecretKey) ToID(code byte) string { _ = "STUB: not implemented"; return "" }

// EncodePublicKey return base64 public key from secret key
func (k *SecretKey) EncodePublicKey() string { _ = "STUB: not implemented"; return "" }

// Sign message with secret key
func (k *SecretKey) Sign(msg []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Decrypt data with sercet key
func (k *SecretKey) Decrypt(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalPublicKey(key *ecdsa.PublicKey) []byte { _ = "STUB: not implemented"; return nil }

func formatToID(code byte, key []byte) string { _ = "STUB: not implemented"; return "" }

func x963KDF(sharedKeySeed []byte, ephemeralPublicKey []byte, hfnc func() hash.Hash) []byte {
	_ = "STUB: not implemented"
	return nil
}

// nolint: errcheck

// nolint: errcheck

// nolint: errcheck

func calcSharedKey(sec *ecdsa.PrivateKey, pub *ecdsa.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type noMaybeReader struct {
	reader io.Reader
	maybe  bool
}

func NewNoMaybeReader(reader io.Reader) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (r *noMaybeReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
