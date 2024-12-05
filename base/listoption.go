package base

import (
	"fmt"
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

type ListOptionValueType int

const (
	ListOptionValueTypeNil ListOptionValueType = iota
	ListOptionValueTypeUint32
	ListOptionValueTypeString
	ListOptionValueTypeStringList
	ListOptionValueTypeUint32List
	ListOptionValueTypeUint64
	ListOptionValueTypeTimeStampRange
	ListOptionValueTypeUint64List
	ListOptionValueTypeBool
	ListOptionValueTypeInt32
)

type ListOptionHandler struct {
	typ              ListOptionValueType
	cbNone           func() error
	cbUint32         func(val uint32) error
	cbString         func(val string) error
	cbStringList     func(val []string) error
	cbUint32List     func(valList []uint32) error
	cbUint64         func(val uint64) error
	cbUint64List     func(valList []uint64) error
	cbTimeStampRange func(beginAt, endAt uint32) error
	cbBool           func(val bool) error
	ignoreZeroValue  bool
	cbInt32          func(val int32) error
}

type ListOptionProcessor struct {
	listOption *ListOption
	handlers   map[int32]*ListOptionHandler
}

func toInt32(typ interface{}) int32 {
	// Assuming typ is an int32 or can be converted to int32
	return typ.(int32)
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

// Process processes the list options and applies the corresponding callbacks.
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
		if v.Value == "" && h.ignoreZeroValue {
			continue
		}
		err = p.processOption(h, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// processOption processes a single option based on its type and applies the callback.
func (p *ListOptionProcessor) processOption(h *ListOptionHandler, v *Option) error {
	switch h.typ {
	case ListOptionValueTypeNil:
		if h.cbNone != nil {
			return h.cbNone()
		}
	case ListOptionValueTypeUint32:
		x, err := strconv.ParseInt(v.Value, 10, 32)
		if err != nil {
			return err
		}
		if h.cbUint32 != nil {
			return h.cbUint32(uint32(x))
		}
	case ListOptionValueTypeString:
		if h.cbString != nil {
			return h.cbString(v.Value)
		}
	case ListOptionValueTypeStringList:
		list := strings.Split(v.Value, ",")
		nonEmptyList := filterNonEmptyStrings(list)
		if len(nonEmptyList) == 0 && h.ignoreZeroValue {
			return nil
		}
		if h.cbStringList != nil {
			return h.cbStringList(nonEmptyList)
		}
	case ListOptionValueTypeUint32List:
		list := strings.Split(v.Value, ",")
		intList, err := parseUint32List(list)
		if err != nil {
			return err
		}
		if h.cbUint32List != nil {
			return h.cbUint32List(intList)
		}
	case ListOptionValueTypeUint64:
		x, err := strconv.ParseUint(v.Value, 10, 64)
		if err != nil {
			return err
		}
		if h.cbUint64 != nil {
			return h.cbUint64(x)
		}
	case ListOptionValueTypeTimeStampRange:
		tStr := splitTimeRange(v.Value)
		if len(tStr) != 2 && h.ignoreZeroValue {
			return nil
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
			return h.cbTimeStampRange(uint32(t1), uint32(t2))
		}
	case ListOptionValueTypeUint64List:
		list := strings.Split(v.Value, ",")
		intList, err := parseUint64List(list)
		if err != nil {
			return err
		}
		if h.cbUint64List != nil {
			return h.cbUint64List(intList)
		}
	case ListOptionValueTypeBool:
		value := strings.ToLower(v.Value)
		if !InSliceStr(value, []string{"1", "true", "0", "false"}) {
			return nil
		}
		x := value == "1" || value == "true"
		if h.cbBool != nil {
			return h.cbBool(x)
		}
	case ListOptionValueTypeInt32:
		x, err := strconv.ParseInt(v.Value, 10, 32)
		if err != nil {
			return err
		}
		if h.cbInt32 != nil {
			return h.cbInt32(int32(x))
		}
	default:
		return fmt.Errorf("unknown option type: %d", h.typ)
	}
	return nil
}

func filterNonEmptyStrings(list []string) []string {
	var nonEmptyList []string
	for _, v := range list {
		if v != "" {
			nonEmptyList = append(nonEmptyList, v)
		}
	}
	return nonEmptyList
}

func parseUint32List(list []string) ([]uint32, error) {
	var intList []uint32
	for _, item := range list {
		x, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			return nil, err
		}
		intList = append(intList, uint32(x))
	}
	return intList, nil
}

func parseUint64List(list []string) ([]uint64, error) {
	var intList []uint64
	for _, item := range list {
		x, err := strconv.ParseUint(item, 10, 64)
		if err != nil {
			return nil, err
		}
		intList = append(intList, x)
	}
	return intList, nil
}

func splitTimeRange(value string) []string {
	if strings.Contains(value, ",") {
		return strings.Split(value, ",")
	}
	return strings.Split(value, "-")
}

func InSliceStr(s string, list []string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}
