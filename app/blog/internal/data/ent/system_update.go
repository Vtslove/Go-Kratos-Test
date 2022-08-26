// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-blog/app/blog/internal/data/ent/predicate"
	"kratos-blog/app/blog/internal/data/ent/system"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SystemUpdate is the builder for updating System entities.
type SystemUpdate struct {
	config
	hooks     []Hook
	mutation  *SystemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SystemUpdate builder.
func (su *SystemUpdate) Where(ps ...predicate.System) *SystemUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *SystemUpdate) SetUpdateTime(i int64) *SystemUpdate {
	su.mutation.ResetUpdateTime()
	su.mutation.SetUpdateTime(i)
	return su
}

// AddUpdateTime adds i to the "update_time" field.
func (su *SystemUpdate) AddUpdateTime(i int64) *SystemUpdate {
	su.mutation.AddUpdateTime(i)
	return su
}

// ClearUpdateTime clears the value of the "update_time" field.
func (su *SystemUpdate) ClearUpdateTime() *SystemUpdate {
	su.mutation.ClearUpdateTime()
	return su
}

// SetTitle sets the "title" field.
func (su *SystemUpdate) SetTitle(s string) *SystemUpdate {
	su.mutation.SetTitle(s)
	return su
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (su *SystemUpdate) SetNillableTitle(s *string) *SystemUpdate {
	if s != nil {
		su.SetTitle(*s)
	}
	return su
}

// ClearTitle clears the value of the "title" field.
func (su *SystemUpdate) ClearTitle() *SystemUpdate {
	su.mutation.ClearTitle()
	return su
}

// SetKeywords sets the "keywords" field.
func (su *SystemUpdate) SetKeywords(s string) *SystemUpdate {
	su.mutation.SetKeywords(s)
	return su
}

// SetNillableKeywords sets the "keywords" field if the given value is not nil.
func (su *SystemUpdate) SetNillableKeywords(s *string) *SystemUpdate {
	if s != nil {
		su.SetKeywords(*s)
	}
	return su
}

// ClearKeywords clears the value of the "keywords" field.
func (su *SystemUpdate) ClearKeywords() *SystemUpdate {
	su.mutation.ClearKeywords()
	return su
}

// SetDescription sets the "description" field.
func (su *SystemUpdate) SetDescription(s string) *SystemUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (su *SystemUpdate) SetNillableDescription(s *string) *SystemUpdate {
	if s != nil {
		su.SetDescription(*s)
	}
	return su
}

// ClearDescription clears the value of the "description" field.
func (su *SystemUpdate) ClearDescription() *SystemUpdate {
	su.mutation.ClearDescription()
	return su
}

// SetRecordNumber sets the "record_number" field.
func (su *SystemUpdate) SetRecordNumber(s string) *SystemUpdate {
	su.mutation.SetRecordNumber(s)
	return su
}

// SetNillableRecordNumber sets the "record_number" field if the given value is not nil.
func (su *SystemUpdate) SetNillableRecordNumber(s *string) *SystemUpdate {
	if s != nil {
		su.SetRecordNumber(*s)
	}
	return su
}

// ClearRecordNumber clears the value of the "record_number" field.
func (su *SystemUpdate) ClearRecordNumber() *SystemUpdate {
	su.mutation.ClearRecordNumber()
	return su
}

// SetTheme sets the "Theme" field.
func (su *SystemUpdate) SetTheme(i int32) *SystemUpdate {
	su.mutation.ResetTheme()
	su.mutation.SetTheme(i)
	return su
}

// SetNillableTheme sets the "Theme" field if the given value is not nil.
func (su *SystemUpdate) SetNillableTheme(i *int32) *SystemUpdate {
	if i != nil {
		su.SetTheme(*i)
	}
	return su
}

// AddTheme adds i to the "Theme" field.
func (su *SystemUpdate) AddTheme(i int32) *SystemUpdate {
	su.mutation.AddTheme(i)
	return su
}

// ClearTheme clears the value of the "Theme" field.
func (su *SystemUpdate) ClearTheme() *SystemUpdate {
	su.mutation.ClearTheme()
	return su
}

// Mutation returns the SystemMutation object of the builder.
func (su *SystemUpdate) Mutation() *SystemMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SystemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SystemUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SystemUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SystemUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SystemUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok && !su.mutation.UpdateTimeCleared() {
		v := system.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SystemUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SystemUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SystemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   system.Table,
			Columns: system.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: system.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if su.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: system.FieldCreateTime,
		})
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: system.FieldUpdateTime,
		})
	}
	if value, ok := su.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: system.FieldUpdateTime,
		})
	}
	if su.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: system.FieldUpdateTime,
		})
	}
	if value, ok := su.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: system.FieldTitle,
		})
	}
	if su.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: system.FieldTitle,
		})
	}
	if value, ok := su.mutation.Keywords(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: system.FieldKeywords,
		})
	}
	if su.mutation.KeywordsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: system.FieldKeywords,
		})
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: system.FieldDescription,
		})
	}
	if su.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: system.FieldDescription,
		})
	}
	if value, ok := su.mutation.RecordNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: system.FieldRecordNumber,
		})
	}
	if su.mutation.RecordNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: system.FieldRecordNumber,
		})
	}
	if value, ok := su.mutation.Theme(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: system.FieldTheme,
		})
	}
	if value, ok := su.mutation.AddedTheme(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: system.FieldTheme,
		})
	}
	if su.mutation.ThemeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: system.FieldTheme,
		})
	}
	_spec.Modifiers = su.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{system.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SystemUpdateOne is the builder for updating a single System entity.
type SystemUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SystemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (suo *SystemUpdateOne) SetUpdateTime(i int64) *SystemUpdateOne {
	suo.mutation.ResetUpdateTime()
	suo.mutation.SetUpdateTime(i)
	return suo
}

// AddUpdateTime adds i to the "update_time" field.
func (suo *SystemUpdateOne) AddUpdateTime(i int64) *SystemUpdateOne {
	suo.mutation.AddUpdateTime(i)
	return suo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (suo *SystemUpdateOne) ClearUpdateTime() *SystemUpdateOne {
	suo.mutation.ClearUpdateTime()
	return suo
}

// SetTitle sets the "title" field.
func (suo *SystemUpdateOne) SetTitle(s string) *SystemUpdateOne {
	suo.mutation.SetTitle(s)
	return suo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableTitle(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetTitle(*s)
	}
	return suo
}

// ClearTitle clears the value of the "title" field.
func (suo *SystemUpdateOne) ClearTitle() *SystemUpdateOne {
	suo.mutation.ClearTitle()
	return suo
}

// SetKeywords sets the "keywords" field.
func (suo *SystemUpdateOne) SetKeywords(s string) *SystemUpdateOne {
	suo.mutation.SetKeywords(s)
	return suo
}

// SetNillableKeywords sets the "keywords" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableKeywords(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetKeywords(*s)
	}
	return suo
}

// ClearKeywords clears the value of the "keywords" field.
func (suo *SystemUpdateOne) ClearKeywords() *SystemUpdateOne {
	suo.mutation.ClearKeywords()
	return suo
}

// SetDescription sets the "description" field.
func (suo *SystemUpdateOne) SetDescription(s string) *SystemUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableDescription(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetDescription(*s)
	}
	return suo
}

// ClearDescription clears the value of the "description" field.
func (suo *SystemUpdateOne) ClearDescription() *SystemUpdateOne {
	suo.mutation.ClearDescription()
	return suo
}

// SetRecordNumber sets the "record_number" field.
func (suo *SystemUpdateOne) SetRecordNumber(s string) *SystemUpdateOne {
	suo.mutation.SetRecordNumber(s)
	return suo
}

// SetNillableRecordNumber sets the "record_number" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableRecordNumber(s *string) *SystemUpdateOne {
	if s != nil {
		suo.SetRecordNumber(*s)
	}
	return suo
}

// ClearRecordNumber clears the value of the "record_number" field.
func (suo *SystemUpdateOne) ClearRecordNumber() *SystemUpdateOne {
	suo.mutation.ClearRecordNumber()
	return suo
}

// SetTheme sets the "Theme" field.
func (suo *SystemUpdateOne) SetTheme(i int32) *SystemUpdateOne {
	suo.mutation.ResetTheme()
	suo.mutation.SetTheme(i)
	return suo
}

// SetNillableTheme sets the "Theme" field if the given value is not nil.
func (suo *SystemUpdateOne) SetNillableTheme(i *int32) *SystemUpdateOne {
	if i != nil {
		suo.SetTheme(*i)
	}
	return suo
}

// AddTheme adds i to the "Theme" field.
func (suo *SystemUpdateOne) AddTheme(i int32) *SystemUpdateOne {
	suo.mutation.AddTheme(i)
	return suo
}

// ClearTheme clears the value of the "Theme" field.
func (suo *SystemUpdateOne) ClearTheme() *SystemUpdateOne {
	suo.mutation.ClearTheme()
	return suo
}

// Mutation returns the SystemMutation object of the builder.
func (suo *SystemUpdateOne) Mutation() *SystemMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SystemUpdateOne) Select(field string, fields ...string) *SystemUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated System entity.
func (suo *SystemUpdateOne) Save(ctx context.Context) (*System, error) {
	var (
		err  error
		node *System
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*System)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SystemMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SystemUpdateOne) SaveX(ctx context.Context) *System {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SystemUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SystemUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SystemUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok && !suo.mutation.UpdateTimeCleared() {
		v := system.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SystemUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SystemUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SystemUpdateOne) sqlSave(ctx context.Context) (_node *System, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   system.Table,
			Columns: system.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: system.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "System.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, system.FieldID)
		for _, f := range fields {
			if !system.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != system.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if suo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: system.FieldCreateTime,
		})
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: system.FieldUpdateTime,
		})
	}
	if value, ok := suo.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: system.FieldUpdateTime,
		})
	}
	if suo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: system.FieldUpdateTime,
		})
	}
	if value, ok := suo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: system.FieldTitle,
		})
	}
	if suo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: system.FieldTitle,
		})
	}
	if value, ok := suo.mutation.Keywords(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: system.FieldKeywords,
		})
	}
	if suo.mutation.KeywordsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: system.FieldKeywords,
		})
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: system.FieldDescription,
		})
	}
	if suo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: system.FieldDescription,
		})
	}
	if value, ok := suo.mutation.RecordNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: system.FieldRecordNumber,
		})
	}
	if suo.mutation.RecordNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: system.FieldRecordNumber,
		})
	}
	if value, ok := suo.mutation.Theme(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: system.FieldTheme,
		})
	}
	if value, ok := suo.mutation.AddedTheme(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: system.FieldTheme,
		})
	}
	if suo.mutation.ThemeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: system.FieldTheme,
		})
	}
	_spec.Modifiers = suo.modifiers
	_node = &System{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{system.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
