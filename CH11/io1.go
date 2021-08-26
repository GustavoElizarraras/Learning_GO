package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

// io and Friends
// io.Reader and io.writer are very used interfaces
// Each of them implement one method

type Writer interface {
	// takes in a slice of bytes and then write them
	// to the interface implmentation
	Write(p []byte) (n int, err error)
}

type Reader interface {
	// A slice input parameter is passed and then modified,
	// the method returns the number of bytes written
	Read(p []byte) (n int, err error)
}

// Example with io.Reader
func countLetters(r io.Reader) (map[string]int, error) {
	// create buffer once and reuse it on every call to r.Read,
	// one memory allocation to read from a potential large data
	// source, if it were written to return a []byte, it would require
	// a new allocation in every call, in the heap.
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf) // memory allocation under control
		// n indicate how many bytes were written to the buffer
		for _, b := range buf[:n] { // processing the read data
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		// this if is first because there might be bytes returned before an error
		// was triggered by the end of the data stream or unexpexted condition
		if err == io.EOF {
			// this error is returned if the reader is done
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

// Because countLetters() depends on a reader, we can use it to count English 
// letters in a gzip compressed file

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	// creation of an *os.File which meets the io.Reader interface
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	// gr is a *gzip.Reader instance
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
	// closer closure that cleans up our resources

func main() {
	
	s := "The quick brown fox"
	// Modifying the interfaces
	// We can created an io.Reader from a string
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		return
	}
	fmt.Println(counts)

	// Since *gzip.Reader implements io.Reader, we can use it with countLetters() 
	r, close, err := buildGZipReader("compressed_data.gz")
	if err != nil {
		return err
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		return err
	}
	fmt.Println(counts)
}

// Other useful functions:
// io.MultiReader(), io.LimitReader(), io.MultiWriter()

// Other one-method interfaces in io:
type Closer interface {
	// implemented when a cleanup is needed after a read or write
	Close() error
}
type Seeker interface {
	// random access to a resource. The valid values for whence are the 
	// constants: io.SeekStart, io.SeekCurrent and io.SeekEnd. It would 
	// have been clearer with a custom type, but at the end is an int
	Seek(offset int64, whence int) (int64, error)
}
// The io package defines interfaces that combine this four interfaces, like
// io.ReadCloser, io.ReadSeeker, io.ReadWriteSeeker, etc. Use them to specify
// what your functions expect to do with the data. For example, use the
// interfaces to specify exactly what your function will do with the parameter

// Pattern for adding a method to a Go type
type nopCloser struct {
	// implements a io.Reader but not io.Closer
	io.Reader
	// this needs to be passed to a function that expects a io.ReadCloser
}
func (nopCloser) Close() error { return nil }

func NopCloser(r io.Reader) io.ReadCloser {
	// passing the io.Reader and returning a io.ReadCloser, thanks to ioutil.NopCloser()
	return nopCloser{r}
}