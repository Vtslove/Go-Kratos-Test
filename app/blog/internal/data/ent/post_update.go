// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-blog/app/blog/internal/data/ent/post"
	"kratos-blog/app/blog/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks     []Hook
	mutation  *PostMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdateTime sets the "update_time" field.
func (pu *PostUpdate) SetUpdateTime(i int64) *PostUpdate {
	pu.mutation.ResetUpdateTime()
	pu.mutation.SetUpdateTime(i)
	return pu
}

// AddUpdateTime adds i to the "update_time" field.
func (pu *PostUpdate) AddUpdateTime(i int64) *PostUpdate {
	pu.mutation.AddUpdateTime(i)
	return pu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (pu *PostUpdate) ClearUpdateTime() *PostUpdate {
	pu.mutation.ClearUpdateTime()
	return pu
}

// SetTitle sets the "title" field.
func (pu *PostUpdate) SetTitle(s string) *PostUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pu *PostUpdate) SetNillableTitle(s *string) *PostUpdate {
	if s != nil {
		pu.SetTitle(*s)
	}
	return pu
}

// ClearTitle clears the value of the "title" field.
func (pu *PostUpdate) ClearTitle() *PostUpdate {
	pu.mutation.ClearTitle()
	return pu
}

// SetSummary sets the "summary" field.
func (pu *PostUpdate) SetSummary(s string) *PostUpdate {
	pu.mutation.SetSummary(s)
	return pu
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (pu *PostUpdate) SetNillableSummary(s *string) *PostUpdate {
	if s != nil {
		pu.SetSummary(*s)
	}
	return pu
}

// ClearSummary clears the value of the "summary" field.
func (pu *PostUpdate) ClearSummary() *PostUpdate {
	pu.mutation.ClearSummary()
	return pu
}

// SetOriginal sets the "original" field.
func (pu *PostUpdate) SetOriginal(s string) *PostUpdate {
	pu.mutation.SetOriginal(s)
	return pu
}

// SetNillableOriginal sets the "original" field if the given value is not nil.
func (pu *PostUpdate) SetNillableOriginal(s *string) *PostUpdate {
	if s != nil {
		pu.SetOriginal(*s)
	}
	return pu
}

// ClearOriginal clears the value of the "original" field.
func (pu *PostUpdate) ClearOriginal() *PostUpdate {
	pu.mutation.ClearOriginal()
	return pu
}

// SetContent sets the "content" field.
func (pu *PostUpdate) SetContent(s string) *PostUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (pu *PostUpdate) SetNillableContent(s *string) *PostUpdate {
	if s != nil {
		pu.SetContent(*s)
	}
	return pu
}

// ClearContent clears the value of the "content" field.
func (pu *PostUpdate) ClearContent() *PostUpdate {
	pu.mutation.ClearContent()
	return pu
}

// SetPassword sets the "password" field.
func (pu *PostUpdate) SetPassword(s string) *PostUpdate {
	pu.mutation.SetPassword(s)
	return pu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (pu *PostUpdate) SetNillablePassword(s *string) *PostUpdate {
	if s != nil {
		pu.SetPassword(*s)
	}
	return pu
}

// ClearPassword clears the value of the "password" field.
func (pu *PostUpdate) ClearPassword() *PostUpdate {
	pu.mutation.ClearPassword()
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PostUpdate) SetUserID(i int64) *PostUpdate {
	pu.mutation.ResetUserID()
	pu.mutation.SetUserID(i)
	return pu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pu *PostUpdate) SetNillableUserID(i *int64) *PostUpdate {
	if i != nil {
		pu.SetUserID(*i)
	}
	return pu
}

// AddUserID adds i to the "user_id" field.
func (pu *PostUpdate) AddUserID(i int64) *PostUpdate {
	pu.mutation.AddUserID(i)
	return pu
}

// ClearUserID clears the value of the "user_id" field.
func (pu *PostUpdate) ClearUserID() *PostUpdate {
	pu.mutation.ClearUserID()
	return pu
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PostUpdate) defaults() {
	if _, ok := pu.mutation.UpdateTime(); !ok && !pu.mutation.UpdateTimeCleared() {
		v := post.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PostUpdate) check() error {
	if v, ok := pu.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PostUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PostUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: post.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pu.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: post.FieldCreateTime,
		})
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUpdateTime,
		})
	}
	if value, ok := pu.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUpdateTime,
		})
	}
	if pu.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: post.FieldUpdateTime,
		})
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldTitle,
		})
	}
	if pu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: post.FieldTitle,
		})
	}
	if value, ok := pu.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldSummary,
		})
	}
	if pu.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: post.FieldSummary,
		})
	}
	if value, ok := pu.mutation.Original(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldOriginal,
		})
	}
	if pu.mutation.OriginalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: post.FieldOriginal,
		})
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldContent,
		})
	}
	if pu.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: post.FieldContent,
		})
	}
	if value, ok := pu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldPassword,
		})
	}
	if pu.mutation.PasswordCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: post.FieldPassword,
		})
	}
	if value, ok := pu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUserID,
		})
	}
	if value, ok := pu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUserID,
		})
	}
	if pu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: post.FieldUserID,
		})
	}
	_spec.Modifiers = pu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PostMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (puo *PostUpdateOne) SetUpdateTime(i int64) *PostUpdateOne {
	puo.mutation.ResetUpdateTime()
	puo.mutation.SetUpdateTime(i)
	return puo
}

// AddUpdateTime adds i to the "update_time" field.
func (puo *PostUpdateOne) AddUpdateTime(i int64) *PostUpdateOne {
	puo.mutation.AddUpdateTime(i)
	return puo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (puo *PostUpdateOne) ClearUpdateTime() *PostUpdateOne {
	puo.mutation.ClearUpdateTime()
	return puo
}

// SetTitle sets the "title" field.
func (puo *PostUpdateOne) SetTitle(s string) *PostUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableTitle(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetTitle(*s)
	}
	return puo
}

// ClearTitle clears the value of the "title" field.
func (puo *PostUpdateOne) ClearTitle() *PostUpdateOne {
	puo.mutation.ClearTitle()
	return puo
}

// SetSummary sets the "summary" field.
func (puo *PostUpdateOne) SetSummary(s string) *PostUpdateOne {
	puo.mutation.SetSummary(s)
	return puo
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableSummary(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetSummary(*s)
	}
	return puo
}

// ClearSummary clears the value of the "summary" field.
func (puo *PostUpdateOne) ClearSummary() *PostUpdateOne {
	puo.mutation.ClearSummary()
	return puo
}

// SetOriginal sets the "original" field.
func (puo *PostUpdateOne) SetOriginal(s string) *PostUpdateOne {
	puo.mutation.SetOriginal(s)
	return puo
}

// SetNillableOriginal sets the "original" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableOriginal(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetOriginal(*s)
	}
	return puo
}

// ClearOriginal clears the value of the "original" field.
func (puo *PostUpdateOne) ClearOriginal() *PostUpdateOne {
	puo.mutation.ClearOriginal()
	return puo
}

// SetContent sets the "content" field.
func (puo *PostUpdateOne) SetContent(s string) *PostUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableContent(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetContent(*s)
	}
	return puo
}

// ClearContent clears the value of the "content" field.
func (puo *PostUpdateOne) ClearContent() *PostUpdateOne {
	puo.mutation.ClearContent()
	return puo
}

// SetPassword sets the "password" field.
func (puo *PostUpdateOne) SetPassword(s string) *PostUpdateOne {
	puo.mutation.SetPassword(s)
	return puo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillablePassword(s *string) *PostUpdateOne {
	if s != nil {
		puo.SetPassword(*s)
	}
	return puo
}

// ClearPassword clears the value of the "password" field.
func (puo *PostUpdateOne) ClearPassword() *PostUpdateOne {
	puo.mutation.ClearPassword()
	return puo
}

// SetUserID sets the "user_id" field.
func (puo *PostUpdateOne) SetUserID(i int64) *PostUpdateOne {
	puo.mutation.ResetUserID()
	puo.mutation.SetUserID(i)
	return puo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableUserID(i *int64) *PostUpdateOne {
	if i != nil {
		puo.SetUserID(*i)
	}
	return puo
}

// AddUserID adds i to the "user_id" field.
func (puo *PostUpdateOne) AddUserID(i int64) *PostUpdateOne {
	puo.mutation.AddUserID(i)
	return puo
}

// ClearUserID clears the value of the "user_id" field.
func (puo *PostUpdateOne) ClearUserID() *PostUpdateOne {
	puo.mutation.ClearUserID()
	return puo
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	var (
		err  error
		node *Post
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Post)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PostMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PostUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateTime(); !ok && !puo.mutation.UpdateTimeCleared() {
		v := post.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PostUpdateOne) check() error {
	if v, ok := puo.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PostUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PostUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: post.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if puo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: post.FieldCreateTime,
		})
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUpdateTime,
		})
	}
	if value, ok := puo.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUpdateTime,
		})
	}
	if puo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: post.FieldUpdateTime,
		})
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldTitle,
		})
	}
	if puo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: post.FieldTitle,
		})
	}
	if value, ok := puo.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldSummary,
		})
	}
	if puo.mutation.SummaryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: post.FieldSummary,
		})
	}
	if value, ok := puo.mutation.Original(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldOriginal,
		})
	}
	if puo.mutation.OriginalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: post.FieldOriginal,
		})
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldContent,
		})
	}
	if puo.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: post.FieldContent,
		})
	}
	if value, ok := puo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldPassword,
		})
	}
	if puo.mutation.PasswordCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: post.FieldPassword,
		})
	}
	if value, ok := puo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUserID,
		})
	}
	if value, ok := puo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: post.FieldUserID,
		})
	}
	if puo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: post.FieldUserID,
		})
	}
	_spec.Modifiers = puo.modifiers
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
