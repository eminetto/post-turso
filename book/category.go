package book

import "bytes"

type Category int

const (
	WantToRead Category = iota + 1
	Reading
	Read
)

func (c Category) String() string {
	switch c {
	case WantToRead:
		return "Want to Read"
	case Read:
		return "Read"
	case Reading:
		return "Reading"
	}
	return "Unknown"
}

func (c Category) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(c.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func NewCategory(s string) Category {
	switch s {
	case "Want to Read":
		return WantToRead
	case "Read":
		return Read
	case "Reading":
		return Reading
	}
	return WantToRead
}
