i-did
============

Based on the [i-did project by VonKavalier](https://github.com/VonKavalier/i-did)

Little script to keep track of what you do everyday.

The idea is to keep track of things you've done, not to do. It's based on the ["Done List"](https://www.lifehack.org/articles/productivity/why-to-do-lists-dont-work-and-done-lists-do.html) principle.

It features a readable list with tasks grouped by days.

Works on Linux (Mac and Windows not tested)

Requirement
-----------

* [Go](https://golang.org/)

Installation
-----------

1. Download project with

```bash
$ go get https://github.com/bihetq/idid
```

2. Build the package with 
```bash
$ cd $GOPATH/src/github.com/bihet/i-did/
$ go install
```


Usage
-----------

```bash
# Show help
$ i-did -h
```

There are 3 flag you can use:

| Flag      | Detail                           | Require or not |
|:---------:|----------------------------------|----------------|
|```-path```|Path for file (default $HOME/idid)|Optional        |
|```-msg``` |Message to write                  |Require         |
|```-show```| Display file                     |Optional        |

```bash
# Write a message to default file
$ i-did -msg="<Your message>"

# Write a message to custom file
$ i-did -path=<path/to/your/file> -msg="<Your message>"
```

```bash
# Show default file
$ i-did -show

# Show custom file
$ i-did -show -path=<path/to/your/file>
```
