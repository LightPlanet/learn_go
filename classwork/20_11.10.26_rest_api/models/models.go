package models

type Book struct {
	ID     string
	Author string
	Title  string
	Avail  bool
}

func CalcWidth(books []Book, getString func(Book) string) (ret int) {
	for _, b := range books {
		l := len(getString(b))
		if l > ret {
			ret = l
		}
	}
	return
}

type Context struct {
	Books     []Book
	UserInput string
	UserWants int
	RowFormat string
}

// "var" is needed for global variables?
var Books = []Book{
	{"1", "Donald Knuth", "The Art of Computer Programming", true},
	{"2", "Steven Skiena", "Algorithm Design Manual", false},
	{"3", "Vladimir Arnold", "A Mathematical Trivium", true},
}

func GetAvailBooks() (ret []Book) {
	ret = make([]Book, 0, len(Books))
	for _, b := range Books {
		if b.Avail {
			ret = append(ret, b)
		}
	}
	return
}

func TakeBookByID(id string) *Book {
	for i := 0; i < len(Books); i++ {
		b := &Books[i]
		if b.ID == id {
			b.Avail = false
			return b
		}
	}
	return nil
}

func FilterByAuthor(author string) (ret []Book) {
	ret = make([]Book, 0, len(Books))
	for _, b := range Books {
		if b.Avail && b.Author == author {
			ret = append(ret, b)
		}
	}
	return
}
