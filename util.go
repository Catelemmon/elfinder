package elfinder

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"
)

const VolumeIDPrefix = "v"

func Decode64(s string) (string, error) {
	t, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(t), nil
}

func Encode64(s string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(s))
}

func CreateHash(volumeId, path string) string {
	return volumeId + "_" + Encode64(path)
}

func GenerateVolumeID() string {
	ctx := md5.New()
	ctx.Write([]byte(time.Now().String()))
	return fmt.Sprintf("%s%x", VolumeIDPrefix, ctx.Sum(nil))
}

func ParseStringToBool(s string) bool {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true", "on":
		return true
	}
	return false
}

func ReadWritePem(pem os.FileMode) (readable, writable byte) {
	if pem&(1<<uint(9-1-0)) != 0 {
		readable = 1
	}
	if pem&(1<<uint(9-1-1)) != 0 {
		writable = 1
	}
	return
}
