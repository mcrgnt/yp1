package gzip

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
)

const (
	CopyBytes = 1024
)

func Compress(data []byte) (io.Reader, error) {
	var buf bytes.Buffer
	g := gzip.NewWriter(&buf)
	if _, err := g.Write(data); err != nil {
		return nil, fmt.Errorf("write failed: %w", err)
	}
	if err := g.Close(); err != nil {
		return nil, fmt.Errorf("close failed: %w", err)
	}
	return &buf, nil
}

type DecompressParams struct {
	Reader io.Reader
}

func Decompress(params *DecompressParams) (reader io.Reader, err error) {
	var (
		gzipReader *gzip.Reader
	)
	if gzipReader, err = gzip.NewReader(params.Reader); err != nil {
		err = fmt.Errorf("new reader failed: %w", err)
		return
	}
	defer func() {
		if e := gzipReader.Close(); e != nil {
			reader = nil
			err = errors.Join(err, fmt.Errorf("gzip reader close failed: %w", e))
		}
	}()

	var b bytes.Buffer
	for {
		if _, e := io.CopyN(&b, gzipReader, CopyBytes); e != nil {
			if errors.Is(e, io.EOF) {
				break
			}
			err = errors.Join(err, fmt.Errorf("copyn failed: %w", e))
			return
		}
	}
	reader = &b
	return
}
