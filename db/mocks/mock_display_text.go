package mocks

import "github.com/satimoto/go-datastore/db"

type DisplayTextsResponse struct {
	DisplayTexts []db.DisplayText
	Error        error
}
