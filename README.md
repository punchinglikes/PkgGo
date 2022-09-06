# PkgGo
## A Go Package that grabs information about Go Packages :)

PkgGo is Package that has many useful functions on obtaining various information about Go Packages through [pkg.go.dev](https://pkg.go.dev). This is actually my first "real" project with Golang, so do not use this for production purposes (for now).

## Features

- Query Go Package information
- Get's the number of packages returned (max 50)
- Get the Package link

The Project is still in it's early stages. Many features are still missing

## Installation

PkgGo requires [Golang](https://go.dev/). Developed on v1.19

Assuming you already have a mod file;
In your current project, run:

```sh
go get github.com/punchinglikes/PkgGo@latest
```
If PkgGo was recently updated, and the command above is not installing the correct version, try setting your GOPROXY environment variable to "direct" (You can thank google for that)

## Usage

Currently PkgGo only has 2 functions:
### PackageLink
```PackageLink(PackageName) - []string```
Returns an array of the top 5 query results.
**Example:**
```go
s := PkgGo.PackageLink("fmt")
fmt.Println(s[0])
```
### PackageCount
```PackageCount(PackageName) - int```
Returns an integer of number of packages that came back from the query (max 50)
**Example:**
```Go
s := PkgGo.PackageCount("fmt")
fmt.Println(s)
```
**Note: The query is capped at 50, if this function returns back 50, good chances are it's more than that.**
## Dependencies

Nothing the user has to worry about, however it should be noted that PkgGo uses goquery for pretty much every function. Goquery can be found [here](https://pkg.go.dev/github.com/PuerkitoBio/goquery).

## License

I call it the "Do what you want" License. Do whatever you want to do with my code, I do not mind.

