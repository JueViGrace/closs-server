// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
;

insert into closs_user (
    id,
    username,
    password,
    codigo,
    role,
    ult_sinc,
    version,
    created_at,
    updated_at
)
values (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING id, username, password, codigo, role, ult_sinc, version, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	ID        string
	Username  string
	Password  string
	Codigo    string
	Role      string
	UltSinc   string
	Version   string
	CreatedAt string
	UpdatedAt string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (ClossUser, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Codigo,
		arg.Role,
		arg.UltSinc,
		arg.Version,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i ClossUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Codigo,
		&i.Role,
		&i.UltSinc,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteUserById = `-- name: DeleteUserById :exec
delete from closs_user
where id = ?
`

func (q *Queries) DeleteUserById(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUserById, id)
	return err
}

const getUserById = `-- name: GetUserById :one
;

select id, username, password, codigo, role, ult_sinc, version, created_at, updated_at, deleted_at
from closs_user
where id = ?
`

func (q *Queries) GetUserById(ctx context.Context, id string) (ClossUser, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i ClossUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Codigo,
		&i.Role,
		&i.UltSinc,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
;

select id, username, password, codigo, role, ult_sinc, version, created_at, updated_at, deleted_at
from closs_user
where username = ? and deleted_at is null
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (ClossUser, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i ClossUser
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Codigo,
		&i.Role,
		&i.UltSinc,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
select id, username, password, codigo, role, ult_sinc, version, created_at, updated_at, deleted_at
from closs_user
`

func (q *Queries) GetUsers(ctx context.Context) ([]ClossUser, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ClossUser
	for rows.Next() {
		var i ClossUser
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Codigo,
			&i.Role,
			&i.UltSinc,
			&i.Version,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const softDeleteUser = `-- name: SoftDeleteUser :exec
update closs_user set 
    updated_at = ?,
    deleted_at = ?
where id = ?
`

type SoftDeleteUserParams struct {
	UpdatedAt string
	DeletedAt sql.NullString
	ID        string
}

func (q *Queries) SoftDeleteUser(ctx context.Context, arg SoftDeleteUserParams) error {
	_, err := q.db.ExecContext(ctx, softDeleteUser, arg.UpdatedAt, arg.DeletedAt, arg.ID)
	return err
}

const updateLastSync = `-- name: UpdateLastSync :exec
update closs_user set
    ult_sinc = ?,
    version = ?,
    updated_at = ?
where id = ?
RETURNING id, username, password, codigo, role, ult_sinc, version, created_at, updated_at, deleted_at
`

type UpdateLastSyncParams struct {
	UltSinc   string
	Version   string
	UpdatedAt string
	ID        string
}

func (q *Queries) UpdateLastSync(ctx context.Context, arg UpdateLastSyncParams) error {
	_, err := q.db.ExecContext(ctx, updateLastSync,
		arg.UltSinc,
		arg.Version,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updatePassword = `-- name: UpdatePassword :exec
update closs_user set
    password = ?,
    updated_at = ?
where id = ?
RETURNING id, username, password, codigo, role, ult_sinc, version, created_at, updated_at, deleted_at
`

type UpdatePasswordParams struct {
	Password  string
	UpdatedAt string
	ID        string
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, arg.Password, arg.UpdatedAt, arg.ID)
	return err
}
