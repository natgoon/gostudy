package main

import "fmt"

type book struct {
	title   string
	author  string
	book_id int
}

func printbooks(books book) {
	fmt.Printf("book name is %s", books.title)
}

type mp struct {
	info map[string]string
}

func (m mp) judge(k string) bool {
	if _, found := m.info[k]; found {
		fmt.Println(found)
		return true
	} else {
		//
		
		fmt.Println(found)
		return false
	}
}

func main() {

	var book1 book = book{"bookname", "bookauthor", 10}
	var book2 book
	var ptr *book
	book2 = book{"book2", "lihua", 20}

	book3 := book{"book3", "hua", 30}
	ptr = &book3

	book3.book_id = 40
	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book3)
	fmt.Println(book3.title, book3.book_id)
	printbooks(book1)

	fmt.Println(*ptr)

	mi := mp{}
	ma := map[string]string{"haha": "hege"}
	mi.info = ma
	fmt.Println(mi.judge("hahe"))

}
