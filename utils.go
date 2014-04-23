package main

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/golang/glog"
	"os"
	"runtime"
)

//returns the first non empty string
//TODO; should probably rename it to something more obvious
func withDefault(args ...string) string {
	for _, arg := range args {
		if arg != "" {
			return arg
		}
	}
	return ""
}

//create a 32 character random hex string
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

//prints an informational message in the log
func logInfo(args ...interface{}) {
	glog.Infoln(args...)
}

//logs stack and error if there is an error
//also stops the app by calling os.Exit(-1)
func logFatal(err error, args ...interface{}) {
	logError(err, args...)
	if err != nil {
		os.Exit(-1)
	}
}

//logs stack and error if there is an error
func logError(err error, args ...interface{}) {
	if err == nil {
		return
	}

	buf := make([]byte, 4096)
	buf = buf[:runtime.Stack(buf, false)]

	glog.Errorln(args...)
	glog.Errorf("ERROR: %q", err)
	glog.Errorf("STACK: %s", buf)
}
