// Code generated by entc, DO NOT EDIT.

package dbpackage

import (
	"fmt"
)

const (
	// Label holds the string label denoting the dbpackage type in the database.
	Label = "db_package"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPkgbase holds the string denoting the pkgbase field in the database.
	FieldPkgbase = "pkgbase"
	// FieldPackages holds the string denoting the packages field in the database.
	FieldPackages = "packages"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSkipReason holds the string denoting the skip_reason field in the database.
	FieldSkipReason = "skip_reason"
	// FieldRepository holds the string denoting the repository field in the database.
	FieldRepository = "repository"
	// FieldMarch holds the string denoting the march field in the database.
	FieldMarch = "march"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldRepoVersion holds the string denoting the repo_version field in the database.
	FieldRepoVersion = "repo_version"
	// FieldBuildTimeStart holds the string denoting the build_time_start field in the database.
	FieldBuildTimeStart = "build_time_start"
	// FieldBuildTimeEnd holds the string denoting the build_time_end field in the database.
	FieldBuildTimeEnd = "build_time_end"
	// FieldUpdated holds the string denoting the updated field in the database.
	FieldUpdated = "updated"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldLto holds the string denoting the lto field in the database.
	FieldLto = "lto"
	// Table holds the table name of the dbpackage in the database.
	Table = "db_packages"
)

// Columns holds all SQL columns for dbpackage fields.
var Columns = []string{
	FieldID,
	FieldPkgbase,
	FieldPackages,
	FieldStatus,
	FieldSkipReason,
	FieldRepository,
	FieldMarch,
	FieldVersion,
	FieldRepoVersion,
	FieldBuildTimeStart,
	FieldBuildTimeEnd,
	FieldUpdated,
	FieldHash,
	FieldLto,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// PkgbaseValidator is a validator for the "pkgbase" field. It is called by the builders before save.
	PkgbaseValidator func(string) error
	// MarchValidator is a validator for the "march" field. It is called by the builders before save.
	MarchValidator func(string) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusUnknown is the default value of the Status enum.
const DefaultStatus = StatusUnknown

// Status values.
const (
	StatusSkipped  Status = "skipped"
	StatusFailed   Status = "failed"
	StatusBuild    Status = "build"
	StatusQueued   Status = "queued"
	StatusBuilding Status = "building"
	StatusLatest   Status = "latest"
	StatusSigning  Status = "signing"
	StatusUnknown  Status = "unknown"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusSkipped, StatusFailed, StatusBuild, StatusQueued, StatusBuilding, StatusLatest, StatusSigning, StatusUnknown:
		return nil
	default:
		return fmt.Errorf("dbpackage: invalid enum value for status field: %q", s)
	}
}

// Repository defines the type for the "repository" enum field.
type Repository string

// Repository values.
const (
	RepositoryExtra     Repository = "extra"
	RepositoryCore      Repository = "core"
	RepositoryCommunity Repository = "community"
)

func (r Repository) String() string {
	return string(r)
}

// RepositoryValidator is a validator for the "repository" field enum values. It is called by the builders before save.
func RepositoryValidator(r Repository) error {
	switch r {
	case RepositoryExtra, RepositoryCore, RepositoryCommunity:
		return nil
	default:
		return fmt.Errorf("dbpackage: invalid enum value for repository field: %q", r)
	}
}

// Lto defines the type for the "lto" enum field.
type Lto string

// LtoUnknown is the default value of the Lto enum.
const DefaultLto = LtoUnknown

// Lto values.
const (
	LtoEnabled      Lto = "enabled"
	LtoUnknown      Lto = "unknown"
	LtoDisabled     Lto = "disabled"
	LtoAutoDisabled Lto = "auto_disabled"
)

func (l Lto) String() string {
	return string(l)
}

// LtoValidator is a validator for the "lto" field enum values. It is called by the builders before save.
func LtoValidator(l Lto) error {
	switch l {
	case LtoEnabled, LtoUnknown, LtoDisabled, LtoAutoDisabled:
		return nil
	default:
		return fmt.Errorf("dbpackage: invalid enum value for lto field: %q", l)
	}
}
