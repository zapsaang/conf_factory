package validates

import (
	"reflect"

	"github.com/zapsaang/conf_factory/pkg/sets"
	"github.com/zapsaang/conf_factory/utils/errors"
)

func GeneralEquals(target ...interface{}) func(args ...interface{}) error {
	return func(args ...interface{}) error {
		tLen, aLen := len(target), len(args)
		if tLen != aLen {
			return errors.ErrArgumentsLengthMisMatch
		} else if !reflect.DeepEqual(target, args) {
			return errors.ErrArgumentsContentMisMatch
		}
		return nil
	}
}

func Int64Equals(target int64) func(args ...interface{}) error {
	return func(args ...interface{}) error {
		aLen := len(args)
		if aLen != 1 {
			return errors.ErrArgumentsLengthMisMatch
		}
		v, ok := args[0].(int64)
		if !ok {
			return errors.ErrArgumentsTypeMisMatch
		}
		if v != target {
			return errors.ErrArgumentsContentMisMatch
		}
		return nil
	}
}

func StringEquals(target string) func(args ...interface{}) error {
	return func(args ...interface{}) error {
		aLen := len(args)
		if aLen != 1 {
			return errors.ErrArgumentsLengthMisMatch
		}
		v, ok := args[0].(string)
		if !ok {
			return errors.ErrArgumentsTypeMisMatch
		}
		if v != target {
			return errors.ErrArgumentsContentMisMatch
		}
		return nil
	}
}

func Int64In(target ...int64) func(args ...interface{}) error {
	targetSet := sets.New(target...)
	return func(args ...interface{}) error {
		for _, a := range args {
			v, ok := a.(int64)
			if !ok {
				return errors.ErrArgumentsTypeMisMatch
			}
			if !targetSet.Exists(v) {
				return errors.ErrArgumentsContentMisMatch
			}
		}
		return nil
	}
}

func StringIn(target ...string) func(args ...interface{}) error {
	targetSet := sets.New(target...)
	return func(args ...interface{}) error {
		for _, a := range args {
			v, ok := a.(string)
			if !ok {
				return errors.ErrArgumentsTypeMisMatch
			}
			if !targetSet.Exists(v) {
				return errors.ErrArgumentsContentMisMatch
			}
		}
		return nil
	}
}

func Equals[T sets.SetElement](target T) func(args ...interface{}) error {
	return func(args ...interface{}) error {
		aLen := len(args)
		if aLen != 1 {
			return errors.ErrArgumentsLengthMisMatch
		}
		v, ok := args[0].(T)
		if !ok {
			return errors.ErrArgumentsTypeMisMatch
		}
		if v != target {
			return errors.ErrArgumentsContentMisMatch
		}
		return nil
	}
}

func In[T sets.SetElement](target ...T) func(args ...interface{}) error {
	targetSet := sets.New(target...)
	return func(args ...interface{}) error {
		for _, a := range args {
			v, ok := a.(T)
			if !ok {
				return errors.ErrArgumentsTypeMisMatch
			}
			if !targetSet.Exists(v) {
				return errors.ErrArgumentsContentMisMatch
			}
		}
		return nil
	}
}
