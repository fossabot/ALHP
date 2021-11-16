// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"ALHP.go/ent/dbpackage"
	"ALHP.go/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DbPackageUpdate is the builder for updating DbPackage entities.
type DbPackageUpdate struct {
	config
	hooks    []Hook
	mutation *DbPackageMutation
}

// Where appends a list predicates to the DbPackageUpdate builder.
func (dpu *DbPackageUpdate) Where(ps ...predicate.DbPackage) *DbPackageUpdate {
	dpu.mutation.Where(ps...)
	return dpu
}

// SetPackages sets the "packages" field.
func (dpu *DbPackageUpdate) SetPackages(s []string) *DbPackageUpdate {
	dpu.mutation.SetPackages(s)
	return dpu
}

// ClearPackages clears the value of the "packages" field.
func (dpu *DbPackageUpdate) ClearPackages() *DbPackageUpdate {
	dpu.mutation.ClearPackages()
	return dpu
}

// SetStatus sets the "status" field.
func (dpu *DbPackageUpdate) SetStatus(d dbpackage.Status) *DbPackageUpdate {
	dpu.mutation.SetStatus(d)
	return dpu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableStatus(d *dbpackage.Status) *DbPackageUpdate {
	if d != nil {
		dpu.SetStatus(*d)
	}
	return dpu
}

// ClearStatus clears the value of the "status" field.
func (dpu *DbPackageUpdate) ClearStatus() *DbPackageUpdate {
	dpu.mutation.ClearStatus()
	return dpu
}

// SetSkipReason sets the "skip_reason" field.
func (dpu *DbPackageUpdate) SetSkipReason(s string) *DbPackageUpdate {
	dpu.mutation.SetSkipReason(s)
	return dpu
}

// SetNillableSkipReason sets the "skip_reason" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableSkipReason(s *string) *DbPackageUpdate {
	if s != nil {
		dpu.SetSkipReason(*s)
	}
	return dpu
}

// ClearSkipReason clears the value of the "skip_reason" field.
func (dpu *DbPackageUpdate) ClearSkipReason() *DbPackageUpdate {
	dpu.mutation.ClearSkipReason()
	return dpu
}

// SetRepository sets the "repository" field.
func (dpu *DbPackageUpdate) SetRepository(d dbpackage.Repository) *DbPackageUpdate {
	dpu.mutation.SetRepository(d)
	return dpu
}

// SetMarch sets the "march" field.
func (dpu *DbPackageUpdate) SetMarch(s string) *DbPackageUpdate {
	dpu.mutation.SetMarch(s)
	return dpu
}

// SetVersion sets the "version" field.
func (dpu *DbPackageUpdate) SetVersion(s string) *DbPackageUpdate {
	dpu.mutation.SetVersion(s)
	return dpu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableVersion(s *string) *DbPackageUpdate {
	if s != nil {
		dpu.SetVersion(*s)
	}
	return dpu
}

// ClearVersion clears the value of the "version" field.
func (dpu *DbPackageUpdate) ClearVersion() *DbPackageUpdate {
	dpu.mutation.ClearVersion()
	return dpu
}

// SetRepoVersion sets the "repo_version" field.
func (dpu *DbPackageUpdate) SetRepoVersion(s string) *DbPackageUpdate {
	dpu.mutation.SetRepoVersion(s)
	return dpu
}

// SetNillableRepoVersion sets the "repo_version" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableRepoVersion(s *string) *DbPackageUpdate {
	if s != nil {
		dpu.SetRepoVersion(*s)
	}
	return dpu
}

// ClearRepoVersion clears the value of the "repo_version" field.
func (dpu *DbPackageUpdate) ClearRepoVersion() *DbPackageUpdate {
	dpu.mutation.ClearRepoVersion()
	return dpu
}

// SetBuildTimeStart sets the "build_time_start" field.
func (dpu *DbPackageUpdate) SetBuildTimeStart(t time.Time) *DbPackageUpdate {
	dpu.mutation.SetBuildTimeStart(t)
	return dpu
}

// SetNillableBuildTimeStart sets the "build_time_start" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableBuildTimeStart(t *time.Time) *DbPackageUpdate {
	if t != nil {
		dpu.SetBuildTimeStart(*t)
	}
	return dpu
}

// ClearBuildTimeStart clears the value of the "build_time_start" field.
func (dpu *DbPackageUpdate) ClearBuildTimeStart() *DbPackageUpdate {
	dpu.mutation.ClearBuildTimeStart()
	return dpu
}

// SetBuildTimeEnd sets the "build_time_end" field.
func (dpu *DbPackageUpdate) SetBuildTimeEnd(t time.Time) *DbPackageUpdate {
	dpu.mutation.SetBuildTimeEnd(t)
	return dpu
}

// SetNillableBuildTimeEnd sets the "build_time_end" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableBuildTimeEnd(t *time.Time) *DbPackageUpdate {
	if t != nil {
		dpu.SetBuildTimeEnd(*t)
	}
	return dpu
}

// ClearBuildTimeEnd clears the value of the "build_time_end" field.
func (dpu *DbPackageUpdate) ClearBuildTimeEnd() *DbPackageUpdate {
	dpu.mutation.ClearBuildTimeEnd()
	return dpu
}

// SetUpdated sets the "updated" field.
func (dpu *DbPackageUpdate) SetUpdated(t time.Time) *DbPackageUpdate {
	dpu.mutation.SetUpdated(t)
	return dpu
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableUpdated(t *time.Time) *DbPackageUpdate {
	if t != nil {
		dpu.SetUpdated(*t)
	}
	return dpu
}

// ClearUpdated clears the value of the "updated" field.
func (dpu *DbPackageUpdate) ClearUpdated() *DbPackageUpdate {
	dpu.mutation.ClearUpdated()
	return dpu
}

// SetHash sets the "hash" field.
func (dpu *DbPackageUpdate) SetHash(s string) *DbPackageUpdate {
	dpu.mutation.SetHash(s)
	return dpu
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableHash(s *string) *DbPackageUpdate {
	if s != nil {
		dpu.SetHash(*s)
	}
	return dpu
}

// ClearHash clears the value of the "hash" field.
func (dpu *DbPackageUpdate) ClearHash() *DbPackageUpdate {
	dpu.mutation.ClearHash()
	return dpu
}

// SetLto sets the "lto" field.
func (dpu *DbPackageUpdate) SetLto(d dbpackage.Lto) *DbPackageUpdate {
	dpu.mutation.SetLto(d)
	return dpu
}

// SetNillableLto sets the "lto" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableLto(d *dbpackage.Lto) *DbPackageUpdate {
	if d != nil {
		dpu.SetLto(*d)
	}
	return dpu
}

// ClearLto clears the value of the "lto" field.
func (dpu *DbPackageUpdate) ClearLto() *DbPackageUpdate {
	dpu.mutation.ClearLto()
	return dpu
}

// Mutation returns the DbPackageMutation object of the builder.
func (dpu *DbPackageUpdate) Mutation() *DbPackageMutation {
	return dpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dpu *DbPackageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dpu.hooks) == 0 {
		if err = dpu.check(); err != nil {
			return 0, err
		}
		affected, err = dpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbPackageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dpu.check(); err != nil {
				return 0, err
			}
			dpu.mutation = mutation
			affected, err = dpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dpu.hooks) - 1; i >= 0; i-- {
			if dpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dpu *DbPackageUpdate) SaveX(ctx context.Context) int {
	affected, err := dpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dpu *DbPackageUpdate) Exec(ctx context.Context) error {
	_, err := dpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpu *DbPackageUpdate) ExecX(ctx context.Context) {
	if err := dpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dpu *DbPackageUpdate) check() error {
	if v, ok := dpu.mutation.Status(); ok {
		if err := dbpackage.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := dpu.mutation.Repository(); ok {
		if err := dbpackage.RepositoryValidator(v); err != nil {
			return &ValidationError{Name: "repository", err: fmt.Errorf("ent: validator failed for field \"repository\": %w", err)}
		}
	}
	if v, ok := dpu.mutation.March(); ok {
		if err := dbpackage.MarchValidator(v); err != nil {
			return &ValidationError{Name: "march", err: fmt.Errorf("ent: validator failed for field \"march\": %w", err)}
		}
	}
	if v, ok := dpu.mutation.Lto(); ok {
		if err := dbpackage.LtoValidator(v); err != nil {
			return &ValidationError{Name: "lto", err: fmt.Errorf("ent: validator failed for field \"lto\": %w", err)}
		}
	}
	return nil
}

func (dpu *DbPackageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbpackage.Table,
			Columns: dbpackage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbpackage.FieldID,
			},
		},
	}
	if ps := dpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dpu.mutation.Packages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbpackage.FieldPackages,
		})
	}
	if dpu.mutation.PackagesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: dbpackage.FieldPackages,
		})
	}
	if value, ok := dpu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: dbpackage.FieldStatus,
		})
	}
	if dpu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: dbpackage.FieldStatus,
		})
	}
	if value, ok := dpu.mutation.SkipReason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldSkipReason,
		})
	}
	if dpu.mutation.SkipReasonCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: dbpackage.FieldSkipReason,
		})
	}
	if value, ok := dpu.mutation.Repository(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: dbpackage.FieldRepository,
		})
	}
	if value, ok := dpu.mutation.March(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldMarch,
		})
	}
	if value, ok := dpu.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldVersion,
		})
	}
	if dpu.mutation.VersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: dbpackage.FieldVersion,
		})
	}
	if value, ok := dpu.mutation.RepoVersion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldRepoVersion,
		})
	}
	if dpu.mutation.RepoVersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: dbpackage.FieldRepoVersion,
		})
	}
	if value, ok := dpu.mutation.BuildTimeStart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldBuildTimeStart,
		})
	}
	if dpu.mutation.BuildTimeStartCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: dbpackage.FieldBuildTimeStart,
		})
	}
	if value, ok := dpu.mutation.BuildTimeEnd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldBuildTimeEnd,
		})
	}
	if dpu.mutation.BuildTimeEndCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: dbpackage.FieldBuildTimeEnd,
		})
	}
	if value, ok := dpu.mutation.Updated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldUpdated,
		})
	}
	if dpu.mutation.UpdatedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: dbpackage.FieldUpdated,
		})
	}
	if value, ok := dpu.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldHash,
		})
	}
	if dpu.mutation.HashCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: dbpackage.FieldHash,
		})
	}
	if value, ok := dpu.mutation.Lto(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: dbpackage.FieldLto,
		})
	}
	if dpu.mutation.LtoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: dbpackage.FieldLto,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dbpackage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DbPackageUpdateOne is the builder for updating a single DbPackage entity.
type DbPackageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DbPackageMutation
}

// SetPackages sets the "packages" field.
func (dpuo *DbPackageUpdateOne) SetPackages(s []string) *DbPackageUpdateOne {
	dpuo.mutation.SetPackages(s)
	return dpuo
}

// ClearPackages clears the value of the "packages" field.
func (dpuo *DbPackageUpdateOne) ClearPackages() *DbPackageUpdateOne {
	dpuo.mutation.ClearPackages()
	return dpuo
}

// SetStatus sets the "status" field.
func (dpuo *DbPackageUpdateOne) SetStatus(d dbpackage.Status) *DbPackageUpdateOne {
	dpuo.mutation.SetStatus(d)
	return dpuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableStatus(d *dbpackage.Status) *DbPackageUpdateOne {
	if d != nil {
		dpuo.SetStatus(*d)
	}
	return dpuo
}

// ClearStatus clears the value of the "status" field.
func (dpuo *DbPackageUpdateOne) ClearStatus() *DbPackageUpdateOne {
	dpuo.mutation.ClearStatus()
	return dpuo
}

// SetSkipReason sets the "skip_reason" field.
func (dpuo *DbPackageUpdateOne) SetSkipReason(s string) *DbPackageUpdateOne {
	dpuo.mutation.SetSkipReason(s)
	return dpuo
}

// SetNillableSkipReason sets the "skip_reason" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableSkipReason(s *string) *DbPackageUpdateOne {
	if s != nil {
		dpuo.SetSkipReason(*s)
	}
	return dpuo
}

// ClearSkipReason clears the value of the "skip_reason" field.
func (dpuo *DbPackageUpdateOne) ClearSkipReason() *DbPackageUpdateOne {
	dpuo.mutation.ClearSkipReason()
	return dpuo
}

// SetRepository sets the "repository" field.
func (dpuo *DbPackageUpdateOne) SetRepository(d dbpackage.Repository) *DbPackageUpdateOne {
	dpuo.mutation.SetRepository(d)
	return dpuo
}

// SetMarch sets the "march" field.
func (dpuo *DbPackageUpdateOne) SetMarch(s string) *DbPackageUpdateOne {
	dpuo.mutation.SetMarch(s)
	return dpuo
}

// SetVersion sets the "version" field.
func (dpuo *DbPackageUpdateOne) SetVersion(s string) *DbPackageUpdateOne {
	dpuo.mutation.SetVersion(s)
	return dpuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableVersion(s *string) *DbPackageUpdateOne {
	if s != nil {
		dpuo.SetVersion(*s)
	}
	return dpuo
}

// ClearVersion clears the value of the "version" field.
func (dpuo *DbPackageUpdateOne) ClearVersion() *DbPackageUpdateOne {
	dpuo.mutation.ClearVersion()
	return dpuo
}

// SetRepoVersion sets the "repo_version" field.
func (dpuo *DbPackageUpdateOne) SetRepoVersion(s string) *DbPackageUpdateOne {
	dpuo.mutation.SetRepoVersion(s)
	return dpuo
}

// SetNillableRepoVersion sets the "repo_version" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableRepoVersion(s *string) *DbPackageUpdateOne {
	if s != nil {
		dpuo.SetRepoVersion(*s)
	}
	return dpuo
}

// ClearRepoVersion clears the value of the "repo_version" field.
func (dpuo *DbPackageUpdateOne) ClearRepoVersion() *DbPackageUpdateOne {
	dpuo.mutation.ClearRepoVersion()
	return dpuo
}

// SetBuildTimeStart sets the "build_time_start" field.
func (dpuo *DbPackageUpdateOne) SetBuildTimeStart(t time.Time) *DbPackageUpdateOne {
	dpuo.mutation.SetBuildTimeStart(t)
	return dpuo
}

// SetNillableBuildTimeStart sets the "build_time_start" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableBuildTimeStart(t *time.Time) *DbPackageUpdateOne {
	if t != nil {
		dpuo.SetBuildTimeStart(*t)
	}
	return dpuo
}

// ClearBuildTimeStart clears the value of the "build_time_start" field.
func (dpuo *DbPackageUpdateOne) ClearBuildTimeStart() *DbPackageUpdateOne {
	dpuo.mutation.ClearBuildTimeStart()
	return dpuo
}

// SetBuildTimeEnd sets the "build_time_end" field.
func (dpuo *DbPackageUpdateOne) SetBuildTimeEnd(t time.Time) *DbPackageUpdateOne {
	dpuo.mutation.SetBuildTimeEnd(t)
	return dpuo
}

// SetNillableBuildTimeEnd sets the "build_time_end" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableBuildTimeEnd(t *time.Time) *DbPackageUpdateOne {
	if t != nil {
		dpuo.SetBuildTimeEnd(*t)
	}
	return dpuo
}

// ClearBuildTimeEnd clears the value of the "build_time_end" field.
func (dpuo *DbPackageUpdateOne) ClearBuildTimeEnd() *DbPackageUpdateOne {
	dpuo.mutation.ClearBuildTimeEnd()
	return dpuo
}

// SetUpdated sets the "updated" field.
func (dpuo *DbPackageUpdateOne) SetUpdated(t time.Time) *DbPackageUpdateOne {
	dpuo.mutation.SetUpdated(t)
	return dpuo
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableUpdated(t *time.Time) *DbPackageUpdateOne {
	if t != nil {
		dpuo.SetUpdated(*t)
	}
	return dpuo
}

// ClearUpdated clears the value of the "updated" field.
func (dpuo *DbPackageUpdateOne) ClearUpdated() *DbPackageUpdateOne {
	dpuo.mutation.ClearUpdated()
	return dpuo
}

// SetHash sets the "hash" field.
func (dpuo *DbPackageUpdateOne) SetHash(s string) *DbPackageUpdateOne {
	dpuo.mutation.SetHash(s)
	return dpuo
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableHash(s *string) *DbPackageUpdateOne {
	if s != nil {
		dpuo.SetHash(*s)
	}
	return dpuo
}

// ClearHash clears the value of the "hash" field.
func (dpuo *DbPackageUpdateOne) ClearHash() *DbPackageUpdateOne {
	dpuo.mutation.ClearHash()
	return dpuo
}

// SetLto sets the "lto" field.
func (dpuo *DbPackageUpdateOne) SetLto(d dbpackage.Lto) *DbPackageUpdateOne {
	dpuo.mutation.SetLto(d)
	return dpuo
}

// SetNillableLto sets the "lto" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableLto(d *dbpackage.Lto) *DbPackageUpdateOne {
	if d != nil {
		dpuo.SetLto(*d)
	}
	return dpuo
}

// ClearLto clears the value of the "lto" field.
func (dpuo *DbPackageUpdateOne) ClearLto() *DbPackageUpdateOne {
	dpuo.mutation.ClearLto()
	return dpuo
}

// Mutation returns the DbPackageMutation object of the builder.
func (dpuo *DbPackageUpdateOne) Mutation() *DbPackageMutation {
	return dpuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dpuo *DbPackageUpdateOne) Select(field string, fields ...string) *DbPackageUpdateOne {
	dpuo.fields = append([]string{field}, fields...)
	return dpuo
}

// Save executes the query and returns the updated DbPackage entity.
func (dpuo *DbPackageUpdateOne) Save(ctx context.Context) (*DbPackage, error) {
	var (
		err  error
		node *DbPackage
	)
	if len(dpuo.hooks) == 0 {
		if err = dpuo.check(); err != nil {
			return nil, err
		}
		node, err = dpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbPackageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dpuo.check(); err != nil {
				return nil, err
			}
			dpuo.mutation = mutation
			node, err = dpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dpuo.hooks) - 1; i >= 0; i-- {
			if dpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dpuo *DbPackageUpdateOne) SaveX(ctx context.Context) *DbPackage {
	node, err := dpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dpuo *DbPackageUpdateOne) Exec(ctx context.Context) error {
	_, err := dpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpuo *DbPackageUpdateOne) ExecX(ctx context.Context) {
	if err := dpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dpuo *DbPackageUpdateOne) check() error {
	if v, ok := dpuo.mutation.Status(); ok {
		if err := dbpackage.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := dpuo.mutation.Repository(); ok {
		if err := dbpackage.RepositoryValidator(v); err != nil {
			return &ValidationError{Name: "repository", err: fmt.Errorf("ent: validator failed for field \"repository\": %w", err)}
		}
	}
	if v, ok := dpuo.mutation.March(); ok {
		if err := dbpackage.MarchValidator(v); err != nil {
			return &ValidationError{Name: "march", err: fmt.Errorf("ent: validator failed for field \"march\": %w", err)}
		}
	}
	if v, ok := dpuo.mutation.Lto(); ok {
		if err := dbpackage.LtoValidator(v); err != nil {
			return &ValidationError{Name: "lto", err: fmt.Errorf("ent: validator failed for field \"lto\": %w", err)}
		}
	}
	return nil
}

func (dpuo *DbPackageUpdateOne) sqlSave(ctx context.Context) (_node *DbPackage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dbpackage.Table,
			Columns: dbpackage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbpackage.FieldID,
			},
		},
	}
	id, ok := dpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DbPackage.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := dpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dbpackage.FieldID)
		for _, f := range fields {
			if !dbpackage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dbpackage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dpuo.mutation.Packages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbpackage.FieldPackages,
		})
	}
	if dpuo.mutation.PackagesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: dbpackage.FieldPackages,
		})
	}
	if value, ok := dpuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: dbpackage.FieldStatus,
		})
	}
	if dpuo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: dbpackage.FieldStatus,
		})
	}
	if value, ok := dpuo.mutation.SkipReason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldSkipReason,
		})
	}
	if dpuo.mutation.SkipReasonCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: dbpackage.FieldSkipReason,
		})
	}
	if value, ok := dpuo.mutation.Repository(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: dbpackage.FieldRepository,
		})
	}
	if value, ok := dpuo.mutation.March(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldMarch,
		})
	}
	if value, ok := dpuo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldVersion,
		})
	}
	if dpuo.mutation.VersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: dbpackage.FieldVersion,
		})
	}
	if value, ok := dpuo.mutation.RepoVersion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldRepoVersion,
		})
	}
	if dpuo.mutation.RepoVersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: dbpackage.FieldRepoVersion,
		})
	}
	if value, ok := dpuo.mutation.BuildTimeStart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldBuildTimeStart,
		})
	}
	if dpuo.mutation.BuildTimeStartCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: dbpackage.FieldBuildTimeStart,
		})
	}
	if value, ok := dpuo.mutation.BuildTimeEnd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldBuildTimeEnd,
		})
	}
	if dpuo.mutation.BuildTimeEndCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: dbpackage.FieldBuildTimeEnd,
		})
	}
	if value, ok := dpuo.mutation.Updated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldUpdated,
		})
	}
	if dpuo.mutation.UpdatedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: dbpackage.FieldUpdated,
		})
	}
	if value, ok := dpuo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldHash,
		})
	}
	if dpuo.mutation.HashCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: dbpackage.FieldHash,
		})
	}
	if value, ok := dpuo.mutation.Lto(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: dbpackage.FieldLto,
		})
	}
	if dpuo.mutation.LtoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: dbpackage.FieldLto,
		})
	}
	_node = &DbPackage{config: dpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dbpackage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
