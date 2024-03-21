package rofi

import (
	"fmt"
	"io"
)

func (r *Application) WriteTo(w io.Writer) (int64, error) {
	b := 0
	bytesWritten := int64(b)
	if r.Prompt != "" {
		b, err := w.Write([]byte(fmt.Sprintf("%c%s%c%s\n", 0x00, "prompt", 0x1f, r.Prompt)))
		if err != nil {
			return bytesWritten, err
		} else {
			bytesWritten += int64(b)
		}
	}
	if r.Message != "" {
		b, err := w.Write([]byte(fmt.Sprintf("%c%s%c%s\n", 0x00, "message", 0x1f, r.Message)))
		if err != nil {
			return bytesWritten, err
		} else {
			bytesWritten += int64(b)
		}
	}
	for _, v := range r.Commands {
		b, err := v.WriteTo(w)
		if err != nil {
			return int64(bytesWritten), err
		} else {
			bytesWritten += b
		}
	}
	return bytesWritten, nil
}

func (e *Command) WriteTo(w io.Writer) (int64, error) {
	b := 0
	bytesWritten := int64(b)
	var err error
	if b, err = w.Write([]byte(e.Name)); err != nil {
		return bytesWritten, err
	} else {
		bytesWritten += int64(b)
	}
	if b, err = w.Write([]byte{0x00}); err != nil {
		return bytesWritten, err
	} else {
		bytesWritten += int64(b)
	}
	if e.Info != "" {
		if b, err = w.Write([]byte(fmt.Sprintf("info%c%s", 0x1f, e.Info))); err != nil {
			return bytesWritten, err
		} else {
			bytesWritten += int64(b)
		}
	}
	b, err = w.Write([]byte("\n"))
	if err != nil {
		return bytesWritten, err
	} else {
		bytesWritten += int64(b)
	}
	return bytesWritten, err
}
