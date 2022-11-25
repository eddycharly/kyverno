package policy

import (
	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
)

type Policies []kyvernov1.PolicyInterface

type Predicate = func(kyvernov1.PolicyInterface) bool

func (policies Policies) Where(predicate Predicate) Policies {
	var result []kyvernov1.PolicyInterface
	for _, policy := range policies {
		if predicate(policy) {
			result = append(result, policy)
		}
	}
	return result
}
