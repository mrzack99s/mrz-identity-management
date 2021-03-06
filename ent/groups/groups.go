// Code generated by entc, DO NOT EDIT.

package groups

import (
	"time"
)

const (
	// Label holds the string label denoting the groups type in the database.
	Label = "groups"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGName holds the string denoting the g_name field in the database.
	FieldGName = "g_name"
	// FieldGIsIntOrg holds the string denoting the g_is_int_org field in the database.
	FieldGIsIntOrg = "g_is_int_org"
	// FieldGIsSuperAdmin holds the string denoting the g_is_super_admin field in the database.
	FieldGIsSuperAdmin = "g_is_super_admin"
	// FieldGCreatedAt holds the string denoting the g_created_at field in the database.
	FieldGCreatedAt = "g_created_at"
	// EdgeUseBandwidth holds the string denoting the use_bandwidth edge name in mutations.
	EdgeUseBandwidth = "use_bandwidth"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the groups in the database.
	Table = "groups"
	// UseBandwidthTable is the table that holds the use_bandwidth relation/edge.
	UseBandwidthTable = "groups"
	// UseBandwidthInverseTable is the table name for the GroupBandwidth entity.
	// It exists in this package in order to avoid circular dependency with the "groupbandwidth" package.
	UseBandwidthInverseTable = "group_bandwidths"
	// UseBandwidthColumn is the table column denoting the use_bandwidth relation/edge.
	UseBandwidthColumn = "group_bandwidth_groups"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the Users entity.
	// It exists in this package in order to avoid circular dependency with the "users" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "groups_users"
)

// Columns holds all SQL columns for groups fields.
var Columns = []string{
	FieldID,
	FieldGName,
	FieldGIsIntOrg,
	FieldGIsSuperAdmin,
	FieldGCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "groups"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_bandwidth_groups",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// GNameValidator is a validator for the "g_name" field. It is called by the builders before save.
	GNameValidator func(string) error
	// DefaultGIsIntOrg holds the default value on creation for the "g_is_int_org" field.
	DefaultGIsIntOrg bool
	// DefaultGIsSuperAdmin holds the default value on creation for the "g_is_super_admin" field.
	DefaultGIsSuperAdmin bool
	// DefaultGCreatedAt holds the default value on creation for the "g_created_at" field.
	DefaultGCreatedAt func() time.Time
)
