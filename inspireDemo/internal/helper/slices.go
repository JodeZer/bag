package helper

type StringSlice struct {
	raw []string
}

func MakeStringSlice(param ...int) *StringSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &StringSlice{
		raw: make([]string, p1, p2),
	}
}

func MakeStringSliceFromRaw(raw []string) *StringSlice {
	if raw == nil {
		raw = make([]string, 0)
	}
	return &StringSlice{
		raw: raw,
	}
}

func (this *StringSlice) Append(vals ...string) {
	this.AppendIf(nil, vals...)
}

func (this *StringSlice) AppendSlice(slice *StringSlice) {
	slice.RangeVal(func(val string) {
		this.Append(val)
	})
}

func (this *StringSlice) AppendIf(filters ...StringFilter, vals ...string) {
	this.AppendSlice(MakeStringSliceFromRaw(vals).Filter(filter...))
}

func (this *StringSlice) Range(rangers ...StringRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *StringSlice) RangeVal(rangers ...StringValRanger) {
	this.Range(StringValRangerBatchToStringRanger(rangers...)...)
}

func (this *StringSlice) GetRaw() []string {
	return this.raw
}

func (this *StringSlice) Filter(filters ...StringFilter) *StringSlice {

	res := MakeStringSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val string) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *StringSlice) ValFilter(filters ...StringValFilter) *StringSlice {
	return this.Filter(StringValFilterBatchToStringFilter(filters...)...)
}

func (this *StringSlice) Map(mappers ...StringMapper) *StringSlice {

	res := MakeStringSlice(0, len(this.GetRaw()))

	this.Range(func(val string) {
		for i, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	res = tmp

	return res
}

func (this *StringSlice) GetValMendedSlice(mappers ...StringMapper) *StringSlice {
	return this.Map(StringValMenderBatchToStringMapper(mappers...)...)
}

// let it crash apis
func (this *StringSlice) First() string {
	return this.At(0)
}

func (this *StringSlice) Last() string {
	return this.At(this.Len() - 1)
}

func (this *StringSlice) Ats(is ...int) *StringSlice {
	res := MakeStringSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *StringSlice) Slice(left, right int) *StringSlice {
	return MakeStringSliceFromRaw(this.GetRaw()[left:right])
}

func (this *StringSlice) At(i int) string {
	return this.GetRaw()[i]
}

func (this *StringSlice) Len() int {
	return len(this.GetRaw())
}

func (this *StringSlice) Cap() int {
	return cap(this.GetRaw())
}
