package main

// import "fmt"

func main(){

}

// In the same way you can embed a type in a struct, you can
// embed an interface within an interface. 

// Ex - the io.ReadCloser is built out of an io.Reader and an io.Closer

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReaderCloser interface {
	Reader
	Closer
}