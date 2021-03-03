package response

import "pep/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
