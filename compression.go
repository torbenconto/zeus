package zeus

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"io"

	"github.com/andybalholm/brotli"
)

func GZIPCompressText(in string) ([]byte, error) {
	var b bytes.Buffer

	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(in)); err != nil {
		return nil, err
	}
	if err := gz.Flush(); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func GZIPDecompressText(in []byte) (string, error) {
	rdata := bytes.NewReader(in)

	r, err := gzip.NewReader(rdata)
	if err != nil {
		return "", err
	}

	s, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(s), nil
}

func FlateCompressText(in string) ([]byte, error) {
	var b bytes.Buffer
	w, err := flate.NewWriter(&b, 9)
	if err != nil {
		return nil, err
	}
	w.Write([]byte(in))
	w.Close()
	return b.Bytes(), nil
}

func FlateDecompressText(in []byte) (string, error) {
	r := flate.NewReader(bytes.NewReader(in))
	defer r.Close()

	decompressed, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(decompressed), nil
}

func BrotliCompressText(data string) []byte {
	var b bytes.Buffer
	w := brotli.NewWriterLevel(&b, brotli.BestCompression)
	w.Write([]byte(data))
	w.Close()
	return b.Bytes()
}

func BrotliDecompressText(in []byte) (string, error) {
	r := bytes.NewReader(in)
	br := brotli.NewReader(r)

	var decompressed bytes.Buffer
	_, err := decompressed.ReadFrom(br)
	if err != nil {
		return "", err
	}

	return decompressed.String(), nil
}
