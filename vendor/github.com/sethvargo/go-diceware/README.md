## Golang Diceware Generator

[![Build Status](https://travis-ci.org/sethvargo/go-diceware.svg?branch=master)](https://travis-ci.org/sethvargo/go-diceware)
[![GoDoc](https://godoc.org/github.com/sethvargo/go-diceware?status.svg)](https://godoc.org/github.com/sethvargo/go-diceware)

This library implements the [Diceware](https://en.wikipedia.org/wiki/Diceware)
algorithm in pure Golang. The algorithm is most-commonly used when generating
human-readable passwords. You may be familiar with the [XKCD
comic](https://xkcd.com/936/).

The list of words are generated from the [EFF's "long"
list](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases).
However, the API's are abstracted so you can roll die and then use your own word
list as-needed.

It uses crypto/rand for rolling die for added randomness.

Sample example words this library may choose:

```text
squirt catchy anatomy storm
patchy replica scholar alkalize
operative shrank lying uncorrupt
confusion studio abstain subdivide chewy ouch password tropical pentagon
```

## Installation

```sh
$ go get -u github.com/sethvargo/go-diceware
```

## Usage

```golang
package main

import (
  "strings"

  "github.com/sethvargo/go-diceware/diceware"
)

func main() {
  // Generate 6 words using the diceware algorithm.
  list, err := diceware.Generate(6)
  if err != nil  {
    log.Fatal(err)
  }
  log.Printf(strings.Join(list, "-"))
}
```

See the [GoDoc](https://godoc.org/github.com/sethvargo/go-diceware) for more
information.

## License

This code is licensed under the MIT license.
