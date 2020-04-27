package circular

import "errors"

// Buffer structure
type Buffer struct {
	head, tail int
	num, size  int
	data       []byte
}

// NewBuffer creates a new Buffer object.
func NewBuffer(s int) *Buffer {
	return &Buffer{
		head: 0, tail: -1,
		num: 0, size: s,
		data: make([]byte, s),
	}
}

// ReadByte reads the oldest value in the buffer.
func (b *Buffer) ReadByte() (byte, error) {
	if b.num == 0 {
		return 0, errors.New("buffer is empty")
	}
	res := b.data[b.head]
	b.head = (b.head + 1) % b.size
	b.num--
	return res, nil
}

// WriteByte writes a byte into the buffer.
// If buffer is full returns a error.
func (b *Buffer) WriteByte(c byte) error {
	if b.num == b.size {
		return errors.New("buffer is full, using Overwrite method to write into buffer")
	}
	b.tail = (b.tail + 1) % b.size
	b.data[b.tail] = c
	b.num++
	return nil
}

// Overwrite writes a byte into the buffer with a forced write.
func (b *Buffer) Overwrite(c byte) {
	b.tail = (b.tail + 1) % b.size
	b.data[b.tail] = c
	if b.num < b.size {
		b.num++
	} else if b.num == b.size {
		b.head = (b.head + 1) % b.size
	}
}

// Reset puts buffer into empty state.
func (b *Buffer) Reset() {
	b.head, b.tail = 0, -1
	b.num = 0
	for i := 0; i < b.size; i++ {
		b.data[i] = 0
	}
}
