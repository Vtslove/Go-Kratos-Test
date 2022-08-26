// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-blog/app/blog/internal/data/ent/menu"
	"kratos-blog/app/blog/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuUpdate is the builder for updating Menu entities.
type MenuUpdate struct {
	config
	hooks     []Hook
	mutation  *MenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MenuUpdate builder.
func (mu *MenuUpdate) Where(ps ...predicate.Menu) *MenuUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdateTime sets the "update_time" field.
func (mu *MenuUpdate) SetUpdateTime(i int64) *MenuUpdate {
	mu.mutation.ResetUpdateTime()
	mu.mutation.SetUpdateTime(i)
	return mu
}

// AddUpdateTime adds i to the "update_time" field.
func (mu *MenuUpdate) AddUpdateTime(i int64) *MenuUpdate {
	mu.mutation.AddUpdateTime(i)
	return mu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (mu *MenuUpdate) ClearUpdateTime() *MenuUpdate {
	mu.mutation.ClearUpdateTime()
	return mu
}

// SetName sets the "name" field.
func (mu *MenuUpdate) SetName(s string) *MenuUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableName(s *string) *MenuUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// ClearName clears the value of the "name" field.
func (mu *MenuUpdate) ClearName() *MenuUpdate {
	mu.mutation.ClearName()
	return mu
}

// SetURL sets the "url" field.
func (mu *MenuUpdate) SetURL(s string) *MenuUpdate {
	mu.mutation.SetURL(s)
	return mu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableURL(s *string) *MenuUpdate {
	if s != nil {
		mu.SetURL(*s)
	}
	return mu
}

// ClearURL clears the value of the "url" field.
func (mu *MenuUpdate) ClearURL() *MenuUpdate {
	mu.mutation.ClearURL()
	return mu
}

// SetPriority sets the "priority" field.
func (mu *MenuUpdate) SetPriority(i int32) *MenuUpdate {
	mu.mutation.ResetPriority()
	mu.mutation.SetPriority(i)
	return mu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (mu *MenuUpdate) SetNillablePriority(i *int32) *MenuUpdate {
	if i != nil {
		mu.SetPriority(*i)
	}
	return mu
}

// AddPriority adds i to the "priority" field.
func (mu *MenuUpdate) AddPriority(i int32) *MenuUpdate {
	mu.mutation.AddPriority(i)
	return mu
}

// ClearPriority clears the value of the "priority" field.
func (mu *MenuUpdate) ClearPriority() *MenuUpdate {
	mu.mutation.ClearPriority()
	return mu
}

// SetTarget sets the "target" field.
func (mu *MenuUpdate) SetTarget(s string) *MenuUpdate {
	mu.mutation.SetTarget(s)
	return mu
}

// SetNillableTarget sets the "target" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableTarget(s *string) *MenuUpdate {
	if s != nil {
		mu.SetTarget(*s)
	}
	return mu
}

// ClearTarget clears the value of the "target" field.
func (mu *MenuUpdate) ClearTarget() *MenuUpdate {
	mu.mutation.ClearTarget()
	return mu
}

// SetIcon sets the "icon" field.
func (mu *MenuUpdate) SetIcon(s string) *MenuUpdate {
	mu.mutation.SetIcon(s)
	return mu
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableIcon(s *string) *MenuUpdate {
	if s != nil {
		mu.SetIcon(*s)
	}
	return mu
}

// ClearIcon clears the value of the "icon" field.
func (mu *MenuUpdate) ClearIcon() *MenuUpdate {
	mu.mutation.ClearIcon()
	return mu
}

// SetParentID sets the "parent_id" field.
func (mu *MenuUpdate) SetParentID(u uint32) *MenuUpdate {
	mu.mutation.ResetParentID()
	mu.mutation.SetParentID(u)
	return mu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableParentID(u *uint32) *MenuUpdate {
	if u != nil {
		mu.SetParentID(*u)
	}
	return mu
}

// AddParentID adds u to the "parent_id" field.
func (mu *MenuUpdate) AddParentID(u int32) *MenuUpdate {
	mu.mutation.AddParentID(u)
	return mu
}

// ClearParentID clears the value of the "parent_id" field.
func (mu *MenuUpdate) ClearParentID() *MenuUpdate {
	mu.mutation.ClearParentID()
	return mu
}

// SetTeam sets the "team" field.
func (mu *MenuUpdate) SetTeam(s string) *MenuUpdate {
	mu.mutation.SetTeam(s)
	return mu
}

// SetNillableTeam sets the "team" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableTeam(s *string) *MenuUpdate {
	if s != nil {
		mu.SetTeam(*s)
	}
	return mu
}

// ClearTeam clears the value of the "team" field.
func (mu *MenuUpdate) ClearTeam() *MenuUpdate {
	mu.mutation.ClearTeam()
	return mu
}

// Mutation returns the MenuMutation object of the builder.
func (mu *MenuUpdate) Mutation() *MenuMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MenuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MenuUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MenuUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MenuUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MenuUpdate) defaults() {
	if _, ok := mu.mutation.UpdateTime(); !ok && !mu.mutation.UpdateTimeCleared() {
		v := menu.UpdateDefaultUpdateTime()
		mu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MenuUpdate) check() error {
	if v, ok := mu.mutation.Name(); ok {
		if err := menu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Menu.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MenuUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MenuUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: menu.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mu.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: menu.FieldCreateTime,
		})
	}
	if value, ok := mu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldUpdateTime,
		})
	}
	if value, ok := mu.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldUpdateTime,
		})
	}
	if mu.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: menu.FieldUpdateTime,
		})
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldName,
		})
	}
	if mu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldName,
		})
	}
	if value, ok := mu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldURL,
		})
	}
	if mu.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldURL,
		})
	}
	if value, ok := mu.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: menu.FieldPriority,
		})
	}
	if value, ok := mu.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: menu.FieldPriority,
		})
	}
	if mu.mutation.PriorityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: menu.FieldPriority,
		})
	}
	if value, ok := mu.mutation.Target(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldTarget,
		})
	}
	if mu.mutation.TargetCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldTarget,
		})
	}
	if value, ok := mu.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldIcon,
		})
	}
	if mu.mutation.IconCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldIcon,
		})
	}
	if value, ok := mu.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: menu.FieldParentID,
		})
	}
	if value, ok := mu.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: menu.FieldParentID,
		})
	}
	if mu.mutation.ParentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: menu.FieldParentID,
		})
	}
	if value, ok := mu.mutation.Team(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldTeam,
		})
	}
	if mu.mutation.TeamCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldTeam,
		})
	}
	_spec.Modifiers = mu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MenuUpdateOne is the builder for updating a single Menu entity.
type MenuUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (muo *MenuUpdateOne) SetUpdateTime(i int64) *MenuUpdateOne {
	muo.mutation.ResetUpdateTime()
	muo.mutation.SetUpdateTime(i)
	return muo
}

// AddUpdateTime adds i to the "update_time" field.
func (muo *MenuUpdateOne) AddUpdateTime(i int64) *MenuUpdateOne {
	muo.mutation.AddUpdateTime(i)
	return muo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (muo *MenuUpdateOne) ClearUpdateTime() *MenuUpdateOne {
	muo.mutation.ClearUpdateTime()
	return muo
}

// SetName sets the "name" field.
func (muo *MenuUpdateOne) SetName(s string) *MenuUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableName(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// ClearName clears the value of the "name" field.
func (muo *MenuUpdateOne) ClearName() *MenuUpdateOne {
	muo.mutation.ClearName()
	return muo
}

// SetURL sets the "url" field.
func (muo *MenuUpdateOne) SetURL(s string) *MenuUpdateOne {
	muo.mutation.SetURL(s)
	return muo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableURL(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetURL(*s)
	}
	return muo
}

// ClearURL clears the value of the "url" field.
func (muo *MenuUpdateOne) ClearURL() *MenuUpdateOne {
	muo.mutation.ClearURL()
	return muo
}

// SetPriority sets the "priority" field.
func (muo *MenuUpdateOne) SetPriority(i int32) *MenuUpdateOne {
	muo.mutation.ResetPriority()
	muo.mutation.SetPriority(i)
	return muo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillablePriority(i *int32) *MenuUpdateOne {
	if i != nil {
		muo.SetPriority(*i)
	}
	return muo
}

// AddPriority adds i to the "priority" field.
func (muo *MenuUpdateOne) AddPriority(i int32) *MenuUpdateOne {
	muo.mutation.AddPriority(i)
	return muo
}

// ClearPriority clears the value of the "priority" field.
func (muo *MenuUpdateOne) ClearPriority() *MenuUpdateOne {
	muo.mutation.ClearPriority()
	return muo
}

// SetTarget sets the "target" field.
func (muo *MenuUpdateOne) SetTarget(s string) *MenuUpdateOne {
	muo.mutation.SetTarget(s)
	return muo
}

// SetNillableTarget sets the "target" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableTarget(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetTarget(*s)
	}
	return muo
}

// ClearTarget clears the value of the "target" field.
func (muo *MenuUpdateOne) ClearTarget() *MenuUpdateOne {
	muo.mutation.ClearTarget()
	return muo
}

// SetIcon sets the "icon" field.
func (muo *MenuUpdateOne) SetIcon(s string) *MenuUpdateOne {
	muo.mutation.SetIcon(s)
	return muo
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableIcon(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetIcon(*s)
	}
	return muo
}

// ClearIcon clears the value of the "icon" field.
func (muo *MenuUpdateOne) ClearIcon() *MenuUpdateOne {
	muo.mutation.ClearIcon()
	return muo
}

// SetParentID sets the "parent_id" field.
func (muo *MenuUpdateOne) SetParentID(u uint32) *MenuUpdateOne {
	muo.mutation.ResetParentID()
	muo.mutation.SetParentID(u)
	return muo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableParentID(u *uint32) *MenuUpdateOne {
	if u != nil {
		muo.SetParentID(*u)
	}
	return muo
}

// AddParentID adds u to the "parent_id" field.
func (muo *MenuUpdateOne) AddParentID(u int32) *MenuUpdateOne {
	muo.mutation.AddParentID(u)
	return muo
}

// ClearParentID clears the value of the "parent_id" field.
func (muo *MenuUpdateOne) ClearParentID() *MenuUpdateOne {
	muo.mutation.ClearParentID()
	return muo
}

// SetTeam sets the "team" field.
func (muo *MenuUpdateOne) SetTeam(s string) *MenuUpdateOne {
	muo.mutation.SetTeam(s)
	return muo
}

// SetNillableTeam sets the "team" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableTeam(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetTeam(*s)
	}
	return muo
}

// ClearTeam clears the value of the "team" field.
func (muo *MenuUpdateOne) ClearTeam() *MenuUpdateOne {
	muo.mutation.ClearTeam()
	return muo
}

// Mutation returns the MenuMutation object of the builder.
func (muo *MenuUpdateOne) Mutation() *MenuMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MenuUpdateOne) Select(field string, fields ...string) *MenuUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Menu entity.
func (muo *MenuUpdateOne) Save(ctx context.Context) (*Menu, error) {
	var (
		err  error
		node *Menu
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Menu)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MenuMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MenuUpdateOne) SaveX(ctx context.Context) *Menu {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MenuUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MenuUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MenuUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdateTime(); !ok && !muo.mutation.UpdateTimeCleared() {
		v := menu.UpdateDefaultUpdateTime()
		muo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MenuUpdateOne) check() error {
	if v, ok := muo.mutation.Name(); ok {
		if err := menu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Menu.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MenuUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MenuUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MenuUpdateOne) sqlSave(ctx context.Context) (_node *Menu, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: menu.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Menu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menu.FieldID)
		for _, f := range fields {
			if !menu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != menu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if muo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: menu.FieldCreateTime,
		})
	}
	if value, ok := muo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldUpdateTime,
		})
	}
	if value, ok := muo.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldUpdateTime,
		})
	}
	if muo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: menu.FieldUpdateTime,
		})
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldName,
		})
	}
	if muo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldName,
		})
	}
	if value, ok := muo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldURL,
		})
	}
	if muo.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldURL,
		})
	}
	if value, ok := muo.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: menu.FieldPriority,
		})
	}
	if value, ok := muo.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: menu.FieldPriority,
		})
	}
	if muo.mutation.PriorityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: menu.FieldPriority,
		})
	}
	if value, ok := muo.mutation.Target(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldTarget,
		})
	}
	if muo.mutation.TargetCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldTarget,
		})
	}
	if value, ok := muo.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldIcon,
		})
	}
	if muo.mutation.IconCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldIcon,
		})
	}
	if value, ok := muo.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: menu.FieldParentID,
		})
	}
	if value, ok := muo.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: menu.FieldParentID,
		})
	}
	if muo.mutation.ParentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: menu.FieldParentID,
		})
	}
	if value, ok := muo.mutation.Team(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldTeam,
		})
	}
	if muo.mutation.TeamCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldTeam,
		})
	}
	_spec.Modifiers = muo.modifiers
	_node = &Menu{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
