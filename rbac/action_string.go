// Code generated by "stringer -type Action ./rbac"; DO NOT EDIT.

package rbac

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WatchAction-0]
	_ = x[CreateOrganizationAction-1]
	_ = x[UpdateOrganizationAction-2]
	_ = x[GetOrganizationAction-3]
	_ = x[ListOrganizationsAction-4]
	_ = x[GetEntitlementsAction-5]
	_ = x[DeleteOrganizationAction-6]
	_ = x[CreateVCSProviderAction-7]
	_ = x[GetVCSProviderAction-8]
	_ = x[ListVCSProvidersAction-9]
	_ = x[DeleteVCSProviderAction-10]
	_ = x[CreateAgentTokenAction-11]
	_ = x[ListAgentTokensAction-12]
	_ = x[DeleteAgentTokenAction-13]
	_ = x[CreateRegistrySessionAction-14]
	_ = x[CreateModuleAction-15]
	_ = x[CreateModuleVersionAction-16]
	_ = x[UpdateModuleAction-17]
	_ = x[ListModulesAction-18]
	_ = x[GetModuleAction-19]
	_ = x[DeleteModuleAction-20]
	_ = x[DeleteModuleVersionAction-21]
	_ = x[CreateVariableAction-22]
	_ = x[UpdateVariableAction-23]
	_ = x[ListVariablesAction-24]
	_ = x[GetVariableAction-25]
	_ = x[DeleteVariableAction-26]
	_ = x[GetRunAction-27]
	_ = x[ListRunsAction-28]
	_ = x[ApplyRunAction-29]
	_ = x[CreateRunAction-30]
	_ = x[DiscardRunAction-31]
	_ = x[DeleteRunAction-32]
	_ = x[CancelRunAction-33]
	_ = x[EnqueuePlanAction-34]
	_ = x[StartPhaseAction-35]
	_ = x[FinishPhaseAction-36]
	_ = x[PutChunkAction-37]
	_ = x[TailLogsAction-38]
	_ = x[GetPlanFileAction-39]
	_ = x[UploadPlanFileAction-40]
	_ = x[GetLockFileAction-41]
	_ = x[UploadLockFileAction-42]
	_ = x[ListWorkspacesAction-43]
	_ = x[GetWorkspaceAction-44]
	_ = x[CreateWorkspaceAction-45]
	_ = x[DeleteWorkspaceAction-46]
	_ = x[SetWorkspacePermissionAction-47]
	_ = x[UnsetWorkspacePermissionAction-48]
	_ = x[UpdateWorkspaceAction-49]
	_ = x[LockWorkspaceAction-50]
	_ = x[UnlockWorkspaceAction-51]
	_ = x[ForceUnlockWorkspaceAction-52]
	_ = x[CreateStateVersionAction-53]
	_ = x[ListStateVersionsAction-54]
	_ = x[GetStateVersionAction-55]
	_ = x[DownloadStateAction-56]
	_ = x[GetStateVersionOutputAction-57]
	_ = x[CreateConfigurationVersionAction-58]
	_ = x[ListConfigurationVersionsAction-59]
	_ = x[GetConfigurationVersionAction-60]
	_ = x[DownloadConfigurationVersionAction-61]
	_ = x[DeleteConfigurationVersionAction-62]
	_ = x[CreateUserAction-63]
	_ = x[ListUsersAction-64]
	_ = x[GetUserAction-65]
	_ = x[DeleteUserAction-66]
	_ = x[CreateTeamAction-67]
	_ = x[UpdateTeamAction-68]
	_ = x[GetTeamAction-69]
	_ = x[ListTeamsAction-70]
	_ = x[AddTeamMembershipAction-71]
	_ = x[RemoveTeamMembershipAction-72]
}

const _Action_name = "WatchActionCreateOrganizationActionUpdateOrganizationActionGetOrganizationActionListOrganizationsActionGetEntitlementsActionDeleteOrganizationActionCreateVCSProviderActionGetVCSProviderActionListVCSProvidersActionDeleteVCSProviderActionCreateAgentTokenActionListAgentTokensActionDeleteAgentTokenActionCreateRegistrySessionActionCreateModuleActionCreateModuleVersionActionUpdateModuleActionListModulesActionGetModuleActionDeleteModuleActionDeleteModuleVersionActionCreateVariableActionUpdateVariableActionListVariablesActionGetVariableActionDeleteVariableActionGetRunActionListRunsActionApplyRunActionCreateRunActionDiscardRunActionDeleteRunActionCancelRunActionEnqueuePlanActionStartPhaseActionFinishPhaseActionPutChunkActionTailLogsActionGetPlanFileActionUploadPlanFileActionGetLockFileActionUploadLockFileActionListWorkspacesActionGetWorkspaceActionCreateWorkspaceActionDeleteWorkspaceActionSetWorkspacePermissionActionUnsetWorkspacePermissionActionUpdateWorkspaceActionLockWorkspaceActionUnlockWorkspaceActionForceUnlockWorkspaceActionCreateStateVersionActionListStateVersionsActionGetStateVersionActionDownloadStateActionGetStateVersionOutputActionCreateConfigurationVersionActionListConfigurationVersionsActionGetConfigurationVersionActionDownloadConfigurationVersionActionDeleteConfigurationVersionActionCreateUserActionListUsersActionGetUserActionDeleteUserActionCreateTeamActionUpdateTeamActionGetTeamActionListTeamsActionAddTeamMembershipActionRemoveTeamMembershipAction"

var _Action_index = [...]uint16{0, 11, 35, 59, 80, 103, 124, 148, 171, 191, 213, 236, 258, 279, 301, 328, 346, 371, 389, 406, 421, 439, 464, 484, 504, 523, 540, 560, 572, 586, 600, 615, 631, 646, 661, 678, 694, 711, 725, 739, 756, 776, 793, 813, 833, 851, 872, 893, 921, 951, 972, 991, 1012, 1038, 1062, 1085, 1106, 1125, 1152, 1184, 1215, 1244, 1278, 1310, 1326, 1341, 1354, 1370, 1386, 1402, 1415, 1430, 1453, 1479}

func (i Action) String() string {
	if i < 0 || i >= Action(len(_Action_index)-1) {
		return "Action(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Action_name[_Action_index[i]:_Action_index[i+1]]
}
