package stats

type ResolvedTicketsStats struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type OpenTicketsStats struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type OperatorInfo struct {
	ID    *string `json:"id"`
	Name  *string  `json:"name"`
	Email *string `json:"email"`
}

type OperatorTicketsStats struct { 
	Operator OperatorInfo `json:"operator"`
	Count    int          `json:"count"`
}



