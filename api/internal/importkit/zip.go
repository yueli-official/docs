package importkit

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime"
	"path"
	"sort"
	"strings"

	"github.com/ulikunitz/xz"
)

const (
	zipMethodXZ        uint16 = 95
	zipXZDictionaryCap        = 64 * 1024 * 1024
)

type UnsupportedCompressionError struct {
	Method uint16
}

func (err *UnsupportedCompressionError) Error() string {
	return fmt.Sprintf("unsupported zip compression method %d", err.Method)
}

type xzReadCloser struct {
	reader io.Reader
	err    error
}

func (r *xzReadCloser) Read(p []byte) (int, error) {
	if r.err != nil {
		err := r.err
		r.err = nil
		return 0, err
	}
	return r.reader.Read(p)
}

func (*xzReadCloser) Close() error { return nil }

func openXZEntry(compressed io.Reader) io.ReadCloser {
	reader, err := (xz.ReaderConfig{
		DictCap:      zipXZDictionaryCap,
		SingleStream: true,
	}).NewReader(compressed)
	return &xzReadCloser{reader: reader, err: err}
}

func readZipFiles(data []byte, opts Options) (map[string][]byte, error) {
	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, fmt.Errorf("invalid zip: %w", err)
	}
	zr.RegisterDecompressor(zipMethodXZ, openXZEntry)
	files := map[string][]byte{}
	var extractedBytes uint64
	for _, f := range zr.File {
		if f.FileInfo().IsDir() {
			continue
		}
		name := cleanZipPath(f.Name)
		if name == "" || name == "." || strings.HasPrefix(name, "../") || strings.Contains(name, "/../") {
			return nil, fmt.Errorf("unsafe zip path: %s", f.Name)
		}
		if len(files) >= opts.MaxEntries {
			return nil, fmt.Errorf("archive exceeds %d file entries", opts.MaxEntries)
		}
		if f.UncompressedSize64 > uint64(opts.MaxEntryBytes) {
			return nil, fmt.Errorf("archive entry exceeds %d bytes: %s", opts.MaxEntryBytes, name)
		}
		if f.UncompressedSize64 > uint64(opts.MaxExtractedBytes) || extractedBytes > uint64(opts.MaxExtractedBytes)-f.UncompressedSize64 {
			return nil, fmt.Errorf("archive expands beyond %d bytes", opts.MaxExtractedBytes)
		}
		if f.Method != zip.Store && f.Method != zip.Deflate && f.Method != zipMethodXZ {
			return nil, &UnsupportedCompressionError{Method: f.Method}
		}
		extractedBytes += f.UncompressedSize64
		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("open archive entry %s: %w", name, err)
		}
		body, readErr := io.ReadAll(io.LimitReader(rc, opts.MaxEntryBytes+1))
		closeErr := rc.Close()
		if readErr != nil {
			return nil, fmt.Errorf("read archive entry %s: %w", name, readErr)
		}
		if closeErr != nil {
			return nil, fmt.Errorf("close archive entry %s: %w", name, closeErr)
		}
		if int64(len(body)) > opts.MaxEntryBytes {
			return nil, fmt.Errorf("archive entry exceeds %d bytes: %s", opts.MaxEntryBytes, name)
		}
		files[name] = body
	}
	return files, nil
}

func sortedKeys(files map[string][]byte) []string {
	keys := make([]string, 0, len(files))
	for key := range files {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func isMarkdownPath(p string) bool {
	ext := strings.ToLower(path.Ext(p))
	if ext != ".md" {
		return false
	}
	for _, segment := range strings.Split(p, "/") {
		switch strings.ToLower(segment) {
		case "images", "image", "img", "assets", "_assets", "_static", "static":
			return false
		}
	}
	return true
}

func isImagePath(p string) bool {
	switch strings.ToLower(path.Ext(p)) {
	case ".png", ".jpg", ".jpeg", ".webp", ".gif", ".svg":
		return true
	default:
		return false
	}
}

func assetFromFile(source string, body []byte) AssetFile {
	sum := sha256.Sum256(body)
	mimeType := mime.TypeByExtension(strings.ToLower(path.Ext(source)))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	return AssetFile{
		SourcePath: source,
		Bytes:      body,
		Mime:       mimeType,
		Size:       int64(len(body)),
		SHA256:     hex.EncodeToString(sum[:]),
	}
}
