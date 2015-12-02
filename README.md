# onerr
[![Build Status](https://travis-ci.org/go-onerr/onerr.svg?branch=v1)](https://travis-ci.org/go-onerr/onerr)
[![Code Coverage](http://gocover.io/_badge/gopkg.in/onerr.v1)](http://gocover.io/gopkg.in/onerr.v1)
[![Documentation](https://godoc.org/gopkg.in/onerr.v1?status.svg)](https://godoc.org/gopkg.in/onerr.v1)

onerr is a very small package that provides functions to quickly handle errors.


## Example

```go
f, err := os.Create("output.txt")
onerr.Panic(err)             // Panics if the file cannot be created.
defer onerr.LogFunc(f.Close) // Logs any error while closing the file.

_, err = f.WriteString("output")
// Logs any write error like: error while writing "output.txt": disk is full
onerr.Logf(err, "error while writing %q", f.Name())
```


## Documentation

https://godoc.org/gopkg.in/onerr.v1


## Download

    go get gopkg.in/onerr.v1


## License

[MIT](LICENSE)
