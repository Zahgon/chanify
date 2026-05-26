package logic

import (
	"crypto/rand"
	"errors"

	"github.com/chanify/chanify/crypto"
	"github.com/chanify/chanify/model"
	"github.com/sideshow/apns2"
)

// variable define
var (
	APIEndpoint            = "https://api.chanify.net"
	MockPusher  APNSPusher = nil

	randReader = rand.Read

	ErrNoSupportMethod = errors.New("no support method")
	ErrNotFound        = errors.New("not found")
	ErrInvalidContent  = errors.New("invalid content")
	ErrSystemLimited   = errors.New("system limited")
)

// Options for init logic
type Options struct {
	Name         string
	Version      string
	Endpoint     string
	DataPath     string
	FilePath     string
	PluginPath   string
	DBUrl        string
	Secret       string
	Registerable bool
	RegUsers     []string
	WebHooks     []map[string]interface{}
}

// Logic instance
type Logic struct {
	srvless      bool
	registerable bool
	db           model.DB
	secKey       *crypto.SecretKey
	Name         string
	NodeID       string
	Version      string
	Endpoint     string
	Features     []string

	infoData      []byte
	infoSign      string
	whitelist     map[string]bool
	filepath      string
	webhookManger *pluginManager

	apnsPClient *apns2.Client
	apnsDClient *apns2.Client
}

// APNSPusher is the interface of APNS2
type APNSPusher interface {
	Push(n *apns2.Notification) (*apns2.Response, error)
}

const authKey = "MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgQ6vCLkUeDj223nfPfKGrjG+Coc53EbKHmO6Oa9YcHiGgCgYIKoZIzj0DAQehRANCAAQNwg3W2eOqNlX0nl9kGbfmMxwSZoO4RmqKoKJnH/vGkU8csJuN5Dg4JiI6ni5PEx+A1rb19DuDm4AzwBVvl8Jt"

func (opts *Options) fixOptions() { _ = "STUB: not implemented"; return }

// NewLogic with options
func NewLogic(opts *Options) (*Logic, error) { _ = "STUB: not implemented"; return nil, nil }

// nolint: errcheck
// nolint: errcheck
// nolint: errcheck

// Close and cleanup logic instance
func (l *Logic) Close() { _ = "STUB: not implemented"; return }

// CanFileStore return file stroage is available
func (l *Logic) CanFileStore() bool { _ = "STUB: not implemented"; return false }

// GetUser find user info with user id
func (l *Logic) GetUser(uid string) (*model.User, error) {
	_ = "STUB: not implemented"
	return nil,

		// GetUserKey find user key with user id
		nil
}

func (l *Logic) GetUserKey(uid string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UpsertUser insert or update user info
func (l *Logic) UpsertUser(uid string, key string, serverless bool) (*model.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BindDevice to user
func (l *Logic) BindDevice(uid string, uuid string, key string, devType int) error {
	_ = "STUB: not implemented"
	return nil
}

// UnbindDevice from user
func (l *Logic) UnbindDevice(uid string, uuid string) error { _ = "STUB: not implemented"; return nil }

// UpdatePushToken for APNS
func (l *Logic) UpdatePushToken(uid string, uuid string, token string, sandbox bool) error {
	_ = "STUB: not implemented"
	return nil
}

// GetDeviceKey return device key with device uuid
func (l *Logic) GetDeviceKey(uuid string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// GetDevices return all devices with user id
}

func (l *Logic) GetDevices(uid string) ([]*model.Device, error) {
	_ = "STUB: not implemented"
	return nil,

		// Decrypt data with node secret key
		nil
}

func (l *Logic) Decrypt(data []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// VerifyToken chekc sender token
func (l *Logic) VerifyToken(tk *model.Token) bool { _ = "STUB: not implemented"; return false }

// LoadFile read with file type & data
func (l *Logic) LoadFile(tname string, name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SaveFile save with file type & data
func (l *Logic) SaveFile(tname string, data []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetWebhook with name
func (l *Logic) GetWebhook(name string) (*Webhook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SendAPNS send message to APNS
func (l *Logic) SendAPNS(uid string, data []byte, devices []*model.Device, priority int, interruptionLevel string, isTimeline bool) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}

// only 10 or 5

// 1: iOS, 2: watchOS, 3: macOS

func (l *Logic) getAPNS(sandbox bool) APNSPusher {
	_ = "STUB: not implemented"
	return *new(APNSPusher)
}

func (l *Logic) loadDB(dburl string) error { _ = "STUB: not implemented"; return nil }

func (l *Logic) fixSecretKey() error { _ = "STUB: not implemented"; return nil }

func (l *Logic) createUser(uid string, pk *crypto.PublicKey, serverless bool) (*model.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *Logic) canRegisterUser(uid string) bool { _ = "STUB: not implemented"; return false }
