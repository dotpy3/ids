# Benchmark between `regexp` and `strings` for splitting strings

This repo holds a small benchmark for splitting strings, using the Go stdlib's `strings` or `regexp`.

These are the results on my machine:

```
/Users/egourlao/go/src/github.com/dotpy3/ids$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/dotpy3/ids
BenchmarkReadIDUsingStrings-4   	100000000	        13.1 ns/op
BenchmarkReadIDUsingRegexp-4    	 2000000	       893 ns/op
PASS
ok  	github.com/dotpy3/ids	3.873s
```

Conclusion: avoid using `regexp` for simple string operations that can be handled by `strings` ¯\_(ツ)_/¯

## Reproduce

Make sure you have a Go environment set up.

```
go get github.com/dotpy3/ids
cd $GOPATH/src/github.com/dotpy3/ids
go test -bench=.
```
