package flight_request

type CreateFlightRequest struct {
	Name          string `json:"name" binding:"required"`
	From          string `json:"from" binding:"required"`
	To            string `json:"to" binding:"required"`
	DepartDate    string `json:"departDate" binding:"required"`
	DepartTime    string `json:"departTime" binding:"required"`
	Status        string `json:"status"`
	AvailableSlot int32  `json:"slot"`
}

type SearchFlightRequest struct {
	Name     string `json:"name"`
	From     string `json:"from"`
	To       string `json:"to"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}
