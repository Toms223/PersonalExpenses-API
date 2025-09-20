package Errors

type APIError struct {
	Message string
	Code    int
}

type ProblemJson struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

var httpReasonPhrases = map[int]string{
	400: "Bad Request",
	404: "Not Found",
	409: "Conflict",
	500: "Internal Server Error",
}

func (err APIError) ToProblemJSON() ProblemJson {
	title := httpReasonPhrases[err.Code]
	return ProblemJson{
		Title:  title,
		Status: err.Code,
		Detail: err.Message,
	}
}
