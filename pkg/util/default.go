package util

import "database/sql"

func DefaultBool(i interface{}, fallback bool) bool {
	switch t := i.(type) {
	case sql.NullBool:
		if t.Valid {
			return t.Bool
		}
	case bool:
		return t
	case *bool:
		return *t
	}

	return fallback
}

func DefaultString(i interface{}, fallback string) string {
	switch t := i.(type) {
	case sql.NullString:
		if t.Valid {
			return t.String
		}
	case string:
		return t
	case *string:
		return *t
	}

	return fallback
}