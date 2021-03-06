// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/mrzack99s/mrz-identity-management/ent/groups"
	"github.com/mrzack99s/mrz-identity-management/ent/users"
)

// Users is the model entity for the Users schema.
type Users struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UPid holds the value of the "u_pid" field.
	UPid string `json:"u_pid,omitempty"`
	// UOrgid holds the value of the "u_orgid" field.
	UOrgid string `json:"u_orgid,omitempty"`
	// UFirstName holds the value of the "u_first_name" field.
	UFirstName string `json:"u_first_name,omitempty"`
	// ULastName holds the value of the "u_last_name" field.
	ULastName string `json:"u_last_name,omitempty"`
	// UIsActive holds the value of the "u_is_active" field.
	UIsActive bool `json:"u_is_active,omitempty"`
	// UCreatedAt holds the value of the "u_created_at" field.
	UCreatedAt time.Time `json:"u_created_at,omitempty"`
	// UPasswordUpdatedAt holds the value of the "u_password_updated_at" field.
	UPasswordUpdatedAt time.Time `json:"u_password_updated_at,omitempty"`
	// UExpiredAt holds the value of the "u_expired_at" field.
	UExpiredAt time.Time `json:"u_expired_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UsersQuery when eager-loading is set.
	Edges        UsersEdges `json:"edges"`
	groups_users *int
}

// UsersEdges holds the relations/edges for other nodes in the graph.
type UsersEdges struct {
	// InGroup holds the value of the in_group edge.
	InGroup *Groups `json:"in_group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InGroupOrErr returns the InGroup value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UsersEdges) InGroupOrErr() (*Groups, error) {
	if e.loadedTypes[0] {
		if e.InGroup == nil {
			// The edge in_group was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: groups.Label}
		}
		return e.InGroup, nil
	}
	return nil, &NotLoadedError{edge: "in_group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Users) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case users.FieldUIsActive:
			values[i] = new(sql.NullBool)
		case users.FieldID:
			values[i] = new(sql.NullInt64)
		case users.FieldUPid, users.FieldUOrgid, users.FieldUFirstName, users.FieldULastName:
			values[i] = new(sql.NullString)
		case users.FieldUCreatedAt, users.FieldUPasswordUpdatedAt, users.FieldUExpiredAt:
			values[i] = new(sql.NullTime)
		case users.ForeignKeys[0]: // groups_users
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Users", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Users fields.
func (u *Users) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case users.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case users.FieldUPid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field u_pid", values[i])
			} else if value.Valid {
				u.UPid = value.String
			}
		case users.FieldUOrgid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field u_orgid", values[i])
			} else if value.Valid {
				u.UOrgid = value.String
			}
		case users.FieldUFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field u_first_name", values[i])
			} else if value.Valid {
				u.UFirstName = value.String
			}
		case users.FieldULastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field u_last_name", values[i])
			} else if value.Valid {
				u.ULastName = value.String
			}
		case users.FieldUIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field u_is_active", values[i])
			} else if value.Valid {
				u.UIsActive = value.Bool
			}
		case users.FieldUCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field u_created_at", values[i])
			} else if value.Valid {
				u.UCreatedAt = value.Time
			}
		case users.FieldUPasswordUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field u_password_updated_at", values[i])
			} else if value.Valid {
				u.UPasswordUpdatedAt = value.Time
			}
		case users.FieldUExpiredAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field u_expired_at", values[i])
			} else if value.Valid {
				u.UExpiredAt = value.Time
			}
		case users.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field groups_users", value)
			} else if value.Valid {
				u.groups_users = new(int)
				*u.groups_users = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryInGroup queries the "in_group" edge of the Users entity.
func (u *Users) QueryInGroup() *GroupsQuery {
	return (&UsersClient{config: u.config}).QueryInGroup(u)
}

// Update returns a builder for updating this Users.
// Note that you need to call Users.Unwrap() before calling this method if this Users
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Users) Update() *UsersUpdateOne {
	return (&UsersClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the Users entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Users) Unwrap() *Users {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Users is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Users) String() string {
	var builder strings.Builder
	builder.WriteString("Users(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", u_pid=")
	builder.WriteString(u.UPid)
	builder.WriteString(", u_orgid=")
	builder.WriteString(u.UOrgid)
	builder.WriteString(", u_first_name=")
	builder.WriteString(u.UFirstName)
	builder.WriteString(", u_last_name=")
	builder.WriteString(u.ULastName)
	builder.WriteString(", u_is_active=")
	builder.WriteString(fmt.Sprintf("%v", u.UIsActive))
	builder.WriteString(", u_created_at=")
	builder.WriteString(u.UCreatedAt.Format(time.ANSIC))
	builder.WriteString(", u_password_updated_at=")
	builder.WriteString(u.UPasswordUpdatedAt.Format(time.ANSIC))
	builder.WriteString(", u_expired_at=")
	builder.WriteString(u.UExpiredAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UsersSlice is a parsable slice of Users.
type UsersSlice []*Users

func (u UsersSlice) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
