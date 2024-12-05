package base

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
