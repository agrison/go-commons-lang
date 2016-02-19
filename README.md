# go-commons-lang

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[license]: https://github.com/agrison/go-commons-lang/blob/master/LICENSE
[godocs]: https://godoc.org/github.com/agrison/go-commons-lang

This is a port of the popular Java Apache Commons where it makes sense in Golang.

| Package | Description |
| ------------- | ------------- |
| `stringUtils` | String Utilities reflecting what's available in [StringUtils](http://commons.apache.org/proper/commons-lang/apidocs/org/apache/commons/lang3/StringUtils.html) |
| `randUtils` | [RandomUtils](http://commons.apache.org/proper/commons-lang/apidocs/index.html?org/apache/commons/lang3/StringUtils.html)  |
| `mathUtils` | `Fraction` implementation of Apache Commons  |

## Usage: `stringUtils`

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

## Join

PRs are more than welcome :)

## Author

[Alexandre Grison](https://github.com/agrison) and obviously all the contributors of the original Apache Commons Lang.
