// Code generated by "go generate"; DO NOT EDIT.

package paths

import "fmt"

func OrganizationRuns(organization string) string {
	return fmt.Sprintf("/app/organizations/%s/runs", organization)
}
