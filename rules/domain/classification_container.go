package domain

import (
	"fmt"
	"main/models"
)

type ClassificationContainer struct {
	Person models.Person
	Result []*Result
}

// AddResult add the classification results
func (cc *ClassificationContainer) AddResult(result *Result) {
	cc.Result = append(cc.Result, result)
}

func (cc *ClassificationContainer) CheckIfError() {
	for _, result := range cc.Result {
		fmt.Println(result.Rule)
		fmt.Println(result.Reason)
	}
}
