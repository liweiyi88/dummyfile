[![GoDoc](https://godoc.org/github.com/liweiyi88/dummyfile?status.svg)](https://godoc.org/github.com/liweiyi88/dummyfile)
![tests](https://github.com/liweiyi88/dummyfile/actions/workflows/tests.yaml/badge.svg)
[![codecov](https://codecov.io/gh/liweiyi88/dummyfile/branch/main/graph/badge.svg?token=ROIDLHX41V)](https://codecov.io/gh/liweiyi88/dummyfile)
[![Go Report Card](https://goreportcard.com/badge/github.com/liweiyi88/dummyfile)](https://goreportcard.com/report/github.com/liweiyi88/dummyfile)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/liweiyi88/dummyfile/blob/main/LICENSE.md)

Dummyfile is a cli tool/package for creating a fixed-size dummy file with random contents.

## Usage

### Install as package
`go get -u github.com/liweiyi88/dummyfile/dummyfile`

Then use it in your code
```
import "github.com/liweiyi88/dummyfile/dummyfile"

func main() {
    file, err := os.Create(path)
    if err != nil {
        return fmt.Errorf("failed to create dummy file: %v", err)
    }

    defer func() {
        err := file.Close()
        if err != nil {
            log.Printf("failed to close file: %s", file.Name())
        }
    }()

    // as the Create method accepts the io.Writer.
    // you have the flexibility to test not only the file but any io.Writer.
    err := dummyfile.Create(file, 100*dummyfile.GB)
    //....
}
```

### Install as binary
`dummyfile` binaries are available in https://github.com/liweiyi88/dummyfile/releases. Use the latest version of the binary that is suitable to your OS.
After downloading the binary and move it to the folder that is in your $PATH env var (e.g. `/usr/local/bin/dummyfile`), give it executable permissions (e.g. `sudo chmod +x /usr/local/bin/dummyfile`). Then you should be able to run it:
```
$ dummyfile /path/to/dummy.txt --size 100gb
```
