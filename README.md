# Random Number Generator

A simple command-line tool written in Go that generates random numbers within a specified range.

## Installation

Build from source: `go build -o rng main.go`

Install to OS: `go install`

## Usage

Run the program with optional arguments specifying the minimum and maximum numbers:

`rng 1 4`

Output: `1` or `2` or `3` or `4`

If no arguments are provided, it will generate a random number between 1 and 2:
`rng`

Output:
`1` or `2`