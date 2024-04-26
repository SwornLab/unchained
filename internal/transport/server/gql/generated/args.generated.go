// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"github.com/TimeleapLabs/unchained/internal/model"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/TimeleapLabs/unchained/internal/transport/server/gql/types"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type EventLogArgResolver interface {
	Value(ctx context.Context, obj *model.EventLogArg) (string, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _EventLogArg_name(ctx context.Context, field graphql.CollectedField, obj *model.EventLogArg) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_EventLogArg_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_EventLogArg_name(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "EventLogArg",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _EventLogArg_value(ctx context.Context, field graphql.CollectedField, obj *model.EventLogArg) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_EventLogArg_value(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.EventLogArg().Value(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_EventLogArg_value(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "EventLogArg",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var eventLogArgImplementors = []string{"EventLogArg"}

func (ec *executionContext) _EventLogArg(ctx context.Context, sel ast.SelectionSet, obj *model.EventLogArg) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, eventLogArgImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("EventLogArg")
		case "name":
			out.Values[i] = ec._EventLogArg_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&out.Invalids, 1)
			}
		case "value":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._EventLogArg_value(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNBytes2githubᚗcomᚋTimeleapLabsᚋunchainedᚋinternalᚋtransportᚋserverᚋgqlᚋtypesᚐBytes(ctx context.Context, v interface{}) (types.Bytes, error) {
	var res types.Bytes
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNBytes2githubᚗcomᚋTimeleapLabsᚋunchainedᚋinternalᚋtransportᚋserverᚋgqlᚋtypesᚐBytes(ctx context.Context, sel ast.SelectionSet, v types.Bytes) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return v
}

func (ec *executionContext) marshalNEventLogArg2githubᚗcomᚋTimeleapLabsᚋunchainedᚋinternalᚋdatasetsᚐEventLogArg(ctx context.Context, sel ast.SelectionSet, v model.EventLogArg) graphql.Marshaler {
	return ec._EventLogArg(ctx, sel, &v)
}

func (ec *executionContext) marshalNEventLogArg2ᚕgithubᚗcomᚋTimeleapLabsᚋunchainedᚋinternalᚋdatasetsᚐEventLogArgᚄ(ctx context.Context, sel ast.SelectionSet, v []model.EventLogArg) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNEventLogArg2githubᚗcomᚋTimeleapLabsᚋunchainedᚋinternalᚋdatasetsᚐEventLogArg(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

// endregion ***************************** type.gotpl *****************************
