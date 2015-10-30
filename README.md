cliTricks
==========

Custom written CLI tools for interfacing with a paginated API endpoint.

It consists of four binaries:

* [`templcat`](templcat/readme.md) -build output using [the go template engine](golang.org/pkg/text/template)
* [`lasercat`](lasercat/readme.md) - pluci items from JSON input
* [`prettycat`](prettycat/readme.md) - pretty format JSON
* [`apicat`](apicat/readme.md) - run JSON input against an API endpoint, incrementing the input to get all pages

Installing
----------

* If you have not, [install go](https://golang.org/doc/install).

* Get the source

```
go get github.com/JackKnifed/cliTricks
```

* Compile each binary

```
go install github.com/JackKnifed/cliTricks/templcat
go install github.com/JackKnifed/cliTricks/lasercat
go install github.com/JackKnifed/cliTricks/prettycat
go install github.com/JackKnifed/cliTricks/apicat
```

* You should now be able to find all of the binaries at `$GOPATH/bin`

Testing
-------

There are no unit tests or similar written for this - at least for the most part. THe binaries are provided as is, without proper testing.

This project is a hack, but it works.