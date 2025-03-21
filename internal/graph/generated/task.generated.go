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
	"github.com/ugabiga/falcon/internal/graph/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Task_id(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_id(ctx, field)
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

func (ec *executionContext) fieldContext_Task_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_tradingAccountID(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_tradingAccountID(ctx, field)
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
		return obj.TradingAccountID, nil
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

func (ec *executionContext) fieldContext_Task_tradingAccountID(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_currency(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_currency(ctx, field)
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
		return obj.Currency, nil
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

func (ec *executionContext) fieldContext_Task_currency(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_size(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_size(ctx, field)
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
		return obj.Size, nil
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
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalNFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Task_size(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Float does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_symbol(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_symbol(ctx, field)
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
		return obj.Symbol, nil
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

func (ec *executionContext) fieldContext_Task_symbol(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_cron(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_cron(ctx, field)
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
		return obj.Cron, nil
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

func (ec *executionContext) fieldContext_Task_cron(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_nextExecutionTime(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_nextExecutionTime(ctx, field)
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
		return obj.NextExecutionTime, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*time.Time)
	fc.Result = res
	return ec.marshalOTime2ᚖtimeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Task_nextExecutionTime(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_isActive(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_isActive(ctx, field)
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
		return obj.IsActive, nil
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

func (ec *executionContext) fieldContext_Task_isActive(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_type(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_type(ctx, field)
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
		return obj.Type, nil
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

func (ec *executionContext) fieldContext_Task_type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_params(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_params(ctx, field)
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
		return obj.Params, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(model.JSON)
	fc.Result = res
	return ec.marshalOJSON2githubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋmodelᚐJSON(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Task_params(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type JSON does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_updatedAt(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_updatedAt(ctx, field)
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

func (ec *executionContext) fieldContext_Task_updatedAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_createdAt(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_createdAt(ctx, field)
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

func (ec *executionContext) fieldContext_Task_createdAt(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_tradingAccount(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_tradingAccount(ctx, field)
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
		return obj.TradingAccount, nil
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
	res := resTmp.(*TradingAccount)
	fc.Result = res
	return ec.marshalNTradingAccount2ᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTradingAccount(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Task_tradingAccount(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_TradingAccount_id(ctx, field)
			case "userID":
				return ec.fieldContext_TradingAccount_userID(ctx, field)
			case "name":
				return ec.fieldContext_TradingAccount_name(ctx, field)
			case "exchange":
				return ec.fieldContext_TradingAccount_exchange(ctx, field)
			case "ip":
				return ec.fieldContext_TradingAccount_ip(ctx, field)
			case "key":
				return ec.fieldContext_TradingAccount_key(ctx, field)
			case "updatedAt":
				return ec.fieldContext_TradingAccount_updatedAt(ctx, field)
			case "createdAt":
				return ec.fieldContext_TradingAccount_createdAt(ctx, field)
			case "user":
				return ec.fieldContext_TradingAccount_user(ctx, field)
			case "tasks":
				return ec.fieldContext_TradingAccount_tasks(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type TradingAccount", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Task_taskHistories(ctx context.Context, field graphql.CollectedField, obj *Task) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Task_taskHistories(ctx, field)
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

func (ec *executionContext) fieldContext_Task_taskHistories(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Task",
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

func (ec *executionContext) _TaskIndex_selectedTradingAccount(ctx context.Context, field graphql.CollectedField, obj *TaskIndex) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskIndex_selectedTradingAccount(ctx, field)
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
		return obj.SelectedTradingAccount, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*TradingAccount)
	fc.Result = res
	return ec.marshalOTradingAccount2ᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTradingAccount(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TaskIndex_selectedTradingAccount(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskIndex",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_TradingAccount_id(ctx, field)
			case "userID":
				return ec.fieldContext_TradingAccount_userID(ctx, field)
			case "name":
				return ec.fieldContext_TradingAccount_name(ctx, field)
			case "exchange":
				return ec.fieldContext_TradingAccount_exchange(ctx, field)
			case "ip":
				return ec.fieldContext_TradingAccount_ip(ctx, field)
			case "key":
				return ec.fieldContext_TradingAccount_key(ctx, field)
			case "updatedAt":
				return ec.fieldContext_TradingAccount_updatedAt(ctx, field)
			case "createdAt":
				return ec.fieldContext_TradingAccount_createdAt(ctx, field)
			case "user":
				return ec.fieldContext_TradingAccount_user(ctx, field)
			case "tasks":
				return ec.fieldContext_TradingAccount_tasks(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type TradingAccount", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _TaskIndex_tradingAccounts(ctx context.Context, field graphql.CollectedField, obj *TaskIndex) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_TaskIndex_tradingAccounts(ctx, field)
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
		return obj.TradingAccounts, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*TradingAccount)
	fc.Result = res
	return ec.marshalOTradingAccount2ᚕᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTradingAccountᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_TaskIndex_tradingAccounts(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "TaskIndex",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_TradingAccount_id(ctx, field)
			case "userID":
				return ec.fieldContext_TradingAccount_userID(ctx, field)
			case "name":
				return ec.fieldContext_TradingAccount_name(ctx, field)
			case "exchange":
				return ec.fieldContext_TradingAccount_exchange(ctx, field)
			case "ip":
				return ec.fieldContext_TradingAccount_ip(ctx, field)
			case "key":
				return ec.fieldContext_TradingAccount_key(ctx, field)
			case "updatedAt":
				return ec.fieldContext_TradingAccount_updatedAt(ctx, field)
			case "createdAt":
				return ec.fieldContext_TradingAccount_createdAt(ctx, field)
			case "user":
				return ec.fieldContext_TradingAccount_user(ctx, field)
			case "tasks":
				return ec.fieldContext_TradingAccount_tasks(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type TradingAccount", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputCreateTaskInput(ctx context.Context, obj interface{}) (CreateTaskInput, error) {
	var it CreateTaskInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	for k, v := range asMap {
		switch k {
		case "tradingAccountID":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("tradingAccountID"))
			it.TradingAccountID, err = ec.unmarshalNID2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "currency":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("currency"))
			it.Currency, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "size":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("size"))
			it.Size, err = ec.unmarshalNFloat2float64(ctx, v)
			if err != nil {
				return it, err
			}
		case "symbol":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("symbol"))
			it.Symbol, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "days":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("days"))
			it.Days, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "hours":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("hours"))
			it.Hours, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "type":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("type"))
			it.Type, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "params":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("params"))
			it.Params, err = ec.unmarshalOJSON2githubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋmodelᚐJSON(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputUpdateTaskInput(ctx context.Context, obj interface{}) (UpdateTaskInput, error) {
	var it UpdateTaskInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	for k, v := range asMap {
		switch k {
		case "currency":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("currency"))
			it.Currency, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "size":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("size"))
			it.Size, err = ec.unmarshalNFloat2float64(ctx, v)
			if err != nil {
				return it, err
			}
		case "symbol":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("symbol"))
			it.Symbol, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "days":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("days"))
			it.Days, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "hours":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("hours"))
			it.Hours, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "type":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("type"))
			it.Type, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "isActive":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("isActive"))
			it.IsActive, err = ec.unmarshalNBoolean2bool(ctx, v)
			if err != nil {
				return it, err
			}
		case "params":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("params"))
			it.Params, err = ec.unmarshalOJSON2githubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋmodelᚐJSON(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var taskImplementors = []string{"Task"}

func (ec *executionContext) _Task(ctx context.Context, sel ast.SelectionSet, obj *Task) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, taskImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Task")
		case "id":

			out.Values[i] = ec._Task_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "tradingAccountID":

			out.Values[i] = ec._Task_tradingAccountID(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "currency":

			out.Values[i] = ec._Task_currency(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "size":

			out.Values[i] = ec._Task_size(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "symbol":

			out.Values[i] = ec._Task_symbol(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "cron":

			out.Values[i] = ec._Task_cron(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "nextExecutionTime":

			out.Values[i] = ec._Task_nextExecutionTime(ctx, field, obj)

		case "isActive":

			out.Values[i] = ec._Task_isActive(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "type":

			out.Values[i] = ec._Task_type(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "params":

			out.Values[i] = ec._Task_params(ctx, field, obj)

		case "updatedAt":

			out.Values[i] = ec._Task_updatedAt(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "createdAt":

			out.Values[i] = ec._Task_createdAt(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "tradingAccount":

			out.Values[i] = ec._Task_tradingAccount(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "taskHistories":

			out.Values[i] = ec._Task_taskHistories(ctx, field, obj)

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

var taskIndexImplementors = []string{"TaskIndex"}

func (ec *executionContext) _TaskIndex(ctx context.Context, sel ast.SelectionSet, obj *TaskIndex) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, taskIndexImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("TaskIndex")
		case "selectedTradingAccount":

			out.Values[i] = ec._TaskIndex_selectedTradingAccount(ctx, field, obj)

		case "tradingAccounts":

			out.Values[i] = ec._TaskIndex_tradingAccounts(ctx, field, obj)

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

func (ec *executionContext) unmarshalNCreateTaskInput2githubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐCreateTaskInput(ctx context.Context, v interface{}) (CreateTaskInput, error) {
	res, err := ec.unmarshalInputCreateTaskInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNTask2githubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTask(ctx context.Context, sel ast.SelectionSet, v Task) graphql.Marshaler {
	return ec._Task(ctx, sel, &v)
}

func (ec *executionContext) marshalNTask2ᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTask(ctx context.Context, sel ast.SelectionSet, v *Task) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Task(ctx, sel, v)
}

func (ec *executionContext) unmarshalNUpdateTaskInput2githubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐUpdateTaskInput(ctx context.Context, v interface{}) (UpdateTaskInput, error) {
	res, err := ec.unmarshalInputUpdateTaskInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOJSON2githubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋmodelᚐJSON(ctx context.Context, v interface{}) (model.JSON, error) {
	if v == nil {
		return nil, nil
	}
	res, err := model.UnmarshalJSON(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOJSON2githubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋmodelᚐJSON(ctx context.Context, sel ast.SelectionSet, v model.JSON) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := model.MarshalJSON(v)
	return res
}

func (ec *executionContext) marshalOTask2ᚕᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTaskᚄ(ctx context.Context, sel ast.SelectionSet, v []*Task) graphql.Marshaler {
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
			ret[i] = ec.marshalNTask2ᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTask(ctx, sel, v[i])
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

func (ec *executionContext) marshalOTaskIndex2ᚖgithubᚗcomᚋugabigaᚋfalconᚋinternalᚋgraphᚋgeneratedᚐTaskIndex(ctx context.Context, sel ast.SelectionSet, v *TaskIndex) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._TaskIndex(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
