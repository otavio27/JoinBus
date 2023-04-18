package structs

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type routes struct {
	RouteID       string `json:"route_id"`
	RouteLongName string `json:"route_long_name"`
}

type timedata struct {
	TripID        string `json:"trip_id"`
	DepartureTime string `json:"departure_time"`
}

type servicedata struct {
	ServiceID string       `json:"service_id"`
	Type      string       `json:"service_name"`
	TimeData  [][]timedata `json:"time_data"`
}

type stopdata struct {
	StopID       string        `json:"stop_id"`
	StopName     string        `json:"stop_name"`
	StopSequence int           `json:"stop_sequence"`
	ServiceData  []servicedata `json:"service_data"`
}

type Stations struct {
	StationName   string   `json:"station_name"`
	Routes        []routes `json:"routes"`
	RouteID       string   `json:"route_id"`
	RouteLongName string   `json:"route_long_name"`
}

type Location struct {
	StopID       string `json:"stop_id"`
	StopName     string `json:"stop_name"`
	StopLat      string `json:"stop_lat"`
	StopLon      string `json:"stop_lon"`
	LocationType string `json:"location_type"`
	Distance     int    `json:"distance"`
}

type StopList struct {
	ShapeID      string `json:"shape_id"`
	StopSequence int    `json:"stop_sequence"`
	TripHeadsign string `json:"trip_headsign"`
	StopTime     string `json:"stop_time"`
	TimeLeft     int    `json:"time_left"`
	NextTrip     int    `json:"next_trip"`
	Trips        []struct {
		Eta            string  `json:"eta"`
		TimeLeft       int     `json:"time_left"`
		StartTimeDiff  int     `json:"start_time_diff"`
		TripID         string  `json:"trip_id,omitempty"`
		TripStatus     string  `json:"trip_status,omitempty"`
		ReportLat      float64 `json:"report_lat,omitempty"`
		ReportLon      float64 `json:"report_lon,omitempty"`
		StopName       string  `json:"stop_name,omitempty"`
		StopDistance   float64 `json:"stop_distance,omitempty"`
		StopOrder      int     `json:"stop_order,omitempty"`
		VehicleID      string  `json:"vehicle_id,omitempty"`
		DistanceDiff   int     `json:"distance_diff,omitempty"`
		ReportTimeDiff int     `json:"report_time_diff,omitempty"`
	} `json:"trips"`
}

type LineStruct struct {
	Routes        []routes   `json:"routes"`
	StopData      []stopdata `json:"stop_data"`
	RouteID       string     `json:"route_id"`
	Direction     string     `json:"direction"`
	ServiceName   string     `json:"service_name"`
	DepartureTime string     `json:"departure_time"`
}
