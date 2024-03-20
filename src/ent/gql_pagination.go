// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/KenshiTech/unchained/ent/assetprice"
	"github.com/KenshiTech/unchained/ent/correctnessreport"
	"github.com/KenshiTech/unchained/ent/eventlog"
	"github.com/KenshiTech/unchained/ent/signer"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// AssetPriceEdge is the edge representation of AssetPrice.
type AssetPriceEdge struct {
	Node   *AssetPrice `json:"node"`
	Cursor Cursor      `json:"cursor"`
}

// AssetPriceConnection is the connection containing edges to AssetPrice.
type AssetPriceConnection struct {
	Edges      []*AssetPriceEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

func (c *AssetPriceConnection) build(nodes []*AssetPrice, pager *assetpricePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *AssetPrice
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *AssetPrice {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *AssetPrice {
			return nodes[i]
		}
	}
	c.Edges = make([]*AssetPriceEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &AssetPriceEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// AssetPricePaginateOption enables pagination customization.
type AssetPricePaginateOption func(*assetpricePager) error

// WithAssetPriceOrder configures pagination ordering.
func WithAssetPriceOrder(order *AssetPriceOrder) AssetPricePaginateOption {
	if order == nil {
		order = DefaultAssetPriceOrder
	}
	o := *order
	return func(pager *assetpricePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultAssetPriceOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithAssetPriceFilter configures pagination filter.
func WithAssetPriceFilter(filter func(*AssetPriceQuery) (*AssetPriceQuery, error)) AssetPricePaginateOption {
	return func(pager *assetpricePager) error {
		if filter == nil {
			return errors.New("AssetPriceQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type assetpricePager struct {
	reverse bool
	order   *AssetPriceOrder
	filter  func(*AssetPriceQuery) (*AssetPriceQuery, error)
}

func newAssetPricePager(opts []AssetPricePaginateOption, reverse bool) (*assetpricePager, error) {
	pager := &assetpricePager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultAssetPriceOrder
	}
	return pager, nil
}

func (p *assetpricePager) applyFilter(query *AssetPriceQuery) (*AssetPriceQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *assetpricePager) toCursor(ap *AssetPrice) Cursor {
	return p.order.Field.toCursor(ap)
}

func (p *assetpricePager) applyCursors(query *AssetPriceQuery, after, before *Cursor) (*AssetPriceQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultAssetPriceOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *assetpricePager) applyOrder(query *AssetPriceQuery) *AssetPriceQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultAssetPriceOrder.Field {
		query = query.Order(DefaultAssetPriceOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *assetpricePager) orderExpr(query *AssetPriceQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultAssetPriceOrder.Field {
			b.Comma().Ident(DefaultAssetPriceOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to AssetPrice.
func (ap *AssetPriceQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...AssetPricePaginateOption,
) (*AssetPriceConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newAssetPricePager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if ap, err = pager.applyFilter(ap); err != nil {
		return nil, err
	}
	conn := &AssetPriceConnection{Edges: []*AssetPriceEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := ap.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if ap, err = pager.applyCursors(ap, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		ap.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := ap.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	ap = pager.applyOrder(ap)
	nodes, err := ap.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// AssetPriceOrderFieldBlock orders AssetPrice by block.
	AssetPriceOrderFieldBlock = &AssetPriceOrderField{
		Value: func(ap *AssetPrice) (ent.Value, error) {
			return ap.Block, nil
		},
		column: assetprice.FieldBlock,
		toTerm: assetprice.ByBlock,
		toCursor: func(ap *AssetPrice) Cursor {
			return Cursor{
				ID:    ap.ID,
				Value: ap.Block,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f AssetPriceOrderField) String() string {
	var str string
	switch f.column {
	case AssetPriceOrderFieldBlock.column:
		str = "BLOCK"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f AssetPriceOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *AssetPriceOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("AssetPriceOrderField %T must be a string", v)
	}
	switch str {
	case "BLOCK":
		*f = *AssetPriceOrderFieldBlock
	default:
		return fmt.Errorf("%s is not a valid AssetPriceOrderField", str)
	}
	return nil
}

// AssetPriceOrderField defines the ordering field of AssetPrice.
type AssetPriceOrderField struct {
	// Value extracts the ordering value from the given AssetPrice.
	Value    func(*AssetPrice) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) assetprice.OrderOption
	toCursor func(*AssetPrice) Cursor
}

// AssetPriceOrder defines the ordering of AssetPrice.
type AssetPriceOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *AssetPriceOrderField `json:"field"`
}

// DefaultAssetPriceOrder is the default ordering of AssetPrice.
var DefaultAssetPriceOrder = &AssetPriceOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &AssetPriceOrderField{
		Value: func(ap *AssetPrice) (ent.Value, error) {
			return ap.ID, nil
		},
		column: assetprice.FieldID,
		toTerm: assetprice.ByID,
		toCursor: func(ap *AssetPrice) Cursor {
			return Cursor{ID: ap.ID}
		},
	},
}

// ToEdge converts AssetPrice into AssetPriceEdge.
func (ap *AssetPrice) ToEdge(order *AssetPriceOrder) *AssetPriceEdge {
	if order == nil {
		order = DefaultAssetPriceOrder
	}
	return &AssetPriceEdge{
		Node:   ap,
		Cursor: order.Field.toCursor(ap),
	}
}

// CorrectnessReportEdge is the edge representation of CorrectnessReport.
type CorrectnessReportEdge struct {
	Node   *CorrectnessReport `json:"node"`
	Cursor Cursor             `json:"cursor"`
}

// CorrectnessReportConnection is the connection containing edges to CorrectnessReport.
type CorrectnessReportConnection struct {
	Edges      []*CorrectnessReportEdge `json:"edges"`
	PageInfo   PageInfo                 `json:"pageInfo"`
	TotalCount int                      `json:"totalCount"`
}

func (c *CorrectnessReportConnection) build(nodes []*CorrectnessReport, pager *correctnessreportPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *CorrectnessReport
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *CorrectnessReport {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *CorrectnessReport {
			return nodes[i]
		}
	}
	c.Edges = make([]*CorrectnessReportEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &CorrectnessReportEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// CorrectnessReportPaginateOption enables pagination customization.
type CorrectnessReportPaginateOption func(*correctnessreportPager) error

// WithCorrectnessReportOrder configures pagination ordering.
func WithCorrectnessReportOrder(order *CorrectnessReportOrder) CorrectnessReportPaginateOption {
	if order == nil {
		order = DefaultCorrectnessReportOrder
	}
	o := *order
	return func(pager *correctnessreportPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultCorrectnessReportOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithCorrectnessReportFilter configures pagination filter.
func WithCorrectnessReportFilter(filter func(*CorrectnessReportQuery) (*CorrectnessReportQuery, error)) CorrectnessReportPaginateOption {
	return func(pager *correctnessreportPager) error {
		if filter == nil {
			return errors.New("CorrectnessReportQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type correctnessreportPager struct {
	reverse bool
	order   *CorrectnessReportOrder
	filter  func(*CorrectnessReportQuery) (*CorrectnessReportQuery, error)
}

func newCorrectnessReportPager(opts []CorrectnessReportPaginateOption, reverse bool) (*correctnessreportPager, error) {
	pager := &correctnessreportPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultCorrectnessReportOrder
	}
	return pager, nil
}

func (p *correctnessreportPager) applyFilter(query *CorrectnessReportQuery) (*CorrectnessReportQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *correctnessreportPager) toCursor(cr *CorrectnessReport) Cursor {
	return p.order.Field.toCursor(cr)
}

func (p *correctnessreportPager) applyCursors(query *CorrectnessReportQuery, after, before *Cursor) (*CorrectnessReportQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultCorrectnessReportOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *correctnessreportPager) applyOrder(query *CorrectnessReportQuery) *CorrectnessReportQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultCorrectnessReportOrder.Field {
		query = query.Order(DefaultCorrectnessReportOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *correctnessreportPager) orderExpr(query *CorrectnessReportQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultCorrectnessReportOrder.Field {
			b.Comma().Ident(DefaultCorrectnessReportOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to CorrectnessReport.
func (cr *CorrectnessReportQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...CorrectnessReportPaginateOption,
) (*CorrectnessReportConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newCorrectnessReportPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if cr, err = pager.applyFilter(cr); err != nil {
		return nil, err
	}
	conn := &CorrectnessReportConnection{Edges: []*CorrectnessReportEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := cr.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if cr, err = pager.applyCursors(cr, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		cr.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := cr.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	cr = pager.applyOrder(cr)
	nodes, err := cr.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// CorrectnessReportOrderFieldTimestamp orders CorrectnessReport by timestamp.
	CorrectnessReportOrderFieldTimestamp = &CorrectnessReportOrderField{
		Value: func(cr *CorrectnessReport) (ent.Value, error) {
			return cr.Timestamp, nil
		},
		column: correctnessreport.FieldTimestamp,
		toTerm: correctnessreport.ByTimestamp,
		toCursor: func(cr *CorrectnessReport) Cursor {
			return Cursor{
				ID:    cr.ID,
				Value: cr.Timestamp,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f CorrectnessReportOrderField) String() string {
	var str string
	switch f.column {
	case CorrectnessReportOrderFieldTimestamp.column:
		str = "TIMESTAMP"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f CorrectnessReportOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *CorrectnessReportOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("CorrectnessReportOrderField %T must be a string", v)
	}
	switch str {
	case "TIMESTAMP":
		*f = *CorrectnessReportOrderFieldTimestamp
	default:
		return fmt.Errorf("%s is not a valid CorrectnessReportOrderField", str)
	}
	return nil
}

// CorrectnessReportOrderField defines the ordering field of CorrectnessReport.
type CorrectnessReportOrderField struct {
	// Value extracts the ordering value from the given CorrectnessReport.
	Value    func(*CorrectnessReport) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) correctnessreport.OrderOption
	toCursor func(*CorrectnessReport) Cursor
}

// CorrectnessReportOrder defines the ordering of CorrectnessReport.
type CorrectnessReportOrder struct {
	Direction OrderDirection               `json:"direction"`
	Field     *CorrectnessReportOrderField `json:"field"`
}

// DefaultCorrectnessReportOrder is the default ordering of CorrectnessReport.
var DefaultCorrectnessReportOrder = &CorrectnessReportOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &CorrectnessReportOrderField{
		Value: func(cr *CorrectnessReport) (ent.Value, error) {
			return cr.ID, nil
		},
		column: correctnessreport.FieldID,
		toTerm: correctnessreport.ByID,
		toCursor: func(cr *CorrectnessReport) Cursor {
			return Cursor{ID: cr.ID}
		},
	},
}

// ToEdge converts CorrectnessReport into CorrectnessReportEdge.
func (cr *CorrectnessReport) ToEdge(order *CorrectnessReportOrder) *CorrectnessReportEdge {
	if order == nil {
		order = DefaultCorrectnessReportOrder
	}
	return &CorrectnessReportEdge{
		Node:   cr,
		Cursor: order.Field.toCursor(cr),
	}
}

// EventLogEdge is the edge representation of EventLog.
type EventLogEdge struct {
	Node   *EventLog `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// EventLogConnection is the connection containing edges to EventLog.
type EventLogConnection struct {
	Edges      []*EventLogEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

func (c *EventLogConnection) build(nodes []*EventLog, pager *eventlogPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *EventLog
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *EventLog {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *EventLog {
			return nodes[i]
		}
	}
	c.Edges = make([]*EventLogEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &EventLogEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// EventLogPaginateOption enables pagination customization.
type EventLogPaginateOption func(*eventlogPager) error

// WithEventLogOrder configures pagination ordering.
func WithEventLogOrder(order *EventLogOrder) EventLogPaginateOption {
	if order == nil {
		order = DefaultEventLogOrder
	}
	o := *order
	return func(pager *eventlogPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultEventLogOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithEventLogFilter configures pagination filter.
func WithEventLogFilter(filter func(*EventLogQuery) (*EventLogQuery, error)) EventLogPaginateOption {
	return func(pager *eventlogPager) error {
		if filter == nil {
			return errors.New("EventLogQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type eventlogPager struct {
	reverse bool
	order   *EventLogOrder
	filter  func(*EventLogQuery) (*EventLogQuery, error)
}

func newEventLogPager(opts []EventLogPaginateOption, reverse bool) (*eventlogPager, error) {
	pager := &eventlogPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultEventLogOrder
	}
	return pager, nil
}

func (p *eventlogPager) applyFilter(query *EventLogQuery) (*EventLogQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *eventlogPager) toCursor(el *EventLog) Cursor {
	return p.order.Field.toCursor(el)
}

func (p *eventlogPager) applyCursors(query *EventLogQuery, after, before *Cursor) (*EventLogQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultEventLogOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *eventlogPager) applyOrder(query *EventLogQuery) *EventLogQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultEventLogOrder.Field {
		query = query.Order(DefaultEventLogOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *eventlogPager) orderExpr(query *EventLogQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultEventLogOrder.Field {
			b.Comma().Ident(DefaultEventLogOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to EventLog.
func (el *EventLogQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...EventLogPaginateOption,
) (*EventLogConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newEventLogPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if el, err = pager.applyFilter(el); err != nil {
		return nil, err
	}
	conn := &EventLogConnection{Edges: []*EventLogEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := el.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if el, err = pager.applyCursors(el, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		el.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := el.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	el = pager.applyOrder(el)
	nodes, err := el.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// EventLogOrderFieldBlock orders EventLog by block.
	EventLogOrderFieldBlock = &EventLogOrderField{
		Value: func(el *EventLog) (ent.Value, error) {
			return el.Block, nil
		},
		column: eventlog.FieldBlock,
		toTerm: eventlog.ByBlock,
		toCursor: func(el *EventLog) Cursor {
			return Cursor{
				ID:    el.ID,
				Value: el.Block,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f EventLogOrderField) String() string {
	var str string
	switch f.column {
	case EventLogOrderFieldBlock.column:
		str = "BLOCK"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f EventLogOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *EventLogOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("EventLogOrderField %T must be a string", v)
	}
	switch str {
	case "BLOCK":
		*f = *EventLogOrderFieldBlock
	default:
		return fmt.Errorf("%s is not a valid EventLogOrderField", str)
	}
	return nil
}

// EventLogOrderField defines the ordering field of EventLog.
type EventLogOrderField struct {
	// Value extracts the ordering value from the given EventLog.
	Value    func(*EventLog) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) eventlog.OrderOption
	toCursor func(*EventLog) Cursor
}

// EventLogOrder defines the ordering of EventLog.
type EventLogOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *EventLogOrderField `json:"field"`
}

// DefaultEventLogOrder is the default ordering of EventLog.
var DefaultEventLogOrder = &EventLogOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &EventLogOrderField{
		Value: func(el *EventLog) (ent.Value, error) {
			return el.ID, nil
		},
		column: eventlog.FieldID,
		toTerm: eventlog.ByID,
		toCursor: func(el *EventLog) Cursor {
			return Cursor{ID: el.ID}
		},
	},
}

// ToEdge converts EventLog into EventLogEdge.
func (el *EventLog) ToEdge(order *EventLogOrder) *EventLogEdge {
	if order == nil {
		order = DefaultEventLogOrder
	}
	return &EventLogEdge{
		Node:   el,
		Cursor: order.Field.toCursor(el),
	}
}

// SignerEdge is the edge representation of Signer.
type SignerEdge struct {
	Node   *Signer `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// SignerConnection is the connection containing edges to Signer.
type SignerConnection struct {
	Edges      []*SignerEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

func (c *SignerConnection) build(nodes []*Signer, pager *signerPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Signer
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Signer {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Signer {
			return nodes[i]
		}
	}
	c.Edges = make([]*SignerEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &SignerEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// SignerPaginateOption enables pagination customization.
type SignerPaginateOption func(*signerPager) error

// WithSignerOrder configures pagination ordering.
func WithSignerOrder(order *SignerOrder) SignerPaginateOption {
	if order == nil {
		order = DefaultSignerOrder
	}
	o := *order
	return func(pager *signerPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultSignerOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithSignerFilter configures pagination filter.
func WithSignerFilter(filter func(*SignerQuery) (*SignerQuery, error)) SignerPaginateOption {
	return func(pager *signerPager) error {
		if filter == nil {
			return errors.New("SignerQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type signerPager struct {
	reverse bool
	order   *SignerOrder
	filter  func(*SignerQuery) (*SignerQuery, error)
}

func newSignerPager(opts []SignerPaginateOption, reverse bool) (*signerPager, error) {
	pager := &signerPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultSignerOrder
	}
	return pager, nil
}

func (p *signerPager) applyFilter(query *SignerQuery) (*SignerQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *signerPager) toCursor(s *Signer) Cursor {
	return p.order.Field.toCursor(s)
}

func (p *signerPager) applyCursors(query *SignerQuery, after, before *Cursor) (*SignerQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultSignerOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *signerPager) applyOrder(query *SignerQuery) *SignerQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultSignerOrder.Field {
		query = query.Order(DefaultSignerOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *signerPager) orderExpr(query *SignerQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultSignerOrder.Field {
			b.Comma().Ident(DefaultSignerOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Signer.
func (s *SignerQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...SignerPaginateOption,
) (*SignerConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newSignerPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if s, err = pager.applyFilter(s); err != nil {
		return nil, err
	}
	conn := &SignerConnection{Edges: []*SignerEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := s.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if s, err = pager.applyCursors(s, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		s.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := s.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	s = pager.applyOrder(s)
	nodes, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// SignerOrderField defines the ordering field of Signer.
type SignerOrderField struct {
	// Value extracts the ordering value from the given Signer.
	Value    func(*Signer) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) signer.OrderOption
	toCursor func(*Signer) Cursor
}

// SignerOrder defines the ordering of Signer.
type SignerOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *SignerOrderField `json:"field"`
}

// DefaultSignerOrder is the default ordering of Signer.
var DefaultSignerOrder = &SignerOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &SignerOrderField{
		Value: func(s *Signer) (ent.Value, error) {
			return s.ID, nil
		},
		column: signer.FieldID,
		toTerm: signer.ByID,
		toCursor: func(s *Signer) Cursor {
			return Cursor{ID: s.ID}
		},
	},
}

// ToEdge converts Signer into SignerEdge.
func (s *Signer) ToEdge(order *SignerOrder) *SignerEdge {
	if order == nil {
		order = DefaultSignerOrder
	}
	return &SignerEdge{
		Node:   s,
		Cursor: order.Field.toCursor(s),
	}
}
