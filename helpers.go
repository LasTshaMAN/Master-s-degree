package main

import "strings"

func removeDuplicates(backTranslations []string) (result []string) {
	backTranslationsUnique := map[string]struct{}{}
	for _, backTranslation := range backTranslations {
		if _, ok := backTranslationsUnique[backTranslation]; !ok {
			result = append(result, backTranslation)
			backTranslationsUnique[backTranslation] = struct{}{}
		}
	}
	return
}

func removeMultiWordPhrases(backTranslations []string) (result []string) {
	for _, backTranslation := range backTranslations {
		if !strings.Contains(backTranslation, " ") {
			result = append(result, backTranslation)
		}
	}
	return
}
