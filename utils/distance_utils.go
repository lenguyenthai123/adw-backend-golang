package utils

import "math"

func HaversineDistance(lat1, lng1, lat2, lng2 float64) float64 {
	r := 6371.0 // Radius of the Earth in kilometers

	// Convert latitude and longitude from degrees to radians
	lat1Rad := lat1 * (math.Pi / 180.0)
	lon1Rad := lng1 * (math.Pi / 180.0)
	lat2Rad := lat2 * (math.Pi / 180.0)
	lon2Rad := lng2 * (math.Pi / 180.0)

	// Calculate differences
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Apply Haversine formula
	a := math.Pow(math.Sin(deltaLat/2), 2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate the distance
	distance := r * c
	return distance
}
