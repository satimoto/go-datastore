// Code generated by sqlc. DO NOT EDIT.
// source: token_authorization_connector.sql

package db

import (
	"context"
)

const listTokenAuthorizationConnectors = `-- name: ListTokenAuthorizationConnectors :many
SELECT c.id, c.evse_id, c.uid, c.identifier, c.standard, c.format, c.power_type, c.voltage, c.amperage, c.wattage, c.tariff_id, c.terms_and_conditions, c.last_updated FROM connectors c
  INNER JOIN token_authorization_connectors tac ON tac.connector_id = c.id
  WHERE tac.token_authorization_id = $1
  ORDER BY c.id
`

func (q *Queries) ListTokenAuthorizationConnectors(ctx context.Context, tokenAuthorizationID int64) ([]Connector, error) {
	rows, err := q.db.QueryContext(ctx, listTokenAuthorizationConnectors, tokenAuthorizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Connector
	for rows.Next() {
		var i Connector
		if err := rows.Scan(
			&i.ID,
			&i.EvseID,
			&i.Uid,
			&i.Identifier,
			&i.Standard,
			&i.Format,
			&i.PowerType,
			&i.Voltage,
			&i.Amperage,
			&i.Wattage,
			&i.TariffID,
			&i.TermsAndConditions,
			&i.LastUpdated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setTokenAuthorizationConnector = `-- name: SetTokenAuthorizationConnector :exec
INSERT INTO token_authorization_connectors (token_authorization_id, connector_id)
  VALUES ($1, $2)
`

type SetTokenAuthorizationConnectorParams struct {
	TokenAuthorizationID int64 `db:"token_authorization_id" json:"tokenAuthorizationID"`
	ConnectorID          int64 `db:"connector_id" json:"connectorID"`
}

func (q *Queries) SetTokenAuthorizationConnector(ctx context.Context, arg SetTokenAuthorizationConnectorParams) error {
	_, err := q.db.ExecContext(ctx, setTokenAuthorizationConnector, arg.TokenAuthorizationID, arg.ConnectorID)
	return err
}
