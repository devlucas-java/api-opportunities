package handler

import "fmt"

func errParamRequerided(field string, typ string) error {
	return fmt.Errorf("Erro required field: %s, ( type: %s )", field, typ)
}

type CreatedOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Remote   *bool  `json:"remote"`
	Salary   int    `json:"salary"`
	Location string `json:"location"`
	Link     string `json:"link"`
}

func (r *CreatedOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Salary <= 0 && r.Remote == nil {
		return fmt.Errorf("Error required fields: role, company, location, link, salary and remote")
	}
	if r.Role == "" {
		return errParamRequerided("role", "string")
	}
	if r.Company == "" {
		return errParamRequerided("company", "string")
	}
	if r.Location == "" {
		return errParamRequerided("location", "string")
	}
	if r.Link == "" {
		return errParamRequerided("link", "string")
	}
	if r.Salary <= 0 {
		return errParamRequerided("salary", "int64")
	}
	if r.Remote == nil {
		return errParamRequerided("remote", "bool")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Remote   *bool  `json:"remote"`
	Salary   int    `json:"salary"`
	Location string `json:"location"`
	Link     string `json:"link"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Salary > 0 || r.Remote != nil {
		return nil
	}
	return fmt.Errorf("At least one valid field must be provieded")

}
