package errs

// EntError describes all server-known errors
type EntError struct {
	Status int               `json:"-"`
	Msg    string            `json:"message,omitempty"`
	Fields map[string]string `json:"fields,omitempty"`
	Advice string            `json:"advice,omitempty"`
	Err    error             `json:"-"`
}

func newEntError(status int, msg string, advice string) EntError {
	return EntError{Status: status, Msg: msg, Advice: advice}
}

func (e EntError) AddError(err error) EntError {
	e.Err = err
	return e
}

func (e EntError) AddFields(fields map[string]string) EntError {
	e.Fields = fields
	return e
}

// Error implements the Error type
func (e EntError) Error() string {
	return e.Msg
}

func (e EntError) GetInfo() *AbstractError {

	return &AbstractError{
		Status: e.Status,
		Msg:    e.Msg,
		Fields: e.Fields,
		Advice: e.Advice,
		Err:    e.Err,
	}
}
