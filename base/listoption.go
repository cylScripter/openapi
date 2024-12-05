package base

import (
	"reflect"
	"strconv"
	"strings"
)

func NewListOptionWithOption(opt ...*Option) *ListOption {
	p := NewListOption()
	option := make([]*Option, 0)
	for _, v := range opt {
		option = append(option, v)
	}
	p.SetOptions(option)
	return p
}

func (l *ListOption) AddOption(name int32, value string) {
	l.Options = append(l.Options, &Option{
		Name:  name,
		Value: value,
	})
}

const (
	ListOptionValueTypeNil            = 0
	ListOptionValueTypeUint32         = 1
	ListOptionValueTypeString         = 2
	ListOptionValueTypeUint32List     = 3
	ListOptionValueTypeUint64         = 4
	ListOptionValueTypeTimeStampRange = 5
	ListOptionValueTypeUint64List     = 6
	ListOptionValueTypeBool           = 7
	ListOptionValueTypeStringList     = 8
	ListOptionValueTypeInt32          = 9
)

type ListOptionHandler struct {
	typ              int
	cbNone           func() error
	cbUint32         func(val uint32) error
	cbInt32          func(val int32) error
	cbString         func(val string) error
	cbUint32List     func(valList []uint32) error
	cbUint64         func(val uint64) error
	cbTimeStampRange func(beginAt, endAt uint32) error
	cbUint64List     func(val []uint64) error
	cbBool           func(val bool) error
	cbStringList     func(val []string) error
	ignoreZeroValue  bool
}
type ListOptionProcessor struct {
	listOption *ListOption
	handlers   map[int32]*ListOptionHandler
}

func toInt32(i interface{}) int32 {
	t := reflect.TypeOf(i)
	k := t.Kind()
	switch k {
	case reflect.Int,
		reflect.Int32,
		reflect.Int64,
		reflect.Int16,
		reflect.Int8:
		return int32(reflect.ValueOf(i).Int())
	case reflect.Uint,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uint16,
		reflect.Uint8:
		return int32(reflect.ValueOf(i).Uint())
	default:
		return 0
	}
}

func NewListOptionProcessor(listOption *ListOption) *ListOptionProcessor {
	return &ListOptionProcessor{
		listOption: listOption,
		handlers:   make(map[int32]*ListOptionHandler),
	}
}

func (p *ListOptionProcessor) AddNone(typ interface{}, cb func() error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:    ListOptionValueTypeNil,
		cbNone: cb,
	}
	return p
}

func (p *ListOptionProcessor) AddUint32(typ interface{}, cb func(val uint32) error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:      ListOptionValueTypeUint32,
		cbUint32: cb,
	}
	return p
}

func (p *ListOptionProcessor) AddInt32(typ interface{}, cb func(val int32) error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:     ListOptionValueTypeInt32,
		cbInt32: cb,
	}
	return p
}

func (p *ListOptionProcessor) AddString(typ interface{}, cb func(val string) error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:             ListOptionValueTypeString,
		cbString:        cb,
		ignoreZeroValue: true,
	}
	return p
}

func (p *ListOptionProcessor) AddStringIgnoreZero(typ interface{}, cb func(val string) error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:      ListOptionValueTypeString,
		cbString: cb,
	}
	return p
}

func (p *ListOptionProcessor) AddStringList(typ interface{}, cb func(val []string) error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:             ListOptionValueTypeStringList,
		cbStringList:    cb,
		ignoreZeroValue: true,
	}
	return p
}

func (p *ListOptionProcessor) AddUint32List(typ interface{}, cb func(valList []uint32) error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:             ListOptionValueTypeUint32List,
		cbUint32List:    cb,
		ignoreZeroValue: true,
	}
	return p
}

func (p *ListOptionProcessor) AddUint64(typ interface{}, cb func(val uint64) error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:      ListOptionValueTypeUint64,
		cbUint64: cb,
	}
	return p
}

func (p *ListOptionProcessor) AddUint64List(typ interface{}, cb func(valList []uint64) error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:             ListOptionValueTypeUint64List,
		cbUint64List:    cb,
		ignoreZeroValue: true,
	}
	return p
}

func (p *ListOptionProcessor) AddTimeStampRange(typ interface{}, cb func(beginAt, endAt uint32) error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:              ListOptionValueTypeTimeStampRange,
		cbTimeStampRange: cb,
		ignoreZeroValue:  true,
	}
	return p
}

func (p *ListOptionProcessor) AddBool(typ interface{}, cb func(val bool) error) *ListOptionProcessor {
	x := toInt32(typ)
	p.handlers[x] = &ListOptionHandler{
		typ:    ListOptionValueTypeBool,
		cbBool: cb,
	}
	return p
}

func (p *ListOptionProcessor) Process() error {
	if p.listOption == nil || p.handlers == nil || len(p.handlers) == 0 {
		return nil
	}
	var err error
	for _, v := range p.listOption.Options {
		h := p.handlers[v.Name]
		if h == nil {
			continue
		}
		switch h.typ {
		case ListOptionValueTypeNil:
			if h.cbNone != nil {
				err = h.cbNone()
				if err != nil {
					return err
				}
			}

		case ListOptionValueTypeUint32:
			if v.Value == "" {
				continue
			}
			x, err := strconv.ParseInt(v.Value, 10, 32)
			if err != nil {
				return err
			}
			if h.cbUint32 != nil {
				err = h.cbUint32(uint32(x))
				if err != nil {
					return err
				}
			}

		case ListOptionValueTypeString:
			if v.Value == "" && h.ignoreZeroValue {
				continue
			}
			if h.cbString != nil {
				if err = h.cbString(v.Value); err != nil {
					return err
				}
			}

		case ListOptionValueTypeStringList:
			if v.Value == "" && h.ignoreZeroValue {
				continue
			}
			if h.cbStringList != nil {
				list := strings.Split(v.Value, ",")
				// 过滤掉空串
				var nonEmptyList []string
				for _, v := range list {
					if v != "" {
						nonEmptyList = append(nonEmptyList, v)
					}
				}
				if len(nonEmptyList) == 0 && h.ignoreZeroValue {
					continue
				}
				if err = h.cbStringList(nonEmptyList); err != nil {
					return err
				}
			}

		case ListOptionValueTypeUint32List:
			if v.Value == "" && h.ignoreZeroValue {
				continue
			}
			list := strings.Split(v.Value, ",")
			var intList []uint32
			for _, item := range list {
				x, err := strconv.ParseInt(item, 10, 32)
				if err != nil {
					return err
				}
				intList = append(intList, uint32(x))
			}
			if h.cbUint32List != nil {
				if err = h.cbUint32List(intList); err != nil {
					return err
				}
			}

		case ListOptionValueTypeUint64:
			if v.Value == "" {
				continue
			}
			x, err := strconv.ParseUint(v.Value, 10, 64)
			if err != nil {
				return err
			}
			if h.cbUint64 != nil {
				err = h.cbUint64(x)
				if err != nil {
					return err
				}
			}

		case ListOptionValueTypeTimeStampRange:
			var tStr []string
			if strings.Index(v.Value, ",") > 0 {
				tStr = strings.Split(v.Value, ",")
				if len(tStr) != 2 && h.ignoreZeroValue {
					continue
				}
			} else {
				tStr = strings.Split(v.Value, "-")
				if len(tStr) != 2 && h.ignoreZeroValue {
					continue
				}
			}
			t1, err := strconv.ParseUint(tStr[0], 10, 64)
			if err != nil {
				return err
			}
			t2, err := strconv.ParseUint(tStr[1], 10, 64)
			if err != nil {
				return err
			}
			if h.cbTimeStampRange != nil {
				if err = h.cbTimeStampRange(uint32(t1), uint32(t2)); err != nil {
					return err
				}
			}

		case ListOptionValueTypeUint64List:
			if v.Value == "" && h.ignoreZeroValue {
				continue
			}
			list := strings.Split(v.Value, ",")
			var intList []uint64
			for _, item := range list {
				x, err := strconv.ParseUint(item, 10, 64)
				if err != nil {
					return err
				}
				intList = append(intList, x)
			}
			if h.cbUint64List != nil {
				if err = h.cbUint64List(intList); err != nil {
					return err
				}
			}

		case ListOptionValueTypeBool:
			//if v.Value != "0" && v.Value != "1" {
			// continue
			//}
			value := strings.ToLower(v.Value)
			var x bool
			if InSliceStr(value, []string{"1", "true"}) {
				x = true
			} else if InSliceStr(value, []string{"0", "false"}) {
				x = false
			} else {
				continue
			}
			if h.cbBool != nil {
				err = h.cbBool(x)
				if err != nil {
					return err
				}
			}

		case ListOptionValueTypeInt32:
			if v.Value == "" {
				continue
			}
			x, err := strconv.ParseInt(v.Value, 10, 32)
			if err != nil {
				return err
			}
			if h.cbInt32 != nil {
				err = h.cbInt32(int32(x))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func InSliceStr(s string, list []string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}
