// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mrzack99s/mrz-identity-management/ent/groupbandwidth"
	"github.com/mrzack99s/mrz-identity-management/ent/groups"
	"github.com/mrzack99s/mrz-identity-management/ent/predicate"
)

// GroupBandwidthUpdate is the builder for updating GroupBandwidth entities.
type GroupBandwidthUpdate struct {
	config
	hooks    []Hook
	mutation *GroupBandwidthMutation
}

// Where appends a list predicates to the GroupBandwidthUpdate builder.
func (gbu *GroupBandwidthUpdate) Where(ps ...predicate.GroupBandwidth) *GroupBandwidthUpdate {
	gbu.mutation.Where(ps...)
	return gbu
}

// SetGbwDownloadSpeed sets the "gbw_download_speed" field.
func (gbu *GroupBandwidthUpdate) SetGbwDownloadSpeed(i int) *GroupBandwidthUpdate {
	gbu.mutation.ResetGbwDownloadSpeed()
	gbu.mutation.SetGbwDownloadSpeed(i)
	return gbu
}

// AddGbwDownloadSpeed adds i to the "gbw_download_speed" field.
func (gbu *GroupBandwidthUpdate) AddGbwDownloadSpeed(i int) *GroupBandwidthUpdate {
	gbu.mutation.AddGbwDownloadSpeed(i)
	return gbu
}

// SetGbwUploadSpeed sets the "gbw_upload_speed" field.
func (gbu *GroupBandwidthUpdate) SetGbwUploadSpeed(i int) *GroupBandwidthUpdate {
	gbu.mutation.ResetGbwUploadSpeed()
	gbu.mutation.SetGbwUploadSpeed(i)
	return gbu
}

// AddGbwUploadSpeed adds i to the "gbw_upload_speed" field.
func (gbu *GroupBandwidthUpdate) AddGbwUploadSpeed(i int) *GroupBandwidthUpdate {
	gbu.mutation.AddGbwUploadSpeed(i)
	return gbu
}

// SetGbwCreatedAt sets the "gbw_created_at" field.
func (gbu *GroupBandwidthUpdate) SetGbwCreatedAt(t time.Time) *GroupBandwidthUpdate {
	gbu.mutation.SetGbwCreatedAt(t)
	return gbu
}

// SetNillableGbwCreatedAt sets the "gbw_created_at" field if the given value is not nil.
func (gbu *GroupBandwidthUpdate) SetNillableGbwCreatedAt(t *time.Time) *GroupBandwidthUpdate {
	if t != nil {
		gbu.SetGbwCreatedAt(*t)
	}
	return gbu
}

// ClearGbwCreatedAt clears the value of the "gbw_created_at" field.
func (gbu *GroupBandwidthUpdate) ClearGbwCreatedAt() *GroupBandwidthUpdate {
	gbu.mutation.ClearGbwCreatedAt()
	return gbu
}

// AddGroupIDs adds the "groups" edge to the Groups entity by IDs.
func (gbu *GroupBandwidthUpdate) AddGroupIDs(ids ...int) *GroupBandwidthUpdate {
	gbu.mutation.AddGroupIDs(ids...)
	return gbu
}

// AddGroups adds the "groups" edges to the Groups entity.
func (gbu *GroupBandwidthUpdate) AddGroups(g ...*Groups) *GroupBandwidthUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gbu.AddGroupIDs(ids...)
}

// Mutation returns the GroupBandwidthMutation object of the builder.
func (gbu *GroupBandwidthUpdate) Mutation() *GroupBandwidthMutation {
	return gbu.mutation
}

// ClearGroups clears all "groups" edges to the Groups entity.
func (gbu *GroupBandwidthUpdate) ClearGroups() *GroupBandwidthUpdate {
	gbu.mutation.ClearGroups()
	return gbu
}

// RemoveGroupIDs removes the "groups" edge to Groups entities by IDs.
func (gbu *GroupBandwidthUpdate) RemoveGroupIDs(ids ...int) *GroupBandwidthUpdate {
	gbu.mutation.RemoveGroupIDs(ids...)
	return gbu
}

// RemoveGroups removes "groups" edges to Groups entities.
func (gbu *GroupBandwidthUpdate) RemoveGroups(g ...*Groups) *GroupBandwidthUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gbu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gbu *GroupBandwidthUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gbu.hooks) == 0 {
		affected, err = gbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupBandwidthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gbu.mutation = mutation
			affected, err = gbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gbu.hooks) - 1; i >= 0; i-- {
			if gbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gbu *GroupBandwidthUpdate) SaveX(ctx context.Context) int {
	affected, err := gbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gbu *GroupBandwidthUpdate) Exec(ctx context.Context) error {
	_, err := gbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gbu *GroupBandwidthUpdate) ExecX(ctx context.Context) {
	if err := gbu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gbu *GroupBandwidthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupbandwidth.Table,
			Columns: groupbandwidth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupbandwidth.FieldID,
			},
		},
	}
	if ps := gbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gbu.mutation.GbwDownloadSpeed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbandwidth.FieldGbwDownloadSpeed,
		})
	}
	if value, ok := gbu.mutation.AddedGbwDownloadSpeed(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbandwidth.FieldGbwDownloadSpeed,
		})
	}
	if value, ok := gbu.mutation.GbwUploadSpeed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbandwidth.FieldGbwUploadSpeed,
		})
	}
	if value, ok := gbu.mutation.AddedGbwUploadSpeed(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbandwidth.FieldGbwUploadSpeed,
		})
	}
	if value, ok := gbu.mutation.GbwCreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupbandwidth.FieldGbwCreatedAt,
		})
	}
	if gbu.mutation.GbwCreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: groupbandwidth.FieldGbwCreatedAt,
		})
	}
	if gbu.mutation.GroupsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gbu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !gbu.mutation.GroupsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gbu.mutation.GroupsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupbandwidth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GroupBandwidthUpdateOne is the builder for updating a single GroupBandwidth entity.
type GroupBandwidthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupBandwidthMutation
}

// SetGbwDownloadSpeed sets the "gbw_download_speed" field.
func (gbuo *GroupBandwidthUpdateOne) SetGbwDownloadSpeed(i int) *GroupBandwidthUpdateOne {
	gbuo.mutation.ResetGbwDownloadSpeed()
	gbuo.mutation.SetGbwDownloadSpeed(i)
	return gbuo
}

// AddGbwDownloadSpeed adds i to the "gbw_download_speed" field.
func (gbuo *GroupBandwidthUpdateOne) AddGbwDownloadSpeed(i int) *GroupBandwidthUpdateOne {
	gbuo.mutation.AddGbwDownloadSpeed(i)
	return gbuo
}

// SetGbwUploadSpeed sets the "gbw_upload_speed" field.
func (gbuo *GroupBandwidthUpdateOne) SetGbwUploadSpeed(i int) *GroupBandwidthUpdateOne {
	gbuo.mutation.ResetGbwUploadSpeed()
	gbuo.mutation.SetGbwUploadSpeed(i)
	return gbuo
}

// AddGbwUploadSpeed adds i to the "gbw_upload_speed" field.
func (gbuo *GroupBandwidthUpdateOne) AddGbwUploadSpeed(i int) *GroupBandwidthUpdateOne {
	gbuo.mutation.AddGbwUploadSpeed(i)
	return gbuo
}

// SetGbwCreatedAt sets the "gbw_created_at" field.
func (gbuo *GroupBandwidthUpdateOne) SetGbwCreatedAt(t time.Time) *GroupBandwidthUpdateOne {
	gbuo.mutation.SetGbwCreatedAt(t)
	return gbuo
}

// SetNillableGbwCreatedAt sets the "gbw_created_at" field if the given value is not nil.
func (gbuo *GroupBandwidthUpdateOne) SetNillableGbwCreatedAt(t *time.Time) *GroupBandwidthUpdateOne {
	if t != nil {
		gbuo.SetGbwCreatedAt(*t)
	}
	return gbuo
}

// ClearGbwCreatedAt clears the value of the "gbw_created_at" field.
func (gbuo *GroupBandwidthUpdateOne) ClearGbwCreatedAt() *GroupBandwidthUpdateOne {
	gbuo.mutation.ClearGbwCreatedAt()
	return gbuo
}

// AddGroupIDs adds the "groups" edge to the Groups entity by IDs.
func (gbuo *GroupBandwidthUpdateOne) AddGroupIDs(ids ...int) *GroupBandwidthUpdateOne {
	gbuo.mutation.AddGroupIDs(ids...)
	return gbuo
}

// AddGroups adds the "groups" edges to the Groups entity.
func (gbuo *GroupBandwidthUpdateOne) AddGroups(g ...*Groups) *GroupBandwidthUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gbuo.AddGroupIDs(ids...)
}

// Mutation returns the GroupBandwidthMutation object of the builder.
func (gbuo *GroupBandwidthUpdateOne) Mutation() *GroupBandwidthMutation {
	return gbuo.mutation
}

// ClearGroups clears all "groups" edges to the Groups entity.
func (gbuo *GroupBandwidthUpdateOne) ClearGroups() *GroupBandwidthUpdateOne {
	gbuo.mutation.ClearGroups()
	return gbuo
}

// RemoveGroupIDs removes the "groups" edge to Groups entities by IDs.
func (gbuo *GroupBandwidthUpdateOne) RemoveGroupIDs(ids ...int) *GroupBandwidthUpdateOne {
	gbuo.mutation.RemoveGroupIDs(ids...)
	return gbuo
}

// RemoveGroups removes "groups" edges to Groups entities.
func (gbuo *GroupBandwidthUpdateOne) RemoveGroups(g ...*Groups) *GroupBandwidthUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gbuo.RemoveGroupIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gbuo *GroupBandwidthUpdateOne) Select(field string, fields ...string) *GroupBandwidthUpdateOne {
	gbuo.fields = append([]string{field}, fields...)
	return gbuo
}

// Save executes the query and returns the updated GroupBandwidth entity.
func (gbuo *GroupBandwidthUpdateOne) Save(ctx context.Context) (*GroupBandwidth, error) {
	var (
		err  error
		node *GroupBandwidth
	)
	if len(gbuo.hooks) == 0 {
		node, err = gbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupBandwidthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gbuo.mutation = mutation
			node, err = gbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gbuo.hooks) - 1; i >= 0; i-- {
			if gbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gbuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gbuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gbuo *GroupBandwidthUpdateOne) SaveX(ctx context.Context) *GroupBandwidth {
	node, err := gbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gbuo *GroupBandwidthUpdateOne) Exec(ctx context.Context) error {
	_, err := gbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gbuo *GroupBandwidthUpdateOne) ExecX(ctx context.Context) {
	if err := gbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gbuo *GroupBandwidthUpdateOne) sqlSave(ctx context.Context) (_node *GroupBandwidth, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupbandwidth.Table,
			Columns: groupbandwidth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupbandwidth.FieldID,
			},
		},
	}
	id, ok := gbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GroupBandwidth.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := gbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupbandwidth.FieldID)
		for _, f := range fields {
			if !groupbandwidth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != groupbandwidth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gbuo.mutation.GbwDownloadSpeed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbandwidth.FieldGbwDownloadSpeed,
		})
	}
	if value, ok := gbuo.mutation.AddedGbwDownloadSpeed(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbandwidth.FieldGbwDownloadSpeed,
		})
	}
	if value, ok := gbuo.mutation.GbwUploadSpeed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbandwidth.FieldGbwUploadSpeed,
		})
	}
	if value, ok := gbuo.mutation.AddedGbwUploadSpeed(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupbandwidth.FieldGbwUploadSpeed,
		})
	}
	if value, ok := gbuo.mutation.GbwCreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupbandwidth.FieldGbwCreatedAt,
		})
	}
	if gbuo.mutation.GbwCreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: groupbandwidth.FieldGbwCreatedAt,
		})
	}
	if gbuo.mutation.GroupsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gbuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !gbuo.mutation.GroupsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gbuo.mutation.GroupsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GroupBandwidth{config: gbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupbandwidth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
