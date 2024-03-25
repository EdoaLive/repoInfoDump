# repoInfoDump

This tools helps you quickly obtain a string containing the repo status (commit hash, time) at the time of build of your binary.

## Usage
```go
fmt.println(repoInfoDump.GetString())
```
Output example:
```
Starting repoInfoDumpTest commit: 12a5db0 (2023-02-09T17:16:14Z) (REPO MODIFIED)
```

Other customized usages are possible.

## TODO
* Better documentation
* Add link-time git tag version info

