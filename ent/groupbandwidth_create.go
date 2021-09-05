// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/mrz-identity-management/ent/groupbandwidth"
	"github.com/mrzack99s/mrz-identity-management/ent/groups"
)

// GroupBandwidthCreate is the builder for creating a GroupBandwidth entity.
type GroupBandwidthCreate struct {
	config
	mutation *GroupBandwidthMutation
	hooks    []Hook
}

// SetGbwDownloadSpeed sets the "gbw_download_speed" field.
func (gbc *GroupBandwidthCreate) SetGbwDownloadSpeed(i int) *GroupBandwidthCreate {
	gbc.mutation.SetGbwDownloadSpeed(i)
	return gbc
}

// SetGbwUploadSpeed sets the "gbw_upload_speed" field.
func (gbc *GroupBandwidthCreate) SetGbwUploadSpeed(i int) *GroupBandwidthCreate {
	gbc.mutation.SetGbwUploadSpeed(i)
	return gbc
}

// SetGbwCreatedAt sets the "gbw_created_at" field.
func (gbc *GroupBandwidthCreate) SetGbwCreatedAt(t time.Time) *GroupBandwidthCreate {
	gbc.mutation.SetGbwCreatedAt(t)
	return gbc
}

// SetNillableGbwCreatedAt sets the "gbw_created_at" field if the given value is not nil.
func (gbc *GroupBandwidthCreate) SetNillableGbwCreatedAt(t *time.Time) *GroupBandwidthCreate {
	if t != nil {
		gbc.SetGbwCreatedAt(*t)
	}
	return gbc
}

// AddGroupIDs adds the "groups" edge to the Groups entity by IDs.
func (gbc *GroupBandwidthCreate) AddGroupIDs(ids ...int) *GroupBandwidthCreate {
	gbc.mutation.AddGroupIDs(ids...)
	return gbc
}

// AddGroups adds the "groups" edges to the Groups entity.
func (gbc *GroupBandwidthCreate) AddGroups(g ...*Groups) *GroupBandwidthCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gbc.AddGroupIDs(ids...)
}

// Mutation returns the GroupBandwidthMutation object of the builder.
func (gbc *GroupBandwidthCreate) Mutation() *GroupBandwidthMutation {
	return gbc.mutation
}

// Save creates the GroupBandwidth in the database.
func (gbc *GroupBandwidthCreate) Save(ctx context.Context) (*GroupBandwidth, error) {
	var (
		err  error
		node *GroupBandwidth
	)
	gbc.defaults()
	if len(gbc.hooks) == 0 {
		if err = gbc.check(); err != nil {
			return nil, err
		}
		node, err = gbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupBandwidthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gbc.check(); err != nil {
				return nil, err
			}
			gbc.mutation = mutation
			if node, err = gbc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gbc.hooks) - 1; i >= 0; i-- {
			if gbc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gbc *GroupBandwidthCreate) SaveX(ctx context.Context) *GroupBandwidth {
	v, err := gbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gbc *GroupBandwidthCreate) Exec(ctx context.Context) error {
	_, err := gbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gbc *GroupBandwidthCreate) ExecX(ctx context.Context) {
	if err := gbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gbc *GroupBandwidthCreate) defaults() {
	if _, ok := gbc.mutation.GbwCreatedAt(); !ok {
		v := groupbandwidth.DefaultGbwCreatedAt()
		gbc.mutation.SetGbwCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gbc *GroupBandwidthCreate) check() error {
	if _, ok := gbc.mutation.GbwDownloadSpeed(); !ok {
		return &ValidationError{Name: "gbw_download_speed", err: errors.New(`ent: missing required field "gbw_download_speed"`)}
	}
	if _, ok := gbc.mutation.GbwUploadSpeed(); !ok {
		return &ValidationError{Name: "gbw_upload_speed", err: errors.New(`ent: missing required field "gbw_upload_speed"`)}
	}
	return nil
}

func (gbc *GroupBandwidthCreate) sqlSave(ctx context.Context) (*GroupBandwidth, error) {
	_node, _spec := gbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gbc *GroupBandwidthCreate) createSpec() (*GroupBandwidth, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupBandwidth{config: gbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: groupbandwidth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupbandwidth.FieldID,
			},
		}
	)
	if value, ok := gbc.mutation.GbwDownloadSpeed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbandwidth.FieldGbwDownloadSpeed,
		})
		_node.GbwDownloadSpeed = value
	}
	if value, ok := gbc.mutation.GbwUploadSpeed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbandwidth.FieldGbwUploadSpeed,
		})
		_node.GbwUploadSpeed = value
	}
	if value, ok := gbc.mutation.GbwCreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupbandwidth.FieldGbwCreatedAt,
		})
		_node.GbwCreatedAt = value
	}
	if nodes := gbc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   groupbandwidth.GroupsTable,
			Columns: []string{groupbandwidth.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: groups.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupBandwidthCreateBulk is the builder for creating many GroupBandwidth entities in bulk.
type GroupBandwidthCreateBulk struct {
	config
	builders []*GroupBandwidthCreate
}

// Save creates the GroupBandwidth entities in the database.
func (gbcb *GroupBandwidthCreateBulk) Save(ctx context.Context) ([]*GroupBandwidth, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gbcb.builders))
	nodes := make([]*GroupBandwidth, len(gbcb.builders))
	mutators := make([]Mutator, len(gbcb.builders))
	for i := range gbcb.builders {
		func(i int, root context.Context) {
			builder := gbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupBandwidthMutation)
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
					_, err = mutators[i+1].Mutate(root, gbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gbcb *GroupBandwidthCreateBulk) SaveX(ctx context.Context) []*GroupBandwidth {
	v, err := gbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gbcb *GroupBandwidthCreateBulk) Exec(ctx context.Context) error {
	_, err := gbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gbcb *GroupBandwidthCreateBulk) ExecX(ctx context.Context) {
	if err := gbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
