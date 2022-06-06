package reponse

type UpdateResponse struct {
	StatusText string `json:"statusText"`
	Status     int    `json:"status"`
	Data       any    `json:"data"`
}
