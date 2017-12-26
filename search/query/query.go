package query

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/swiftdiaries/phone-lookup/search/store"
	"github.com/swiftdiaries/phone-lookup/search/util"
)

//CheckAndFetch checks if entry exists in Redis and fetches the result
//if it doesn't exist
func CheckAndFetch(person *util.Person) *util.Person {
	response := store.Get(person.Phonenumber)

	if response == "" {
		urls := ResultURLScrape(person.Phonenumber)
		person = FindUsernameExists(person, urls)
		data, err := json.Marshal(person)
		if err != nil {
			fmt.Printf("Error in marshalling person: %s", err)
		}
		err = store.Set(person.Phonenumber, string(data))
		if err != nil {
			log.Fatalf("Error in storing to redis: %s", err)
		}
	} else {
		var responsePerson *util.Person
		err := json.Unmarshal([]byte(response), &responsePerson)
		if err != nil {
			fmt.Printf("Error in Unmarshalling JSON string: %s", err)
		}
		if strings.Contains(responsePerson.Name, person.Name) {
			person.Address = responsePerson.Address
			person.Name = responsePerson.Name
		} else {
			fmt.Printf("%s", responsePerson.Name)
			return nil
		}

	}
	return person
}

//ResultURLScrape gives URLs of results to scrape
func ResultURLScrape(phonenumber string) []string {
	baseurl := "https://www.truepeoplesearch.com"
	resultURLs := make(map[string]bool)
	doc, err := goquery.NewDocument("https://www.truepeoplesearch.com/results?phoneno=" + phonenumber)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "View All Details") {
			href, _ := s.Attr("href")
			resultURLs[baseurl+href] = true
		}
	})
	keys := make([]string, 0, len(resultURLs))
	for k := range resultURLs {
		keys = append(keys, k)
	}
	return keys
}

//FindUsernameExists check if username has an exact match for results generated
func FindUsernameExists(person *util.Person, urls []string) *util.Person {
	for _, url := range urls {
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("title").Each(func(i int, s *goquery.Selection) {
			titleString := doc.Find("title").First().Text()
			nameCheck := strings.Split(titleString, "-")[1]
			if strings.Contains(nameCheck, person.Name) {
				person.Name = nameCheck
				//fmt.Printf("name check: %s", nameCheck)
			}
			if strings.Contains(s.Text(), person.Name) {
				sel := doc.Find("a.link-to-more").First()
				person.Address = sel.Text()
			}
		})
	}
	return person
}
