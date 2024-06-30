package response 

type Response struct {
	StatusCode int
	Data interface{}
	Error string
	Message string
}

func CreateResponse(statusCode int, data interface{}, error error, message string) Response {
	if error != nil {
		var res Response
		res.StatusCode = statusCode
		res.Error = error.Error()
		if message == "" {
			res.Message = createErrorMessage(error)
		} else {
			res.Message = message
		}
		return res
	}

	return Response{
		StatusCode: statusCode,
		Data: data,
		Error: "",
		Message: message,
	}
}

func createErrorMessage(err error) string {
	return err.Error()
}

