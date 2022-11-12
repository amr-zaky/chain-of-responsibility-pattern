package rules

import (
	"fmt"
	"main/models"
	"main/rules/domain"
)

type VisaCheck struct {
	Visa     models.Visa
	nextStep ClassificationRule
}

func (vc *VisaCheck) next(next ClassificationRule) {
	vc.nextStep = next
}

func (vc *VisaCheck) Execute(container *domain.ClassificationContainer) {
	fmt.Println("checking your visa .....")
	if vc.Visa.Balance > container.Person.Visa.Balance {

		container.AddResult(&domain.Result{
			Rule:    "Visa Rule",
			Success: false,
			Reason:  "Visa not accepted balance not enough",
		})
	}

	if vc.Visa.CreditRating > container.Person.Visa.CreditRating {
		container.AddResult(&domain.Result{
			Rule:    "Visa Rule",
			Success: false,
			Reason:  "Visa not accepted CreditRating is low ",
		})
	}

	if vc.nextStep != nil {
		fmt.Println("visa  passed successfully.....")
		vc.nextStep.Execute(container)
	}
}
