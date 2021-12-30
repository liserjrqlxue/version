package version

import (
	"fmt"
	"log"
	"os"
)

var (
	buildStamp    string
	gitDescribe   string
	golangVersion string
)

// LogVersion log out version info
func LogVersion() {
	log.Printf("Git Describe   : %s\n", gitDescribe)
	log.Printf("Build Time     : %s\n", buildStamp)
	log.Printf("Golang Version : %s\n", golangVersion)
	var hostName, err = os.Hostname()
	log.Printf("Current Host   : %s%v\n", hostName, err)
}

// Version print out version info
func Version() {
	fmt.Printf("Git Describe   : %s\n", gitDescribe)
	fmt.Printf("Build Time     : %s\n", buildStamp)
	fmt.Printf("Golang Version : %s\n", golangVersion)
	var hostName, err = os.Hostname()
	fmt.Printf("Current Host   : %s%v\n", hostName, err)
}
