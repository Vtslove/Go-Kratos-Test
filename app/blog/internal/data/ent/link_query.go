// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kratos-blog/app/blog/internal/data/ent/link"
	"kratos-blog/app/blog/internal/data/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkQuery is the builder for querying Link entities.
type LinkQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Link
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LinkQuery builder.
func (lq *LinkQuery) Where(ps ...predicate.Link) *LinkQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit adds a limit step to the query.
func (lq *LinkQuery) Limit(limit int) *LinkQuery {
	lq.limit = &limit
	return lq
}

// Offset adds an offset step to the query.
func (lq *LinkQuery) Offset(offset int) *LinkQuery {
	lq.offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LinkQuery) Unique(unique bool) *LinkQuery {
	lq.unique = &unique
	return lq
}

// Order adds an order step to the query.
func (lq *LinkQuery) Order(o ...OrderFunc) *LinkQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// First returns the first Link entity from the query.
// Returns a *NotFoundError when no Link was found.
func (lq *LinkQuery) First(ctx context.Context) (*Link, error) {
	nodes, err := lq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{link.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LinkQuery) FirstX(ctx context.Context) *Link {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Link ID from the query.
// Returns a *NotFoundError when no Link ID was found.
func (lq *LinkQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = lq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{link.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LinkQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Link entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Link entity is found.
// Returns a *NotFoundError when no Link entities are found.
func (lq *LinkQuery) Only(ctx context.Context) (*Link, error) {
	nodes, err := lq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{link.Label}
	default:
		return nil, &NotSingularError{link.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LinkQuery) OnlyX(ctx context.Context) *Link {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Link ID in the query.
// Returns a *NotSingularError when more than one Link ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LinkQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = lq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{link.Label}
	default:
		err = &NotSingularError{link.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LinkQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Links.
func (lq *LinkQuery) All(ctx context.Context) ([]*Link, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lq *LinkQuery) AllX(ctx context.Context) []*Link {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Link IDs.
func (lq *LinkQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := lq.Select(link.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LinkQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LinkQuery) Count(ctx context.Context) (int, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LinkQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LinkQuery) Exist(ctx context.Context) (bool, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LinkQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LinkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LinkQuery) Clone() *LinkQuery {
	if lq == nil {
		return nil
	}
	return &LinkQuery{
		config:     lq.config,
		limit:      lq.limit,
		offset:     lq.offset,
		order:      append([]OrderFunc{}, lq.order...),
		predicates: append([]predicate.Link{}, lq.predicates...),
		// clone intermediate query.
		sql:    lq.sql.Clone(),
		path:   lq.path,
		unique: lq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime int64 `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Link.Query().
//		GroupBy(link.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (lq *LinkQuery) GroupBy(field string, fields ...string) *LinkGroupBy {
	grbuild := &LinkGroupBy{config: lq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lq.sqlQuery(ctx), nil
	}
	grbuild.label = link.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime int64 `json:"create_time,omitempty"`
//	}
//
//	client.Link.Query().
//		Select(link.FieldCreateTime).
//		Scan(ctx, &v)
//
func (lq *LinkQuery) Select(fields ...string) *LinkSelect {
	lq.fields = append(lq.fields, fields...)
	selbuild := &LinkSelect{LinkQuery: lq}
	selbuild.label = link.Label
	selbuild.flds, selbuild.scan = &lq.fields, selbuild.Scan
	return selbuild
}

func (lq *LinkQuery) prepareQuery(ctx context.Context) error {
	for _, f := range lq.fields {
		if !link.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LinkQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Link, error) {
	var (
		nodes = []*Link{}
		_spec = lq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Link).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Link{config: lq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(lq.modifiers) > 0 {
		_spec.Modifiers = lq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lq *LinkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	if len(lq.modifiers) > 0 {
		_spec.Modifiers = lq.modifiers
	}
	_spec.Node.Columns = lq.fields
	if len(lq.fields) > 0 {
		_spec.Unique = lq.unique != nil && *lq.unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LinkQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := lq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (lq *LinkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   link.Table,
			Columns: link.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: link.FieldID,
			},
		},
		From:   lq.sql,
		Unique: true,
	}
	if unique := lq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, link.FieldID)
		for i := range fields {
			if fields[i] != link.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LinkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(link.Table)
	columns := lq.fields
	if len(columns) == 0 {
		columns = link.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.unique != nil && *lq.unique {
		selector.Distinct()
	}
	for _, m := range lq.modifiers {
		m(selector)
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lq *LinkQuery) Modify(modifiers ...func(s *sql.Selector)) *LinkSelect {
	lq.modifiers = append(lq.modifiers, modifiers...)
	return lq.Select()
}

// LinkGroupBy is the group-by builder for Link entities.
type LinkGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LinkGroupBy) Aggregate(fns ...AggregateFunc) *LinkGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the group-by query and scans the result into the given value.
func (lgb *LinkGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lgb.path(ctx)
	if err != nil {
		return err
	}
	lgb.sql = query
	return lgb.sqlScan(ctx, v)
}

func (lgb *LinkGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lgb.fields {
		if !link.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lgb *LinkGroupBy) sqlQuery() *sql.Selector {
	selector := lgb.sql.Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lgb.fields)+len(lgb.fns))
		for _, f := range lgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lgb.fields...)...)
}

// LinkSelect is the builder for selecting fields of Link entities.
type LinkSelect struct {
	*LinkQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LinkSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	ls.sql = ls.LinkQuery.sqlQuery(ctx)
	return ls.sqlScan(ctx, v)
}

func (ls *LinkSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ls.sql.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ls *LinkSelect) Modify(modifiers ...func(s *sql.Selector)) *LinkSelect {
	ls.modifiers = append(ls.modifiers, modifiers...)
	return ls
}
