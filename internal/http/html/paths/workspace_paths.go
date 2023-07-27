// Code generated by "go generate"; DO NOT EDIT.

package paths

import "fmt"

func Workspaces(organization string) string {
	return fmt.Sprintf("/app/organizations/%s/workspaces", organization)
}

func CreateWorkspace(organization string) string {
	return fmt.Sprintf("/app/organizations/%s/workspaces/create", organization)
}

func NewWorkspace(organization string) string {
	return fmt.Sprintf("/app/organizations/%s/workspaces/new", organization)
}

func Workspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s", workspace)
}

func EditWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/edit", workspace)
}

func UpdateWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/update", workspace)
}

func DeleteWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/delete", workspace)
}

func LockWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/lock", workspace)
}

func UnlockWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/unlock", workspace)
}

func ForceUnlockWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/force-unlock", workspace)
}

func SetPermissionWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/set-permission", workspace)
}

func UnsetPermissionWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/unset-permission", workspace)
}

func WatchWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/watch", workspace)
}

func ConnectWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/connect", workspace)
}

func DisconnectWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/disconnect", workspace)
}

func StartRunWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/start-run", workspace)
}

func SetupConnectionProviderWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/setup-connection-provider", workspace)
}

func SetupConnectionRepoWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/setup-connection-repo", workspace)
}

func CreateTagWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/create-tag", workspace)
}

func DeleteTagWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/delete-tag", workspace)
}

func StateWorkspace(workspace string) string {
	return fmt.Sprintf("/app/workspaces/%s/state", workspace)
}
