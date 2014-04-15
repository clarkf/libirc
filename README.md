# libirc

[![Build Status](http://img.shields.io/travis/clarkf/libirc.svg)](https://travis-ci.org/clarkf/libirc)

libirc is an IRC library for [Go](http://golang.org/).

## Examples

A simple example:

```go
package main

import (
	"github.com/clarkf/libirc"
	"time"
)

func main() {
	c := libirc.NewClient("nickname", "username", "Real Name")

	err := c.ConnectAndListen("irc.myserver.tld:6667")

	// libirc will block until the connection is ready

	if err != nil {
		panic(err)
	}

	c.Join("#hello")

	c.Write(libirc.NewMessage("PRIVMSG", []string{"#hello", "Hello world!"}))

	time.Sleep(time.Duration(5) * time.Second)

	// Client will disconnect now.
}
```

## License

The MIT License (MIT)

Copyright (c) 2014 Clark Fischer

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
