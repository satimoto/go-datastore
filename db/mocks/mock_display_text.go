package mocks

import "github.com/satimoto/go-datastore/db"

type DisplayTextsPayload struct {
	DisplayTexts []db.DisplayText
	Error        error
}
