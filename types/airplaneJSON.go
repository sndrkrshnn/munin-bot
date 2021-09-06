package types

type Airline struct {
	Name string `json:"name"`
}

type FlightInfo struct {
	IATACode string `json:"iata"`
}

type Destination struct {
	Airport  string `json:"airport"`
	IATACode string `json:"iata"`
}
type Flight struct {
	Airline    Airline     `json:"airline"`
	FlightInfo FlightInfo  `json:"flight"`
	Departure  Destination `json:"departure"`
	Arrival    Destination `json:"arrival"`
}

type Response struct {
	Flights []Flight `json:"results"`
}
