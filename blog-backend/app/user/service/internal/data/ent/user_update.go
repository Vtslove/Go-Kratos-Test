// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-blog/app/user/service/internal/data/ent/predicate"
	"kratos-blog/app/user/service/internal/data/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *UserUpdate) SetUpdateTime(i int64) *UserUpdate {
	uu.mutation.ResetUpdateTime()
	uu.mutation.SetUpdateTime(i)
	return uu
}

// AddUpdateTime adds i to the "update_time" field.
func (uu *UserUpdate) AddUpdateTime(i int64) *UserUpdate {
	uu.mutation.AddUpdateTime(i)
	return uu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (uu *UserUpdate) ClearUpdateTime() *UserUpdate {
	uu.mutation.ClearUpdateTime()
	return uu
}

// SetDeleteTime sets the "delete_time" field.
func (uu *UserUpdate) SetDeleteTime(i int64) *UserUpdate {
	uu.mutation.ResetDeleteTime()
	uu.mutation.SetDeleteTime(i)
	return uu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeleteTime(i *int64) *UserUpdate {
	if i != nil {
		uu.SetDeleteTime(*i)
	}
	return uu
}

// AddDeleteTime adds i to the "delete_time" field.
func (uu *UserUpdate) AddDeleteTime(i int64) *UserUpdate {
	uu.mutation.AddDeleteTime(i)
	return uu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (uu *UserUpdate) ClearDeleteTime() *UserUpdate {
	uu.mutation.ClearDeleteTime()
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickname(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickname(*s)
	}
	return uu
}

// ClearNickname clears the value of the "nickname" field.
func (uu *UserUpdate) ClearNickname() *UserUpdate {
	uu.mutation.ClearNickname()
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "email" field.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// SetDescription sets the "description" field.
func (uu *UserUpdate) SetDescription(s string) *UserUpdate {
	uu.mutation.SetDescription(s)
	return uu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDescription(s *string) *UserUpdate {
	if s != nil {
		uu.SetDescription(*s)
	}
	return uu
}

// ClearDescription clears the value of the "description" field.
func (uu *UserUpdate) ClearDescription() *UserUpdate {
	uu.mutation.ClearDescription()
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// ClearPassword clears the value of the "password" field.
func (uu *UserUpdate) ClearPassword() *UserUpdate {
	uu.mutation.ClearPassword()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdateTime(); !ok && !uu.mutation.UpdateTimeCleared() {
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Avatar(); ok {
		if err := user.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "User.avatar": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Description(); ok {
		if err := user.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "User.description": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uu *UserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdate {
	uu.modifiers = append(uu.modifiers, modifiers...)
	return uu
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uu.mutation.CreateTimeCleared() {
		_spec.ClearField(user.FieldCreateTime, field.TypeInt64)
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedUpdateTime(); ok {
		_spec.AddField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if uu.mutation.UpdateTimeCleared() {
		_spec.ClearField(user.FieldUpdateTime, field.TypeInt64)
	}
	if value, ok := uu.mutation.DeleteTime(); ok {
		_spec.SetField(user.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedDeleteTime(); ok {
		_spec.AddField(user.FieldDeleteTime, field.TypeInt64, value)
	}
	if uu.mutation.DeleteTimeCleared() {
		_spec.ClearField(user.FieldDeleteTime, field.TypeInt64)
	}
	if uu.mutation.UsernameCleared() {
		_spec.ClearField(user.FieldUsername, field.TypeString)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if uu.mutation.NicknameCleared() {
		_spec.ClearField(user.FieldNickname, field.TypeString)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uu.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uu.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
	}
	if uu.mutation.DescriptionCleared() {
		_spec.ClearField(user.FieldDescription, field.TypeString)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uu.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	_spec.AddModifiers(uu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UserUpdateOne) SetUpdateTime(i int64) *UserUpdateOne {
	uuo.mutation.ResetUpdateTime()
	uuo.mutation.SetUpdateTime(i)
	return uuo
}

// AddUpdateTime adds i to the "update_time" field.
func (uuo *UserUpdateOne) AddUpdateTime(i int64) *UserUpdateOne {
	uuo.mutation.AddUpdateTime(i)
	return uuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (uuo *UserUpdateOne) ClearUpdateTime() *UserUpdateOne {
	uuo.mutation.ClearUpdateTime()
	return uuo
}

// SetDeleteTime sets the "delete_time" field.
func (uuo *UserUpdateOne) SetDeleteTime(i int64) *UserUpdateOne {
	uuo.mutation.ResetDeleteTime()
	uuo.mutation.SetDeleteTime(i)
	return uuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeleteTime(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetDeleteTime(*i)
	}
	return uuo
}

// AddDeleteTime adds i to the "delete_time" field.
func (uuo *UserUpdateOne) AddDeleteTime(i int64) *UserUpdateOne {
	uuo.mutation.AddDeleteTime(i)
	return uuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (uuo *UserUpdateOne) ClearDeleteTime() *UserUpdateOne {
	uuo.mutation.ClearDeleteTime()
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickname(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickname(*s)
	}
	return uuo
}

// ClearNickname clears the value of the "nickname" field.
func (uuo *UserUpdateOne) ClearNickname() *UserUpdateOne {
	uuo.mutation.ClearNickname()
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "email" field.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// SetDescription sets the "description" field.
func (uuo *UserUpdateOne) SetDescription(s string) *UserUpdateOne {
	uuo.mutation.SetDescription(s)
	return uuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDescription(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDescription(*s)
	}
	return uuo
}

// ClearDescription clears the value of the "description" field.
func (uuo *UserUpdateOne) ClearDescription() *UserUpdateOne {
	uuo.mutation.ClearDescription()
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// ClearPassword clears the value of the "password" field.
func (uuo *UserUpdateOne) ClearPassword() *UserUpdateOne {
	uuo.mutation.ClearPassword()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdateTime(); !ok && !uuo.mutation.UpdateTimeCleared() {
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Avatar(); ok {
		if err := user.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "User.avatar": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Description(); ok {
		if err := user.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "User.description": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uuo *UserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdateOne {
	uuo.modifiers = append(uuo.modifiers, modifiers...)
	return uuo
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uuo.mutation.CreateTimeCleared() {
		_spec.ClearField(user.FieldCreateTime, field.TypeInt64)
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if uuo.mutation.UpdateTimeCleared() {
		_spec.ClearField(user.FieldUpdateTime, field.TypeInt64)
	}
	if value, ok := uuo.mutation.DeleteTime(); ok {
		_spec.SetField(user.FieldDeleteTime, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedDeleteTime(); ok {
		_spec.AddField(user.FieldDeleteTime, field.TypeInt64, value)
	}
	if uuo.mutation.DeleteTimeCleared() {
		_spec.ClearField(user.FieldDeleteTime, field.TypeInt64)
	}
	if uuo.mutation.UsernameCleared() {
		_spec.ClearField(user.FieldUsername, field.TypeString)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if uuo.mutation.NicknameCleared() {
		_spec.ClearField(user.FieldNickname, field.TypeString)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uuo.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uuo.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
	}
	if uuo.mutation.DescriptionCleared() {
		_spec.ClearField(user.FieldDescription, field.TypeString)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uuo.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	_spec.AddModifiers(uuo.modifiers...)
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
