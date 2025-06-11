package simple

import (
	"bytes"
	"compress/gzip"
	"io"
)

func EnGzip(data []byte) []byte {

	var buffer bytes.Buffer
	var gZipper *gzip.Writer

	gZipper, err := gzip.NewWriterLevel(&buffer, gzip.BestCompression)
	if err != nil {
		gZipper = gzip.NewWriter(&buffer)
	}

	_, _ = gZipper.Write(data)
	_ = gZipper.Close()

	return buffer.Bytes()
}

func UnGzip(data []byte) ([]byte, error) {

	buffer := bytes.NewReader(data)

	gZipper, err := gzip.NewReader(buffer)

	if err != nil {
		return nil, err
	}

	unzippedData, err := io.ReadAll(gZipper)

	_ = gZipper.Close()

	return unzippedData, err
}
