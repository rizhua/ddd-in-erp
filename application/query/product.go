package query

import "errors"

type Product struct {
	ID int64 `json:"id" form:"id"`
}

func (t Product) Validate() error {
	if t.ID <= 0 {
		return errors.New("id is required")
	}
	return nil
}
