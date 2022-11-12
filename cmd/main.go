package main

import (
	"main/models"
	"main/rules"
	"main/rules/domain"
)

func main() {
	visa := models.Visa{
		Bank:         "HSBC",
		Balance:      10000,
		CreditRating: 4,
	}

	acceptedVisa := models.Visa{
		Bank:         "HSBC",
		Balance:      5000,
		CreditRating: 7,
	}
	person := models.Person{
		Age:     22,
		Country: "EG",
		Gender:  "male",
		Visa:    visa,
	}
	classificationContainer := domain.ClassificationContainer{
		Person: person,
	}

	visaRule := rules.VisaCheck{
		Visa: acceptedVisa,
	}
	countryRule := rules.CountryCheck{
		CountyList: models.CountryProhibited,
		NextStep:   &visaRule,
	}
	genderRule := rules.GenderRule{
		Gender:   "male",
		NextStep: &countryRule,
	}
	ageRule := rules.AgeGreaterThan{
		Age:      20,
		NextStep: &genderRule,
	}
	ageRule.Execute(&classificationContainer)
	classificationContainer.CheckIfError()
}
