package utils

import (
	"errors"
	"io"

	"github.com/charmbracelet/ssh"
)

// copied from snips.sh
func ReadInput(sesh ssh.Session, maxSize uint64) ([]byte, error) {
	content := make([]byte, 0)
	size := uint64(0)
	for {
		buf := make([]byte, UploadBufferSize)
		n, err := sesh.Read(buf)
		isEOF := errors.Is(err, io.EOF)
		if err != nil && !isEOF {
			return nil, err
		}

		size += uint64(n)
		content = append(content, buf[:n]...)

		if size > maxSize {
			return nil, ErrFileTooLarge
		}

		if isEOF {
			if size == 0 {
				return nil, ErrEmptyContent
			}
			return content, nil
		}
	}
}
