package bitbucket

import (
    "os"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func ValidatePayload(payload []byte, signature string) bool {

    secret := os.Getenv("GIT_SECRET")
	if len(payload) == 0 {
        return false
	}

	if len(secret) > 0 {
		if len(signature) == 0 {
            return false
		}
		mac := hmac.New(sha256.New, []byte(secret))
		_, _ = mac.Write(payload)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature[7:]), []byte(expectedMAC)) {
            return false
		}
	}

    return true

}
