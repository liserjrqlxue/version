# version

[![GoDoc](https://godoc.org/github.com/liserjrqlxue/version?status.svg)](https://pkg.go.dev/github.com/liserjrqlxue/version) [![Go Report Card](https://goreportcard.com/badge/github.com/liserjrqlxue/version)](https://goreportcard.com/report/github.com/liserjrqlxue/version)

After assign version info when `go build`, log out version info when run

```
gitDescribe=$(git branch --show-current):$(git describe --tags --dirty)
golangVersion=$(go version)
buildStamp=$(date '+%Y-%m-%d_%I:%M:%S%p')
go build -ldflags "-X 'github.com/liserjrqlxue/version.gitDescribe=$gitDescribe' -X 'github.com/liserjrqlxue/version.buildStamp=$buildStamp -X 'github.com/liserjrqlxue/version.golangVersion=$golangVersion'"
```

## vb
`vb` do the `go build` commands above