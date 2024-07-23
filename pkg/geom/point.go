package geom

import (
	"strconv"

	"github.com/paulmach/orb"
)

func NewPoint(longitude, latitude string) (orb.Point, error) {
	var latitudeFloat, longitudeFloat float64
	var err error

	if latitudeFloat, err = strconv.ParseFloat(latitude, 64); err != nil {
		return orb.Point{}, err
	}

	if longitudeFloat, err = strconv.ParseFloat(longitude, 64); err != nil {
		return orb.Point{}, err
	}

	return orb.Point{longitudeFloat, latitudeFloat}, nil
}
