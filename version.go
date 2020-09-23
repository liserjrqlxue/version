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
	log.Printf("Git Describe\t: %s\n", gitDescribe)
	log.Printf("Build Time\t: %s\n", buildStamp)
	log.Printf("Golang Version\t: %s\n", golangVersion)
	var hostName, err = os.Hostname()
	log.Printf("Hostname\t: %s%v\n", hostName, err)
}
