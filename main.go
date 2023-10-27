package main

import (
	"schedule/GO/schedule/db"
	"schedule/GO/schedule/scrapper"
)

func main() {
	URL := "https://cfuv.ru/raspisanie-fakultativov-fiziko-tekhnicheskogo-instituta"
	json_data := scrapper.Parse_to_json(URL)
	db.Make_db(json_data)
}
