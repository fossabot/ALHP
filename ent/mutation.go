// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"ALHP.go/ent/dbpackage"
	"ALHP.go/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeDbPackage = "DbPackage"
)

// DbPackageMutation represents an operation that mutates the DbPackage nodes in the graph.
type DbPackageMutation struct {
	config
	op                Op
	typ               string
	id                *int
	pkgbase           *string
	packages          *[]string
	status            *int
	addstatus         *int
	skip_reason       *string
	repository        *string
	march             *string
	version           *string
	repo_version      *string
	build_time        *time.Time
	build_duration    *uint64
	addbuild_duration *uint64
	updated           *time.Time
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*DbPackage, error)
	predicates        []predicate.DbPackage
}

var _ ent.Mutation = (*DbPackageMutation)(nil)

// dbpackageOption allows management of the mutation configuration using functional options.
type dbpackageOption func(*DbPackageMutation)

// newDbPackageMutation creates new mutation for the DbPackage entity.
func newDbPackageMutation(c config, op Op, opts ...dbpackageOption) *DbPackageMutation {
	m := &DbPackageMutation{
		config:        c,
		op:            op,
		typ:           TypeDbPackage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDbPackageID sets the ID field of the mutation.
func withDbPackageID(id int) dbpackageOption {
	return func(m *DbPackageMutation) {
		var (
			err   error
			once  sync.Once
			value *DbPackage
		)
		m.oldValue = func(ctx context.Context) (*DbPackage, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().DbPackage.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDbPackage sets the old DbPackage of the mutation.
func withDbPackage(node *DbPackage) dbpackageOption {
	return func(m *DbPackageMutation) {
		m.oldValue = func(context.Context) (*DbPackage, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DbPackageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DbPackageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *DbPackageMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetPkgbase sets the "pkgbase" field.
func (m *DbPackageMutation) SetPkgbase(s string) {
	m.pkgbase = &s
}

// Pkgbase returns the value of the "pkgbase" field in the mutation.
func (m *DbPackageMutation) Pkgbase() (r string, exists bool) {
	v := m.pkgbase
	if v == nil {
		return
	}
	return *v, true
}

// OldPkgbase returns the old "pkgbase" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldPkgbase(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPkgbase is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPkgbase requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPkgbase: %w", err)
	}
	return oldValue.Pkgbase, nil
}

// ResetPkgbase resets all changes to the "pkgbase" field.
func (m *DbPackageMutation) ResetPkgbase() {
	m.pkgbase = nil
}

// SetPackages sets the "packages" field.
func (m *DbPackageMutation) SetPackages(s []string) {
	m.packages = &s
}

// Packages returns the value of the "packages" field in the mutation.
func (m *DbPackageMutation) Packages() (r []string, exists bool) {
	v := m.packages
	if v == nil {
		return
	}
	return *v, true
}

// OldPackages returns the old "packages" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldPackages(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPackages is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPackages requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPackages: %w", err)
	}
	return oldValue.Packages, nil
}

// ClearPackages clears the value of the "packages" field.
func (m *DbPackageMutation) ClearPackages() {
	m.packages = nil
	m.clearedFields[dbpackage.FieldPackages] = struct{}{}
}

// PackagesCleared returns if the "packages" field was cleared in this mutation.
func (m *DbPackageMutation) PackagesCleared() bool {
	_, ok := m.clearedFields[dbpackage.FieldPackages]
	return ok
}

// ResetPackages resets all changes to the "packages" field.
func (m *DbPackageMutation) ResetPackages() {
	m.packages = nil
	delete(m.clearedFields, dbpackage.FieldPackages)
}

// SetStatus sets the "status" field.
func (m *DbPackageMutation) SetStatus(i int) {
	m.status = &i
	m.addstatus = nil
}

// Status returns the value of the "status" field in the mutation.
func (m *DbPackageMutation) Status() (r int, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldStatus(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// AddStatus adds i to the "status" field.
func (m *DbPackageMutation) AddStatus(i int) {
	if m.addstatus != nil {
		*m.addstatus += i
	} else {
		m.addstatus = &i
	}
}

// AddedStatus returns the value that was added to the "status" field in this mutation.
func (m *DbPackageMutation) AddedStatus() (r int, exists bool) {
	v := m.addstatus
	if v == nil {
		return
	}
	return *v, true
}

// ResetStatus resets all changes to the "status" field.
func (m *DbPackageMutation) ResetStatus() {
	m.status = nil
	m.addstatus = nil
}

// SetSkipReason sets the "skip_reason" field.
func (m *DbPackageMutation) SetSkipReason(s string) {
	m.skip_reason = &s
}

// SkipReason returns the value of the "skip_reason" field in the mutation.
func (m *DbPackageMutation) SkipReason() (r string, exists bool) {
	v := m.skip_reason
	if v == nil {
		return
	}
	return *v, true
}

// OldSkipReason returns the old "skip_reason" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldSkipReason(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSkipReason is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSkipReason requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSkipReason: %w", err)
	}
	return oldValue.SkipReason, nil
}

// ClearSkipReason clears the value of the "skip_reason" field.
func (m *DbPackageMutation) ClearSkipReason() {
	m.skip_reason = nil
	m.clearedFields[dbpackage.FieldSkipReason] = struct{}{}
}

// SkipReasonCleared returns if the "skip_reason" field was cleared in this mutation.
func (m *DbPackageMutation) SkipReasonCleared() bool {
	_, ok := m.clearedFields[dbpackage.FieldSkipReason]
	return ok
}

// ResetSkipReason resets all changes to the "skip_reason" field.
func (m *DbPackageMutation) ResetSkipReason() {
	m.skip_reason = nil
	delete(m.clearedFields, dbpackage.FieldSkipReason)
}

// SetRepository sets the "repository" field.
func (m *DbPackageMutation) SetRepository(s string) {
	m.repository = &s
}

// Repository returns the value of the "repository" field in the mutation.
func (m *DbPackageMutation) Repository() (r string, exists bool) {
	v := m.repository
	if v == nil {
		return
	}
	return *v, true
}

// OldRepository returns the old "repository" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldRepository(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRepository is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRepository requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRepository: %w", err)
	}
	return oldValue.Repository, nil
}

// ResetRepository resets all changes to the "repository" field.
func (m *DbPackageMutation) ResetRepository() {
	m.repository = nil
}

// SetMarch sets the "march" field.
func (m *DbPackageMutation) SetMarch(s string) {
	m.march = &s
}

// March returns the value of the "march" field in the mutation.
func (m *DbPackageMutation) March() (r string, exists bool) {
	v := m.march
	if v == nil {
		return
	}
	return *v, true
}

// OldMarch returns the old "march" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldMarch(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMarch is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMarch requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMarch: %w", err)
	}
	return oldValue.March, nil
}

// ResetMarch resets all changes to the "march" field.
func (m *DbPackageMutation) ResetMarch() {
	m.march = nil
}

// SetVersion sets the "version" field.
func (m *DbPackageMutation) SetVersion(s string) {
	m.version = &s
}

// Version returns the value of the "version" field in the mutation.
func (m *DbPackageMutation) Version() (r string, exists bool) {
	v := m.version
	if v == nil {
		return
	}
	return *v, true
}

// OldVersion returns the old "version" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldVersion(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldVersion is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldVersion requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldVersion: %w", err)
	}
	return oldValue.Version, nil
}

// ClearVersion clears the value of the "version" field.
func (m *DbPackageMutation) ClearVersion() {
	m.version = nil
	m.clearedFields[dbpackage.FieldVersion] = struct{}{}
}

// VersionCleared returns if the "version" field was cleared in this mutation.
func (m *DbPackageMutation) VersionCleared() bool {
	_, ok := m.clearedFields[dbpackage.FieldVersion]
	return ok
}

// ResetVersion resets all changes to the "version" field.
func (m *DbPackageMutation) ResetVersion() {
	m.version = nil
	delete(m.clearedFields, dbpackage.FieldVersion)
}

// SetRepoVersion sets the "repo_version" field.
func (m *DbPackageMutation) SetRepoVersion(s string) {
	m.repo_version = &s
}

// RepoVersion returns the value of the "repo_version" field in the mutation.
func (m *DbPackageMutation) RepoVersion() (r string, exists bool) {
	v := m.repo_version
	if v == nil {
		return
	}
	return *v, true
}

// OldRepoVersion returns the old "repo_version" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldRepoVersion(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRepoVersion is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRepoVersion requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRepoVersion: %w", err)
	}
	return oldValue.RepoVersion, nil
}

// ClearRepoVersion clears the value of the "repo_version" field.
func (m *DbPackageMutation) ClearRepoVersion() {
	m.repo_version = nil
	m.clearedFields[dbpackage.FieldRepoVersion] = struct{}{}
}

// RepoVersionCleared returns if the "repo_version" field was cleared in this mutation.
func (m *DbPackageMutation) RepoVersionCleared() bool {
	_, ok := m.clearedFields[dbpackage.FieldRepoVersion]
	return ok
}

// ResetRepoVersion resets all changes to the "repo_version" field.
func (m *DbPackageMutation) ResetRepoVersion() {
	m.repo_version = nil
	delete(m.clearedFields, dbpackage.FieldRepoVersion)
}

// SetBuildTime sets the "build_time" field.
func (m *DbPackageMutation) SetBuildTime(t time.Time) {
	m.build_time = &t
}

// BuildTime returns the value of the "build_time" field in the mutation.
func (m *DbPackageMutation) BuildTime() (r time.Time, exists bool) {
	v := m.build_time
	if v == nil {
		return
	}
	return *v, true
}

// OldBuildTime returns the old "build_time" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldBuildTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldBuildTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldBuildTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBuildTime: %w", err)
	}
	return oldValue.BuildTime, nil
}

// ClearBuildTime clears the value of the "build_time" field.
func (m *DbPackageMutation) ClearBuildTime() {
	m.build_time = nil
	m.clearedFields[dbpackage.FieldBuildTime] = struct{}{}
}

// BuildTimeCleared returns if the "build_time" field was cleared in this mutation.
func (m *DbPackageMutation) BuildTimeCleared() bool {
	_, ok := m.clearedFields[dbpackage.FieldBuildTime]
	return ok
}

// ResetBuildTime resets all changes to the "build_time" field.
func (m *DbPackageMutation) ResetBuildTime() {
	m.build_time = nil
	delete(m.clearedFields, dbpackage.FieldBuildTime)
}

// SetBuildDuration sets the "build_duration" field.
func (m *DbPackageMutation) SetBuildDuration(u uint64) {
	m.build_duration = &u
	m.addbuild_duration = nil
}

// BuildDuration returns the value of the "build_duration" field in the mutation.
func (m *DbPackageMutation) BuildDuration() (r uint64, exists bool) {
	v := m.build_duration
	if v == nil {
		return
	}
	return *v, true
}

// OldBuildDuration returns the old "build_duration" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldBuildDuration(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldBuildDuration is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldBuildDuration requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBuildDuration: %w", err)
	}
	return oldValue.BuildDuration, nil
}

// AddBuildDuration adds u to the "build_duration" field.
func (m *DbPackageMutation) AddBuildDuration(u uint64) {
	if m.addbuild_duration != nil {
		*m.addbuild_duration += u
	} else {
		m.addbuild_duration = &u
	}
}

// AddedBuildDuration returns the value that was added to the "build_duration" field in this mutation.
func (m *DbPackageMutation) AddedBuildDuration() (r uint64, exists bool) {
	v := m.addbuild_duration
	if v == nil {
		return
	}
	return *v, true
}

// ClearBuildDuration clears the value of the "build_duration" field.
func (m *DbPackageMutation) ClearBuildDuration() {
	m.build_duration = nil
	m.addbuild_duration = nil
	m.clearedFields[dbpackage.FieldBuildDuration] = struct{}{}
}

// BuildDurationCleared returns if the "build_duration" field was cleared in this mutation.
func (m *DbPackageMutation) BuildDurationCleared() bool {
	_, ok := m.clearedFields[dbpackage.FieldBuildDuration]
	return ok
}

// ResetBuildDuration resets all changes to the "build_duration" field.
func (m *DbPackageMutation) ResetBuildDuration() {
	m.build_duration = nil
	m.addbuild_duration = nil
	delete(m.clearedFields, dbpackage.FieldBuildDuration)
}

// SetUpdated sets the "updated" field.
func (m *DbPackageMutation) SetUpdated(t time.Time) {
	m.updated = &t
}

// Updated returns the value of the "updated" field in the mutation.
func (m *DbPackageMutation) Updated() (r time.Time, exists bool) {
	v := m.updated
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdated returns the old "updated" field's value of the DbPackage entity.
// If the DbPackage object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DbPackageMutation) OldUpdated(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdated is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdated requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdated: %w", err)
	}
	return oldValue.Updated, nil
}

// ClearUpdated clears the value of the "updated" field.
func (m *DbPackageMutation) ClearUpdated() {
	m.updated = nil
	m.clearedFields[dbpackage.FieldUpdated] = struct{}{}
}

// UpdatedCleared returns if the "updated" field was cleared in this mutation.
func (m *DbPackageMutation) UpdatedCleared() bool {
	_, ok := m.clearedFields[dbpackage.FieldUpdated]
	return ok
}

// ResetUpdated resets all changes to the "updated" field.
func (m *DbPackageMutation) ResetUpdated() {
	m.updated = nil
	delete(m.clearedFields, dbpackage.FieldUpdated)
}

// Op returns the operation name.
func (m *DbPackageMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (DbPackage).
func (m *DbPackageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DbPackageMutation) Fields() []string {
	fields := make([]string, 0, 11)
	if m.pkgbase != nil {
		fields = append(fields, dbpackage.FieldPkgbase)
	}
	if m.packages != nil {
		fields = append(fields, dbpackage.FieldPackages)
	}
	if m.status != nil {
		fields = append(fields, dbpackage.FieldStatus)
	}
	if m.skip_reason != nil {
		fields = append(fields, dbpackage.FieldSkipReason)
	}
	if m.repository != nil {
		fields = append(fields, dbpackage.FieldRepository)
	}
	if m.march != nil {
		fields = append(fields, dbpackage.FieldMarch)
	}
	if m.version != nil {
		fields = append(fields, dbpackage.FieldVersion)
	}
	if m.repo_version != nil {
		fields = append(fields, dbpackage.FieldRepoVersion)
	}
	if m.build_time != nil {
		fields = append(fields, dbpackage.FieldBuildTime)
	}
	if m.build_duration != nil {
		fields = append(fields, dbpackage.FieldBuildDuration)
	}
	if m.updated != nil {
		fields = append(fields, dbpackage.FieldUpdated)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DbPackageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case dbpackage.FieldPkgbase:
		return m.Pkgbase()
	case dbpackage.FieldPackages:
		return m.Packages()
	case dbpackage.FieldStatus:
		return m.Status()
	case dbpackage.FieldSkipReason:
		return m.SkipReason()
	case dbpackage.FieldRepository:
		return m.Repository()
	case dbpackage.FieldMarch:
		return m.March()
	case dbpackage.FieldVersion:
		return m.Version()
	case dbpackage.FieldRepoVersion:
		return m.RepoVersion()
	case dbpackage.FieldBuildTime:
		return m.BuildTime()
	case dbpackage.FieldBuildDuration:
		return m.BuildDuration()
	case dbpackage.FieldUpdated:
		return m.Updated()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DbPackageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case dbpackage.FieldPkgbase:
		return m.OldPkgbase(ctx)
	case dbpackage.FieldPackages:
		return m.OldPackages(ctx)
	case dbpackage.FieldStatus:
		return m.OldStatus(ctx)
	case dbpackage.FieldSkipReason:
		return m.OldSkipReason(ctx)
	case dbpackage.FieldRepository:
		return m.OldRepository(ctx)
	case dbpackage.FieldMarch:
		return m.OldMarch(ctx)
	case dbpackage.FieldVersion:
		return m.OldVersion(ctx)
	case dbpackage.FieldRepoVersion:
		return m.OldRepoVersion(ctx)
	case dbpackage.FieldBuildTime:
		return m.OldBuildTime(ctx)
	case dbpackage.FieldBuildDuration:
		return m.OldBuildDuration(ctx)
	case dbpackage.FieldUpdated:
		return m.OldUpdated(ctx)
	}
	return nil, fmt.Errorf("unknown DbPackage field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DbPackageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case dbpackage.FieldPkgbase:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPkgbase(v)
		return nil
	case dbpackage.FieldPackages:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPackages(v)
		return nil
	case dbpackage.FieldStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case dbpackage.FieldSkipReason:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSkipReason(v)
		return nil
	case dbpackage.FieldRepository:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRepository(v)
		return nil
	case dbpackage.FieldMarch:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMarch(v)
		return nil
	case dbpackage.FieldVersion:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetVersion(v)
		return nil
	case dbpackage.FieldRepoVersion:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRepoVersion(v)
		return nil
	case dbpackage.FieldBuildTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBuildTime(v)
		return nil
	case dbpackage.FieldBuildDuration:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBuildDuration(v)
		return nil
	case dbpackage.FieldUpdated:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdated(v)
		return nil
	}
	return fmt.Errorf("unknown DbPackage field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DbPackageMutation) AddedFields() []string {
	var fields []string
	if m.addstatus != nil {
		fields = append(fields, dbpackage.FieldStatus)
	}
	if m.addbuild_duration != nil {
		fields = append(fields, dbpackage.FieldBuildDuration)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DbPackageMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case dbpackage.FieldStatus:
		return m.AddedStatus()
	case dbpackage.FieldBuildDuration:
		return m.AddedBuildDuration()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DbPackageMutation) AddField(name string, value ent.Value) error {
	switch name {
	case dbpackage.FieldStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStatus(v)
		return nil
	case dbpackage.FieldBuildDuration:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddBuildDuration(v)
		return nil
	}
	return fmt.Errorf("unknown DbPackage numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DbPackageMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(dbpackage.FieldPackages) {
		fields = append(fields, dbpackage.FieldPackages)
	}
	if m.FieldCleared(dbpackage.FieldSkipReason) {
		fields = append(fields, dbpackage.FieldSkipReason)
	}
	if m.FieldCleared(dbpackage.FieldVersion) {
		fields = append(fields, dbpackage.FieldVersion)
	}
	if m.FieldCleared(dbpackage.FieldRepoVersion) {
		fields = append(fields, dbpackage.FieldRepoVersion)
	}
	if m.FieldCleared(dbpackage.FieldBuildTime) {
		fields = append(fields, dbpackage.FieldBuildTime)
	}
	if m.FieldCleared(dbpackage.FieldBuildDuration) {
		fields = append(fields, dbpackage.FieldBuildDuration)
	}
	if m.FieldCleared(dbpackage.FieldUpdated) {
		fields = append(fields, dbpackage.FieldUpdated)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DbPackageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DbPackageMutation) ClearField(name string) error {
	switch name {
	case dbpackage.FieldPackages:
		m.ClearPackages()
		return nil
	case dbpackage.FieldSkipReason:
		m.ClearSkipReason()
		return nil
	case dbpackage.FieldVersion:
		m.ClearVersion()
		return nil
	case dbpackage.FieldRepoVersion:
		m.ClearRepoVersion()
		return nil
	case dbpackage.FieldBuildTime:
		m.ClearBuildTime()
		return nil
	case dbpackage.FieldBuildDuration:
		m.ClearBuildDuration()
		return nil
	case dbpackage.FieldUpdated:
		m.ClearUpdated()
		return nil
	}
	return fmt.Errorf("unknown DbPackage nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DbPackageMutation) ResetField(name string) error {
	switch name {
	case dbpackage.FieldPkgbase:
		m.ResetPkgbase()
		return nil
	case dbpackage.FieldPackages:
		m.ResetPackages()
		return nil
	case dbpackage.FieldStatus:
		m.ResetStatus()
		return nil
	case dbpackage.FieldSkipReason:
		m.ResetSkipReason()
		return nil
	case dbpackage.FieldRepository:
		m.ResetRepository()
		return nil
	case dbpackage.FieldMarch:
		m.ResetMarch()
		return nil
	case dbpackage.FieldVersion:
		m.ResetVersion()
		return nil
	case dbpackage.FieldRepoVersion:
		m.ResetRepoVersion()
		return nil
	case dbpackage.FieldBuildTime:
		m.ResetBuildTime()
		return nil
	case dbpackage.FieldBuildDuration:
		m.ResetBuildDuration()
		return nil
	case dbpackage.FieldUpdated:
		m.ResetUpdated()
		return nil
	}
	return fmt.Errorf("unknown DbPackage field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DbPackageMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DbPackageMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DbPackageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DbPackageMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DbPackageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DbPackageMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DbPackageMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown DbPackage unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DbPackageMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown DbPackage edge %s", name)
}