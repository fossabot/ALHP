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

// Where adds a new predicate for the DbPackageUpdate builder.
func (dpu *DbPackageUpdate) Where(ps ...predicate.DbPackage) *DbPackageUpdate {
	dpu.mutation.predicates = append(dpu.mutation.predicates, ps...)
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
func (dpu *DbPackageUpdate) SetStatus(i int) *DbPackageUpdate {
	dpu.mutation.ResetStatus()
	dpu.mutation.SetStatus(i)
	return dpu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableStatus(i *int) *DbPackageUpdate {
	if i != nil {
		dpu.SetStatus(*i)
	}
	return dpu
}

// AddStatus adds i to the "status" field.
func (dpu *DbPackageUpdate) AddStatus(i int) *DbPackageUpdate {
	dpu.mutation.AddStatus(i)
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
func (dpu *DbPackageUpdate) SetRepository(s string) *DbPackageUpdate {
	dpu.mutation.SetRepository(s)
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

// SetBuildTime sets the "build_time" field.
func (dpu *DbPackageUpdate) SetBuildTime(t time.Time) *DbPackageUpdate {
	dpu.mutation.SetBuildTime(t)
	return dpu
}

// SetNillableBuildTime sets the "build_time" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableBuildTime(t *time.Time) *DbPackageUpdate {
	if t != nil {
		dpu.SetBuildTime(*t)
	}
	return dpu
}

// ClearBuildTime clears the value of the "build_time" field.
func (dpu *DbPackageUpdate) ClearBuildTime() *DbPackageUpdate {
	dpu.mutation.ClearBuildTime()
	return dpu
}

// SetBuildDuration sets the "build_duration" field.
func (dpu *DbPackageUpdate) SetBuildDuration(u uint64) *DbPackageUpdate {
	dpu.mutation.ResetBuildDuration()
	dpu.mutation.SetBuildDuration(u)
	return dpu
}

// SetNillableBuildDuration sets the "build_duration" field if the given value is not nil.
func (dpu *DbPackageUpdate) SetNillableBuildDuration(u *uint64) *DbPackageUpdate {
	if u != nil {
		dpu.SetBuildDuration(*u)
	}
	return dpu
}

// AddBuildDuration adds u to the "build_duration" field.
func (dpu *DbPackageUpdate) AddBuildDuration(u uint64) *DbPackageUpdate {
	dpu.mutation.AddBuildDuration(u)
	return dpu
}

// ClearBuildDuration clears the value of the "build_duration" field.
func (dpu *DbPackageUpdate) ClearBuildDuration() *DbPackageUpdate {
	dpu.mutation.ClearBuildDuration()
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
	if v, ok := dpu.mutation.BuildDuration(); ok {
		if err := dbpackage.BuildDurationValidator(v); err != nil {
			return &ValidationError{Name: "build_duration", err: fmt.Errorf("ent: validator failed for field \"build_duration\": %w", err)}
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
			Type:   field.TypeInt,
			Value:  value,
			Column: dbpackage.FieldStatus,
		})
	}
	if value, ok := dpu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
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
			Type:   field.TypeString,
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
	if value, ok := dpu.mutation.BuildTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldBuildTime,
		})
	}
	if dpu.mutation.BuildTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: dbpackage.FieldBuildTime,
		})
	}
	if value, ok := dpu.mutation.BuildDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: dbpackage.FieldBuildDuration,
		})
	}
	if value, ok := dpu.mutation.AddedBuildDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: dbpackage.FieldBuildDuration,
		})
	}
	if dpu.mutation.BuildDurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: dbpackage.FieldBuildDuration,
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
	if n, err = sqlgraph.UpdateNodes(ctx, dpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dbpackage.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
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
func (dpuo *DbPackageUpdateOne) SetStatus(i int) *DbPackageUpdateOne {
	dpuo.mutation.ResetStatus()
	dpuo.mutation.SetStatus(i)
	return dpuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableStatus(i *int) *DbPackageUpdateOne {
	if i != nil {
		dpuo.SetStatus(*i)
	}
	return dpuo
}

// AddStatus adds i to the "status" field.
func (dpuo *DbPackageUpdateOne) AddStatus(i int) *DbPackageUpdateOne {
	dpuo.mutation.AddStatus(i)
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
func (dpuo *DbPackageUpdateOne) SetRepository(s string) *DbPackageUpdateOne {
	dpuo.mutation.SetRepository(s)
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

// SetBuildTime sets the "build_time" field.
func (dpuo *DbPackageUpdateOne) SetBuildTime(t time.Time) *DbPackageUpdateOne {
	dpuo.mutation.SetBuildTime(t)
	return dpuo
}

// SetNillableBuildTime sets the "build_time" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableBuildTime(t *time.Time) *DbPackageUpdateOne {
	if t != nil {
		dpuo.SetBuildTime(*t)
	}
	return dpuo
}

// ClearBuildTime clears the value of the "build_time" field.
func (dpuo *DbPackageUpdateOne) ClearBuildTime() *DbPackageUpdateOne {
	dpuo.mutation.ClearBuildTime()
	return dpuo
}

// SetBuildDuration sets the "build_duration" field.
func (dpuo *DbPackageUpdateOne) SetBuildDuration(u uint64) *DbPackageUpdateOne {
	dpuo.mutation.ResetBuildDuration()
	dpuo.mutation.SetBuildDuration(u)
	return dpuo
}

// SetNillableBuildDuration sets the "build_duration" field if the given value is not nil.
func (dpuo *DbPackageUpdateOne) SetNillableBuildDuration(u *uint64) *DbPackageUpdateOne {
	if u != nil {
		dpuo.SetBuildDuration(*u)
	}
	return dpuo
}

// AddBuildDuration adds u to the "build_duration" field.
func (dpuo *DbPackageUpdateOne) AddBuildDuration(u uint64) *DbPackageUpdateOne {
	dpuo.mutation.AddBuildDuration(u)
	return dpuo
}

// ClearBuildDuration clears the value of the "build_duration" field.
func (dpuo *DbPackageUpdateOne) ClearBuildDuration() *DbPackageUpdateOne {
	dpuo.mutation.ClearBuildDuration()
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
	if v, ok := dpuo.mutation.BuildDuration(); ok {
		if err := dbpackage.BuildDurationValidator(v); err != nil {
			return &ValidationError{Name: "build_duration", err: fmt.Errorf("ent: validator failed for field \"build_duration\": %w", err)}
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
			Type:   field.TypeInt,
			Value:  value,
			Column: dbpackage.FieldStatus,
		})
	}
	if value, ok := dpuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
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
			Type:   field.TypeString,
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
	if value, ok := dpuo.mutation.BuildTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldBuildTime,
		})
	}
	if dpuo.mutation.BuildTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: dbpackage.FieldBuildTime,
		})
	}
	if value, ok := dpuo.mutation.BuildDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: dbpackage.FieldBuildDuration,
		})
	}
	if value, ok := dpuo.mutation.AddedBuildDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: dbpackage.FieldBuildDuration,
		})
	}
	if dpuo.mutation.BuildDurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: dbpackage.FieldBuildDuration,
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
	_node = &DbPackage{config: dpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dbpackage.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}