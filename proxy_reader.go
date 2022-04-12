package bombardier

import "io"

type proxyReader struct {
	io.Reader
}
