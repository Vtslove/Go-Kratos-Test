// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-blog/app/blog/internal/data/ent/predicate"
	"kratos-blog/app/blog/internal/data/ent/tag"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks     []Hook
	mutation  *TagMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdateTime sets the "update_time" field.
func (tu *TagUpdate) SetUpdateTime(i int64) *TagUpdate {
	tu.mutation.ResetUpdateTime()
	tu.mutation.SetUpdateTime(i)
	return tu
}

// AddUpdateTime adds i to the "update_time" field.
func (tu *TagUpdate) AddUpdateTime(i int64) *TagUpdate {
	tu.mutation.AddUpdateTime(i)
	return tu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (tu *TagUpdate) ClearUpdateTime() *TagUpdate {
	tu.mutation.ClearUpdateTime()
	return tu
}

// SetName sets the "name" field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TagUpdate) SetNillableName(s *string) *TagUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// ClearName clears the value of the "name" field.
func (tu *TagUpdate) ClearName() *TagUpdate {
	tu.mutation.ClearName()
	return tu
}

// SetSlug sets the "slug" field.
func (tu *TagUpdate) SetSlug(s string) *TagUpdate {
	tu.mutation.SetSlug(s)
	return tu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (tu *TagUpdate) SetNillableSlug(s *string) *TagUpdate {
	if s != nil {
		tu.SetSlug(*s)
	}
	return tu
}

// ClearSlug clears the value of the "slug" field.
func (tu *TagUpdate) ClearSlug() *TagUpdate {
	tu.mutation.ClearSlug()
	return tu
}

// SetColor sets the "color" field.
func (tu *TagUpdate) SetColor(s string) *TagUpdate {
	tu.mutation.SetColor(s)
	return tu
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (tu *TagUpdate) SetNillableColor(s *string) *TagUpdate {
	if s != nil {
		tu.SetColor(*s)
	}
	return tu
}

// ClearColor clears the value of the "color" field.
func (tu *TagUpdate) ClearColor() *TagUpdate {
	tu.mutation.ClearColor()
	return tu
}

// SetThumbnail sets the "thumbnail" field.
func (tu *TagUpdate) SetThumbnail(s string) *TagUpdate {
	tu.mutation.SetThumbnail(s)
	return tu
}

// SetNillableThumbnail sets the "thumbnail" field if the given value is not nil.
func (tu *TagUpdate) SetNillableThumbnail(s *string) *TagUpdate {
	if s != nil {
		tu.SetThumbnail(*s)
	}
	return tu
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (tu *TagUpdate) ClearThumbnail() *TagUpdate {
	tu.mutation.ClearThumbnail()
	return tu
}

// SetSlugName sets the "slug_name" field.
func (tu *TagUpdate) SetSlugName(s string) *TagUpdate {
	tu.mutation.SetSlugName(s)
	return tu
}

// SetNillableSlugName sets the "slug_name" field if the given value is not nil.
func (tu *TagUpdate) SetNillableSlugName(s *string) *TagUpdate {
	if s != nil {
		tu.SetSlugName(*s)
	}
	return tu
}

// ClearSlugName clears the value of the "slug_name" field.
func (tu *TagUpdate) ClearSlugName() *TagUpdate {
	tu.mutation.ClearSlugName()
	return tu
}

// SetPostCount sets the "post_count" field.
func (tu *TagUpdate) SetPostCount(u uint32) *TagUpdate {
	tu.mutation.ResetPostCount()
	tu.mutation.SetPostCount(u)
	return tu
}

// SetNillablePostCount sets the "post_count" field if the given value is not nil.
func (tu *TagUpdate) SetNillablePostCount(u *uint32) *TagUpdate {
	if u != nil {
		tu.SetPostCount(*u)
	}
	return tu
}

// AddPostCount adds u to the "post_count" field.
func (tu *TagUpdate) AddPostCount(u int32) *TagUpdate {
	tu.mutation.AddPostCount(u)
	return tu
}

// ClearPostCount clears the value of the "post_count" field.
func (tu *TagUpdate) ClearPostCount() *TagUpdate {
	tu.mutation.ClearPostCount()
	return tu
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TagUpdate) defaults() {
	if _, ok := tu.mutation.UpdateTime(); !ok && !tu.mutation.UpdateTimeCleared() {
		v := tag.UpdateDefaultUpdateTime()
		tu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TagUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tag.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TagUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TagUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: tag.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tu.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: tag.FieldCreateTime,
		})
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tag.FieldUpdateTime,
		})
	}
	if value, ok := tu.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tag.FieldUpdateTime,
		})
	}
	if tu.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: tag.FieldUpdateTime,
		})
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldName,
		})
	}
	if tu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tag.FieldName,
		})
	}
	if value, ok := tu.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldSlug,
		})
	}
	if tu.mutation.SlugCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tag.FieldSlug,
		})
	}
	if value, ok := tu.mutation.Color(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldColor,
		})
	}
	if tu.mutation.ColorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tag.FieldColor,
		})
	}
	if value, ok := tu.mutation.Thumbnail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldThumbnail,
		})
	}
	if tu.mutation.ThumbnailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tag.FieldThumbnail,
		})
	}
	if value, ok := tu.mutation.SlugName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldSlugName,
		})
	}
	if tu.mutation.SlugNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tag.FieldSlugName,
		})
	}
	if value, ok := tu.mutation.PostCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tag.FieldPostCount,
		})
	}
	if value, ok := tu.mutation.AddedPostCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tag.FieldPostCount,
		})
	}
	if tu.mutation.PostCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: tag.FieldPostCount,
		})
	}
	_spec.Modifiers = tu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TagMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (tuo *TagUpdateOne) SetUpdateTime(i int64) *TagUpdateOne {
	tuo.mutation.ResetUpdateTime()
	tuo.mutation.SetUpdateTime(i)
	return tuo
}

// AddUpdateTime adds i to the "update_time" field.
func (tuo *TagUpdateOne) AddUpdateTime(i int64) *TagUpdateOne {
	tuo.mutation.AddUpdateTime(i)
	return tuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (tuo *TagUpdateOne) ClearUpdateTime() *TagUpdateOne {
	tuo.mutation.ClearUpdateTime()
	return tuo
}

// SetName sets the "name" field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableName(s *string) *TagUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// ClearName clears the value of the "name" field.
func (tuo *TagUpdateOne) ClearName() *TagUpdateOne {
	tuo.mutation.ClearName()
	return tuo
}

// SetSlug sets the "slug" field.
func (tuo *TagUpdateOne) SetSlug(s string) *TagUpdateOne {
	tuo.mutation.SetSlug(s)
	return tuo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableSlug(s *string) *TagUpdateOne {
	if s != nil {
		tuo.SetSlug(*s)
	}
	return tuo
}

// ClearSlug clears the value of the "slug" field.
func (tuo *TagUpdateOne) ClearSlug() *TagUpdateOne {
	tuo.mutation.ClearSlug()
	return tuo
}

// SetColor sets the "color" field.
func (tuo *TagUpdateOne) SetColor(s string) *TagUpdateOne {
	tuo.mutation.SetColor(s)
	return tuo
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableColor(s *string) *TagUpdateOne {
	if s != nil {
		tuo.SetColor(*s)
	}
	return tuo
}

// ClearColor clears the value of the "color" field.
func (tuo *TagUpdateOne) ClearColor() *TagUpdateOne {
	tuo.mutation.ClearColor()
	return tuo
}

// SetThumbnail sets the "thumbnail" field.
func (tuo *TagUpdateOne) SetThumbnail(s string) *TagUpdateOne {
	tuo.mutation.SetThumbnail(s)
	return tuo
}

// SetNillableThumbnail sets the "thumbnail" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableThumbnail(s *string) *TagUpdateOne {
	if s != nil {
		tuo.SetThumbnail(*s)
	}
	return tuo
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (tuo *TagUpdateOne) ClearThumbnail() *TagUpdateOne {
	tuo.mutation.ClearThumbnail()
	return tuo
}

// SetSlugName sets the "slug_name" field.
func (tuo *TagUpdateOne) SetSlugName(s string) *TagUpdateOne {
	tuo.mutation.SetSlugName(s)
	return tuo
}

// SetNillableSlugName sets the "slug_name" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableSlugName(s *string) *TagUpdateOne {
	if s != nil {
		tuo.SetSlugName(*s)
	}
	return tuo
}

// ClearSlugName clears the value of the "slug_name" field.
func (tuo *TagUpdateOne) ClearSlugName() *TagUpdateOne {
	tuo.mutation.ClearSlugName()
	return tuo
}

// SetPostCount sets the "post_count" field.
func (tuo *TagUpdateOne) SetPostCount(u uint32) *TagUpdateOne {
	tuo.mutation.ResetPostCount()
	tuo.mutation.SetPostCount(u)
	return tuo
}

// SetNillablePostCount sets the "post_count" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillablePostCount(u *uint32) *TagUpdateOne {
	if u != nil {
		tuo.SetPostCount(*u)
	}
	return tuo
}

// AddPostCount adds u to the "post_count" field.
func (tuo *TagUpdateOne) AddPostCount(u int32) *TagUpdateOne {
	tuo.mutation.AddPostCount(u)
	return tuo
}

// ClearPostCount clears the value of the "post_count" field.
func (tuo *TagUpdateOne) ClearPostCount() *TagUpdateOne {
	tuo.mutation.ClearPostCount()
	return tuo
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	var (
		err  error
		node *Tag
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Tag)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TagMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TagUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdateTime(); !ok && !tuo.mutation.UpdateTimeCleared() {
		v := tag.UpdateDefaultUpdateTime()
		tuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TagUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tag.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TagUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TagUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: tag.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tuo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: tag.FieldCreateTime,
		})
	}
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tag.FieldUpdateTime,
		})
	}
	if value, ok := tuo.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tag.FieldUpdateTime,
		})
	}
	if tuo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: tag.FieldUpdateTime,
		})
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldName,
		})
	}
	if tuo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tag.FieldName,
		})
	}
	if value, ok := tuo.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldSlug,
		})
	}
	if tuo.mutation.SlugCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tag.FieldSlug,
		})
	}
	if value, ok := tuo.mutation.Color(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldColor,
		})
	}
	if tuo.mutation.ColorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tag.FieldColor,
		})
	}
	if value, ok := tuo.mutation.Thumbnail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldThumbnail,
		})
	}
	if tuo.mutation.ThumbnailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tag.FieldThumbnail,
		})
	}
	if value, ok := tuo.mutation.SlugName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldSlugName,
		})
	}
	if tuo.mutation.SlugNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tag.FieldSlugName,
		})
	}
	if value, ok := tuo.mutation.PostCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tag.FieldPostCount,
		})
	}
	if value, ok := tuo.mutation.AddedPostCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: tag.FieldPostCount,
		})
	}
	if tuo.mutation.PostCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: tag.FieldPostCount,
		})
	}
	_spec.Modifiers = tuo.modifiers
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
