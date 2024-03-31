# go race

[go race](https://go.dev/blog/race-detector)

开启 race detect

```shell
$ go test -race mypkg    // test the package
$ go run -race mysrc.go  // compile and run the program
$ go build -race mycmd   // build the command
$ go install -race mypkg // install the package
```
