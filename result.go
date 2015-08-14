package toolkit

import (
	"time"
)

type ResultStatus string

const (
	Status_OK  ResultStatus = "OK"
	Status_NOK ResultStatus = "NOK"
)

type Result struct {
	Status   ResultStatus
	Message  string
	Duration time.Duration
	Data     interface{}
}

func (a *Result) Run(f func(data interface{}) (interface{}, error), parm interface{}) *Result {
	t0 := time.Now()
	a.Status = Status_OK
	a.Message = ""
	if f == nil {
		a.Data = nil
	} else {
		r, e := f(parm)
		if e != nil {
			a.Status = Status_NOK
			a.Message = e.Error()
			a.Data = nil
		} else {
			a.Data = r
		}
	}
	a.Duration = time.Since(t0)
	return a
}