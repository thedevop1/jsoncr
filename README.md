# JSONCR
[![PkgGoDev](https://pkg.go.dev/badge/thedevop1/jsoncr)](https://pkg.go.dev/thedevop1/jsoncr)   

JSONCR is a Go package that provides method to remove comments and indentations in JSON. Both block and line comments are supported. Common use is to process the JSON config file.

Getting Started
===============

Installing
----------

```sh
$ go get github.com/thedevop1/jsoncr
```

Example
-------
```go
    const j = `/* Comment */
{
    "example": "value" // more comment
}`
	condense, err := jsoncr.Remove(strings.NewReader(j))
	if err != nil {
        // error handling
    }
    var z interface{}
	err = json.Unmarshal(condense, &z)
	if err != nil {
        // error handling
	}
```

