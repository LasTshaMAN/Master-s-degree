package main

import (
	"log"
	"strings"

	"fmt"

	"github.com/fedesog/webdriver"
)

func extractMeanings(word string) (result []meaning) {
	url := fmt.Sprintf("https://translate.google.com/#en/ru/%s", word)
	err := setSessionURL(url)
	if err != nil {
		log.Println(err)
		return
	}

	tableEntries, err := extractTableEntries()
	if err != nil {
		log.Println(err)
		return
	}
	for _, elem := range tableEntries {
		translation := extractTranslation(elem)
		rank := extractStyle(elem)
		backTranslations := extractBackTranslations(elem)
		if backTranslations != nil {
			backTranslations = removeWord(backTranslations, word)
			backTranslations = removeDuplicates(backTranslations)
			backTranslations = removeMultiWordPhrases(backTranslations)
			meaning := newMeaning(backTranslations, rank, translation)
			result = append(result, meaning)
		}
	}
	return
}

func extractTableEntries() ([]webdriver.WebElement, error) {
	result, err := session.FindElements(webdriver.CSS_Selector, "#gt-lc > div.gt-cc > div.gt-cc-r > div > div > div.gt-cd-c > table > tbody > tr")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func extractTranslation(elem webdriver.WebElement) string {
	translationDiv, err := elem.FindElement(webdriver.CSS_Selector, "td:nth-child(2) > div > span")
	if err != nil {
		log.Println(err)
		return ""
	}
	translation, err := translationDiv.Text()
	if err != nil {
		log.Println(err)
		return ""
	}
	return translation
}

func extractBackTranslations(elem webdriver.WebElement) (result []string) {
	backTranslationSpans, err := elem.FindElements(webdriver.CSS_Selector, "td:nth-child(3) > div > span")
	if err != nil {
		log.Println(err)
		return
	}
	for _, backTranslationSpan := range backTranslationSpans {
		backTranslation, err := backTranslationSpan.Text()
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, backTranslation)
	}
	return
}

func extractStyle(elem webdriver.WebElement) string {
	rankDiv, err := elem.FindElement(webdriver.CSS_Selector, "td:nth-child(1) > div > div")
	if err != nil {
		log.Println(err)
		return ""
	}
	rank, err := rankDiv.GetAttribute("style")
	if err != nil {
		log.Println(err)
		return ""
	}
	rank = strings.Replace(rank, "width: ", "", -1)
	rank = strings.Replace(rank, "px;", "", -1)
	return rank
}

func removeWord(backTranslations []string, word string) (result []string) {
	for _, backTranslation := range backTranslations {
		if backTranslation != word {
			result = append(result, backTranslation)
		}
	}
	return
}
