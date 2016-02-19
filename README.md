# go-commons-lang

This is a port of the popular Java Apache Commons where it makes sense in Golang.

| Package | Description |
| ------------- | ------------- |
| `stringUtils` | String Utilities reflecting what's available in [StringUtils](http://commons.apache.org/proper/commons-lang/apidocs/org/apache/commons/lang3/StringUtils.html) |
| `randUtils` | [RandomUtils](http://commons.apache.org/proper/commons-lang/apidocs/index.html?org/apache/commons/lang3/StringUtils.html)  |
| `mathUtils` | `Fraction` implementation of Apache Commons  |

## Usage

```go
package main

import (
	"fmt"
	strUtil "github.com/agrison/go-commons-lang/stringUtils"
)

func main() {
	lib := strUtil.Join([]string{"lang", "commons", "go"}, "-")
	fmt.Print(strUtil.Capitalize(strUtil.Reverse("olleh")))
	fmt.Print(strUtil.SwapCase(strUtil.Chomp(" fROM \n")))
	fmt.Println(strUtil.ReverseDelimited(lib, "-") + strUtil.Right("Go!! ", 2))
	// Output: Hello From go-commons-lang!
}
```

## Installation

```sh
go get github.com/agrison/go-commons-lang
```

## Documentation

Can be seen at [https://godoc.org/github.com/agrison/go-commons-lang](https://godoc.org/github.com/agrison/go-commons-lang)
