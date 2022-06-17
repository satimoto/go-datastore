package geom

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/ewkb"
	"github.com/paulmach/orb/geojson"
)

type Geometry4326 geojson.Geometry

func NewGeometry4326(coordinates orb.Geometry) Geometry4326 {
	return Geometry4326{
		Coordinates: coordinates,
		Type: coordinates.GeoJSONType(),
	}
}

func (g *Geometry4326) Scan(i interface{}) error {
	switch t := i.(type) {
	case Geometry4326:
		ig := i.(Geometry4326)
		g.Coordinates = ig.Coordinates
		g.Geometries = ig.Geometries
		g.Type = ig.Type
	case []byte:
		s := ewkb.Scanner(nil)
		n, err := hex.Decode(t, t)

		if err != nil {
			return err
		}
	
		if err := s.Scan(t[:n]); err != nil {
			return err
		}
	
		if s.Geometry == nil {
			return fmt.Errorf("Geometry is nil")
		}
	
		g.Coordinates = s.Geometry
		g.Type = s.Geometry.GeoJSONType()
	default:
		return fmt.Errorf("Connot convert from %s", t)
	}

	return nil
}

func (g Geometry4326) Value() (driver.Value, error) {
	return ewkb.Value(g.Coordinates, 4326).Value()
}

type NullGeometry4326 struct {
	Geometry4326 Geometry4326
	Valid        bool
}

func (n *NullGeometry4326) Scan(value interface{}) error {
	n.Geometry4326, n.Valid = Geometry4326{}, false

	if value == nil {
		return nil
	}

	err := n.Geometry4326.Scan(value)

	if err != nil {
		return err
	}

	n.Valid = true
	return nil
}

func (n NullGeometry4326) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.Geometry4326.Value()
}
