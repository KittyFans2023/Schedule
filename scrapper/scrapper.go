package scrapper

import (
	"schedule/GO/schedule/excel_scrapper"

	"strings"

	"github.com/gocolly/colly"
)

func getDowloandLink(URL string) string {
	scrapper := colly.NewCollector()
	var link string
	scrapper.OnHTML(".card-body", func(elem *colly.HTMLElement) {

		if strings.Contains(elem.Text, "Программная") {
			elem.ForEach("a", func(_ int, f_elem *colly.HTMLElement) {
				link = f_elem.Attr("href")
			})
		}
	})
	scrapper.Visit(URL)
	return link
}

func dowloand(Link string) {
	URL := getDowloandLink(Link)
	scrapper := colly.NewCollector()
	var link string
	scrapper.OnHTML("#downloadFile", func(elem *colly.HTMLElement) {
		link = elem.Attr("href")
		dowl_scrapper := colly.NewCollector()
		dowl_scrapper.OnResponse(func(r *colly.Response) {
			if strings.Contains(link, "xlsx") {
				r.Save("schedule/excel_scrapper/PI.xlsx")
			}
		})
		dowl_scrapper.Visit(link)
	})
	scrapper.Visit(URL)
}

func Parse_to_json(Link string) map[string]map[int][][]string {
	dowloand(Link)
	schedule := excel_scrapper.Read_schedule()
	// jsformat.To_json(schedule)
	// return jsformat.To_json(schedule)
	return schedule
}
