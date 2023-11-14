package API

import (
	"schedule/GO/schedule/db"
	"schedule/GO/schedule/scrapper"
)

func Update(URL string) {
	schedule := scrapper.Parse(URL)
	db.Make_db(schedule)
}

func Get_info_about(group string, year int, month int, day int) string {
	return db.Info_about(group, year, month, day)
}
func Next(group string) string {
	return db.Next_pair(group)
}
