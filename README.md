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

```bash
$ go get github.com/BtKent/idid
```

Update
-----------

```bash
$ go get -u github.com/BtKent/idid
```

Usage
-----------

```bash
# Show help
$ idid -h
```

There are 4 flags you can use:

| Flag      | Detail                           | Require or not |
|:---------:|----------------------------------|----------------|
|```-path```|Path for file (default $HOME/idid)|Optional        |
|```-msg``` |Message to write                  |Require         |
|```-show```| Display file                     |Optional        |
|```-version```| Display version number                     |Optional        |

```bash
# Write a message to default file
$ idid -msg="<Your message>"

# Write a message to custom file
$ idid -path=<path/to/your/file> -msg="<Your message>"
```

```bash
# Show default file
$ idid -show

# Show custom file
$ idid -show -path=<path/to/your/file>
```
