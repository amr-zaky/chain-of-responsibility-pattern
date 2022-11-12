package rules

import (
	"fmt"
	"main/rules/domain"
)

type GenderRule struct {
	Gender   string
	NextStep ClassificationRule
}

func (gr GenderRule) Execute(container *domain.ClassificationContainer) {
	fmt.Println("checking your gender .....")
	if gr.Gender != container.Person.Gender {
		container.AddResult(&domain.Result{
			Rule:    "Gender Rule",
			Success: false,
			Reason:  "Gender not allowed to ender",
		})
	}
	fmt.Println("gender  checking successfully .....")
	if gr.NextStep != nil {
		gr.NextStep.Execute(container)
		return
	}

}

func (gr *GenderRule) next(rule ClassificationRule) {
	gr.NextStep = rule
}
