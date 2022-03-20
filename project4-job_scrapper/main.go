package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// var baseURL string = "https://kr.indeed.com/jobs?=q=python&limit=50"

var baseURL string = "https://kr.indeed.com/jobs?=q=python"
var LastpageURL string = "https://kr.indeed.com/jobs?q=python&limit=50&start=9999"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)

	// for i := 0; i < totalPages; i++ {
	// 	getPage(i)
	// }
}

func getPage(page int) {

}

//baseURL 참조 변경 및 페이지 추출 함수
func getPages() int {
	pages := 0
	res, err := http.Get(LastpageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages, err = strconv.Atoi(s.Find("b").Text())
	})
	return pages
}

// func getPages() int {
// 	pages := 0

// 	res, err := http.Get(baseURL)
// 	checkErr(err)
// 	checkCode(res)

// 	defer res.Body.Close()

// 	doc, err := goquery.NewDocumentFromReader(res.Body)
// 	checkErr(err)

// 	// doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
// 	// 	// fmt.Println(s.Html())
// 	// 	// fmt.Println(s.Find("a").Length())
// 	// 	pages = s.Find("a").Length()
// 	// })

// 	doc.Find("#searchCountPages").Each(func(i int, s *goquery.Selection) {
// 		result := strings.TrimSpace(s.Text()) //Remove Empty Space
// 		r, _ := regexp.Compile("([0-9])")     // Get Numbers in result
// 		//Slice the List Becuase base URL page is always 0
// 		pageList := r.FindAllString(result, -1)[1:]
// 		//Make list to string and change type
// 		new, _ := strconv.Atoi(strings.Join(pageList, ""))
// 		pages = new / 50 //Set each page show 50 results
// 	})

// 	return pages
// }

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
