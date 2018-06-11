package errors

import (
	"errors"
)

var (
	ERR_BAD_VALUE                = errors.New("bad value")
	ERR_BAD_TYPE                 = errors.New("bad type")
	ERR_OVER_STACK_LEN           = errors.New("the count over the stack length")
	ERR_OVER_CODE_LEN            = errors.New("the count over the code length")
	ERR_UNDER_STACK_LEN          = errors.New("the count under the stack length")
	ERR_FAULT                    = errors.New("the exeution meet fault")
	ERR_NOT_SUPPORT_SERVICE      = errors.New("the service is not registered")
	ERR_NOT_SUPPORT_OPCODE       = errors.New("does not support the operation code")
	ERR_OVER_LIMIT_STACK         = errors.New("the stack over max size")
	ERR_OVER_MAX_ITEM_SIZE       = errors.New("the item over max size")
	ERR_OVER_MAX_ARRAY_SIZE      = errors.New("the array over max size")
	ERR_OVER_MAX_BIGINTEGER_SIZE = errors.New("the biginteger over max size 32bit")
	ERR_OUT_OF_GAS               = errors.New("out of gas")
	ERR_NOT_ARRAY                = errors.New("not array")
	ERR_TABLE_IS_NIL             = errors.New("table is nil")
	ERR_SERVICE_IS_NIL           = errors.New("service is nil")
	ERR_DIV_MOD_BY_ZERO          = errors.New("div or mod by zero")
	ERR_SHIFT_BY_NEG             = errors.New("shift by negtive value")
	ERR_EXECUTION_CONTEXT_NIL    = errors.New("execution context is nil")
	ERR_CURRENT_CONTEXT_NIL      = errors.New("current context is nil")
	ERR_CALLING_CONTEXT_NIL      = errors.New("calling context is nil")
	ERR_ENTRY_CONTEXT_NIL        = errors.New("entry context is nil")
	ERR_APPEND_NOT_ARRAY         = errors.New("append not array")
	ERR_NOT_SUPPORT_TYPE         = errors.New("not a supported type")
)

const callStackDepth = 10

type DetailError interface {
	error
	ErrCoder
	CallStacker
	GetRoot() error
}

func NewErr(errmsg string) error {
	return errors.New(errmsg)
}

func NewDetailErr(err error, errcode ErrCode, errmsg string) DetailError {
	if err == nil {
		return nil
	}

	onterr, ok := err.(ontError)
	if !ok {
		onterr.root = err
		onterr.errmsg = err.Error()
		onterr.callstack = getCallStack(0, callStackDepth)
		onterr.code = errcode

	}
	if errmsg != "" {
		onterr.errmsg = errmsg + ": " + onterr.errmsg
	}

	return onterr
}

func RootErr(err error) error {
	if err, ok := err.(DetailError); ok {
		return err.GetRoot()
	}
	return err
}
