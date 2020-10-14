package html

type FlightEntityHTML struct {
	// for numbering cases
	Index      int                        `json:"index"`
	PNR        string                     `json:"pnr"`
	BookID     string                     `json:"book_id"`
	FlightType string                     `json:"book_id"`
	Airline    FlightAirlineEntityHTML    `json:"airline"`
	Schedules  []FlightScheduleEntityHTML `json:"schedules"`
	Passengers []PassengerEntityHTML      `json:"passengers"`
}

type FlightAirlineEntityHTML struct {
	Logo          string `json:"logo"`
	Class         string `json:"class"`
	FlightNumber  string `json:"flight_number"`
	DepartureDate string `json:"departure_date"`
	DepartureTime string `json:"departure_time"`
	ArrivedTime   string `json:"arrived_time"`
	DurationTime  string `json:"duration_time"`
	FlightDirect  string `json:"flight_direct"`
}

type FlightScheduleEntityHTML struct {
	DepartureTime            string `json:"departure_time"`
	DepartureAirportIATA     string `json:"departure_airport_iata"`
	DepartureAirportName     string `json:"departure_airport_name"`
	DepartureAirportCity     string `json:"departure_airport_city"`
	DepartureAirportTerminal string `json:"departure_airport_terminal"`
	ArrivedTime              string `json:"arrived_time"`
	ArrivedAirportIATA       string `json:"arrived_airport_iata"`
	ArrivedAirportName       string `json:"arrived_airport_name"`
	ArrivedAirportCity       string `json:"arrived_airport_city"`
	ArrivedAirportTerminal   string `json:"arrived_airport_terminal"`
	TransitDurationTime      string `json:"transit_duration_time"`
	TransitCityName          string `json:"transit_city_name"`
	IsOvernight              bool   `json:"is_overnight"`
}

type PassengerEntityHTML struct {
	Name    string
	Type    string
	Baggage int
}

type ReceiptEntityHTML struct {
	BookID    string
	CreatedAt string
	Customer  struct {
		Name  string
		Email string
		Phone string
	}
	Payment struct {
		Method string
		Status string
	}
	Passengers   []PassengerEntityHTML
	Items        []ReceiptItemEntity
	Total        int
	AdminFee     int
	TotalPayment int
}

type ReceiptItemEntity struct {
	Type       string
	Desc       string
	Qty        int
	Price      int
	TotalPrice int
}
