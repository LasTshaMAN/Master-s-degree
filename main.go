package main

import (
	"fmt"
	"log"

	"strings"

	"io/ioutil"

	"github.com/fedesog/webdriver"
)

func main() {
	// -D
	//defer profile(time.Now(), "main")

	// TODO
	// Fix these warnings
	// Some web-driver logs pollute output. We better turn them off.
	log.SetOutput(ioutil.Discard)

	chromeDriver := webdriver.NewChromeDriver("/home/lastshaman/chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
		return
	}
	defer chromeDriver.Stop()

	err = initSession(chromeDriver)
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Delete()

	//query := "put plate on a table"
	//query := "table with data"

	//query := "surf the web"
	//query := "search the web"
	//query := "silky spider web"

	//query := "leave a like"
	//query := "to be like him"
	//query := "Jack likes running"

	//query := "hard metal"
	//query := "hard work"

	//query := "flash of light"
	query := "light shoes"

	// -D
	fmt.Printf("query: '%s'\n", query)

	ctxCache := newCtxCache()

	queryByWords := strings.Split(query, " ")
	for i, word := range queryByWords {

		// -D
		fmt.Printf("word: '%s'\n", word)

		for _, meaning := range extractMeanings(word) {
			// TODO
			// Optimize code in order to be able to consider less likely meaning-candidates as well
			if meaning.IsNotRare() {
				for _, representation := range meaning.Representations() {
					rCtx := ctxCache.GetCtx(representation)
					queryWithoutWord := cutOutWord(queryByWords, i)
					for _, hWord := range queryWithoutWord {
						hCtx := ctxCache.GetCtx(hWord)

						// TODO
						// Increase rank proportionally to conceptnet relevance rating

						meaning.IncreaseRank(compareContexts(rCtx, hCtx))
					}
				}

				// -D
				fmt.Printf("synonyms: '%v' rus: '%v' rank: '%v'\n", meaning.Representations(), meaning.Translation(), meaning.Rank())
			}

			// TODO
			// Add filtering by checking that all the representatives for a meaning exist in word's context
			// (then and only then they can be treated as synonyms)
		}
	}
}

func cutOutWord(queryByWords []string, idx int) (result []string) {
	result = append(result, queryByWords[:idx]...)
	result = append(result, queryByWords[idx+1:]...)
	return
}

func compareContexts(rCtx Context, hCtx Context) (result int) {
	rCtxIndex := map[string]struct{}{}
	for _, word := range rCtx.Words() {
		rCtxIndex[word] = struct{}{}
	}
	for _, word := range hCtx.Words() {
		if _, ok := rCtxIndex[word]; ok {
			result++
		}
	}
	return
}
