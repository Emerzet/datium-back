package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
	"strconv"
	"errors"
)
type library struct {
		ID    int
		title string
		author string
		theme string
		pages int 
}


var (
	ErrExit = errors.New ("exit")
	books = make ([]library, 0, 400)
 	nextID = 1 
)
func readLine(r *bufio.Reader) string {
	line, _ := r.ReadString('\n')
	return strings.TrimSpace(line)


}

func AskForBook (r *bufio.Reader) (library,error) {

var ti, au, th string 
var pg int

	fmt.Print ("Podaj tytuł ksiązki lub wpisz `exit`,aby wyjść: ")

	ti = readLine(r)
	if    strings.ToLower(ti) == "exit" { 

	return library{}, ErrExit

}

	fmt.Print ("Podaj autora ksiązki ")

	au = readLine(r)

	fmt.Print ("Podaj motyw ksiązki ")
	th = readLine(r)


	fmt.Print ("Podaj ilość stron: ")
	input := readLine(r)

	pg, err := strconv.Atoi (input)
	if err != nil{
	log.Println("Błąd podaj liczbę stron")
	return err}

	book := library {title:ti, author:au, theme:th, pages:pg}

	return book, nil



}


func AddBook (book  library)(library,error)  {
	book.ID := nextID 

	nextID ++ 

	books = append(books,book)

	return book, nil


}




} 
}




func main() {
	
reader := bufio.NewReader(os.Stdin)


ksiazki = append (ksiazki, k)

fmt.Println("Dodano:", k.tytul,"-", k.autor , k.strony)

fmt.Printf( "\nRazem zapisanych książek: %d\n", len(ksiazki))

} 




