// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ALHP.go/ent/dbpackage"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DbPackageCreate is the builder for creating a DbPackage entity.
type DbPackageCreate struct {
	config
	mutation *DbPackageMutation
	hooks    []Hook
}

// SetPkgbase sets the "pkgbase" field.
func (dpc *DbPackageCreate) SetPkgbase(s string) *DbPackageCreate {
	dpc.mutation.SetPkgbase(s)
	return dpc
}

// SetPackages sets the "packages" field.
func (dpc *DbPackageCreate) SetPackages(s []string) *DbPackageCreate {
	dpc.mutation.SetPackages(s)
	return dpc
}

// SetStatus sets the "status" field.
func (dpc *DbPackageCreate) SetStatus(d dbpackage.Status) *DbPackageCreate {
	dpc.mutation.SetStatus(d)
	return dpc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dpc *DbPackageCreate) SetNillableStatus(d *dbpackage.Status) *DbPackageCreate {
	if d != nil {
		dpc.SetStatus(*d)
	}
	return dpc
}

// SetSkipReason sets the "skip_reason" field.
func (dpc *DbPackageCreate) SetSkipReason(s string) *DbPackageCreate {
	dpc.mutation.SetSkipReason(s)
	return dpc
}

// SetNillableSkipReason sets the "skip_reason" field if the given value is not nil.
func (dpc *DbPackageCreate) SetNillableSkipReason(s *string) *DbPackageCreate {
	if s != nil {
		dpc.SetSkipReason(*s)
	}
	return dpc
}

// SetRepository sets the "repository" field.
func (dpc *DbPackageCreate) SetRepository(d dbpackage.Repository) *DbPackageCreate {
	dpc.mutation.SetRepository(d)
	return dpc
}

// SetMarch sets the "march" field.
func (dpc *DbPackageCreate) SetMarch(s string) *DbPackageCreate {
	dpc.mutation.SetMarch(s)
	return dpc
}

// SetVersion sets the "version" field.
func (dpc *DbPackageCreate) SetVersion(s string) *DbPackageCreate {
	dpc.mutation.SetVersion(s)
	return dpc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (dpc *DbPackageCreate) SetNillableVersion(s *string) *DbPackageCreate {
	if s != nil {
		dpc.SetVersion(*s)
	}
	return dpc
}

// SetRepoVersion sets the "repo_version" field.
func (dpc *DbPackageCreate) SetRepoVersion(s string) *DbPackageCreate {
	dpc.mutation.SetRepoVersion(s)
	return dpc
}

// SetNillableRepoVersion sets the "repo_version" field if the given value is not nil.
func (dpc *DbPackageCreate) SetNillableRepoVersion(s *string) *DbPackageCreate {
	if s != nil {
		dpc.SetRepoVersion(*s)
	}
	return dpc
}

// SetBuildTimeStart sets the "build_time_start" field.
func (dpc *DbPackageCreate) SetBuildTimeStart(t time.Time) *DbPackageCreate {
	dpc.mutation.SetBuildTimeStart(t)
	return dpc
}

// SetNillableBuildTimeStart sets the "build_time_start" field if the given value is not nil.
func (dpc *DbPackageCreate) SetNillableBuildTimeStart(t *time.Time) *DbPackageCreate {
	if t != nil {
		dpc.SetBuildTimeStart(*t)
	}
	return dpc
}

// SetBuildTimeEnd sets the "build_time_end" field.
func (dpc *DbPackageCreate) SetBuildTimeEnd(t time.Time) *DbPackageCreate {
	dpc.mutation.SetBuildTimeEnd(t)
	return dpc
}

// SetNillableBuildTimeEnd sets the "build_time_end" field if the given value is not nil.
func (dpc *DbPackageCreate) SetNillableBuildTimeEnd(t *time.Time) *DbPackageCreate {
	if t != nil {
		dpc.SetBuildTimeEnd(*t)
	}
	return dpc
}

// SetUpdated sets the "updated" field.
func (dpc *DbPackageCreate) SetUpdated(t time.Time) *DbPackageCreate {
	dpc.mutation.SetUpdated(t)
	return dpc
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (dpc *DbPackageCreate) SetNillableUpdated(t *time.Time) *DbPackageCreate {
	if t != nil {
		dpc.SetUpdated(*t)
	}
	return dpc
}

// SetHash sets the "hash" field.
func (dpc *DbPackageCreate) SetHash(s string) *DbPackageCreate {
	dpc.mutation.SetHash(s)
	return dpc
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (dpc *DbPackageCreate) SetNillableHash(s *string) *DbPackageCreate {
	if s != nil {
		dpc.SetHash(*s)
	}
	return dpc
}

// SetLto sets the "lto" field.
func (dpc *DbPackageCreate) SetLto(d dbpackage.Lto) *DbPackageCreate {
	dpc.mutation.SetLto(d)
	return dpc
}

// SetNillableLto sets the "lto" field if the given value is not nil.
func (dpc *DbPackageCreate) SetNillableLto(d *dbpackage.Lto) *DbPackageCreate {
	if d != nil {
		dpc.SetLto(*d)
	}
	return dpc
}

// Mutation returns the DbPackageMutation object of the builder.
func (dpc *DbPackageCreate) Mutation() *DbPackageMutation {
	return dpc.mutation
}

// Save creates the DbPackage in the database.
func (dpc *DbPackageCreate) Save(ctx context.Context) (*DbPackage, error) {
	var (
		err  error
		node *DbPackage
	)
	dpc.defaults()
	if len(dpc.hooks) == 0 {
		if err = dpc.check(); err != nil {
			return nil, err
		}
		node, err = dpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DbPackageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dpc.check(); err != nil {
				return nil, err
			}
			dpc.mutation = mutation
			if node, err = dpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dpc.hooks) - 1; i >= 0; i-- {
			if dpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dpc *DbPackageCreate) SaveX(ctx context.Context) *DbPackage {
	v, err := dpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dpc *DbPackageCreate) Exec(ctx context.Context) error {
	_, err := dpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpc *DbPackageCreate) ExecX(ctx context.Context) {
	if err := dpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dpc *DbPackageCreate) defaults() {
	if _, ok := dpc.mutation.Status(); !ok {
		v := dbpackage.DefaultStatus
		dpc.mutation.SetStatus(v)
	}
	if _, ok := dpc.mutation.Lto(); !ok {
		v := dbpackage.DefaultLto
		dpc.mutation.SetLto(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dpc *DbPackageCreate) check() error {
	if _, ok := dpc.mutation.Pkgbase(); !ok {
		return &ValidationError{Name: "pkgbase", err: errors.New(`ent: missing required field "pkgbase"`)}
	}
	if v, ok := dpc.mutation.Pkgbase(); ok {
		if err := dbpackage.PkgbaseValidator(v); err != nil {
			return &ValidationError{Name: "pkgbase", err: fmt.Errorf(`ent: validator failed for field "pkgbase": %w`, err)}
		}
	}
	if v, ok := dpc.mutation.Status(); ok {
		if err := dbpackage.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if _, ok := dpc.mutation.Repository(); !ok {
		return &ValidationError{Name: "repository", err: errors.New(`ent: missing required field "repository"`)}
	}
	if v, ok := dpc.mutation.Repository(); ok {
		if err := dbpackage.RepositoryValidator(v); err != nil {
			return &ValidationError{Name: "repository", err: fmt.Errorf(`ent: validator failed for field "repository": %w`, err)}
		}
	}
	if _, ok := dpc.mutation.March(); !ok {
		return &ValidationError{Name: "march", err: errors.New(`ent: missing required field "march"`)}
	}
	if v, ok := dpc.mutation.March(); ok {
		if err := dbpackage.MarchValidator(v); err != nil {
			return &ValidationError{Name: "march", err: fmt.Errorf(`ent: validator failed for field "march": %w`, err)}
		}
	}
	if v, ok := dpc.mutation.Lto(); ok {
		if err := dbpackage.LtoValidator(v); err != nil {
			return &ValidationError{Name: "lto", err: fmt.Errorf(`ent: validator failed for field "lto": %w`, err)}
		}
	}
	return nil
}

func (dpc *DbPackageCreate) sqlSave(ctx context.Context) (*DbPackage, error) {
	_node, _spec := dpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dpc *DbPackageCreate) createSpec() (*DbPackage, *sqlgraph.CreateSpec) {
	var (
		_node = &DbPackage{config: dpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dbpackage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dbpackage.FieldID,
			},
		}
	)
	if value, ok := dpc.mutation.Pkgbase(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldPkgbase,
		})
		_node.Pkgbase = value
	}
	if value, ok := dpc.mutation.Packages(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dbpackage.FieldPackages,
		})
		_node.Packages = value
	}
	if value, ok := dpc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: dbpackage.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := dpc.mutation.SkipReason(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldSkipReason,
		})
		_node.SkipReason = value
	}
	if value, ok := dpc.mutation.Repository(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: dbpackage.FieldRepository,
		})
		_node.Repository = value
	}
	if value, ok := dpc.mutation.March(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldMarch,
		})
		_node.March = value
	}
	if value, ok := dpc.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldVersion,
		})
		_node.Version = value
	}
	if value, ok := dpc.mutation.RepoVersion(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldRepoVersion,
		})
		_node.RepoVersion = value
	}
	if value, ok := dpc.mutation.BuildTimeStart(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldBuildTimeStart,
		})
		_node.BuildTimeStart = value
	}
	if value, ok := dpc.mutation.BuildTimeEnd(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldBuildTimeEnd,
		})
		_node.BuildTimeEnd = value
	}
	if value, ok := dpc.mutation.Updated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dbpackage.FieldUpdated,
		})
		_node.Updated = value
	}
	if value, ok := dpc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dbpackage.FieldHash,
		})
		_node.Hash = value
	}
	if value, ok := dpc.mutation.Lto(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: dbpackage.FieldLto,
		})
		_node.Lto = value
	}
	return _node, _spec
}

// DbPackageCreateBulk is the builder for creating many DbPackage entities in bulk.
type DbPackageCreateBulk struct {
	config
	builders []*DbPackageCreate
}

// Save creates the DbPackage entities in the database.
func (dpcb *DbPackageCreateBulk) Save(ctx context.Context) ([]*DbPackage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dpcb.builders))
	nodes := make([]*DbPackage, len(dpcb.builders))
	mutators := make([]Mutator, len(dpcb.builders))
	for i := range dpcb.builders {
		func(i int, root context.Context) {
			builder := dpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DbPackageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dpcb *DbPackageCreateBulk) SaveX(ctx context.Context) []*DbPackage {
	v, err := dpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dpcb *DbPackageCreateBulk) Exec(ctx context.Context) error {
	_, err := dpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpcb *DbPackageCreateBulk) ExecX(ctx context.Context) {
	if err := dpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
