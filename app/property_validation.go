package app

import (
	"fmt"
	"time"

	c "github.com/agustin-sarasua/rs-common"
	m "github.com/agustin-sarasua/rs-model"
)

func validateProperty(p *m.Property) []error {
	var errs []error

	//errs = c.ValidateExistInMap(m.PropertyTypes, p.Type, "Type is incorrect", errs)
	errs = c.ValidateExistInMap(m.Orientation, p.Orientation, "Orientation is incorrect", errs)
	errs = c.ValidateRangeCondition(0, 10, p.State, fmt.Sprintf("State should be between %v and %v", 0, 10), errs)
	errs = c.ValidateCondition(func() bool { return p.Address != nil }, "Address can not be empty", errs)
	errs = c.ValidatePositiveNumber(p.Bedrooms, "Bedrooms should be positive", errs)
	errs = c.ValidatePositiveNumber(p.Bathrooms, "Bathrooms should be positive", errs)
	errs = c.ValidatePositiveNumber(p.Kitchens, "Kitchens should be positive", errs)
	errs = c.ValidateRangeCondition(1800, time.Now().Year(), p.ConstructionYear, "ConstructionYear should be between 1800 and Now", errs)
	errs = c.ValidatePositiveNumber(p.Elevators, "Elevators should be positive", errs)
	errs = c.ValidatePositiveNumber(p.CourtyardSize, "CourtyardSize should be positive", errs)
	errs = c.ValidatePositiveNumber(p.Size, "Size should be positive", errs)
	errs = c.ValidatePositiveNumber(p.Floors, "Floors should be positive", errs)
	errs = c.ValidatePositiveNumber(p.GarageSize, "GarageSize should be positive", errs)
	errs = c.ValidatePositiveNumber(p.ConstructionSize, "ConstructionSize should be positive", errs)
	errs = c.ValidatePositiveNumber(p.Showers, "Showers should be positive", errs)
	return errs
}
