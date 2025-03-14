// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: repo_connections.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/leg100/otf/internal/resource"
)

const deleteModuleConnectionByID = `-- name: DeleteModuleConnectionByID :one
DELETE
FROM repo_connections
WHERE module_id = $1
RETURNING module_id, workspace_id, repo_path, vcs_provider_id
`

func (q *Queries) DeleteModuleConnectionByID(ctx context.Context, moduleID *resource.ID) (RepoConnection, error) {
	row := q.db.QueryRow(ctx, deleteModuleConnectionByID, moduleID)
	var i RepoConnection
	err := row.Scan(
		&i.ModuleID,
		&i.WorkspaceID,
		&i.RepoPath,
		&i.VCSProviderID,
	)
	return i, err
}

const deleteWorkspaceConnectionByID = `-- name: DeleteWorkspaceConnectionByID :one
DELETE
FROM repo_connections
WHERE workspace_id = $1
RETURNING module_id, workspace_id, repo_path, vcs_provider_id
`

func (q *Queries) DeleteWorkspaceConnectionByID(ctx context.Context, workspaceID *resource.ID) (RepoConnection, error) {
	row := q.db.QueryRow(ctx, deleteWorkspaceConnectionByID, workspaceID)
	var i RepoConnection
	err := row.Scan(
		&i.ModuleID,
		&i.WorkspaceID,
		&i.RepoPath,
		&i.VCSProviderID,
	)
	return i, err
}

const insertRepoConnection = `-- name: InsertRepoConnection :exec
INSERT INTO repo_connections (
    vcs_provider_id,
    repo_path,
    workspace_id,
    module_id
) VALUES (
    $1,
    $2,
    $3,
    $4
)
`

type InsertRepoConnectionParams struct {
	VCSProviderID resource.ID
	RepoPath      pgtype.Text
	WorkspaceID   *resource.ID
	ModuleID      *resource.ID
}

func (q *Queries) InsertRepoConnection(ctx context.Context, arg InsertRepoConnectionParams) error {
	_, err := q.db.Exec(ctx, insertRepoConnection,
		arg.VCSProviderID,
		arg.RepoPath,
		arg.WorkspaceID,
		arg.ModuleID,
	)
	return err
}
