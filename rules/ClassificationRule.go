package rules

import "main/rules/domain"

type ClassificationRule interface {
	Execute(container *domain.ClassificationContainer)
	next(next ClassificationRule)
}
