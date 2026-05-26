package logic

// Info for node server
type Info struct {
	NodeID    string   `json:"nodeid"`
	Name      string   `json:"name,omitempty"`
	Version   string   `json:"version"`
	PublicKey string   `json:"pubkey"`
	Endpoint  string   `json:"endpoint,omitempty"`
	Features  []string `json:"features,omitempty"`
}

// InitInfo calc all info data for node
func (l *Logic) InitInfo() { _ = "STUB: not implemented"; return }

// GetInfo return signed info data
func (l *Logic) GetInfo() ([]byte, string) { _ = "STUB: not implemented"; return nil, "" }

// GetQRCode return QRCode png data
func (l *Logic) GetQRCode() []byte { _ = "STUB: not implemented"; return nil }
