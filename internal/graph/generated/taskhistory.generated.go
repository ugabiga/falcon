// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _TaskHistory_id(ctx context.Context, field graphql.CollectedField, obj *TaskHistory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskHistory_id(ctx, field)
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
		return obj.ID, nil
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
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TaskHistory_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskHistory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _TaskHistory_taskID(ctx context.Context, field graphql.CollectedField, obj *TaskHistory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskHistory_taskID(ctx, field)
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
		return obj.TaskID, nil
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
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TaskHistory_taskID(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskHistory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _TaskHistory_isSuccess(ctx context.Context, field graphql.CollectedField, obj *TaskHistory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskHistory_isSuccess(ctx, field)
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
		return obj.IsSuccess, nil
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
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TaskHistory_isSuccess(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskHistory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _TaskHistory_log(ctx context.Context, field graphql.CollectedField, obj *TaskHistory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskHistory_log(ctx, field)
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
		return obj.Log, nil
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

func (ec *executionContext) fieldContext_TaskHistory_log(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskHistory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _TaskHistory_updatedAt(ctx context.Context, field graphql.CollectedField, obj *TaskHistory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskHistory_updatedAt(ctx, field)
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
		return obj.UpdatedAt, nil
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
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TaskHistory_updatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskHistory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _TaskHistory_createdAt(ctx context.Context, field graphql.CollectedField, obj *TaskHistory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskHistory_createdAt(ctx, field)
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
		return obj.CreatedAt, nil
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
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TaskHistory_createdAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskHistory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _TaskHistory_task(ctx context.Context, field graphql.CollectedField, obj *TaskHistory) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskHistory_task(ctx, field)
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
		return obj.Task, nil
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
	res := resTmp.(*Task)
	fc.Result = res
	return ec.marshalNTask2ᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTask(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TaskHistory_task(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskHistory",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Task_id(ctx, field)
			case "tradingAccountID":
				return ec.fieldContext_Task_tradingAccountID(ctx, field)
			case "currency":
				return ec.fieldContext_Task_currency(ctx, field)
			case "size":
				return ec.fieldContext_Task_size(ctx, field)
			case "symbol":
				return ec.fieldContext_Task_symbol(ctx, field)
			case "cron":
				return ec.fieldContext_Task_cron(ctx, field)
			case "nextExecutionTime":
				return ec.fieldContext_Task_nextExecutionTime(ctx, field)
			case "isActive":
				return ec.fieldContext_Task_isActive(ctx, field)
			case "type":
				return ec.fieldContext_Task_type(ctx, field)
			case "params":
				return ec.fieldContext_Task_params(ctx, field)
			case "updatedAt":
				return ec.fieldContext_Task_updatedAt(ctx, field)
			case "createdAt":
				return ec.fieldContext_Task_createdAt(ctx, field)
			case "tradingAccount":
				return ec.fieldContext_Task_tradingAccount(ctx, field)
			case "taskHistories":
				return ec.fieldContext_Task_taskHistories(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Task", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _TaskHistoryIndex_task(ctx context.Context, field graphql.CollectedField, obj *TaskHistoryIndex) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskHistoryIndex_task(ctx, field)
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
		return obj.Task, nil
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
	res := resTmp.(*Task)
	fc.Result = res
	return ec.marshalNTask2ᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTask(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TaskHistoryIndex_task(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskHistoryIndex",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Task_id(ctx, field)
			case "tradingAccountID":
				return ec.fieldContext_Task_tradingAccountID(ctx, field)
			case "currency":
				return ec.fieldContext_Task_currency(ctx, field)
			case "size":
				return ec.fieldContext_Task_size(ctx, field)
			case "symbol":
				return ec.fieldContext_Task_symbol(ctx, field)
			case "cron":
				return ec.fieldContext_Task_cron(ctx, field)
			case "nextExecutionTime":
				return ec.fieldContext_Task_nextExecutionTime(ctx, field)
			case "isActive":
				return ec.fieldContext_Task_isActive(ctx, field)
			case "type":
				return ec.fieldContext_Task_type(ctx, field)
			case "params":
				return ec.fieldContext_Task_params(ctx, field)
			case "updatedAt":
				return ec.fieldContext_Task_updatedAt(ctx, field)
			case "createdAt":
				return ec.fieldContext_Task_createdAt(ctx, field)
			case "tradingAccount":
				return ec.fieldContext_Task_tradingAccount(ctx, field)
			case "taskHistories":
				return ec.fieldContext_Task_taskHistories(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Task", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _TaskHistoryIndex_taskHistories(ctx context.Context, field graphql.CollectedField, obj *TaskHistoryIndex) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskHistoryIndex_taskHistories(ctx, field)
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
		return obj.TaskHistories, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*TaskHistory)
	fc.Result = res
	return ec.marshalOTaskHistory2ᚕᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTaskHistoryᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TaskHistoryIndex_taskHistories(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskHistoryIndex",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_TaskHistory_id(ctx, field)
			case "taskID":
				return ec.fieldContext_TaskHistory_taskID(ctx, field)
			case "isSuccess":
				return ec.fieldContext_TaskHistory_isSuccess(ctx, field)
			case "log":
				return ec.fieldContext_TaskHistory_log(ctx, field)
			case "updatedAt":
				return ec.fieldContext_TaskHistory_updatedAt(ctx, field)
			case "createdAt":
				return ec.fieldContext_TaskHistory_createdAt(ctx, field)
			case "task":
				return ec.fieldContext_TaskHistory_task(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type TaskHistory", field.Name)
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

var taskHistoryImplementors = []string{"TaskHistory"}

func (ec *executionContext) _TaskHistory(ctx context.Context, sel ast.SelectionSet, obj *TaskHistory) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, taskHistoryImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("TaskHistory")
		case "id":

			out.Values[i] = ec._TaskHistory_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "taskID":

			out.Values[i] = ec._TaskHistory_taskID(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "isSuccess":

			out.Values[i] = ec._TaskHistory_isSuccess(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "log":

			out.Values[i] = ec._TaskHistory_log(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "updatedAt":

			out.Values[i] = ec._TaskHistory_updatedAt(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "createdAt":

			out.Values[i] = ec._TaskHistory_createdAt(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "task":

			out.Values[i] = ec._TaskHistory_task(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var taskHistoryIndexImplementors = []string{"TaskHistoryIndex"}

func (ec *executionContext) _TaskHistoryIndex(ctx context.Context, sel ast.SelectionSet, obj *TaskHistoryIndex) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, taskHistoryIndexImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("TaskHistoryIndex")
		case "task":

			out.Values[i] = ec._TaskHistoryIndex_task(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "taskHistories":

			out.Values[i] = ec._TaskHistoryIndex_taskHistories(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNTaskHistory2ᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTaskHistory(ctx context.Context, sel ast.SelectionSet, v *TaskHistory) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._TaskHistory(ctx, sel, v)
}

func (ec *executionContext) marshalNTaskHistoryIndex2githubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTaskHistoryIndex(ctx context.Context, sel ast.SelectionSet, v TaskHistoryIndex) graphql.Marshaler {
	return ec._TaskHistoryIndex(ctx, sel, &v)
}

func (ec *executionContext) marshalNTaskHistoryIndex2ᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTaskHistoryIndex(ctx context.Context, sel ast.SelectionSet, v *TaskHistoryIndex) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._TaskHistoryIndex(ctx, sel, v)
}

func (ec *executionContext) marshalOTaskHistory2ᚕᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTaskHistoryᚄ(ctx context.Context, sel ast.SelectionSet, v []*TaskHistory) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
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
			ret[i] = ec.marshalNTaskHistory2ᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTaskHistory(ctx, sel, v[i])
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
