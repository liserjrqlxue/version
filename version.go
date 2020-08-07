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
	log.Printf("Git Commit Hash  : %s\n", gitDescribe)
	log.Printf("UTC Build Time   : %s\n", buildStamp)
	log.Printf("Golang Version   : %s\n", golangVersion)
	var hostName, err = os.Hostname()
	log.Printf("Runtime hostname : %s%v\n", hostName, err)
}
