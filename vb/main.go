package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var gitBranch = strings.Trim(string(HandleError(exec.Command("git", "branch", "--show-current").Output()).([]byte)), "\n")
	var gitDescribe = gitBranch + ":" + strings.Trim(string(HandleError(exec.Command("git", "describe", "--tags").Output()).([]byte)), "\n")
	var date = time.Now().Format("2006-01-02_15:04:05PM")
	var goVersion = strings.Trim(string(HandleError(exec.Command("go", "version").Output()).([]byte)), "\n")
	fmt.Println(strings.Join([]string{"go", "build", "-x", "-ldflags", fmt.Sprintf("-X 'github.com/liserjrqlxue/version.gitDescribe=%s' -X 'github.com/liserjrqlxue/version.buildStamp=%s' -X 'github.com/liserjrqlxue/version.golangVersion=%s'", gitDescribe, date, goVersion)}, " "))
	var build = exec.Command("go", "build", "-x", "-ldflags", fmt.Sprintf("-X 'github.com/liserjrqlxue/version.gitDescribe=%s' -X 'github.com/liserjrqlxue/version.buildStamp=%s' -X 'github.com/liserjrqlxue/version.golangVersion=%s'", gitDescribe, date, goVersion))
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	CheckErr(build.Run())
}

// CheckErr log.Fatal if error
func CheckErr(err error, msg ...string) {
	if err != nil {
		//panic(err)
		log.Fatal(err, msg)
	}
}

// HandleError CheckErr and return as interface
func HandleError(a interface{}, err error) interface{} {
	CheckErr(err)
	return a
}
