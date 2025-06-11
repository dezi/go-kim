package simple

import "io"

func ReadAll(reader io.Reader, size int64) (data []byte, err error) {

	if size <= 0 {
		size = 10 * 1024
	}

	data = make([]byte, 0, size)
	xfer := 0

	for {

		if len(data) == cap(data) {
			data = append(data, 0)[:len(data)]
		}

		xfer, err = reader.Read(data[len(data):cap(data)])
		data = data[:len(data)+xfer]

		if err != nil {

			if err == io.EOF {
				err = nil
			}

			return
		}
	}
}
