package zeus

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"io"

	"github.com/andybalholm/brotli"
)

// GZIPCompressText compresses a given text using the GZIP compression algorithm.
//
// Parameters:
//   - in: The input text to be compressed.
//
// Returns:
//   - The compressed data as a byte slice.
//   - An error if compression fails.
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

// GZIPDecompressText decompresses GZIP compressed data into a string.
//
// Parameters:
//   - in: The GZIP compressed data as a byte slice.
//
// Returns:
//   - The decompressed text as a string.
//   - An error if decompression fails.
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

// FlateCompressText compresses a given text using the Flate (DEFLATE) compression algorithm.
//
// Parameters:
//   - in: The input text to be compressed.
//
// Returns:
//   - The compressed data as a byte slice.
//   - An error if compression fails.
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

// FlateDecompressText decompresses Flate (DEFLATE) compressed data into a string.
//
// Parameters:
//   - in: The Flate compressed data as a byte slice.
//
// Returns:
//   - The decompressed text as a string.
//   - An error if decompression fails.
func FlateDecompressText(in []byte) (string, error) {
	r := flate.NewReader(bytes.NewReader(in))
	defer r.Close()

	decompressed, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(decompressed), nil
}

// BrotliCompressText compresses a given text using the Brotli compression algorithm.
//
// Parameters:
//   - data: The input text to be compressed.
//
// Returns:
//   - The compressed data as a byte slice.
func BrotliCompressText(data string) []byte {
	var b bytes.Buffer
	w := brotli.NewWriterLevel(&b, brotli.BestCompression)
	w.Write([]byte(data))
	w.Close()
	return b.Bytes()
}

// BrotliDecompressText decompresses Brotli compressed data into a string.
//
// Parameters:
//   - in: The Brotli compressed data as a byte slice.
//
// Returns:
//   - The decompressed text as a string.
//   - An error if decompression fails.
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
