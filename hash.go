// Package pearson implements Pearson hash as defined by RFC 3074
package pearson

import "hash"

type digest struct {
	sum byte
	tab [256]byte
}

// New returns a new 8-bit Pearson hash.Hash
func New() hash.Hash {
	return &digest{tab: rfc3074}
}

func (d *digest) Write(p []byte) (n int, err error) {
	d.sum = byte(len(p))
	for i := len(p) - 1; i >= 0; i-- {
		d.sum = d.tab[d.sum^p[i]]
	}
	return len(p), nil
}

func (d *digest) Sum(b []byte) []byte {
	return append(b, d.sum)
}

func (d *digest) Reset() {
	d.sum = 0
}

func (d *digest) Size() int {
	return 1
}

func (d *digest) BlockSize() int {
	return 1
}
