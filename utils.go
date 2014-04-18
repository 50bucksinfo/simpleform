package main

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/golang/glog"
)

func withDefault(args ...string) string {

	for _, arg := range args {
		if arg != "" {
			return arg
		}
	}

	return ""
}

//random
func secureHex() string {
	hexBytes := make([]byte, 32)
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		glog.Errorln(err)
		return ""
	}
	hex.Encode(hexBytes, randomBytes)
	return string(hexBytes)
}
