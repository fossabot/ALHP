// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"ALHP.go/ent/dbpackage"
	"entgo.io/ent/dialect/sql"
)

// DbPackage is the model entity for the DbPackage schema.
type DbPackage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Pkgbase holds the value of the "pkgbase" field.
	Pkgbase string `json:"pkgbase,omitempty"`
	// Packages holds the value of the "packages" field.
	Packages []string `json:"packages,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// SkipReason holds the value of the "skip_reason" field.
	SkipReason string `json:"skip_reason,omitempty"`
	// Repository holds the value of the "repository" field.
	Repository string `json:"repository,omitempty"`
	// March holds the value of the "march" field.
	March string `json:"march,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// RepoVersion holds the value of the "repo_version" field.
	RepoVersion string `json:"repo_version,omitempty"`
	// BuildTime holds the value of the "build_time" field.
	BuildTime time.Time `json:"build_time,omitempty"`
	// BuildDuration holds the value of the "build_duration" field.
	BuildDuration uint64 `json:"build_duration,omitempty"`
	// Updated holds the value of the "updated" field.
	Updated time.Time `json:"updated,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DbPackage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dbpackage.FieldPackages:
			values[i] = new([]byte)
		case dbpackage.FieldID, dbpackage.FieldStatus, dbpackage.FieldBuildDuration:
			values[i] = new(sql.NullInt64)
		case dbpackage.FieldPkgbase, dbpackage.FieldSkipReason, dbpackage.FieldRepository, dbpackage.FieldMarch, dbpackage.FieldVersion, dbpackage.FieldRepoVersion:
			values[i] = new(sql.NullString)
		case dbpackage.FieldBuildTime, dbpackage.FieldUpdated:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DbPackage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DbPackage fields.
func (dp *DbPackage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dbpackage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dp.ID = int(value.Int64)
		case dbpackage.FieldPkgbase:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pkgbase", values[i])
			} else if value.Valid {
				dp.Pkgbase = value.String
			}
		case dbpackage.FieldPackages:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field packages", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dp.Packages); err != nil {
					return fmt.Errorf("unmarshal field packages: %w", err)
				}
			}
		case dbpackage.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				dp.Status = int(value.Int64)
			}
		case dbpackage.FieldSkipReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field skip_reason", values[i])
			} else if value.Valid {
				dp.SkipReason = value.String
			}
		case dbpackage.FieldRepository:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repository", values[i])
			} else if value.Valid {
				dp.Repository = value.String
			}
		case dbpackage.FieldMarch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field march", values[i])
			} else if value.Valid {
				dp.March = value.String
			}
		case dbpackage.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				dp.Version = value.String
			}
		case dbpackage.FieldRepoVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repo_version", values[i])
			} else if value.Valid {
				dp.RepoVersion = value.String
			}
		case dbpackage.FieldBuildTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field build_time", values[i])
			} else if value.Valid {
				dp.BuildTime = value.Time
			}
		case dbpackage.FieldBuildDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field build_duration", values[i])
			} else if value.Valid {
				dp.BuildDuration = uint64(value.Int64)
			}
		case dbpackage.FieldUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated", values[i])
			} else if value.Valid {
				dp.Updated = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this DbPackage.
// Note that you need to call DbPackage.Unwrap() before calling this method if this DbPackage
// was returned from a transaction, and the transaction was committed or rolled back.
func (dp *DbPackage) Update() *DbPackageUpdateOne {
	return (&DbPackageClient{config: dp.config}).UpdateOne(dp)
}

// Unwrap unwraps the DbPackage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dp *DbPackage) Unwrap() *DbPackage {
	tx, ok := dp.config.driver.(*txDriver)
	if !ok {
		panic("ent: DbPackage is not a transactional entity")
	}
	dp.config.driver = tx.drv
	return dp
}

// String implements the fmt.Stringer.
func (dp *DbPackage) String() string {
	var builder strings.Builder
	builder.WriteString("DbPackage(")
	builder.WriteString(fmt.Sprintf("id=%v", dp.ID))
	builder.WriteString(", pkgbase=")
	builder.WriteString(dp.Pkgbase)
	builder.WriteString(", packages=")
	builder.WriteString(fmt.Sprintf("%v", dp.Packages))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", dp.Status))
	builder.WriteString(", skip_reason=")
	builder.WriteString(dp.SkipReason)
	builder.WriteString(", repository=")
	builder.WriteString(dp.Repository)
	builder.WriteString(", march=")
	builder.WriteString(dp.March)
	builder.WriteString(", version=")
	builder.WriteString(dp.Version)
	builder.WriteString(", repo_version=")
	builder.WriteString(dp.RepoVersion)
	builder.WriteString(", build_time=")
	builder.WriteString(dp.BuildTime.Format(time.ANSIC))
	builder.WriteString(", build_duration=")
	builder.WriteString(fmt.Sprintf("%v", dp.BuildDuration))
	builder.WriteString(", updated=")
	builder.WriteString(dp.Updated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DbPackages is a parsable slice of DbPackage.
type DbPackages []*DbPackage

func (dp DbPackages) config(cfg config) {
	for _i := range dp {
		dp[_i].config = cfg
	}
}