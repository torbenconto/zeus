package zeus_test

import (
	"testing"

	"github.com/torbenconto/zeus"
)

func TestGZIPCompressAndDecompressText(t *testing.T) {
	originalText := "Hello, world!"

	// Test GZIP compression
	compressed, err := zeus.GZIPCompressText(originalText)
	if err != nil {
		t.Fatalf("GZIPCompressText failed: %v", err)
	}

	// Test GZIP decompression
	decompressed, err := zeus.GZIPDecompressText(compressed)
	if err != nil {
		t.Fatalf("GZIPDecompressText failed: %v", err)
	}

	if originalText != decompressed {
		t.Fatalf("GZIP decompressed text is not equal to original text")
	}
}

func TestFlateCompressAndDecompressText(t *testing.T) {
	originalText := "Hello, world!"

	// Test Flate compression
	compressed, err := zeus.FlateCompressText(originalText)
	if err != nil {
		t.Fatalf("FlateCompressText failed: %v", err)
	}

	// Test Flate decompression
	decompressed, err := zeus.FlateDecompressText(compressed)
	if err != nil {
		t.Fatalf("FlateDecompressText failed: %v", err)
	}

	if originalText != decompressed {
		t.Fatalf("Flate decompressed text is not equal to original text")
	}
}

func TestBrotliCompressAndDecompressText(t *testing.T) {
	originalText := "Hello, world!"

	// Test Brotli compression
	compressed := zeus.BrotliCompressText(originalText)

	// Test Brotli decompression
	decompressed, err := zeus.BrotliDecompressText(compressed)
	if err != nil {
		t.Fatalf("BrotliDecompressText failed: %v", err)
	}

	if originalText != decompressed {
		t.Fatalf("Brotli decompressed text is not equal to original text")
	}
}
