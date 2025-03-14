// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: workspace.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/leg100/otf/internal/resource"
)

const countWorkspaces = `-- name: CountWorkspaces :one
WITH
    workspaces AS (
        SELECT w.workspace_id
        FROM workspaces w
        LEFT JOIN (workspace_tags wt JOIN tags t USING (tag_id)) ON w.workspace_id = wt.workspace_id
		LEFT JOIN runs r ON w.latest_run_id = r.run_id
        WHERE w.name              LIKE '%' || $1 || '%'
        AND   w.organization_name LIKE ANY($2::text[])
		AND (($3::text[] IS NULL) OR (r.status = ANY($3::text[])))
        GROUP BY w.workspace_id
        HAVING array_agg(t.name) @> $4::text[]
    )
SELECT count(*)
FROM workspaces
`

type CountWorkspacesParams struct {
	Search            pgtype.Text
	OrganizationNames []pgtype.Text
	Status            []pgtype.Text
	Tags              []pgtype.Text
}

func (q *Queries) CountWorkspaces(ctx context.Context, arg CountWorkspacesParams) (int64, error) {
	row := q.db.QueryRow(ctx, countWorkspaces,
		arg.Search,
		arg.OrganizationNames,
		arg.Status,
		arg.Tags,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countWorkspacesByUsername = `-- name: CountWorkspacesByUsername :one
SELECT count(*)
FROM workspaces w
JOIN workspace_permissions p USING (workspace_id)
JOIN teams t USING (team_id)
JOIN team_memberships tm USING (team_id)
JOIN users u USING (username)
WHERE w.organization_name = $1
AND   u.username          = $2
`

type CountWorkspacesByUsernameParams struct {
	OrganizationName pgtype.Text
	Username         pgtype.Text
}

func (q *Queries) CountWorkspacesByUsername(ctx context.Context, arg CountWorkspacesByUsernameParams) (int64, error) {
	row := q.db.QueryRow(ctx, countWorkspacesByUsername, arg.OrganizationName, arg.Username)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteWorkspaceByID = `-- name: DeleteWorkspaceByID :exec
DELETE
FROM workspaces
WHERE workspace_id = $1
`

func (q *Queries) DeleteWorkspaceByID(ctx context.Context, workspaceID resource.ID) error {
	_, err := q.db.Exec(ctx, deleteWorkspaceByID, workspaceID)
	return err
}

const findWorkspaceByID = `-- name: FindWorkspaceByID :one
SELECT
    w.workspace_id, w.created_at, w.updated_at, w.allow_destroy_plan, w.auto_apply, w.can_queue_destroy_plan, w.description, w.environment, w.execution_mode, w.global_remote_state, w.migration_environment, w.name, w.queue_all_runs, w.speculative_enabled, w.source_name, w.source_url, w.structured_run_output_enabled, w.terraform_version, w.trigger_prefixes, w.working_directory, w.lock_run_id, w.latest_run_id, w.organization_name, w.branch, w.current_state_version_id, w.trigger_patterns, w.vcs_tags_regex, w.allow_cli_apply, w.agent_pool_id, w.lock_user_id,
    (
        SELECT array_agg(name)::text[]
        FROM tags
        JOIN workspace_tags wt USING (tag_id)
        WHERE wt.workspace_id = w.workspace_id
    ) AS tags,
    r.status AS latest_run_status,
    rc.vcs_provider_id,
    rc.repo_path
FROM workspaces w
LEFT JOIN runs r ON w.latest_run_id = r.run_id
LEFT JOIN repo_connections rc ON w.workspace_id = rc.workspace_id
WHERE w.workspace_id = $1
`

type FindWorkspaceByIDRow struct {
	WorkspaceID                resource.ID
	CreatedAt                  pgtype.Timestamptz
	UpdatedAt                  pgtype.Timestamptz
	AllowDestroyPlan           pgtype.Bool
	AutoApply                  pgtype.Bool
	CanQueueDestroyPlan        pgtype.Bool
	Description                pgtype.Text
	Environment                pgtype.Text
	ExecutionMode              pgtype.Text
	GlobalRemoteState          pgtype.Bool
	MigrationEnvironment       pgtype.Text
	Name                       pgtype.Text
	QueueAllRuns               pgtype.Bool
	SpeculativeEnabled         pgtype.Bool
	SourceName                 pgtype.Text
	SourceURL                  pgtype.Text
	StructuredRunOutputEnabled pgtype.Bool
	TerraformVersion           pgtype.Text
	TriggerPrefixes            []pgtype.Text
	WorkingDirectory           pgtype.Text
	LockRunID                  *resource.ID
	LatestRunID                *resource.ID
	OrganizationName           pgtype.Text
	Branch                     pgtype.Text
	CurrentStateVersionID      *resource.ID
	TriggerPatterns            []pgtype.Text
	VCSTagsRegex               pgtype.Text
	AllowCLIApply              pgtype.Bool
	AgentPoolID                *resource.ID
	LockUserID                 *resource.ID
	Tags                       []pgtype.Text
	LatestRunStatus            pgtype.Text
	VCSProviderID              resource.ID
	RepoPath                   pgtype.Text
}

func (q *Queries) FindWorkspaceByID(ctx context.Context, id resource.ID) (FindWorkspaceByIDRow, error) {
	row := q.db.QueryRow(ctx, findWorkspaceByID, id)
	var i FindWorkspaceByIDRow
	err := row.Scan(
		&i.WorkspaceID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AllowDestroyPlan,
		&i.AutoApply,
		&i.CanQueueDestroyPlan,
		&i.Description,
		&i.Environment,
		&i.ExecutionMode,
		&i.GlobalRemoteState,
		&i.MigrationEnvironment,
		&i.Name,
		&i.QueueAllRuns,
		&i.SpeculativeEnabled,
		&i.SourceName,
		&i.SourceURL,
		&i.StructuredRunOutputEnabled,
		&i.TerraformVersion,
		&i.TriggerPrefixes,
		&i.WorkingDirectory,
		&i.LockRunID,
		&i.LatestRunID,
		&i.OrganizationName,
		&i.Branch,
		&i.CurrentStateVersionID,
		&i.TriggerPatterns,
		&i.VCSTagsRegex,
		&i.AllowCLIApply,
		&i.AgentPoolID,
		&i.LockUserID,
		&i.Tags,
		&i.LatestRunStatus,
		&i.VCSProviderID,
		&i.RepoPath,
	)
	return i, err
}

const findWorkspaceByIDForUpdate = `-- name: FindWorkspaceByIDForUpdate :one
SELECT
    w.workspace_id, w.created_at, w.updated_at, w.allow_destroy_plan, w.auto_apply, w.can_queue_destroy_plan, w.description, w.environment, w.execution_mode, w.global_remote_state, w.migration_environment, w.name, w.queue_all_runs, w.speculative_enabled, w.source_name, w.source_url, w.structured_run_output_enabled, w.terraform_version, w.trigger_prefixes, w.working_directory, w.lock_run_id, w.latest_run_id, w.organization_name, w.branch, w.current_state_version_id, w.trigger_patterns, w.vcs_tags_regex, w.allow_cli_apply, w.agent_pool_id, w.lock_user_id,
    (
        SELECT array_agg(name)::text[]
        FROM tags
        JOIN workspace_tags wt USING (tag_id)
        WHERE wt.workspace_id = w.workspace_id
    ) AS tags,
    r.status AS latest_run_status,
    rc.vcs_provider_id,
    rc.repo_path
FROM workspaces w
LEFT JOIN runs r ON w.latest_run_id = r.run_id
LEFT JOIN repo_connections rc ON w.workspace_id = rc.workspace_id
WHERE w.workspace_id = $1
FOR UPDATE OF w
`

type FindWorkspaceByIDForUpdateRow struct {
	WorkspaceID                resource.ID
	CreatedAt                  pgtype.Timestamptz
	UpdatedAt                  pgtype.Timestamptz
	AllowDestroyPlan           pgtype.Bool
	AutoApply                  pgtype.Bool
	CanQueueDestroyPlan        pgtype.Bool
	Description                pgtype.Text
	Environment                pgtype.Text
	ExecutionMode              pgtype.Text
	GlobalRemoteState          pgtype.Bool
	MigrationEnvironment       pgtype.Text
	Name                       pgtype.Text
	QueueAllRuns               pgtype.Bool
	SpeculativeEnabled         pgtype.Bool
	SourceName                 pgtype.Text
	SourceURL                  pgtype.Text
	StructuredRunOutputEnabled pgtype.Bool
	TerraformVersion           pgtype.Text
	TriggerPrefixes            []pgtype.Text
	WorkingDirectory           pgtype.Text
	LockRunID                  *resource.ID
	LatestRunID                *resource.ID
	OrganizationName           pgtype.Text
	Branch                     pgtype.Text
	CurrentStateVersionID      *resource.ID
	TriggerPatterns            []pgtype.Text
	VCSTagsRegex               pgtype.Text
	AllowCLIApply              pgtype.Bool
	AgentPoolID                *resource.ID
	LockUserID                 *resource.ID
	Tags                       []pgtype.Text
	LatestRunStatus            pgtype.Text
	VCSProviderID              resource.ID
	RepoPath                   pgtype.Text
}

func (q *Queries) FindWorkspaceByIDForUpdate(ctx context.Context, id resource.ID) (FindWorkspaceByIDForUpdateRow, error) {
	row := q.db.QueryRow(ctx, findWorkspaceByIDForUpdate, id)
	var i FindWorkspaceByIDForUpdateRow
	err := row.Scan(
		&i.WorkspaceID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AllowDestroyPlan,
		&i.AutoApply,
		&i.CanQueueDestroyPlan,
		&i.Description,
		&i.Environment,
		&i.ExecutionMode,
		&i.GlobalRemoteState,
		&i.MigrationEnvironment,
		&i.Name,
		&i.QueueAllRuns,
		&i.SpeculativeEnabled,
		&i.SourceName,
		&i.SourceURL,
		&i.StructuredRunOutputEnabled,
		&i.TerraformVersion,
		&i.TriggerPrefixes,
		&i.WorkingDirectory,
		&i.LockRunID,
		&i.LatestRunID,
		&i.OrganizationName,
		&i.Branch,
		&i.CurrentStateVersionID,
		&i.TriggerPatterns,
		&i.VCSTagsRegex,
		&i.AllowCLIApply,
		&i.AgentPoolID,
		&i.LockUserID,
		&i.Tags,
		&i.LatestRunStatus,
		&i.VCSProviderID,
		&i.RepoPath,
	)
	return i, err
}

const findWorkspaceByName = `-- name: FindWorkspaceByName :one
SELECT
    w.workspace_id, w.created_at, w.updated_at, w.allow_destroy_plan, w.auto_apply, w.can_queue_destroy_plan, w.description, w.environment, w.execution_mode, w.global_remote_state, w.migration_environment, w.name, w.queue_all_runs, w.speculative_enabled, w.source_name, w.source_url, w.structured_run_output_enabled, w.terraform_version, w.trigger_prefixes, w.working_directory, w.lock_run_id, w.latest_run_id, w.organization_name, w.branch, w.current_state_version_id, w.trigger_patterns, w.vcs_tags_regex, w.allow_cli_apply, w.agent_pool_id, w.lock_user_id,
    (
        SELECT array_agg(name)::text[]
        FROM tags
        JOIN workspace_tags wt USING (tag_id)
        WHERE wt.workspace_id = w.workspace_id
    ) AS tags,
    r.status AS latest_run_status,
    rc.vcs_provider_id,
    rc.repo_path
FROM workspaces w
LEFT JOIN runs r ON w.latest_run_id = r.run_id
LEFT JOIN repo_connections rc ON w.workspace_id = rc.workspace_id
LEFT JOIN (workspace_tags wt JOIN tags t USING (tag_id)) ON w.workspace_id = rc.workspace_id
WHERE w.name              = $1
AND   w.organization_name = $2
`

type FindWorkspaceByNameParams struct {
	Name             pgtype.Text
	OrganizationName pgtype.Text
}

type FindWorkspaceByNameRow struct {
	WorkspaceID                resource.ID
	CreatedAt                  pgtype.Timestamptz
	UpdatedAt                  pgtype.Timestamptz
	AllowDestroyPlan           pgtype.Bool
	AutoApply                  pgtype.Bool
	CanQueueDestroyPlan        pgtype.Bool
	Description                pgtype.Text
	Environment                pgtype.Text
	ExecutionMode              pgtype.Text
	GlobalRemoteState          pgtype.Bool
	MigrationEnvironment       pgtype.Text
	Name                       pgtype.Text
	QueueAllRuns               pgtype.Bool
	SpeculativeEnabled         pgtype.Bool
	SourceName                 pgtype.Text
	SourceURL                  pgtype.Text
	StructuredRunOutputEnabled pgtype.Bool
	TerraformVersion           pgtype.Text
	TriggerPrefixes            []pgtype.Text
	WorkingDirectory           pgtype.Text
	LockRunID                  *resource.ID
	LatestRunID                *resource.ID
	OrganizationName           pgtype.Text
	Branch                     pgtype.Text
	CurrentStateVersionID      *resource.ID
	TriggerPatterns            []pgtype.Text
	VCSTagsRegex               pgtype.Text
	AllowCLIApply              pgtype.Bool
	AgentPoolID                *resource.ID
	LockUserID                 *resource.ID
	Tags                       []pgtype.Text
	LatestRunStatus            pgtype.Text
	VCSProviderID              resource.ID
	RepoPath                   pgtype.Text
}

func (q *Queries) FindWorkspaceByName(ctx context.Context, arg FindWorkspaceByNameParams) (FindWorkspaceByNameRow, error) {
	row := q.db.QueryRow(ctx, findWorkspaceByName, arg.Name, arg.OrganizationName)
	var i FindWorkspaceByNameRow
	err := row.Scan(
		&i.WorkspaceID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AllowDestroyPlan,
		&i.AutoApply,
		&i.CanQueueDestroyPlan,
		&i.Description,
		&i.Environment,
		&i.ExecutionMode,
		&i.GlobalRemoteState,
		&i.MigrationEnvironment,
		&i.Name,
		&i.QueueAllRuns,
		&i.SpeculativeEnabled,
		&i.SourceName,
		&i.SourceURL,
		&i.StructuredRunOutputEnabled,
		&i.TerraformVersion,
		&i.TriggerPrefixes,
		&i.WorkingDirectory,
		&i.LockRunID,
		&i.LatestRunID,
		&i.OrganizationName,
		&i.Branch,
		&i.CurrentStateVersionID,
		&i.TriggerPatterns,
		&i.VCSTagsRegex,
		&i.AllowCLIApply,
		&i.AgentPoolID,
		&i.LockUserID,
		&i.Tags,
		&i.LatestRunStatus,
		&i.VCSProviderID,
		&i.RepoPath,
	)
	return i, err
}

const findWorkspaces = `-- name: FindWorkspaces :many
SELECT
    w.workspace_id, w.created_at, w.updated_at, w.allow_destroy_plan, w.auto_apply, w.can_queue_destroy_plan, w.description, w.environment, w.execution_mode, w.global_remote_state, w.migration_environment, w.name, w.queue_all_runs, w.speculative_enabled, w.source_name, w.source_url, w.structured_run_output_enabled, w.terraform_version, w.trigger_prefixes, w.working_directory, w.lock_run_id, w.latest_run_id, w.organization_name, w.branch, w.current_state_version_id, w.trigger_patterns, w.vcs_tags_regex, w.allow_cli_apply, w.agent_pool_id, w.lock_user_id,
    (
        SELECT array_agg(name)::text[]
        FROM tags
        JOIN workspace_tags wt USING (tag_id)
        WHERE wt.workspace_id = w.workspace_id
        GROUP BY wt.workspace_id
    ) AS tags,
    r.status AS latest_run_status,
    rc.vcs_provider_id,
    rc.repo_path
FROM workspaces w
LEFT JOIN runs r ON w.latest_run_id = r.run_id
LEFT JOIN repo_connections rc ON w.workspace_id = rc.workspace_id
LEFT JOIN (workspace_tags wt JOIN tags t USING (tag_id)) ON wt.workspace_id = w.workspace_id
WHERE w.name                LIKE '%' || $1 || '%'
AND   w.organization_name   LIKE ANY($2::text[])
AND   (($3::text[] IS NULL) OR (r.status = ANY($3::text[])))
GROUP BY w.workspace_id, r.status, rc.vcs_provider_id, rc.repo_path
HAVING array_agg(t.name) @> $4::text[]
ORDER BY w.name ASC
LIMIT $6::int
OFFSET $5::int
`

type FindWorkspacesParams struct {
	Search            pgtype.Text
	OrganizationNames []pgtype.Text
	Status            []pgtype.Text
	Tags              []pgtype.Text
	Offset            pgtype.Int4
	Limit             pgtype.Int4
}

type FindWorkspacesRow struct {
	WorkspaceID                resource.ID
	CreatedAt                  pgtype.Timestamptz
	UpdatedAt                  pgtype.Timestamptz
	AllowDestroyPlan           pgtype.Bool
	AutoApply                  pgtype.Bool
	CanQueueDestroyPlan        pgtype.Bool
	Description                pgtype.Text
	Environment                pgtype.Text
	ExecutionMode              pgtype.Text
	GlobalRemoteState          pgtype.Bool
	MigrationEnvironment       pgtype.Text
	Name                       pgtype.Text
	QueueAllRuns               pgtype.Bool
	SpeculativeEnabled         pgtype.Bool
	SourceName                 pgtype.Text
	SourceURL                  pgtype.Text
	StructuredRunOutputEnabled pgtype.Bool
	TerraformVersion           pgtype.Text
	TriggerPrefixes            []pgtype.Text
	WorkingDirectory           pgtype.Text
	LockRunID                  *resource.ID
	LatestRunID                *resource.ID
	OrganizationName           pgtype.Text
	Branch                     pgtype.Text
	CurrentStateVersionID      *resource.ID
	TriggerPatterns            []pgtype.Text
	VCSTagsRegex               pgtype.Text
	AllowCLIApply              pgtype.Bool
	AgentPoolID                *resource.ID
	LockUserID                 *resource.ID
	Tags                       []pgtype.Text
	LatestRunStatus            pgtype.Text
	VCSProviderID              resource.ID
	RepoPath                   pgtype.Text
}

func (q *Queries) FindWorkspaces(ctx context.Context, arg FindWorkspacesParams) ([]FindWorkspacesRow, error) {
	rows, err := q.db.Query(ctx, findWorkspaces,
		arg.Search,
		arg.OrganizationNames,
		arg.Status,
		arg.Tags,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindWorkspacesRow
	for rows.Next() {
		var i FindWorkspacesRow
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AllowDestroyPlan,
			&i.AutoApply,
			&i.CanQueueDestroyPlan,
			&i.Description,
			&i.Environment,
			&i.ExecutionMode,
			&i.GlobalRemoteState,
			&i.MigrationEnvironment,
			&i.Name,
			&i.QueueAllRuns,
			&i.SpeculativeEnabled,
			&i.SourceName,
			&i.SourceURL,
			&i.StructuredRunOutputEnabled,
			&i.TerraformVersion,
			&i.TriggerPrefixes,
			&i.WorkingDirectory,
			&i.LockRunID,
			&i.LatestRunID,
			&i.OrganizationName,
			&i.Branch,
			&i.CurrentStateVersionID,
			&i.TriggerPatterns,
			&i.VCSTagsRegex,
			&i.AllowCLIApply,
			&i.AgentPoolID,
			&i.LockUserID,
			&i.Tags,
			&i.LatestRunStatus,
			&i.VCSProviderID,
			&i.RepoPath,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findWorkspacesByConnection = `-- name: FindWorkspacesByConnection :many
SELECT
    w.workspace_id, w.created_at, w.updated_at, w.allow_destroy_plan, w.auto_apply, w.can_queue_destroy_plan, w.description, w.environment, w.execution_mode, w.global_remote_state, w.migration_environment, w.name, w.queue_all_runs, w.speculative_enabled, w.source_name, w.source_url, w.structured_run_output_enabled, w.terraform_version, w.trigger_prefixes, w.working_directory, w.lock_run_id, w.latest_run_id, w.organization_name, w.branch, w.current_state_version_id, w.trigger_patterns, w.vcs_tags_regex, w.allow_cli_apply, w.agent_pool_id, w.lock_user_id,
    (
        SELECT array_agg(name)::text[]
        FROM tags
        JOIN workspace_tags wt USING (tag_id)
        WHERE wt.workspace_id = w.workspace_id
    ) AS tags,
    r.status AS latest_run_status,
    rc.vcs_provider_id,
    rc.repo_path
FROM workspaces w
LEFT JOIN runs r ON w.latest_run_id = r.run_id
JOIN repo_connections rc ON w.workspace_id = rc.workspace_id
WHERE rc.vcs_provider_id = $1
AND   rc.repo_path = $2
`

type FindWorkspacesByConnectionParams struct {
	VCSProviderID resource.ID
	RepoPath      pgtype.Text
}

type FindWorkspacesByConnectionRow struct {
	WorkspaceID                resource.ID
	CreatedAt                  pgtype.Timestamptz
	UpdatedAt                  pgtype.Timestamptz
	AllowDestroyPlan           pgtype.Bool
	AutoApply                  pgtype.Bool
	CanQueueDestroyPlan        pgtype.Bool
	Description                pgtype.Text
	Environment                pgtype.Text
	ExecutionMode              pgtype.Text
	GlobalRemoteState          pgtype.Bool
	MigrationEnvironment       pgtype.Text
	Name                       pgtype.Text
	QueueAllRuns               pgtype.Bool
	SpeculativeEnabled         pgtype.Bool
	SourceName                 pgtype.Text
	SourceURL                  pgtype.Text
	StructuredRunOutputEnabled pgtype.Bool
	TerraformVersion           pgtype.Text
	TriggerPrefixes            []pgtype.Text
	WorkingDirectory           pgtype.Text
	LockRunID                  *resource.ID
	LatestRunID                *resource.ID
	OrganizationName           pgtype.Text
	Branch                     pgtype.Text
	CurrentStateVersionID      *resource.ID
	TriggerPatterns            []pgtype.Text
	VCSTagsRegex               pgtype.Text
	AllowCLIApply              pgtype.Bool
	AgentPoolID                *resource.ID
	LockUserID                 *resource.ID
	Tags                       []pgtype.Text
	LatestRunStatus            pgtype.Text
	VCSProviderID              resource.ID
	RepoPath                   pgtype.Text
}

func (q *Queries) FindWorkspacesByConnection(ctx context.Context, arg FindWorkspacesByConnectionParams) ([]FindWorkspacesByConnectionRow, error) {
	rows, err := q.db.Query(ctx, findWorkspacesByConnection, arg.VCSProviderID, arg.RepoPath)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindWorkspacesByConnectionRow
	for rows.Next() {
		var i FindWorkspacesByConnectionRow
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AllowDestroyPlan,
			&i.AutoApply,
			&i.CanQueueDestroyPlan,
			&i.Description,
			&i.Environment,
			&i.ExecutionMode,
			&i.GlobalRemoteState,
			&i.MigrationEnvironment,
			&i.Name,
			&i.QueueAllRuns,
			&i.SpeculativeEnabled,
			&i.SourceName,
			&i.SourceURL,
			&i.StructuredRunOutputEnabled,
			&i.TerraformVersion,
			&i.TriggerPrefixes,
			&i.WorkingDirectory,
			&i.LockRunID,
			&i.LatestRunID,
			&i.OrganizationName,
			&i.Branch,
			&i.CurrentStateVersionID,
			&i.TriggerPatterns,
			&i.VCSTagsRegex,
			&i.AllowCLIApply,
			&i.AgentPoolID,
			&i.LockUserID,
			&i.Tags,
			&i.LatestRunStatus,
			&i.VCSProviderID,
			&i.RepoPath,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findWorkspacesByUsername = `-- name: FindWorkspacesByUsername :many
SELECT
    w.workspace_id, w.created_at, w.updated_at, w.allow_destroy_plan, w.auto_apply, w.can_queue_destroy_plan, w.description, w.environment, w.execution_mode, w.global_remote_state, w.migration_environment, w.name, w.queue_all_runs, w.speculative_enabled, w.source_name, w.source_url, w.structured_run_output_enabled, w.terraform_version, w.trigger_prefixes, w.working_directory, w.lock_run_id, w.latest_run_id, w.organization_name, w.branch, w.current_state_version_id, w.trigger_patterns, w.vcs_tags_regex, w.allow_cli_apply, w.agent_pool_id, w.lock_user_id,
    (
        SELECT array_agg(name)::text[]
        FROM tags
        JOIN workspace_tags wt USING (tag_id)
        WHERE wt.workspace_id = w.workspace_id
    ) AS tags,
    r.status AS latest_run_status,
    rc.vcs_provider_id,
    rc.repo_path
FROM workspaces w
JOIN workspace_permissions p USING (workspace_id)
LEFT JOIN runs r ON w.latest_run_id = r.run_id
LEFT JOIN repo_connections rc ON w.workspace_id = rc.workspace_id
JOIN teams t USING (team_id)
JOIN team_memberships tm USING (team_id)
JOIN users u ON tm.username = u.username
WHERE w.organization_name  = $1
AND   u.username           = $2
ORDER BY w.updated_at DESC
LIMIT $4::int
OFFSET $3::int
`

type FindWorkspacesByUsernameParams struct {
	OrganizationName pgtype.Text
	Username         pgtype.Text
	Offset           pgtype.Int4
	Limit            pgtype.Int4
}

type FindWorkspacesByUsernameRow struct {
	WorkspaceID                resource.ID
	CreatedAt                  pgtype.Timestamptz
	UpdatedAt                  pgtype.Timestamptz
	AllowDestroyPlan           pgtype.Bool
	AutoApply                  pgtype.Bool
	CanQueueDestroyPlan        pgtype.Bool
	Description                pgtype.Text
	Environment                pgtype.Text
	ExecutionMode              pgtype.Text
	GlobalRemoteState          pgtype.Bool
	MigrationEnvironment       pgtype.Text
	Name                       pgtype.Text
	QueueAllRuns               pgtype.Bool
	SpeculativeEnabled         pgtype.Bool
	SourceName                 pgtype.Text
	SourceURL                  pgtype.Text
	StructuredRunOutputEnabled pgtype.Bool
	TerraformVersion           pgtype.Text
	TriggerPrefixes            []pgtype.Text
	WorkingDirectory           pgtype.Text
	LockRunID                  *resource.ID
	LatestRunID                *resource.ID
	OrganizationName           pgtype.Text
	Branch                     pgtype.Text
	CurrentStateVersionID      *resource.ID
	TriggerPatterns            []pgtype.Text
	VCSTagsRegex               pgtype.Text
	AllowCLIApply              pgtype.Bool
	AgentPoolID                *resource.ID
	LockUserID                 *resource.ID
	Tags                       []pgtype.Text
	LatestRunStatus            pgtype.Text
	VCSProviderID              resource.ID
	RepoPath                   pgtype.Text
}

func (q *Queries) FindWorkspacesByUsername(ctx context.Context, arg FindWorkspacesByUsernameParams) ([]FindWorkspacesByUsernameRow, error) {
	rows, err := q.db.Query(ctx, findWorkspacesByUsername,
		arg.OrganizationName,
		arg.Username,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindWorkspacesByUsernameRow
	for rows.Next() {
		var i FindWorkspacesByUsernameRow
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AllowDestroyPlan,
			&i.AutoApply,
			&i.CanQueueDestroyPlan,
			&i.Description,
			&i.Environment,
			&i.ExecutionMode,
			&i.GlobalRemoteState,
			&i.MigrationEnvironment,
			&i.Name,
			&i.QueueAllRuns,
			&i.SpeculativeEnabled,
			&i.SourceName,
			&i.SourceURL,
			&i.StructuredRunOutputEnabled,
			&i.TerraformVersion,
			&i.TriggerPrefixes,
			&i.WorkingDirectory,
			&i.LockRunID,
			&i.LatestRunID,
			&i.OrganizationName,
			&i.Branch,
			&i.CurrentStateVersionID,
			&i.TriggerPatterns,
			&i.VCSTagsRegex,
			&i.AllowCLIApply,
			&i.AgentPoolID,
			&i.LockUserID,
			&i.Tags,
			&i.LatestRunStatus,
			&i.VCSProviderID,
			&i.RepoPath,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertWorkspace = `-- name: InsertWorkspace :exec
INSERT INTO workspaces (
    workspace_id,
    created_at,
    updated_at,
    agent_pool_id,
    allow_cli_apply,
    allow_destroy_plan,
    auto_apply,
    branch,
    can_queue_destroy_plan,
    description,
    environment,
    execution_mode,
    global_remote_state,
    migration_environment,
    name,
    queue_all_runs,
    speculative_enabled,
    source_name,
    source_url,
    structured_run_output_enabled,
    terraform_version,
    trigger_prefixes,
    trigger_patterns,
    vcs_tags_regex,
    working_directory,
    organization_name
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15,
    $16,
    $17,
    $18,
    $19,
    $20,
    $21,
    $22,
    $23,
    $24,
    $25,
    $26
)
`

type InsertWorkspaceParams struct {
	ID                         resource.ID
	CreatedAt                  pgtype.Timestamptz
	UpdatedAt                  pgtype.Timestamptz
	AgentPoolID                *resource.ID
	AllowCLIApply              pgtype.Bool
	AllowDestroyPlan           pgtype.Bool
	AutoApply                  pgtype.Bool
	Branch                     pgtype.Text
	CanQueueDestroyPlan        pgtype.Bool
	Description                pgtype.Text
	Environment                pgtype.Text
	ExecutionMode              pgtype.Text
	GlobalRemoteState          pgtype.Bool
	MigrationEnvironment       pgtype.Text
	Name                       pgtype.Text
	QueueAllRuns               pgtype.Bool
	SpeculativeEnabled         pgtype.Bool
	SourceName                 pgtype.Text
	SourceURL                  pgtype.Text
	StructuredRunOutputEnabled pgtype.Bool
	TerraformVersion           pgtype.Text
	TriggerPrefixes            []pgtype.Text
	TriggerPatterns            []pgtype.Text
	VCSTagsRegex               pgtype.Text
	WorkingDirectory           pgtype.Text
	OrganizationName           pgtype.Text
}

func (q *Queries) InsertWorkspace(ctx context.Context, arg InsertWorkspaceParams) error {
	_, err := q.db.Exec(ctx, insertWorkspace,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.AgentPoolID,
		arg.AllowCLIApply,
		arg.AllowDestroyPlan,
		arg.AutoApply,
		arg.Branch,
		arg.CanQueueDestroyPlan,
		arg.Description,
		arg.Environment,
		arg.ExecutionMode,
		arg.GlobalRemoteState,
		arg.MigrationEnvironment,
		arg.Name,
		arg.QueueAllRuns,
		arg.SpeculativeEnabled,
		arg.SourceName,
		arg.SourceURL,
		arg.StructuredRunOutputEnabled,
		arg.TerraformVersion,
		arg.TriggerPrefixes,
		arg.TriggerPatterns,
		arg.VCSTagsRegex,
		arg.WorkingDirectory,
		arg.OrganizationName,
	)
	return err
}

const updateWorkspaceByID = `-- name: UpdateWorkspaceByID :one
UPDATE workspaces
SET
    agent_pool_id                 = $1,
    allow_destroy_plan            = $2,
    allow_cli_apply               = $3,
    auto_apply                    = $4,
    branch                        = $5,
    description                   = $6,
    execution_mode                = $7,
    global_remote_state           = $8,
    name                          = $9,
    queue_all_runs                = $10,
    speculative_enabled           = $11,
    structured_run_output_enabled = $12,
    terraform_version             = $13,
    trigger_prefixes              = $14,
    trigger_patterns              = $15,
    vcs_tags_regex                = $16,
    working_directory             = $17,
    updated_at                    = $18
WHERE workspace_id = $19
RETURNING workspace_id
`

type UpdateWorkspaceByIDParams struct {
	AgentPoolID                *resource.ID
	AllowDestroyPlan           pgtype.Bool
	AllowCLIApply              pgtype.Bool
	AutoApply                  pgtype.Bool
	Branch                     pgtype.Text
	Description                pgtype.Text
	ExecutionMode              pgtype.Text
	GlobalRemoteState          pgtype.Bool
	Name                       pgtype.Text
	QueueAllRuns               pgtype.Bool
	SpeculativeEnabled         pgtype.Bool
	StructuredRunOutputEnabled pgtype.Bool
	TerraformVersion           pgtype.Text
	TriggerPrefixes            []pgtype.Text
	TriggerPatterns            []pgtype.Text
	VCSTagsRegex               pgtype.Text
	WorkingDirectory           pgtype.Text
	UpdatedAt                  pgtype.Timestamptz
	ID                         resource.ID
}

func (q *Queries) UpdateWorkspaceByID(ctx context.Context, arg UpdateWorkspaceByIDParams) (resource.ID, error) {
	row := q.db.QueryRow(ctx, updateWorkspaceByID,
		arg.AgentPoolID,
		arg.AllowDestroyPlan,
		arg.AllowCLIApply,
		arg.AutoApply,
		arg.Branch,
		arg.Description,
		arg.ExecutionMode,
		arg.GlobalRemoteState,
		arg.Name,
		arg.QueueAllRuns,
		arg.SpeculativeEnabled,
		arg.StructuredRunOutputEnabled,
		arg.TerraformVersion,
		arg.TriggerPrefixes,
		arg.TriggerPatterns,
		arg.VCSTagsRegex,
		arg.WorkingDirectory,
		arg.UpdatedAt,
		arg.ID,
	)
	var workspace_id resource.ID
	err := row.Scan(&workspace_id)
	return workspace_id, err
}

const updateWorkspaceCurrentStateVersionID = `-- name: UpdateWorkspaceCurrentStateVersionID :one
UPDATE workspaces
SET current_state_version_id = $1
WHERE workspace_id = $2
RETURNING workspace_id
`

type UpdateWorkspaceCurrentStateVersionIDParams struct {
	StateVersionID *resource.ID
	WorkspaceID    resource.ID
}

func (q *Queries) UpdateWorkspaceCurrentStateVersionID(ctx context.Context, arg UpdateWorkspaceCurrentStateVersionIDParams) (resource.ID, error) {
	row := q.db.QueryRow(ctx, updateWorkspaceCurrentStateVersionID, arg.StateVersionID, arg.WorkspaceID)
	var workspace_id resource.ID
	err := row.Scan(&workspace_id)
	return workspace_id, err
}

const updateWorkspaceLatestRun = `-- name: UpdateWorkspaceLatestRun :exec
UPDATE workspaces
SET latest_run_id = $1
WHERE workspace_id = $2
`

type UpdateWorkspaceLatestRunParams struct {
	RunID       *resource.ID
	WorkspaceID resource.ID
}

func (q *Queries) UpdateWorkspaceLatestRun(ctx context.Context, arg UpdateWorkspaceLatestRunParams) error {
	_, err := q.db.Exec(ctx, updateWorkspaceLatestRun, arg.RunID, arg.WorkspaceID)
	return err
}

const updateWorkspaceLockByID = `-- name: UpdateWorkspaceLockByID :exec
UPDATE workspaces
SET
    lock_user_id = $1,
    lock_run_id = $2
WHERE workspace_id = $3
`

type UpdateWorkspaceLockByIDParams struct {
	UserID      *resource.ID
	RunID       *resource.ID
	WorkspaceID resource.ID
}

func (q *Queries) UpdateWorkspaceLockByID(ctx context.Context, arg UpdateWorkspaceLockByIDParams) error {
	_, err := q.db.Exec(ctx, updateWorkspaceLockByID, arg.UserID, arg.RunID, arg.WorkspaceID)
	return err
}
