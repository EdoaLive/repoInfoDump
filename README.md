# repoInfoDump

This tools helps you quickly obtain a string containing the repo status (commit hash, time) at the time of build of your binary.

## Usage
```go
fmt.println(repoInfoDump.GetString())
```
Output example:
```
repoInfoDump v0.1.3 commit: cf5013a (2024-03-28T11:53:42Z) (REPO MODIFIED)
```

Other customized usages are possible.

## TODO
* Improve documentation

