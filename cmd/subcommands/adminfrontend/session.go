package adminfrontend

import "github.com/gorilla/sessions"

func reverseBytes(bytes []byte) *[]byte {
	result := make([]byte, len(bytes))
	for i := len(bytes) - 1; i >= 0; i-- {
		result[(len(bytes)-1)-i] = bytes[i]
	}
	return &result
}

var store *sessions.FilesystemStore

func (p *adminFrontendCmd) initSession() {
	var (
		hashKey   = []byte(p.hashKey)
		cryptoKey = []byte(p.cryptoKey)
	)
	store = sessions.NewFilesystemStore("/go/sessions/", hashKey, cryptoKey)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   (86400 * 365),
		HttpOnly: true,
	}
}
