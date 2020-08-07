# version

```
gitDescribe=$(git describe --tags)
golangVersion=$(go version)
buildStamp=$(date -u '+%Y-%m-%d_%I:%M:%S%p')
go build -ldflags "-X 'github.com/liserjrqlxue/version.gitDescribe=$gitDescribe' -X 'github.com/liserjrqlxue/version.buildStamp=$buildStamp -X 'github.com/liserjrqlxue/version.golangVersion=$golangVersion'"
```
