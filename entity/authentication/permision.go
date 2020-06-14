package authentication

type Permission string

const (
	PermissionViewUser   = "PERMISSION_VIEW_USER"
	PermissionUpdateUser = "PERMISSION_UPDATE_USER"

	PermissionAssignRole = "PERMISSION_ASSIGN_ROLE"
	PermissionCreateRole = "PERMISSION_CREATE_ROLE"
	PermissionViewRole   = "PERMISSION_VIEW_ROLE"
	PermissionUpdateRole = "PERMISSION_UPDATE_ROLE"
	PermissionDeleteRole = "PERMISSION_DELETE_ROLE"
)

var Permissions = []Permission{
	PermissionViewUser,
	PermissionCreateRole,
	PermissionViewRole,
	PermissionUpdateRole,
	PermissionDeleteRole,
}
