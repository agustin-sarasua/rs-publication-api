package app

import (
	c "github.com/agustin-sarasua/rs-common"
	m "github.com/agustin-sarasua/rs-model"
)

func validatePublication(p *m.Publication) []error {
	var errs []error

	errs = c.ValidateExistInMap(m.Operation, p.Operation, "Operation is incorrect", errs)
	errs = c.ValidateCondition(func() bool { return p.PropertyID != 0 }, "PropertyID can not be empty", errs)
	errs = c.ValidateCondition(func() bool { return p.Price <= 0 }, "Price is not valid", errs)
	errs = c.ValidateCondition(func() bool { return p.OwnerID != 0 }, "OwnerID can not be empty", errs)
	errs = c.ValidateCondition(func() bool { return p.StartDate.Before(p.EndDate) }, "StartDate is before EndDate", errs)
	errs = c.ValidateCondition(func() bool { return p.MinContractTime <= p.MaxContractTime }, "MinContractTime <= MaxContractTime", errs)

	return errs
}
