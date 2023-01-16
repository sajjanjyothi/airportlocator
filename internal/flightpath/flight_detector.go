// flightpath to detect the entry and exit airport codes from a list of itinerary
package flightpath

// FlightPath interface to detect the
type FlightPath interface {
	GetSourceDestinationAirports(flightData [][]string) ([]string, error)
}

type flights struct {
}

// NewAirportDetector factory function to create airport detector instance
func NewAirportDetector() FlightPath {
	return &flights{}
}

// GetSourceDestinationAirports Get source and destination airport from the path
func (f *flights) GetSourceDestinationAirports(flightData [][]string) ([]string, error) {
	var start, end string

	for idx, path := range flightData {
		if len(start) == 0 {
			start = path[0]
		}
		if len(end) == 0 {
			end = path[1]
		}

		for iidx, path := range flightData { //go through the remaining array once again
			if idx == iidx {
				continue //we are on this path now, no need to calculate path in this
			}
			//if any end match with start, new starting point will be starting point of that path
			if path[1] == start {
				start = path[0]
			}
			//if any start airport match with end, new end point will be end point of that path
			if path[0] == end {
				end = path[1]
			}
		}
	}
	return []string{start, end}, nil
}
