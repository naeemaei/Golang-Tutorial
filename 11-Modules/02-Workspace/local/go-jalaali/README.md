# Jalaali

Golang implementation of [Jalaali JS](https://github.com/jalaali/jalaali-js) and [Jalaali Python](https://github.com/jalaali/jalaali-python) implementations of Jalaali (Jalali, Persian, Khayyami, Khorshidi, Shamsi) convertion to Gregorian calendar system and vice-versa.

This implementation is based on an [algorithm by Kazimierz M. Borkowski](http://www.astro.uni.torun.pl/~kb/Papers/EMP/PersianC-EMP.htm). Borkowski claims that this algorithm works correctly for 3000 years!

Documentation on API is available [here](https://pkg.go.dev/github.com/jalaali/go-jalaali) at Go official documentation site.

## Installation

Use `go get` on this repository:

```sh
$ go get -u github.com/jalaali/go-jalaali
```

## Usage

* Wrapper around Golang [time package](https://golang.org/pkg/time):
  * Call `Jalaali.Now()` to get instance of current time. You can use all function from `time` package with this wrapper.
  * Call `Jalaali.From(t)` and pass a `time` instance to it. The you can work with it the same way you work with `time` package.
* Jalaali Formatting:
  * Call `JFormat` method of a Jalaali instance and pass it the same formatting options that is used for Golang `time` package. The output will be in Jalaali date and use persian digits and words.
