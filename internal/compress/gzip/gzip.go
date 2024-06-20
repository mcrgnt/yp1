package gzip

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func Compress(data []byte) (io.Reader, error) {
	var buf *bytes.Buffer
	g := gzip.NewWriter(buf)

	if _, err := g.Write(data); err != nil {
		return nil, fmt.Errorf("write failed: %w", err)
	}

	if err := g.Close(); err != nil {
		return nil, fmt.Errorf("close failed: %w", err)
	}
	return buf, nil
}

func Decompress()
