# web-scrape-go
Software Engineering Test

A `go` web scraper that extracts specific product data from the Sainsbury's grocery site - Ripe Fruits page.

#### Pre-requisites
Built using `go 1.8`. You will need to install and configure `go` using the instructions [here](https://golang.org/dl/)


#### Installation
1. By using `Go get`:
 ```
go get -v github.com/gordonrehling2/web-scrape-go
```

2. Or you can manually clone the repo in your `$GOPATH/src/gordonrehling2` directory by:
 ```
mkdir -p $GOPATH/src/gordonrehling2
cd $GOPATH/src/gordonrehling2/
git clone git@gordonrehling2/web-scrape-go.git
```

#### Dependencies
The scraper relies on the `net/html` package which should be downloaded using `go get` as follows:
```
go get -v golang.org/x/net/html
```

#### Correct Directory
The following instructions rely on you being in the correct directory, get there as follows:
```
cd $GOPATH/src/gordonrehling2/web-scrape-go
```

#### Running the code
Run using
```
go run main.go
```

#### Testing the code
Test using
```
go test ./...
```

#### Linting the code
Lint using
```
golint ./...
```


