# Golang-based Answers to Cryptopals Challenges

c.f. [The Cryptopals Crypto Challenges](https://cryptopals.com/)

## Instal from GitHub repository and build

Set up Go properly (c.f. [How to start a Go project in 2018](https://boyter.org/posts/how-to-start-go-project-2018/) or the more official [Getting Started](https://golang.org/doc/install))

```
$ cd $GOPATH
$ go get github.com/ptdecker/gocryptopals
$ go install
```
## Solutions

### Set 1 Challenge 1 - Convert hex to base64

```
$ cd $GOPATH
$ bin\gocryptopals < data\sc01c01.txt
```