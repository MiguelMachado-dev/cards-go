# Cards Go
[![Go Report Card](https://goreportcard.com/badge/github.com/MiguelMachado-dev/cards-go)](https://goreportcard.com/report/github.com/MiguelMachado-dev/cards-go)
[![Go](https://github.com/MiguelMachado-dev/cards-go/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/MiguelMachado-dev/cards-go/actions/workflows/go.yml)
![GitHub](https://img.shields.io/github/license/MiguelMachado-dev/cards-go)

This is a simple Go program that simulates a deck of cards. It includes functions for creating a new deck, shuffling the deck, and dealing cards from the deck.

## Installation
To use this program, you must have Go installed on your computer. You can download and install Go from the official website: https://golang.org/dl/

Once you have Go installed, you can download and install this program using the following command:

```shell
go get github.com/MiguelMachado-dev/cards-go
```

## Usage
To use this program, you can import the `cards` package into your own Go program and use the functions provided by the package. Here's an example:

```go
package main

import (
    "fmt"
    "github.com/MiguelMachado-dev/cards-go/cards"
)

func main() {
    deck := cards.NewDeck()
    deck.Shuffle()
    hand, remainingDeck := deck.Deal(5)
    fmt.Println("Hand:", hand)
    fmt.Println("Remaining deck:", remainingDeck)
}
```

This program creates a new deck of cards, shuffles the deck, deals a hand of 5 cards, and prints the hand and remaining deck to the console.

## Contributing
If you find a bug or have a suggestion for this program, please open an issue or submit a pull request on GitHub.

## License
This program is licensed under the MIT License. See the LICENSE file for details.
