package components

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Hash struct{}

func NewMd5Hash() *md5Hash {
	return &md5Hash{}
}

func (h *md5Hash) Hash(data string) string {
	hashder := md5.New()
	hashder.Write([]byte(data))
	return hex.EncodeToString(hashder.Sum(nil))
}