package output

import "io"

type Output interface {
	io.Writer
}
