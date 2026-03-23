package appimagetoolgo

import (
	"encoding/hex"

	"github.com/ProtonMail/gopenpgp/v2/helper"
)

func GenerateSigningKey(name string, email string, passphrase string) (string, error) {
	return helper.GenerateKey(name, email, []byte(passphrase), "rsa", 4096)
}

func SignSha256(hash []byte, privateKey string, passphrase string) (string, error) {
	hexlifiedHash := hex.EncodeToString(hash)
	return helper.SignCleartextMessageArmored(privateKey, []byte(passphrase), hexlifiedHash)
}
