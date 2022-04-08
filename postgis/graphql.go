package postgis

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/paulmach/orb/geojson"
)

func MarshalGeometry4326(g Geometry4326) graphql.Marshaler {
	ng := geojson.NewGeometry(g.Coordinates)
	bytes, err := ng.MarshalJSON()

	if err != nil {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, string(bytes))
	})
}

func MarshalNullGeometry4326(g NullGeometry4326) graphql.Marshaler {
	if g.Valid {
		return MarshalGeometry4326(g.Geometry4326)
	}
	
	return graphql.Null
}

func UnmarshalGeometry4326(v interface{}) (Geometry4326, error) {
	g := Geometry4326{}
	err := g.Scan(v)

	if err != nil {
		return Geometry4326{}, err
	}

	return g, nil
}

func UnmarshalNullGeometry4326(v interface{}) (NullGeometry4326, error) {
	g := NullGeometry4326{}
	err := g.Scan(v)

	if err != nil {
		return NullGeometry4326{}, err
	}

	return g, nil
}
