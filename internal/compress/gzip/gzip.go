package gzip

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"

	"github.com/rs/zerolog"
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
	Logger *zerolog.Logger
}

func Decompress(params *DecompressParams) (io.Reader, error) {
	if gzipReader, err := gzip.NewReader(params.Reader); err != nil {
		return nil, fmt.Errorf("new reader failed: %w", err)
	} else {
		defer func() {
			if err := gzipReader.Close(); err != nil {
				params.Logger.Error().Msg(err.Error())
			}
		}()

		var b bytes.Buffer
		for {
			_, err := io.CopyN(&b, gzipReader, CopyBytes)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return nil, fmt.Errorf("copy failed: %w", err)
			}
		}
		return &b, nil
	}
}
