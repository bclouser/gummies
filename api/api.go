package api 


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}


func getAllStats(writer http.ResponseWriter, request *http.Request) {

}

func setRoomLights(writer http.ResponseWriter, request *http.Request) {

}

func setRoomFan(writer http.ResponseWriter, request *http.Request) {

}

func setProjector(writer http.ResponseWriter, request *http.Request) {

}



func viewHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/view/"):]
	time.Sleep(5 * 1000 * 1000 * 1000)
	page, _ := loadPage(title)
	fmt.Fprintf(writer, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func editHandler(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(5 * 1000 * 1000 * 1000)
	fmt.Fprintf(writer, "<h1>Editing</h1><div>Doing some editing</div>")
}

func saveHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Saving</h1><div>Doing some saving!</div>")
}



func main() {
	fmt.Println("Creating server")
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}