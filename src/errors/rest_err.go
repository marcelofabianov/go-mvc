package rest_err

type RestErr struct {
	Message string   `json:"message"`
	Status  int      `json:"status"`
	Err     string   `json:"error"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *RestErr) Error() string {
	return e.Message
}

func NewRestErr(message string, status int, err string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Err:     err,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  400,
		Err:     "bad_request",
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  401,
		Err:     "unauthorized",
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  403,
		Err:     "forbidden",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  404,
		Err:     "not_found",
	}
}

func NewMethodNotAllowedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  405,
		Err:     "method_not_allowed",
	}
}

func NewUnprocessableEntityError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Status:  422,
		Err:     "unprocessable_entity",
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  500,
		Err:     "internal_server_error",
	}
}
