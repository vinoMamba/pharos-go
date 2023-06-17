// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: validation_codes.sql

package queries

import (
	"context"
)

const countValidationCodes = `-- name: CountValidationCodes :one
SELECT count(*) FROM validation_codes
`

func (q *Queries) CountValidationCodes(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countValidationCodes)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createValidaitonCode = `-- name: CreateValidaitonCode :one
INSERT INTO validation_codes(
  email, code
) VALUES (
  $1,$2
)

RETURNING id, code, email, use_at, created_at, updated_at
`

type CreateValidaitonCodeParams struct {
	Email string
	Code  string
}

func (q *Queries) CreateValidaitonCode(ctx context.Context, arg CreateValidaitonCodeParams) (ValidationCode, error) {
	row := q.db.QueryRowContext(ctx, createValidaitonCode, arg.Email, arg.Code)
	var i ValidationCode
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Email,
		&i.UseAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAllValidationCodes = `-- name: DeleteAllValidationCodes :exec
DELETE FROM validation_codes
`

func (q *Queries) DeleteAllValidationCodes(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllValidationCodes)
	return err
}

const findValidationCode = `-- name: FindValidationCode :one
SELECT id, code, email, use_at, created_at, updated_at FROM validation_codes 
WHERE email = $1 
AND code = $2 
AND use_at IS NULL 
ORDER BY created_at DESC 
LIMIT 1
`

type FindValidationCodeParams struct {
	Email string
	Code  string
}

func (q *Queries) FindValidationCode(ctx context.Context, arg FindValidationCodeParams) (ValidationCode, error) {
	row := q.db.QueryRowContext(ctx, findValidationCode, arg.Email, arg.Code)
	var i ValidationCode
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Email,
		&i.UseAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
