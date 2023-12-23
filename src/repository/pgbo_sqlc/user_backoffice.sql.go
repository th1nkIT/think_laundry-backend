// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user_backoffice.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const deleteUserBackoffice = `-- name: DeleteUserBackoffice :exec
UPDATE user_backoffice
SET
    deleted_at = (now() at time zone 'UTC')::TIMESTAMP,
    deleted_by = $1
WHERE
    guid = $2
    AND deleted_at IS NULL
`

type DeleteUserBackofficeParams struct {
	DeletedBy sql.NullString `json:"deleted_by"`
	Guid      string         `json:"guid"`
}

func (q *Queries) DeleteUserBackoffice(ctx context.Context, arg DeleteUserBackofficeParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserBackoffice, arg.DeletedBy, arg.Guid)
	return err
}

const getCountListUserBackoffice = `-- name: GetCountListUserBackoffice :one
SELECT count(ub.id) FROM user_backoffice ub
WHERE
    (CASE WHEN $1::bool THEN LOWER(ub.name) LIKE LOWER($2) ELSE TRUE END)
    AND(CASE WHEN $3::bool THEN LOWER(ub.phone) LIKE LOWER($4) ELSE TRUE END)
    AND(CASE WHEN $5::bool THEN LOWER(ub.email) LIKE LOWER($6) ELSE TRUE END)
    AND (CASE WHEN $7::bool THEN ub.role_id = $8 ELSE TRUE END)
    AND (CASE WHEN $9::bool THEN ub.is_active = $10 ELSE TRUE END)
    AND ub.deleted_at IS NULL
`

type GetCountListUserBackofficeParams struct {
	SetName     bool         `json:"set_name"`
	Name        string       `json:"name"`
	SetPhone    bool         `json:"set_phone"`
	Phone       string       `json:"phone"`
	SetEmail    bool         `json:"set_email"`
	Email       string       `json:"email"`
	SetRoleID   bool         `json:"set_role_id"`
	RoleID      int32        `json:"role_id"`
	SetIsActive bool         `json:"set_is_active"`
	IsActive    sql.NullBool `json:"is_active"`
}

func (q *Queries) GetCountListUserBackoffice(ctx context.Context, arg GetCountListUserBackofficeParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCountListUserBackoffice,
		arg.SetName,
		arg.Name,
		arg.SetPhone,
		arg.Phone,
		arg.SetEmail,
		arg.Email,
		arg.SetRoleID,
		arg.RoleID,
		arg.SetIsActive,
		arg.IsActive,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUserBackoffice = `-- name: GetUserBackoffice :one
SELECT
       ub.id, ub.guid, ub.name, ub.profile_picture_image_url, ub.phone, ub.email, ub.role_id, ub.password, ub.salt, ub.is_active, ub.created_at, ub.created_by, ub.updated_at, ub.updated_by, ub.deleted_at, ub.deleted_by, ub.last_login,
       ubr.name as role_name,
       ubr.access as role_access,
       ubr.is_all_access as is_all_access
FROM user_backoffice ub
    JOIN user_backoffice_role ubr ON ubr.id = ub.role_id
WHERE
    ub.guid = $1
    AND ub.deleted_at IS NULL
`

type GetUserBackofficeRow struct {
	ID                     int64          `json:"id"`
	Guid                   string         `json:"guid"`
	Name                   sql.NullString `json:"name"`
	ProfilePictureImageUrl sql.NullString `json:"profile_picture_image_url"`
	Phone                  string         `json:"phone"`
	Email                  string         `json:"email"`
	RoleID                 int32          `json:"role_id"`
	Password               string         `json:"password"`
	Salt                   string         `json:"salt"`
	IsActive               sql.NullBool   `json:"is_active"`
	CreatedAt              time.Time      `json:"created_at"`
	CreatedBy              string         `json:"created_by"`
	UpdatedAt              sql.NullTime   `json:"updated_at"`
	UpdatedBy              sql.NullString `json:"updated_by"`
	DeletedAt              sql.NullTime   `json:"deleted_at"`
	DeletedBy              sql.NullString `json:"deleted_by"`
	LastLogin              sql.NullTime   `json:"last_login"`
	RoleName               string         `json:"role_name"`
	RoleAccess             sql.NullString `json:"role_access"`
	IsAllAccess            sql.NullBool   `json:"is_all_access"`
}

func (q *Queries) GetUserBackoffice(ctx context.Context, guid string) (GetUserBackofficeRow, error) {
	row := q.db.QueryRowContext(ctx, getUserBackoffice, guid)
	var i GetUserBackofficeRow
	err := row.Scan(
		&i.ID,
		&i.Guid,
		&i.Name,
		&i.ProfilePictureImageUrl,
		&i.Phone,
		&i.Email,
		&i.RoleID,
		&i.Password,
		&i.Salt,
		&i.IsActive,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.LastLogin,
		&i.RoleName,
		&i.RoleAccess,
		&i.IsAllAccess,
	)
	return i, err
}

const getUserBackofficeByEmail = `-- name: GetUserBackofficeByEmail :one
select
    ub.id, ub.guid, ub.name, ub.profile_picture_image_url, ub.phone, ub.email, ub.role_id, ub.password, ub.salt, ub.is_active, ub.created_at, ub.created_by, ub.updated_at, ub.updated_by, ub.deleted_at, ub.deleted_by, ub.last_login,
    ubr.name as role_name,
    ubr.access as role_access,
    ubr.is_all_access as is_all_access
FROM user_backoffice ub
    JOIN user_backoffice_role ubr ON ubr.id = ub.role_id
WHERE
    ub.email = $1
`

type GetUserBackofficeByEmailRow struct {
	ID                     int64          `json:"id"`
	Guid                   string         `json:"guid"`
	Name                   sql.NullString `json:"name"`
	ProfilePictureImageUrl sql.NullString `json:"profile_picture_image_url"`
	Phone                  string         `json:"phone"`
	Email                  string         `json:"email"`
	RoleID                 int32          `json:"role_id"`
	Password               string         `json:"password"`
	Salt                   string         `json:"salt"`
	IsActive               sql.NullBool   `json:"is_active"`
	CreatedAt              time.Time      `json:"created_at"`
	CreatedBy              string         `json:"created_by"`
	UpdatedAt              sql.NullTime   `json:"updated_at"`
	UpdatedBy              sql.NullString `json:"updated_by"`
	DeletedAt              sql.NullTime   `json:"deleted_at"`
	DeletedBy              sql.NullString `json:"deleted_by"`
	LastLogin              sql.NullTime   `json:"last_login"`
	RoleName               string         `json:"role_name"`
	RoleAccess             sql.NullString `json:"role_access"`
	IsAllAccess            sql.NullBool   `json:"is_all_access"`
}

func (q *Queries) GetUserBackofficeByEmail(ctx context.Context, email string) (GetUserBackofficeByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserBackofficeByEmail, email)
	var i GetUserBackofficeByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Guid,
		&i.Name,
		&i.ProfilePictureImageUrl,
		&i.Phone,
		&i.Email,
		&i.RoleID,
		&i.Password,
		&i.Salt,
		&i.IsActive,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.LastLogin,
		&i.RoleName,
		&i.RoleAccess,
		&i.IsAllAccess,
	)
	return i, err
}

const insertUserBackoffice = `-- name: InsertUserBackoffice :one
INSERT INTO user_backoffice
    (guid, name, profile_picture_image_url, phone, email, role_id, password, salt, is_active, created_at, created_by)
    VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9,  (now() at time zone 'UTC')::TIMESTAMP, $10)
RETURNING user_backoffice.id, user_backoffice.guid, user_backoffice.name, user_backoffice.profile_picture_image_url, user_backoffice.phone, user_backoffice.email, user_backoffice.role_id, user_backoffice.password, user_backoffice.salt, user_backoffice.is_active, user_backoffice.created_at, user_backoffice.created_by, user_backoffice.updated_at, user_backoffice.updated_by, user_backoffice.deleted_at, user_backoffice.deleted_by, user_backoffice.last_login
`

type InsertUserBackofficeParams struct {
	Guid                   string         `json:"guid"`
	Name                   sql.NullString `json:"name"`
	ProfilePictureImageUrl sql.NullString `json:"profile_picture_image_url"`
	Phone                  string         `json:"phone"`
	Email                  string         `json:"email"`
	RoleID                 int32          `json:"role_id"`
	Password               string         `json:"password"`
	Salt                   string         `json:"salt"`
	IsActive               sql.NullBool   `json:"is_active"`
	CreatedBy              string         `json:"created_by"`
}

func (q *Queries) InsertUserBackoffice(ctx context.Context, arg InsertUserBackofficeParams) (UserBackoffice, error) {
	row := q.db.QueryRowContext(ctx, insertUserBackoffice,
		arg.Guid,
		arg.Name,
		arg.ProfilePictureImageUrl,
		arg.Phone,
		arg.Email,
		arg.RoleID,
		arg.Password,
		arg.Salt,
		arg.IsActive,
		arg.CreatedBy,
	)
	var i UserBackoffice
	err := row.Scan(
		&i.ID,
		&i.Guid,
		&i.Name,
		&i.ProfilePictureImageUrl,
		&i.Phone,
		&i.Email,
		&i.RoleID,
		&i.Password,
		&i.Salt,
		&i.IsActive,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.LastLogin,
	)
	return i, err
}

const listUserBackoffice = `-- name: ListUserBackoffice :many
SELECT
    ub.id, ub.guid, ub.name, ub.profile_picture_image_url, ub.phone, ub.email, ub.role_id, ub.password, ub.salt, ub.is_active, ub.created_at, ub.created_by, ub.updated_at, ub.updated_by, ub.deleted_at, ub.deleted_by, ub.last_login,
    ubr.name as role_name,
    ubr.access as role_access,
    ubr.is_all_access as is_all_access
FROM
    user_backoffice ub
    JOIN user_backoffice_role ubr ON ubr.id = ub.role_id
WHERE
    (CASE WHEN $1::bool THEN LOWER(ub.name) LIKE LOWER($2) ELSE TRUE END)
    AND(CASE WHEN $3::bool THEN LOWER(ub.phone) LIKE LOWER($4) ELSE TRUE END)
    AND(CASE WHEN $5::bool THEN LOWER(ub.email) LIKE LOWER($6) ELSE TRUE END)
    AND (CASE WHEN $7::bool THEN ub.role_id = $8 ELSE TRUE END)
    AND (CASE WHEN $9::bool THEN ub.is_active = $10 ELSE TRUE END)
    AND ub.deleted_at IS NULL
ORDER BY (CASE WHEN $11 = 'id ASC' THEN ub.guid END) ASC,
         (CASE WHEN $11 = 'id DESC' THEN ub.guid  END) DESC,
         (CASE WHEN $11 = 'name ASC' THEN ub.name END) ASC,
         (CASE WHEN $11 = 'name DESC' THEN ub.name  END) DESC,
         (CASE WHEN $11 = 'role_id ASC' THEN ub.role_id END) ASC,
         (CASE WHEN $11 = 'role_id DESC' THEN ub.role_id  END) DESC,
         (CASE WHEN $11 = 'created_at ASC' THEN ub.created_at END) ASC,
         (CASE WHEN $11 = 'created_at DESC' THEN ub.created_at  END) DESC,
         ub.created_at DESC
LIMIT $13
OFFSET $12
`

type ListUserBackofficeParams struct {
	SetName     bool         `json:"set_name"`
	Name        string       `json:"name"`
	SetPhone    bool         `json:"set_phone"`
	Phone       string       `json:"phone"`
	SetEmail    bool         `json:"set_email"`
	Email       string       `json:"email"`
	SetRoleID   bool         `json:"set_role_id"`
	RoleID      int32        `json:"role_id"`
	SetIsActive bool         `json:"set_is_active"`
	IsActive    sql.NullBool `json:"is_active"`
	OrderParam  interface{}  `json:"order_param"`
	OffsetPage  int32        `json:"offset_page"`
	LimitData   int32        `json:"limit_data"`
}

type ListUserBackofficeRow struct {
	ID                     int64          `json:"id"`
	Guid                   string         `json:"guid"`
	Name                   sql.NullString `json:"name"`
	ProfilePictureImageUrl sql.NullString `json:"profile_picture_image_url"`
	Phone                  string         `json:"phone"`
	Email                  string         `json:"email"`
	RoleID                 int32          `json:"role_id"`
	Password               string         `json:"password"`
	Salt                   string         `json:"salt"`
	IsActive               sql.NullBool   `json:"is_active"`
	CreatedAt              time.Time      `json:"created_at"`
	CreatedBy              string         `json:"created_by"`
	UpdatedAt              sql.NullTime   `json:"updated_at"`
	UpdatedBy              sql.NullString `json:"updated_by"`
	DeletedAt              sql.NullTime   `json:"deleted_at"`
	DeletedBy              sql.NullString `json:"deleted_by"`
	LastLogin              sql.NullTime   `json:"last_login"`
	RoleName               string         `json:"role_name"`
	RoleAccess             sql.NullString `json:"role_access"`
	IsAllAccess            sql.NullBool   `json:"is_all_access"`
}

func (q *Queries) ListUserBackoffice(ctx context.Context, arg ListUserBackofficeParams) ([]ListUserBackofficeRow, error) {
	rows, err := q.db.QueryContext(ctx, listUserBackoffice,
		arg.SetName,
		arg.Name,
		arg.SetPhone,
		arg.Phone,
		arg.SetEmail,
		arg.Email,
		arg.SetRoleID,
		arg.RoleID,
		arg.SetIsActive,
		arg.IsActive,
		arg.OrderParam,
		arg.OffsetPage,
		arg.LimitData,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUserBackofficeRow
	for rows.Next() {
		var i ListUserBackofficeRow
		if err := rows.Scan(
			&i.ID,
			&i.Guid,
			&i.Name,
			&i.ProfilePictureImageUrl,
			&i.Phone,
			&i.Email,
			&i.RoleID,
			&i.Password,
			&i.Salt,
			&i.IsActive,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.DeletedAt,
			&i.DeletedBy,
			&i.LastLogin,
			&i.RoleName,
			&i.RoleAccess,
			&i.IsAllAccess,
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

const recordUserBackofficeLastLogin = `-- name: RecordUserBackofficeLastLogin :exec
UPDATE user_backoffice
SET
    last_login = (now() at time zone 'UTC')::TIMESTAMP
WHERE
    guid = $1
    AND deleted_at IS NULL
`

func (q *Queries) RecordUserBackofficeLastLogin(ctx context.Context, guid string) error {
	_, err := q.db.ExecContext(ctx, recordUserBackofficeLastLogin, guid)
	return err
}

const updateUserBackoffice = `-- name: UpdateUserBackoffice :one
UPDATE user_backoffice
SET
    name = $1,
    phone = $2,
    profile_picture_image_url = $3,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = $4
WHERE
    guid = $5
    AND deleted_at IS NULL
RETURNING user_backoffice.id, user_backoffice.guid, user_backoffice.name, user_backoffice.profile_picture_image_url, user_backoffice.phone, user_backoffice.email, user_backoffice.role_id, user_backoffice.password, user_backoffice.salt, user_backoffice.is_active, user_backoffice.created_at, user_backoffice.created_by, user_backoffice.updated_at, user_backoffice.updated_by, user_backoffice.deleted_at, user_backoffice.deleted_by, user_backoffice.last_login
`

type UpdateUserBackofficeParams struct {
	Name                   sql.NullString `json:"name"`
	Phone                  string         `json:"phone"`
	ProfilePictureImageUrl sql.NullString `json:"profile_picture_image_url"`
	UpdatedBy              sql.NullString `json:"updated_by"`
	Guid                   string         `json:"guid"`
}

func (q *Queries) UpdateUserBackoffice(ctx context.Context, arg UpdateUserBackofficeParams) (UserBackoffice, error) {
	row := q.db.QueryRowContext(ctx, updateUserBackoffice,
		arg.Name,
		arg.Phone,
		arg.ProfilePictureImageUrl,
		arg.UpdatedBy,
		arg.Guid,
	)
	var i UserBackoffice
	err := row.Scan(
		&i.ID,
		&i.Guid,
		&i.Name,
		&i.ProfilePictureImageUrl,
		&i.Phone,
		&i.Email,
		&i.RoleID,
		&i.Password,
		&i.Salt,
		&i.IsActive,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
		&i.LastLogin,
	)
	return i, err
}

const updateUserBackofficeIsActive = `-- name: UpdateUserBackofficeIsActive :exec
UPDATE user_backoffice
SET
    is_active = $1,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = $2
WHERE
    guid = $3
    AND deleted_at IS NULL
`

type UpdateUserBackofficeIsActiveParams struct {
	IsActive  sql.NullBool   `json:"is_active"`
	UpdatedBy sql.NullString `json:"updated_by"`
	Guid      string         `json:"guid"`
}

func (q *Queries) UpdateUserBackofficeIsActive(ctx context.Context, arg UpdateUserBackofficeIsActiveParams) error {
	_, err := q.db.ExecContext(ctx, updateUserBackofficeIsActive, arg.IsActive, arg.UpdatedBy, arg.Guid)
	return err
}

const updateUserBackofficePassword = `-- name: UpdateUserBackofficePassword :exec
UPDATE user_backoffice
SET
    password = $1,
    salt = $2,
    updated_at = (now() at time zone 'UTC')::TIMESTAMP,
    updated_by = $3
WHERE
    guid = $4
    AND deleted_at IS NULL
RETURNING user_backoffice.id, user_backoffice.guid, user_backoffice.name, user_backoffice.profile_picture_image_url, user_backoffice.phone, user_backoffice.email, user_backoffice.role_id, user_backoffice.password, user_backoffice.salt, user_backoffice.is_active, user_backoffice.created_at, user_backoffice.created_by, user_backoffice.updated_at, user_backoffice.updated_by, user_backoffice.deleted_at, user_backoffice.deleted_by, user_backoffice.last_login
`

type UpdateUserBackofficePasswordParams struct {
	Password  string         `json:"password"`
	Salt      string         `json:"salt"`
	UpdatedBy sql.NullString `json:"updated_by"`
	Guid      string         `json:"guid"`
}

func (q *Queries) UpdateUserBackofficePassword(ctx context.Context, arg UpdateUserBackofficePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserBackofficePassword,
		arg.Password,
		arg.Salt,
		arg.UpdatedBy,
		arg.Guid,
	)
	return err
}
