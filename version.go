package version

import (
	"log"
	"os"
)

var (
	buildStamp    string
	gitDescribe   string
	golangVersion string
)

func LogVersion() {
	log.Printf("Git Describe   : %s\n", gitDescribe)
	log.Printf("Build Time     : %s\n", buildStamp)
	log.Printf("Golang Version : %s\n", golangVersion)
	var hostName, err = os.Hostname()
	log.Printf("Hostname       : %s%v\n", hostName, err)
}
