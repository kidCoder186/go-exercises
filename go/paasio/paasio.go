package paasio

import (
	"io"
	"sync"
)

// ReadWrapper structure
type ReadWrapper struct {
	rd         io.Reader
	numBytes   int64
	numReadOps int
	mux        sync.Mutex
}

func (w *ReadWrapper) Read(p []byte) (n int, err error) {
	w.mux.Lock()
	defer w.mux.Unlock()
	w.numReadOps++
	n, err = w.rd.Read(p)
	w.numBytes += int64(n)
	return
}

// ReadCount returns number of bytes successfully read and number of read operations.
func (w *ReadWrapper) ReadCount() (n int64, nops int) {
	w.mux.Lock()
	defer w.mux.Unlock()
	return w.numBytes, w.numReadOps
}

// WriteWrapper structure
type WriteWrapper struct {
	wd          io.Writer
	numBytes    int64
	numWriteOps int
	mux         sync.Mutex
}

func (w *WriteWrapper) Write(p []byte) (n int, err error) {
	w.mux.Lock()
	defer w.mux.Unlock()
	w.numWriteOps++
	n, err = w.wd.Write(p)
	w.numBytes += int64(n)
	return
}

// WriteCount returns number of bytes successfully write and number of write operations.
func (w *WriteWrapper) WriteCount() (n int64, nops int) {
	w.mux.Lock()
	defer w.mux.Unlock()
	return w.numBytes, w.numWriteOps
}

// ReadWriteWrapper structure
type ReadWriteWrapper struct {
	ReadCounter
	WriteCounter
}

// NewReadCounter creates a new ReadWrapper object.
func NewReadCounter(r io.Reader) ReadCounter {
	return &ReadWrapper{r, 0, 0, sync.Mutex{}}
}

// NewWriteCounter creates a new WriteWrapper object.
func NewWriteCounter(w io.Writer) WriteCounter {
	return &WriteWrapper{w, 0, 0, sync.Mutex{}}
}

// NewReadWriteCounter creates a new ReadWriteWrapper object.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteWrapper {
	r := NewReadCounter(rw)
	w := NewWriteCounter(rw)
	return ReadWriteWrapper{r, w}
}
