package main

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/golang/glog"
	"runtime"
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

//logs stack and error if there is an error
func log(err error, args ...interface{}) {
	if err == nil {
		return
	}

	buf := make([]byte, 4096)
	buf = buf[:runtime.Stack(buf, false)]

	glog.Errorln(args...)
	glog.Errorf("ERROR: %q", err)
	glog.Errorf("STACK: %s", buf)
}
