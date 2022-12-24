package atcodergo

import (
	"io"
)

func readAllClose(body io.ReadCloser) {
	io.Copy(io.Discard, body)
	body.Close()
}
