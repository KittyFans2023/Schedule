package main

import (
	"net/http"
	"schedule/GO/schedule/API"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/schedule", homeHendler)
	router.HandleFunc("/schedule/update", updateHendler)
	router.HandleFunc("/schedule/get_info/{group_number}", groupHendler)
	// router.HandleFunc("/schedule/comment/")
	http.ListenAndServe(":8080", router)
}

func homeHendler(rw http.ResponseWriter, _ *http.Request) {
	rw.Write([]byte(`Пустая`))
}

//	func teacherHendler(rw http.ResponseWriter, req *http.Request){
//		value := mux.Vars(req)
//		name := value{"name"}
//		json_data := API
//	}
func groupHendler(rw http.ResponseWriter, req *http.Request) {
	value := mux.Vars(req)
	group := value["group_number"]
	json_data := API.Get_info_about(group)              //получаем наш документ формата json
	rw.Header().Set("Content-Type", "application/json") //устанавливаем какой будет контент страницы
	rw.Write([]byte(json_data))                         //переводим в массив байтов

}
func updateHendler(rw http.ResponseWriter, req *http.Request) {
	API.Update("https://cfuv.ru/raspisanie-fakultativov-fiziko-tekhnicheskogo-instituta")
}
