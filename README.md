errfunc
=======

[![Go Reference](https://pkg.go.dev/badge/deedles.dev/errfunc.svg)](https://pkg.go.dev/deedles.dev/errfunc)

The errfunc package provides a generic implementation of the pattern described in the [Errors are values](https://go.dev/blog/errors-are-values) Go blog post where a function becomes a no-op once an error has occured, allowing the user to simply call it repeatedly and then check for an error at the bottom.

For example, the original code from the blog post,

```go
ew := &errWriter{w: fd}
ew.write(p0[a:b])
ew.write(p1[c:d])
ew.write(p2[e:f])
// and so on
if ew.err != nil {
  return ew.err
}
```

would become the following with this package:

```go
ef := errfunc.New(fd.Write)
ef.Call(p0[a:b])
ef.Call(p1[c:d])
ef.Call(p2[e:f])
// and so on
if ef.Err() != nil {
  return ef.Err()
 }
 ```
 
 Unlike the original implementation, this version can cleanly wrap any function that takes a single argument and returns a single value and an error. This means that it'll work well for `io.Writer`s, `io.Reader`s, and many other types.
