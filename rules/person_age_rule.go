package rules

import (
	"fmt"
	"main/rules/domain"
)

type AgeGreaterThan struct {
	Age      int
	NextStep ClassificationRule
}

func (agt *AgeGreaterThan) Execute(container *domain.ClassificationContainer) {
	fmt.Println("checking your age .....")
	if agt.Age > container.Person.Age {

		fmt.Println("you are too young .....")
		container.AddResult(&domain.Result{
			Rule:    "Age Rule",
			Success: false,
			Reason:  "Age is less than expected",
		})
	}

	if agt.NextStep != nil {
		fmt.Println("you can pass .....")
		agt.NextStep.Execute(container)
	}

}

func (agt *AgeGreaterThan) next(next ClassificationRule) {
	agt.NextStep = next
}
