package book

type Book struct {
	ID       int64
	Title    string
	Author   string
	Category Category
}

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
