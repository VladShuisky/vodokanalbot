package parsing

import (
	"fmt"
	"regexp"
	"strings"
)

func GetContentByDate(dateStr string, sourceStrings []string) ([]string, error) {
	dateRegex := regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])\.(0[1-9]|1[0-2])\.([0-9]{4})$`)
	if !dateRegex.MatchString(dateStr) {
		return []string{}, fmt.Errorf("неверный формат даты. Ожидается: DD.MM.YYYY")
	}
	filteredStrings := []string{}
	for _, str_el := range sourceStrings {
		if strings.Contains(str_el, dateStr) {
			filteredStrings = append(filteredStrings, str_el)
		}
	}
	if (len(filteredStrings) == 0) {
		filteredStrings = append(filteredStrings, "Нет ни одного совпадения по дате!")
	}
	return filteredStrings, nil
}