package gzip

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"

	"github.com/mcrgnt/yp1/internal/common"
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

func Decompress(r io.Reader) (io.Reader, error) {
	if zr, err := gzip.NewReader(r); err != nil {
		return nil, fmt.Errorf("new reader failed: %w", err)
	} else {
		defer func() {
			if err := zr.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		var b bytes.Buffer

		for {
			_, err := io.CopyN(&b, zr, common.CopyBytes)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return nil, fmt.Errorf("copyn failed: %w", err)
			}
		}
		return &b, nil
	}
}
