package services

// Implement roles for users

// User roles
const (
	RoleSuperAdmin = "superadmin"
	RoleAdmin      = "admin"
	RoleUser       = "user"
	RoleGuest      = "guest"
)

// Role hierarchy
var roleHierarchy = map[string]int{
	RoleSuperAdmin: 3,
	RoleAdmin:      2,
	RoleUser:       1,
	RoleGuest:      0,
}

// Check if a user has sufficient role access
func HasAccess(userRole, requiredRole string) bool {
	return roleHierarchy[userRole] >= roleHierarchy[requiredRole]
}
