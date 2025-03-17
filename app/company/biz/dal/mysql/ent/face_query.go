// Code generated by ent, DO NOT EDIT.

package ent

import (
	"company/biz/dal/mysql/ent/face"
	"company/biz/dal/mysql/ent/internal"
	"company/biz/dal/mysql/ent/member"
	"company/biz/dal/mysql/ent/predicate"
	"company/biz/dal/mysql/ent/user"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FaceQuery is the builder for querying Face entities.
type FaceQuery struct {
	config
	ctx             *QueryContext
	order           []face.OrderOption
	inters          []Interceptor
	predicates      []predicate.Face
	withMemberFaces *MemberQuery
	withUserFaces   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FaceQuery builder.
func (fq *FaceQuery) Where(ps ...predicate.Face) *FaceQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FaceQuery) Limit(limit int) *FaceQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FaceQuery) Offset(offset int) *FaceQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FaceQuery) Unique(unique bool) *FaceQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FaceQuery) Order(o ...face.OrderOption) *FaceQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryMemberFaces chains the current query on the "member_faces" edge.
func (fq *FaceQuery) QueryMemberFaces() *MemberQuery {
	query := (&MemberClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(face.Table, face.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, face.MemberFacesTable, face.MemberFacesColumn),
		)
		schemaConfig := fq.schemaConfig
		step.To.Schema = schemaConfig.Member
		step.Edge.Schema = schemaConfig.Face
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserFaces chains the current query on the "user_faces" edge.
func (fq *FaceQuery) QueryUserFaces() *UserQuery {
	query := (&UserClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(face.Table, face.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, face.UserFacesTable, face.UserFacesColumn),
		)
		schemaConfig := fq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Face
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Face entity from the query.
// Returns a *NotFoundError when no Face was found.
func (fq *FaceQuery) First(ctx context.Context) (*Face, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{face.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FaceQuery) FirstX(ctx context.Context) *Face {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Face ID from the query.
// Returns a *NotFoundError when no Face ID was found.
func (fq *FaceQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{face.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FaceQuery) FirstIDX(ctx context.Context) int64 {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Face entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Face entity is found.
// Returns a *NotFoundError when no Face entities are found.
func (fq *FaceQuery) Only(ctx context.Context) (*Face, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{face.Label}
	default:
		return nil, &NotSingularError{face.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FaceQuery) OnlyX(ctx context.Context) *Face {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Face ID in the query.
// Returns a *NotSingularError when more than one Face ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FaceQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{face.Label}
	default:
		err = &NotSingularError{face.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FaceQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Faces.
func (fq *FaceQuery) All(ctx context.Context) ([]*Face, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryAll)
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Face, *FaceQuery]()
	return withInterceptors[[]*Face](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FaceQuery) AllX(ctx context.Context) []*Face {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Face IDs.
func (fq *FaceQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryIDs)
	if err = fq.Select(face.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FaceQuery) IDsX(ctx context.Context) []int64 {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FaceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryCount)
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FaceQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FaceQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FaceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryExist)
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FaceQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FaceQuery) Clone() *FaceQuery {
	if fq == nil {
		return nil
	}
	return &FaceQuery{
		config:          fq.config,
		ctx:             fq.ctx.Clone(),
		order:           append([]face.OrderOption{}, fq.order...),
		inters:          append([]Interceptor{}, fq.inters...),
		predicates:      append([]predicate.Face{}, fq.predicates...),
		withMemberFaces: fq.withMemberFaces.Clone(),
		withUserFaces:   fq.withUserFaces.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithMemberFaces tells the query-builder to eager-load the nodes that are connected to
// the "member_faces" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FaceQuery) WithMemberFaces(opts ...func(*MemberQuery)) *FaceQuery {
	query := (&MemberClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withMemberFaces = query
	return fq
}

// WithUserFaces tells the query-builder to eager-load the nodes that are connected to
// the "user_faces" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FaceQuery) WithUserFaces(opts ...func(*UserQuery)) *FaceQuery {
	query := (&UserClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withUserFaces = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Face.Query().
//		GroupBy(face.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FaceQuery) GroupBy(field string, fields ...string) *FaceGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FaceGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = face.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Face.Query().
//		Select(face.FieldCreatedAt).
//		Scan(ctx, &v)
func (fq *FaceQuery) Select(fields ...string) *FaceSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FaceSelect{FaceQuery: fq}
	sbuild.label = face.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FaceSelect configured with the given aggregations.
func (fq *FaceQuery) Aggregate(fns ...AggregateFunc) *FaceSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FaceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !face.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FaceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Face, error) {
	var (
		nodes       = []*Face{}
		_spec       = fq.querySpec()
		loadedTypes = [2]bool{
			fq.withMemberFaces != nil,
			fq.withUserFaces != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Face).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Face{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = fq.schemaConfig.Face
	ctx = internal.NewSchemaConfigContext(ctx, fq.schemaConfig)
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withMemberFaces; query != nil {
		if err := fq.loadMemberFaces(ctx, query, nodes, nil,
			func(n *Face, e *Member) { n.Edges.MemberFaces = e }); err != nil {
			return nil, err
		}
	}
	if query := fq.withUserFaces; query != nil {
		if err := fq.loadUserFaces(ctx, query, nodes, nil,
			func(n *Face, e *User) { n.Edges.UserFaces = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FaceQuery) loadMemberFaces(ctx context.Context, query *MemberQuery, nodes []*Face, init func(*Face), assign func(*Face, *Member)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Face)
	for i := range nodes {
		fk := nodes[i].MemberID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(member.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fq *FaceQuery) loadUserFaces(ctx context.Context, query *UserQuery, nodes []*Face, init func(*Face), assign func(*Face, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*Face)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fq *FaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Schema = fq.schemaConfig.Face
	ctx = internal.NewSchemaConfigContext(ctx, fq.schemaConfig)
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(face.Table, face.Columns, sqlgraph.NewFieldSpec(face.FieldID, field.TypeInt64))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, face.FieldID)
		for i := range fields {
			if fields[i] != face.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if fq.withMemberFaces != nil {
			_spec.Node.AddColumnOnce(face.FieldMemberID)
		}
		if fq.withUserFaces != nil {
			_spec.Node.AddColumnOnce(face.FieldUserID)
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(face.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = face.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(fq.schemaConfig.Face)
	ctx = internal.NewSchemaConfigContext(ctx, fq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FaceGroupBy is the group-by builder for Face entities.
type FaceGroupBy struct {
	selector
	build *FaceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FaceGroupBy) Aggregate(fns ...AggregateFunc) *FaceGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FaceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, ent.OpQueryGroupBy)
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FaceQuery, *FaceGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FaceGroupBy) sqlScan(ctx context.Context, root *FaceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FaceSelect is the builder for selecting fields of Face entities.
type FaceSelect struct {
	*FaceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FaceSelect) Aggregate(fns ...AggregateFunc) *FaceSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FaceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, ent.OpQuerySelect)
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FaceQuery, *FaceSelect](ctx, fs.FaceQuery, fs, fs.inters, v)
}

func (fs *FaceSelect) sqlScan(ctx context.Context, root *FaceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
