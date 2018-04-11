package gcore

import "fmt"

type GCoreError struct {
	Code    int                  `json:"-"`
	Errors  *map[string][]string `json:"errors"`
	Message *string              `json:"message"`
}

func (g *GCoreError) Error() string {
	errorString := fmt.Sprintf("response code: %d ", g.Code)
	if g.Errors != nil {
		errorString += fmt.Sprintf("error(s): %+v ", *g.Errors)
	}

	if g.Message != nil {
		errorString += fmt.Sprintf("message: %s", *g.Message)
	}

	return errorString
}
