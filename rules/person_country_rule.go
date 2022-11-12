package rules

import (
	"fmt"
	"main/rules/domain"
)

type CountryCheck struct {
	CountyList []string
	NextStep   ClassificationRule
}

func (ct *CountryCheck) Execute(container *domain.ClassificationContainer) {
	fmt.Println("checking your County  .....")
	if isRejectedCountries(ct.CountyList, container.Person.Country) {
		container.AddResult(&domain.Result{
			Rule:    "Country Rule",
			Success: false,
			Reason:  "your County is  in the Block List",
		})
	}
	fmt.Println("Country  checking successfully .....")
	if ct.NextStep != nil {
		ct.NextStep.Execute(container)
	}

}

func (ct *CountryCheck) next(rule ClassificationRule) {
	ct.NextStep = rule
}

func isRejectedCountries(countyList []string, country string) bool {

	for _, v := range countyList {
		if v == country {
			return true
		}
	}
	return false
}
