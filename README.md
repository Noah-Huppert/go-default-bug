# Go Default Bug
Reproduces a bug with go-default and bool fields.

# Table Of Contents
- [Overview](#overview)

# Overview
A bug exists in the 
[go-default library](https://github.com/mcuadros/go-defaults) which does not 
allow you to set a boolean field to false.  

See the `defaults_test.go` file. The `TestSetFalse` case will always fail.  

Run these tests by executing:

```
go test ./...
```
