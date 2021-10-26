package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var currentBranchAargs = [][]string{
	{"branch", "--show-current"},
	{"rev-parse", "--abbrev-ref", "HEAD"},
	{"symbolic-ref", "--short", "HEAD"},
	{"name-rev", "--name-only", "HEAD"},
}

var tagsArgs = [][]string{
	{"describe", "--tags", "--dirty"},
	{"rev-parse", "HEAD"},
}

var ldflags = flag.String(
	"ldflags",
	"",
	"go build ldflags",
)

func main() {
	flag.Parse()
	var gitBranch string
	for _, args := range currentBranchAargs {
		var currentBranch, err = exec.Command("git", args...).Output()
		if err == nil {
			gitBranch = strings.Trim(string(currentBranch), "\n")
			break
		}
	}
	var gitTag string
	for _, args := range tagsArgs {
		var tag, err = exec.Command("git", args...).Output()
		if err == nil {
			gitTag = strings.Trim(string(tag), "\n")
			break
		}
	}
	var gitDescribe = gitBranch + ":" + gitTag
	var date = time.Now().Format("2006-01-02_15:04:05PM")
	var goVersion = strings.Trim(string(HandleError(exec.Command("go", "version").Output()).([]byte)), "\n")
	var build = exec.Command("go", "build", "-x", "-ldflags", fmt.Sprintf("-X 'github.com/liserjrqlxue/version.gitDescribe=%s' -X 'github.com/liserjrqlxue/version.buildStamp=%s' -X 'github.com/liserjrqlxue/version.golangVersion=%s' %s", gitDescribe, date, goVersion, *ldflags))
	fmt.Println(build.String())
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
