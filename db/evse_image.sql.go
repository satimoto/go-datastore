// Code generated by sqlc. DO NOT EDIT.
// source: evse_image.sql

package db

import (
	"context"
)

const deleteEvseImages = `-- name: DeleteEvseImages :exec
DELETE FROM images im
  USING evse_images ei
  WHERE ei.image_id = im.id AND ei.evse_id = $1
`

func (q *Queries) DeleteEvseImages(ctx context.Context, evseID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEvseImages, evseID)
	return err
}

const listEvseImages = `-- name: ListEvseImages :many
SELECT im.id, im.url, im.thumbnail, im.category, im.type, im.width, im.height FROM images im
  INNER JOIN evse_images ei ON ei.image_id = im.id
  WHERE ei.evse_id = $1
  ORDER BY im.id
`

func (q *Queries) ListEvseImages(ctx context.Context, evseID int64) ([]Image, error) {
	rows, err := q.db.QueryContext(ctx, listEvseImages, evseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Image
	for rows.Next() {
		var i Image
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.Thumbnail,
			&i.Category,
			&i.Type,
			&i.Width,
			&i.Height,
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

const setEvseImage = `-- name: SetEvseImage :exec
INSERT INTO evse_images (evse_id, image_id)
  VALUES ($1, $2)
`

type SetEvseImageParams struct {
	EvseID  int64 `db:"evse_id" json:"evseID"`
	ImageID int64 `db:"image_id" json:"imageID"`
}

func (q *Queries) SetEvseImage(ctx context.Context, arg SetEvseImageParams) error {
	_, err := q.db.ExecContext(ctx, setEvseImage, arg.EvseID, arg.ImageID)
	return err
}
