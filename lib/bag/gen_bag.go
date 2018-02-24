package bag

import ()

// auto generate by bag

// generate slice

// autogen @type: string  @aliasPreix: String

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
	this.raw = append(this.raw, vals...)
}

func (this *StringSlice) AppendSlice(slice *StringSlice) {
	slice.RangeVal(func(val string) {
		this.Append(val)
	})
}

func (this *StringSlice) AppendIf(filter StringFilter, vals ...string) {
	this.AppendSlice(MakeStringSliceFromRaw(vals).Filter(filter))
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

func (this *StringSlice) FilterVal(filters ...StringValFilter) *StringSlice {
	return this.Filter(StringValFilterBatchToStringFilter(filters...)...)
}

func (this *StringSlice) Map(mappers ...StringMapper) *StringSlice {

	res := MakeStringSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val string) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *StringSlice) MapVal(mappers ...StringValMapper) *StringSlice {
	return this.Map(StringValMapperBatchToStringMapper(mappers...)...)
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

// autogen @type: int  @aliasPreix: Int

type IntSlice struct {
	raw []int
}

func MakeIntSlice(param ...int) *IntSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &IntSlice{
		raw: make([]int, p1, p2),
	}
}

func MakeIntSliceFromRaw(raw []int) *IntSlice {
	if raw == nil {
		raw = make([]int, 0)
	}
	return &IntSlice{
		raw: raw,
	}
}

func (this *IntSlice) Append(vals ...int) {
	this.raw = append(this.raw, vals...)
}

func (this *IntSlice) AppendSlice(slice *IntSlice) {
	slice.RangeVal(func(val int) {
		this.Append(val)
	})
}

func (this *IntSlice) AppendIf(filter IntFilter, vals ...int) {
	this.AppendSlice(MakeIntSliceFromRaw(vals).Filter(filter))
}

func (this *IntSlice) Range(rangers ...IntRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *IntSlice) RangeVal(rangers ...IntValRanger) {
	this.Range(IntValRangerBatchToIntRanger(rangers...)...)
}

func (this *IntSlice) GetRaw() []int {
	return this.raw
}

func (this *IntSlice) Filter(filters ...IntFilter) *IntSlice {

	res := MakeIntSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val int) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *IntSlice) FilterVal(filters ...IntValFilter) *IntSlice {
	return this.Filter(IntValFilterBatchToIntFilter(filters...)...)
}

func (this *IntSlice) Map(mappers ...IntMapper) *IntSlice {

	res := MakeIntSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val int) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *IntSlice) MapVal(mappers ...IntValMapper) *IntSlice {
	return this.Map(IntValMapperBatchToIntMapper(mappers...)...)
}

// let it crash apis
func (this *IntSlice) First() int {
	return this.At(0)
}

func (this *IntSlice) Last() int {
	return this.At(this.Len() - 1)
}

func (this *IntSlice) Ats(is ...int) *IntSlice {
	res := MakeIntSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *IntSlice) Slice(left, right int) *IntSlice {
	return MakeIntSliceFromRaw(this.GetRaw()[left:right])
}

func (this *IntSlice) At(i int) int {
	return this.GetRaw()[i]
}

func (this *IntSlice) Len() int {
	return len(this.GetRaw())
}

func (this *IntSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: int64  @aliasPreix: Int64

type Int64Slice struct {
	raw []int64
}

func MakeInt64Slice(param ...int) *Int64Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Int64Slice{
		raw: make([]int64, p1, p2),
	}
}

func MakeInt64SliceFromRaw(raw []int64) *Int64Slice {
	if raw == nil {
		raw = make([]int64, 0)
	}
	return &Int64Slice{
		raw: raw,
	}
}

func (this *Int64Slice) Append(vals ...int64) {
	this.raw = append(this.raw, vals...)
}

func (this *Int64Slice) AppendSlice(slice *Int64Slice) {
	slice.RangeVal(func(val int64) {
		this.Append(val)
	})
}

func (this *Int64Slice) AppendIf(filter Int64Filter, vals ...int64) {
	this.AppendSlice(MakeInt64SliceFromRaw(vals).Filter(filter))
}

func (this *Int64Slice) Range(rangers ...Int64Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Int64Slice) RangeVal(rangers ...Int64ValRanger) {
	this.Range(Int64ValRangerBatchToInt64Ranger(rangers...)...)
}

func (this *Int64Slice) GetRaw() []int64 {
	return this.raw
}

func (this *Int64Slice) Filter(filters ...Int64Filter) *Int64Slice {

	res := MakeInt64Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val int64) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Int64Slice) FilterVal(filters ...Int64ValFilter) *Int64Slice {
	return this.Filter(Int64ValFilterBatchToInt64Filter(filters...)...)
}

func (this *Int64Slice) Map(mappers ...Int64Mapper) *Int64Slice {

	res := MakeInt64Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val int64) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Int64Slice) MapVal(mappers ...Int64ValMapper) *Int64Slice {
	return this.Map(Int64ValMapperBatchToInt64Mapper(mappers...)...)
}

// let it crash apis
func (this *Int64Slice) First() int64 {
	return this.At(0)
}

func (this *Int64Slice) Last() int64 {
	return this.At(this.Len() - 1)
}

func (this *Int64Slice) Ats(is ...int) *Int64Slice {
	res := MakeInt64Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Int64Slice) Slice(left, right int) *Int64Slice {
	return MakeInt64SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Int64Slice) At(i int) int64 {
	return this.GetRaw()[i]
}

func (this *Int64Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Int64Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: int32  @aliasPreix: Int32

type Int32Slice struct {
	raw []int32
}

func MakeInt32Slice(param ...int) *Int32Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Int32Slice{
		raw: make([]int32, p1, p2),
	}
}

func MakeInt32SliceFromRaw(raw []int32) *Int32Slice {
	if raw == nil {
		raw = make([]int32, 0)
	}
	return &Int32Slice{
		raw: raw,
	}
}

func (this *Int32Slice) Append(vals ...int32) {
	this.raw = append(this.raw, vals...)
}

func (this *Int32Slice) AppendSlice(slice *Int32Slice) {
	slice.RangeVal(func(val int32) {
		this.Append(val)
	})
}

func (this *Int32Slice) AppendIf(filter Int32Filter, vals ...int32) {
	this.AppendSlice(MakeInt32SliceFromRaw(vals).Filter(filter))
}

func (this *Int32Slice) Range(rangers ...Int32Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Int32Slice) RangeVal(rangers ...Int32ValRanger) {
	this.Range(Int32ValRangerBatchToInt32Ranger(rangers...)...)
}

func (this *Int32Slice) GetRaw() []int32 {
	return this.raw
}

func (this *Int32Slice) Filter(filters ...Int32Filter) *Int32Slice {

	res := MakeInt32Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val int32) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Int32Slice) FilterVal(filters ...Int32ValFilter) *Int32Slice {
	return this.Filter(Int32ValFilterBatchToInt32Filter(filters...)...)
}

func (this *Int32Slice) Map(mappers ...Int32Mapper) *Int32Slice {

	res := MakeInt32Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val int32) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Int32Slice) MapVal(mappers ...Int32ValMapper) *Int32Slice {
	return this.Map(Int32ValMapperBatchToInt32Mapper(mappers...)...)
}

// let it crash apis
func (this *Int32Slice) First() int32 {
	return this.At(0)
}

func (this *Int32Slice) Last() int32 {
	return this.At(this.Len() - 1)
}

func (this *Int32Slice) Ats(is ...int) *Int32Slice {
	res := MakeInt32Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Int32Slice) Slice(left, right int) *Int32Slice {
	return MakeInt32SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Int32Slice) At(i int) int32 {
	return this.GetRaw()[i]
}

func (this *Int32Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Int32Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: int16  @aliasPreix: Int16

type Int16Slice struct {
	raw []int16
}

func MakeInt16Slice(param ...int) *Int16Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Int16Slice{
		raw: make([]int16, p1, p2),
	}
}

func MakeInt16SliceFromRaw(raw []int16) *Int16Slice {
	if raw == nil {
		raw = make([]int16, 0)
	}
	return &Int16Slice{
		raw: raw,
	}
}

func (this *Int16Slice) Append(vals ...int16) {
	this.raw = append(this.raw, vals...)
}

func (this *Int16Slice) AppendSlice(slice *Int16Slice) {
	slice.RangeVal(func(val int16) {
		this.Append(val)
	})
}

func (this *Int16Slice) AppendIf(filter Int16Filter, vals ...int16) {
	this.AppendSlice(MakeInt16SliceFromRaw(vals).Filter(filter))
}

func (this *Int16Slice) Range(rangers ...Int16Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Int16Slice) RangeVal(rangers ...Int16ValRanger) {
	this.Range(Int16ValRangerBatchToInt16Ranger(rangers...)...)
}

func (this *Int16Slice) GetRaw() []int16 {
	return this.raw
}

func (this *Int16Slice) Filter(filters ...Int16Filter) *Int16Slice {

	res := MakeInt16Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val int16) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Int16Slice) FilterVal(filters ...Int16ValFilter) *Int16Slice {
	return this.Filter(Int16ValFilterBatchToInt16Filter(filters...)...)
}

func (this *Int16Slice) Map(mappers ...Int16Mapper) *Int16Slice {

	res := MakeInt16Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val int16) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Int16Slice) MapVal(mappers ...Int16ValMapper) *Int16Slice {
	return this.Map(Int16ValMapperBatchToInt16Mapper(mappers...)...)
}

// let it crash apis
func (this *Int16Slice) First() int16 {
	return this.At(0)
}

func (this *Int16Slice) Last() int16 {
	return this.At(this.Len() - 1)
}

func (this *Int16Slice) Ats(is ...int) *Int16Slice {
	res := MakeInt16Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Int16Slice) Slice(left, right int) *Int16Slice {
	return MakeInt16SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Int16Slice) At(i int) int16 {
	return this.GetRaw()[i]
}

func (this *Int16Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Int16Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: int8  @aliasPreix: Int8

type Int8Slice struct {
	raw []int8
}

func MakeInt8Slice(param ...int) *Int8Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Int8Slice{
		raw: make([]int8, p1, p2),
	}
}

func MakeInt8SliceFromRaw(raw []int8) *Int8Slice {
	if raw == nil {
		raw = make([]int8, 0)
	}
	return &Int8Slice{
		raw: raw,
	}
}

func (this *Int8Slice) Append(vals ...int8) {
	this.raw = append(this.raw, vals...)
}

func (this *Int8Slice) AppendSlice(slice *Int8Slice) {
	slice.RangeVal(func(val int8) {
		this.Append(val)
	})
}

func (this *Int8Slice) AppendIf(filter Int8Filter, vals ...int8) {
	this.AppendSlice(MakeInt8SliceFromRaw(vals).Filter(filter))
}

func (this *Int8Slice) Range(rangers ...Int8Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Int8Slice) RangeVal(rangers ...Int8ValRanger) {
	this.Range(Int8ValRangerBatchToInt8Ranger(rangers...)...)
}

func (this *Int8Slice) GetRaw() []int8 {
	return this.raw
}

func (this *Int8Slice) Filter(filters ...Int8Filter) *Int8Slice {

	res := MakeInt8Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val int8) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Int8Slice) FilterVal(filters ...Int8ValFilter) *Int8Slice {
	return this.Filter(Int8ValFilterBatchToInt8Filter(filters...)...)
}

func (this *Int8Slice) Map(mappers ...Int8Mapper) *Int8Slice {

	res := MakeInt8Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val int8) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Int8Slice) MapVal(mappers ...Int8ValMapper) *Int8Slice {
	return this.Map(Int8ValMapperBatchToInt8Mapper(mappers...)...)
}

// let it crash apis
func (this *Int8Slice) First() int8 {
	return this.At(0)
}

func (this *Int8Slice) Last() int8 {
	return this.At(this.Len() - 1)
}

func (this *Int8Slice) Ats(is ...int) *Int8Slice {
	res := MakeInt8Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Int8Slice) Slice(left, right int) *Int8Slice {
	return MakeInt8SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Int8Slice) At(i int) int8 {
	return this.GetRaw()[i]
}

func (this *Int8Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Int8Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: uint  @aliasPreix: Uint

type UintSlice struct {
	raw []uint
}

func MakeUintSlice(param ...int) *UintSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &UintSlice{
		raw: make([]uint, p1, p2),
	}
}

func MakeUintSliceFromRaw(raw []uint) *UintSlice {
	if raw == nil {
		raw = make([]uint, 0)
	}
	return &UintSlice{
		raw: raw,
	}
}

func (this *UintSlice) Append(vals ...uint) {
	this.raw = append(this.raw, vals...)
}

func (this *UintSlice) AppendSlice(slice *UintSlice) {
	slice.RangeVal(func(val uint) {
		this.Append(val)
	})
}

func (this *UintSlice) AppendIf(filter UintFilter, vals ...uint) {
	this.AppendSlice(MakeUintSliceFromRaw(vals).Filter(filter))
}

func (this *UintSlice) Range(rangers ...UintRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *UintSlice) RangeVal(rangers ...UintValRanger) {
	this.Range(UintValRangerBatchToUintRanger(rangers...)...)
}

func (this *UintSlice) GetRaw() []uint {
	return this.raw
}

func (this *UintSlice) Filter(filters ...UintFilter) *UintSlice {

	res := MakeUintSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val uint) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *UintSlice) FilterVal(filters ...UintValFilter) *UintSlice {
	return this.Filter(UintValFilterBatchToUintFilter(filters...)...)
}

func (this *UintSlice) Map(mappers ...UintMapper) *UintSlice {

	res := MakeUintSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val uint) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *UintSlice) MapVal(mappers ...UintValMapper) *UintSlice {
	return this.Map(UintValMapperBatchToUintMapper(mappers...)...)
}

// let it crash apis
func (this *UintSlice) First() uint {
	return this.At(0)
}

func (this *UintSlice) Last() uint {
	return this.At(this.Len() - 1)
}

func (this *UintSlice) Ats(is ...int) *UintSlice {
	res := MakeUintSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *UintSlice) Slice(left, right int) *UintSlice {
	return MakeUintSliceFromRaw(this.GetRaw()[left:right])
}

func (this *UintSlice) At(i int) uint {
	return this.GetRaw()[i]
}

func (this *UintSlice) Len() int {
	return len(this.GetRaw())
}

func (this *UintSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: uint8  @aliasPreix: Uint8

type Uint8Slice struct {
	raw []uint8
}

func MakeUint8Slice(param ...int) *Uint8Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Uint8Slice{
		raw: make([]uint8, p1, p2),
	}
}

func MakeUint8SliceFromRaw(raw []uint8) *Uint8Slice {
	if raw == nil {
		raw = make([]uint8, 0)
	}
	return &Uint8Slice{
		raw: raw,
	}
}

func (this *Uint8Slice) Append(vals ...uint8) {
	this.raw = append(this.raw, vals...)
}

func (this *Uint8Slice) AppendSlice(slice *Uint8Slice) {
	slice.RangeVal(func(val uint8) {
		this.Append(val)
	})
}

func (this *Uint8Slice) AppendIf(filter Uint8Filter, vals ...uint8) {
	this.AppendSlice(MakeUint8SliceFromRaw(vals).Filter(filter))
}

func (this *Uint8Slice) Range(rangers ...Uint8Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Uint8Slice) RangeVal(rangers ...Uint8ValRanger) {
	this.Range(Uint8ValRangerBatchToUint8Ranger(rangers...)...)
}

func (this *Uint8Slice) GetRaw() []uint8 {
	return this.raw
}

func (this *Uint8Slice) Filter(filters ...Uint8Filter) *Uint8Slice {

	res := MakeUint8Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val uint8) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Uint8Slice) FilterVal(filters ...Uint8ValFilter) *Uint8Slice {
	return this.Filter(Uint8ValFilterBatchToUint8Filter(filters...)...)
}

func (this *Uint8Slice) Map(mappers ...Uint8Mapper) *Uint8Slice {

	res := MakeUint8Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val uint8) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Uint8Slice) MapVal(mappers ...Uint8ValMapper) *Uint8Slice {
	return this.Map(Uint8ValMapperBatchToUint8Mapper(mappers...)...)
}

// let it crash apis
func (this *Uint8Slice) First() uint8 {
	return this.At(0)
}

func (this *Uint8Slice) Last() uint8 {
	return this.At(this.Len() - 1)
}

func (this *Uint8Slice) Ats(is ...int) *Uint8Slice {
	res := MakeUint8Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Uint8Slice) Slice(left, right int) *Uint8Slice {
	return MakeUint8SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Uint8Slice) At(i int) uint8 {
	return this.GetRaw()[i]
}

func (this *Uint8Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Uint8Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: uint16  @aliasPreix: Uint16

type Uint16Slice struct {
	raw []uint16
}

func MakeUint16Slice(param ...int) *Uint16Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Uint16Slice{
		raw: make([]uint16, p1, p2),
	}
}

func MakeUint16SliceFromRaw(raw []uint16) *Uint16Slice {
	if raw == nil {
		raw = make([]uint16, 0)
	}
	return &Uint16Slice{
		raw: raw,
	}
}

func (this *Uint16Slice) Append(vals ...uint16) {
	this.raw = append(this.raw, vals...)
}

func (this *Uint16Slice) AppendSlice(slice *Uint16Slice) {
	slice.RangeVal(func(val uint16) {
		this.Append(val)
	})
}

func (this *Uint16Slice) AppendIf(filter Uint16Filter, vals ...uint16) {
	this.AppendSlice(MakeUint16SliceFromRaw(vals).Filter(filter))
}

func (this *Uint16Slice) Range(rangers ...Uint16Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Uint16Slice) RangeVal(rangers ...Uint16ValRanger) {
	this.Range(Uint16ValRangerBatchToUint16Ranger(rangers...)...)
}

func (this *Uint16Slice) GetRaw() []uint16 {
	return this.raw
}

func (this *Uint16Slice) Filter(filters ...Uint16Filter) *Uint16Slice {

	res := MakeUint16Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val uint16) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Uint16Slice) FilterVal(filters ...Uint16ValFilter) *Uint16Slice {
	return this.Filter(Uint16ValFilterBatchToUint16Filter(filters...)...)
}

func (this *Uint16Slice) Map(mappers ...Uint16Mapper) *Uint16Slice {

	res := MakeUint16Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val uint16) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Uint16Slice) MapVal(mappers ...Uint16ValMapper) *Uint16Slice {
	return this.Map(Uint16ValMapperBatchToUint16Mapper(mappers...)...)
}

// let it crash apis
func (this *Uint16Slice) First() uint16 {
	return this.At(0)
}

func (this *Uint16Slice) Last() uint16 {
	return this.At(this.Len() - 1)
}

func (this *Uint16Slice) Ats(is ...int) *Uint16Slice {
	res := MakeUint16Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Uint16Slice) Slice(left, right int) *Uint16Slice {
	return MakeUint16SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Uint16Slice) At(i int) uint16 {
	return this.GetRaw()[i]
}

func (this *Uint16Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Uint16Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: uint32  @aliasPreix: Uint32

type Uint32Slice struct {
	raw []uint32
}

func MakeUint32Slice(param ...int) *Uint32Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Uint32Slice{
		raw: make([]uint32, p1, p2),
	}
}

func MakeUint32SliceFromRaw(raw []uint32) *Uint32Slice {
	if raw == nil {
		raw = make([]uint32, 0)
	}
	return &Uint32Slice{
		raw: raw,
	}
}

func (this *Uint32Slice) Append(vals ...uint32) {
	this.raw = append(this.raw, vals...)
}

func (this *Uint32Slice) AppendSlice(slice *Uint32Slice) {
	slice.RangeVal(func(val uint32) {
		this.Append(val)
	})
}

func (this *Uint32Slice) AppendIf(filter Uint32Filter, vals ...uint32) {
	this.AppendSlice(MakeUint32SliceFromRaw(vals).Filter(filter))
}

func (this *Uint32Slice) Range(rangers ...Uint32Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Uint32Slice) RangeVal(rangers ...Uint32ValRanger) {
	this.Range(Uint32ValRangerBatchToUint32Ranger(rangers...)...)
}

func (this *Uint32Slice) GetRaw() []uint32 {
	return this.raw
}

func (this *Uint32Slice) Filter(filters ...Uint32Filter) *Uint32Slice {

	res := MakeUint32Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val uint32) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Uint32Slice) FilterVal(filters ...Uint32ValFilter) *Uint32Slice {
	return this.Filter(Uint32ValFilterBatchToUint32Filter(filters...)...)
}

func (this *Uint32Slice) Map(mappers ...Uint32Mapper) *Uint32Slice {

	res := MakeUint32Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val uint32) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Uint32Slice) MapVal(mappers ...Uint32ValMapper) *Uint32Slice {
	return this.Map(Uint32ValMapperBatchToUint32Mapper(mappers...)...)
}

// let it crash apis
func (this *Uint32Slice) First() uint32 {
	return this.At(0)
}

func (this *Uint32Slice) Last() uint32 {
	return this.At(this.Len() - 1)
}

func (this *Uint32Slice) Ats(is ...int) *Uint32Slice {
	res := MakeUint32Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Uint32Slice) Slice(left, right int) *Uint32Slice {
	return MakeUint32SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Uint32Slice) At(i int) uint32 {
	return this.GetRaw()[i]
}

func (this *Uint32Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Uint32Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: uint64  @aliasPreix: Uint64

type Uint64Slice struct {
	raw []uint64
}

func MakeUint64Slice(param ...int) *Uint64Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Uint64Slice{
		raw: make([]uint64, p1, p2),
	}
}

func MakeUint64SliceFromRaw(raw []uint64) *Uint64Slice {
	if raw == nil {
		raw = make([]uint64, 0)
	}
	return &Uint64Slice{
		raw: raw,
	}
}

func (this *Uint64Slice) Append(vals ...uint64) {
	this.raw = append(this.raw, vals...)
}

func (this *Uint64Slice) AppendSlice(slice *Uint64Slice) {
	slice.RangeVal(func(val uint64) {
		this.Append(val)
	})
}

func (this *Uint64Slice) AppendIf(filter Uint64Filter, vals ...uint64) {
	this.AppendSlice(MakeUint64SliceFromRaw(vals).Filter(filter))
}

func (this *Uint64Slice) Range(rangers ...Uint64Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Uint64Slice) RangeVal(rangers ...Uint64ValRanger) {
	this.Range(Uint64ValRangerBatchToUint64Ranger(rangers...)...)
}

func (this *Uint64Slice) GetRaw() []uint64 {
	return this.raw
}

func (this *Uint64Slice) Filter(filters ...Uint64Filter) *Uint64Slice {

	res := MakeUint64Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val uint64) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Uint64Slice) FilterVal(filters ...Uint64ValFilter) *Uint64Slice {
	return this.Filter(Uint64ValFilterBatchToUint64Filter(filters...)...)
}

func (this *Uint64Slice) Map(mappers ...Uint64Mapper) *Uint64Slice {

	res := MakeUint64Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val uint64) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Uint64Slice) MapVal(mappers ...Uint64ValMapper) *Uint64Slice {
	return this.Map(Uint64ValMapperBatchToUint64Mapper(mappers...)...)
}

// let it crash apis
func (this *Uint64Slice) First() uint64 {
	return this.At(0)
}

func (this *Uint64Slice) Last() uint64 {
	return this.At(this.Len() - 1)
}

func (this *Uint64Slice) Ats(is ...int) *Uint64Slice {
	res := MakeUint64Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Uint64Slice) Slice(left, right int) *Uint64Slice {
	return MakeUint64SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Uint64Slice) At(i int) uint64 {
	return this.GetRaw()[i]
}

func (this *Uint64Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Uint64Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: float32  @aliasPreix: Float32

type Float32Slice struct {
	raw []float32
}

func MakeFloat32Slice(param ...int) *Float32Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Float32Slice{
		raw: make([]float32, p1, p2),
	}
}

func MakeFloat32SliceFromRaw(raw []float32) *Float32Slice {
	if raw == nil {
		raw = make([]float32, 0)
	}
	return &Float32Slice{
		raw: raw,
	}
}

func (this *Float32Slice) Append(vals ...float32) {
	this.raw = append(this.raw, vals...)
}

func (this *Float32Slice) AppendSlice(slice *Float32Slice) {
	slice.RangeVal(func(val float32) {
		this.Append(val)
	})
}

func (this *Float32Slice) AppendIf(filter Float32Filter, vals ...float32) {
	this.AppendSlice(MakeFloat32SliceFromRaw(vals).Filter(filter))
}

func (this *Float32Slice) Range(rangers ...Float32Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Float32Slice) RangeVal(rangers ...Float32ValRanger) {
	this.Range(Float32ValRangerBatchToFloat32Ranger(rangers...)...)
}

func (this *Float32Slice) GetRaw() []float32 {
	return this.raw
}

func (this *Float32Slice) Filter(filters ...Float32Filter) *Float32Slice {

	res := MakeFloat32Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val float32) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Float32Slice) FilterVal(filters ...Float32ValFilter) *Float32Slice {
	return this.Filter(Float32ValFilterBatchToFloat32Filter(filters...)...)
}

func (this *Float32Slice) Map(mappers ...Float32Mapper) *Float32Slice {

	res := MakeFloat32Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val float32) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Float32Slice) MapVal(mappers ...Float32ValMapper) *Float32Slice {
	return this.Map(Float32ValMapperBatchToFloat32Mapper(mappers...)...)
}

// let it crash apis
func (this *Float32Slice) First() float32 {
	return this.At(0)
}

func (this *Float32Slice) Last() float32 {
	return this.At(this.Len() - 1)
}

func (this *Float32Slice) Ats(is ...int) *Float32Slice {
	res := MakeFloat32Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Float32Slice) Slice(left, right int) *Float32Slice {
	return MakeFloat32SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Float32Slice) At(i int) float32 {
	return this.GetRaw()[i]
}

func (this *Float32Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Float32Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: float64  @aliasPreix: Float64

type Float64Slice struct {
	raw []float64
}

func MakeFloat64Slice(param ...int) *Float64Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Float64Slice{
		raw: make([]float64, p1, p2),
	}
}

func MakeFloat64SliceFromRaw(raw []float64) *Float64Slice {
	if raw == nil {
		raw = make([]float64, 0)
	}
	return &Float64Slice{
		raw: raw,
	}
}

func (this *Float64Slice) Append(vals ...float64) {
	this.raw = append(this.raw, vals...)
}

func (this *Float64Slice) AppendSlice(slice *Float64Slice) {
	slice.RangeVal(func(val float64) {
		this.Append(val)
	})
}

func (this *Float64Slice) AppendIf(filter Float64Filter, vals ...float64) {
	this.AppendSlice(MakeFloat64SliceFromRaw(vals).Filter(filter))
}

func (this *Float64Slice) Range(rangers ...Float64Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Float64Slice) RangeVal(rangers ...Float64ValRanger) {
	this.Range(Float64ValRangerBatchToFloat64Ranger(rangers...)...)
}

func (this *Float64Slice) GetRaw() []float64 {
	return this.raw
}

func (this *Float64Slice) Filter(filters ...Float64Filter) *Float64Slice {

	res := MakeFloat64Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val float64) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Float64Slice) FilterVal(filters ...Float64ValFilter) *Float64Slice {
	return this.Filter(Float64ValFilterBatchToFloat64Filter(filters...)...)
}

func (this *Float64Slice) Map(mappers ...Float64Mapper) *Float64Slice {

	res := MakeFloat64Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val float64) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Float64Slice) MapVal(mappers ...Float64ValMapper) *Float64Slice {
	return this.Map(Float64ValMapperBatchToFloat64Mapper(mappers...)...)
}

// let it crash apis
func (this *Float64Slice) First() float64 {
	return this.At(0)
}

func (this *Float64Slice) Last() float64 {
	return this.At(this.Len() - 1)
}

func (this *Float64Slice) Ats(is ...int) *Float64Slice {
	res := MakeFloat64Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Float64Slice) Slice(left, right int) *Float64Slice {
	return MakeFloat64SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Float64Slice) At(i int) float64 {
	return this.GetRaw()[i]
}

func (this *Float64Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Float64Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: bool  @aliasPreix: Bool

type BoolSlice struct {
	raw []bool
}

func MakeBoolSlice(param ...int) *BoolSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &BoolSlice{
		raw: make([]bool, p1, p2),
	}
}

func MakeBoolSliceFromRaw(raw []bool) *BoolSlice {
	if raw == nil {
		raw = make([]bool, 0)
	}
	return &BoolSlice{
		raw: raw,
	}
}

func (this *BoolSlice) Append(vals ...bool) {
	this.raw = append(this.raw, vals...)
}

func (this *BoolSlice) AppendSlice(slice *BoolSlice) {
	slice.RangeVal(func(val bool) {
		this.Append(val)
	})
}

func (this *BoolSlice) AppendIf(filter BoolFilter, vals ...bool) {
	this.AppendSlice(MakeBoolSliceFromRaw(vals).Filter(filter))
}

func (this *BoolSlice) Range(rangers ...BoolRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *BoolSlice) RangeVal(rangers ...BoolValRanger) {
	this.Range(BoolValRangerBatchToBoolRanger(rangers...)...)
}

func (this *BoolSlice) GetRaw() []bool {
	return this.raw
}

func (this *BoolSlice) Filter(filters ...BoolFilter) *BoolSlice {

	res := MakeBoolSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val bool) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *BoolSlice) FilterVal(filters ...BoolValFilter) *BoolSlice {
	return this.Filter(BoolValFilterBatchToBoolFilter(filters...)...)
}

func (this *BoolSlice) Map(mappers ...BoolMapper) *BoolSlice {

	res := MakeBoolSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val bool) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *BoolSlice) MapVal(mappers ...BoolValMapper) *BoolSlice {
	return this.Map(BoolValMapperBatchToBoolMapper(mappers...)...)
}

// let it crash apis
func (this *BoolSlice) First() bool {
	return this.At(0)
}

func (this *BoolSlice) Last() bool {
	return this.At(this.Len() - 1)
}

func (this *BoolSlice) Ats(is ...int) *BoolSlice {
	res := MakeBoolSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *BoolSlice) Slice(left, right int) *BoolSlice {
	return MakeBoolSliceFromRaw(this.GetRaw()[left:right])
}

func (this *BoolSlice) At(i int) bool {
	return this.GetRaw()[i]
}

func (this *BoolSlice) Len() int {
	return len(this.GetRaw())
}

func (this *BoolSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: rune  @aliasPreix: Rune

type RuneSlice struct {
	raw []rune
}

func MakeRuneSlice(param ...int) *RuneSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &RuneSlice{
		raw: make([]rune, p1, p2),
	}
}

func MakeRuneSliceFromRaw(raw []rune) *RuneSlice {
	if raw == nil {
		raw = make([]rune, 0)
	}
	return &RuneSlice{
		raw: raw,
	}
}

func (this *RuneSlice) Append(vals ...rune) {
	this.raw = append(this.raw, vals...)
}

func (this *RuneSlice) AppendSlice(slice *RuneSlice) {
	slice.RangeVal(func(val rune) {
		this.Append(val)
	})
}

func (this *RuneSlice) AppendIf(filter RuneFilter, vals ...rune) {
	this.AppendSlice(MakeRuneSliceFromRaw(vals).Filter(filter))
}

func (this *RuneSlice) Range(rangers ...RuneRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *RuneSlice) RangeVal(rangers ...RuneValRanger) {
	this.Range(RuneValRangerBatchToRuneRanger(rangers...)...)
}

func (this *RuneSlice) GetRaw() []rune {
	return this.raw
}

func (this *RuneSlice) Filter(filters ...RuneFilter) *RuneSlice {

	res := MakeRuneSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val rune) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *RuneSlice) FilterVal(filters ...RuneValFilter) *RuneSlice {
	return this.Filter(RuneValFilterBatchToRuneFilter(filters...)...)
}

func (this *RuneSlice) Map(mappers ...RuneMapper) *RuneSlice {

	res := MakeRuneSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val rune) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *RuneSlice) MapVal(mappers ...RuneValMapper) *RuneSlice {
	return this.Map(RuneValMapperBatchToRuneMapper(mappers...)...)
}

// let it crash apis
func (this *RuneSlice) First() rune {
	return this.At(0)
}

func (this *RuneSlice) Last() rune {
	return this.At(this.Len() - 1)
}

func (this *RuneSlice) Ats(is ...int) *RuneSlice {
	res := MakeRuneSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *RuneSlice) Slice(left, right int) *RuneSlice {
	return MakeRuneSliceFromRaw(this.GetRaw()[left:right])
}

func (this *RuneSlice) At(i int) rune {
	return this.GetRaw()[i]
}

func (this *RuneSlice) Len() int {
	return len(this.GetRaw())
}

func (this *RuneSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: complex64  @aliasPreix: Complex64

type Complex64Slice struct {
	raw []complex64
}

func MakeComplex64Slice(param ...int) *Complex64Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Complex64Slice{
		raw: make([]complex64, p1, p2),
	}
}

func MakeComplex64SliceFromRaw(raw []complex64) *Complex64Slice {
	if raw == nil {
		raw = make([]complex64, 0)
	}
	return &Complex64Slice{
		raw: raw,
	}
}

func (this *Complex64Slice) Append(vals ...complex64) {
	this.raw = append(this.raw, vals...)
}

func (this *Complex64Slice) AppendSlice(slice *Complex64Slice) {
	slice.RangeVal(func(val complex64) {
		this.Append(val)
	})
}

func (this *Complex64Slice) AppendIf(filter Complex64Filter, vals ...complex64) {
	this.AppendSlice(MakeComplex64SliceFromRaw(vals).Filter(filter))
}

func (this *Complex64Slice) Range(rangers ...Complex64Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Complex64Slice) RangeVal(rangers ...Complex64ValRanger) {
	this.Range(Complex64ValRangerBatchToComplex64Ranger(rangers...)...)
}

func (this *Complex64Slice) GetRaw() []complex64 {
	return this.raw
}

func (this *Complex64Slice) Filter(filters ...Complex64Filter) *Complex64Slice {

	res := MakeComplex64Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val complex64) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Complex64Slice) FilterVal(filters ...Complex64ValFilter) *Complex64Slice {
	return this.Filter(Complex64ValFilterBatchToComplex64Filter(filters...)...)
}

func (this *Complex64Slice) Map(mappers ...Complex64Mapper) *Complex64Slice {

	res := MakeComplex64Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val complex64) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Complex64Slice) MapVal(mappers ...Complex64ValMapper) *Complex64Slice {
	return this.Map(Complex64ValMapperBatchToComplex64Mapper(mappers...)...)
}

// let it crash apis
func (this *Complex64Slice) First() complex64 {
	return this.At(0)
}

func (this *Complex64Slice) Last() complex64 {
	return this.At(this.Len() - 1)
}

func (this *Complex64Slice) Ats(is ...int) *Complex64Slice {
	res := MakeComplex64Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Complex64Slice) Slice(left, right int) *Complex64Slice {
	return MakeComplex64SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Complex64Slice) At(i int) complex64 {
	return this.GetRaw()[i]
}

func (this *Complex64Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Complex64Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: complex32  @aliasPreix: Complex32

type Complex32Slice struct {
	raw []complex32
}

func MakeComplex32Slice(param ...int) *Complex32Slice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Complex32Slice{
		raw: make([]complex32, p1, p2),
	}
}

func MakeComplex32SliceFromRaw(raw []complex32) *Complex32Slice {
	if raw == nil {
		raw = make([]complex32, 0)
	}
	return &Complex32Slice{
		raw: raw,
	}
}

func (this *Complex32Slice) Append(vals ...complex32) {
	this.raw = append(this.raw, vals...)
}

func (this *Complex32Slice) AppendSlice(slice *Complex32Slice) {
	slice.RangeVal(func(val complex32) {
		this.Append(val)
	})
}

func (this *Complex32Slice) AppendIf(filter Complex32Filter, vals ...complex32) {
	this.AppendSlice(MakeComplex32SliceFromRaw(vals).Filter(filter))
}

func (this *Complex32Slice) Range(rangers ...Complex32Ranger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Complex32Slice) RangeVal(rangers ...Complex32ValRanger) {
	this.Range(Complex32ValRangerBatchToComplex32Ranger(rangers...)...)
}

func (this *Complex32Slice) GetRaw() []complex32 {
	return this.raw
}

func (this *Complex32Slice) Filter(filters ...Complex32Filter) *Complex32Slice {

	res := MakeComplex32Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val complex32) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Complex32Slice) FilterVal(filters ...Complex32ValFilter) *Complex32Slice {
	return this.Filter(Complex32ValFilterBatchToComplex32Filter(filters...)...)
}

func (this *Complex32Slice) Map(mappers ...Complex32Mapper) *Complex32Slice {

	res := MakeComplex32Slice(0, len(this.GetRaw()))

	this.Range(func(i int, val complex32) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Complex32Slice) MapVal(mappers ...Complex32ValMapper) *Complex32Slice {
	return this.Map(Complex32ValMapperBatchToComplex32Mapper(mappers...)...)
}

// let it crash apis
func (this *Complex32Slice) First() complex32 {
	return this.At(0)
}

func (this *Complex32Slice) Last() complex32 {
	return this.At(this.Len() - 1)
}

func (this *Complex32Slice) Ats(is ...int) *Complex32Slice {
	res := MakeComplex32Slice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Complex32Slice) Slice(left, right int) *Complex32Slice {
	return MakeComplex32SliceFromRaw(this.GetRaw()[left:right])
}

func (this *Complex32Slice) At(i int) complex32 {
	return this.GetRaw()[i]
}

func (this *Complex32Slice) Len() int {
	return len(this.GetRaw())
}

func (this *Complex32Slice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *string  @aliasPreix: StringPtr

type StringPtrSlice struct {
	raw []*string
}

func MakeStringPtrSlice(param ...int) *StringPtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &StringPtrSlice{
		raw: make([]*string, p1, p2),
	}
}

func MakeStringPtrSliceFromRaw(raw []*string) *StringPtrSlice {
	if raw == nil {
		raw = make([]*string, 0)
	}
	return &StringPtrSlice{
		raw: raw,
	}
}

func (this *StringPtrSlice) Append(vals ...*string) {
	this.raw = append(this.raw, vals...)
}

func (this *StringPtrSlice) AppendSlice(slice *StringPtrSlice) {
	slice.RangeVal(func(val *string) {
		this.Append(val)
	})
}

func (this *StringPtrSlice) AppendIf(filter StringPtrFilter, vals ...*string) {
	this.AppendSlice(MakeStringPtrSliceFromRaw(vals).Filter(filter))
}

func (this *StringPtrSlice) Range(rangers ...StringPtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *StringPtrSlice) RangeVal(rangers ...StringPtrValRanger) {
	this.Range(StringPtrValRangerBatchToStringPtrRanger(rangers...)...)
}

func (this *StringPtrSlice) GetRaw() []*string {
	return this.raw
}

func (this *StringPtrSlice) Filter(filters ...StringPtrFilter) *StringPtrSlice {

	res := MakeStringPtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *string) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *StringPtrSlice) FilterVal(filters ...StringPtrValFilter) *StringPtrSlice {
	return this.Filter(StringPtrValFilterBatchToStringPtrFilter(filters...)...)
}

func (this *StringPtrSlice) Map(mappers ...StringPtrMapper) *StringPtrSlice {

	res := MakeStringPtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *string) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *StringPtrSlice) MapVal(mappers ...StringPtrValMapper) *StringPtrSlice {
	return this.Map(StringPtrValMapperBatchToStringPtrMapper(mappers...)...)
}

// let it crash apis
func (this *StringPtrSlice) First() *string {
	return this.At(0)
}

func (this *StringPtrSlice) Last() *string {
	return this.At(this.Len() - 1)
}

func (this *StringPtrSlice) Ats(is ...int) *StringPtrSlice {
	res := MakeStringPtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *StringPtrSlice) Slice(left, right int) *StringPtrSlice {
	return MakeStringPtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *StringPtrSlice) At(i int) *string {
	return this.GetRaw()[i]
}

func (this *StringPtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *StringPtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *int  @aliasPreix: IntPtr

type IntPtrSlice struct {
	raw []*int
}

func MakeIntPtrSlice(param ...int) *IntPtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &IntPtrSlice{
		raw: make([]*int, p1, p2),
	}
}

func MakeIntPtrSliceFromRaw(raw []*int) *IntPtrSlice {
	if raw == nil {
		raw = make([]*int, 0)
	}
	return &IntPtrSlice{
		raw: raw,
	}
}

func (this *IntPtrSlice) Append(vals ...*int) {
	this.raw = append(this.raw, vals...)
}

func (this *IntPtrSlice) AppendSlice(slice *IntPtrSlice) {
	slice.RangeVal(func(val *int) {
		this.Append(val)
	})
}

func (this *IntPtrSlice) AppendIf(filter IntPtrFilter, vals ...*int) {
	this.AppendSlice(MakeIntPtrSliceFromRaw(vals).Filter(filter))
}

func (this *IntPtrSlice) Range(rangers ...IntPtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *IntPtrSlice) RangeVal(rangers ...IntPtrValRanger) {
	this.Range(IntPtrValRangerBatchToIntPtrRanger(rangers...)...)
}

func (this *IntPtrSlice) GetRaw() []*int {
	return this.raw
}

func (this *IntPtrSlice) Filter(filters ...IntPtrFilter) *IntPtrSlice {

	res := MakeIntPtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *int) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *IntPtrSlice) FilterVal(filters ...IntPtrValFilter) *IntPtrSlice {
	return this.Filter(IntPtrValFilterBatchToIntPtrFilter(filters...)...)
}

func (this *IntPtrSlice) Map(mappers ...IntPtrMapper) *IntPtrSlice {

	res := MakeIntPtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *int) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *IntPtrSlice) MapVal(mappers ...IntPtrValMapper) *IntPtrSlice {
	return this.Map(IntPtrValMapperBatchToIntPtrMapper(mappers...)...)
}

// let it crash apis
func (this *IntPtrSlice) First() *int {
	return this.At(0)
}

func (this *IntPtrSlice) Last() *int {
	return this.At(this.Len() - 1)
}

func (this *IntPtrSlice) Ats(is ...int) *IntPtrSlice {
	res := MakeIntPtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *IntPtrSlice) Slice(left, right int) *IntPtrSlice {
	return MakeIntPtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *IntPtrSlice) At(i int) *int {
	return this.GetRaw()[i]
}

func (this *IntPtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *IntPtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *int64  @aliasPreix: Int64Ptr

type Int64PtrSlice struct {
	raw []*int64
}

func MakeInt64PtrSlice(param ...int) *Int64PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Int64PtrSlice{
		raw: make([]*int64, p1, p2),
	}
}

func MakeInt64PtrSliceFromRaw(raw []*int64) *Int64PtrSlice {
	if raw == nil {
		raw = make([]*int64, 0)
	}
	return &Int64PtrSlice{
		raw: raw,
	}
}

func (this *Int64PtrSlice) Append(vals ...*int64) {
	this.raw = append(this.raw, vals...)
}

func (this *Int64PtrSlice) AppendSlice(slice *Int64PtrSlice) {
	slice.RangeVal(func(val *int64) {
		this.Append(val)
	})
}

func (this *Int64PtrSlice) AppendIf(filter Int64PtrFilter, vals ...*int64) {
	this.AppendSlice(MakeInt64PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Int64PtrSlice) Range(rangers ...Int64PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Int64PtrSlice) RangeVal(rangers ...Int64PtrValRanger) {
	this.Range(Int64PtrValRangerBatchToInt64PtrRanger(rangers...)...)
}

func (this *Int64PtrSlice) GetRaw() []*int64 {
	return this.raw
}

func (this *Int64PtrSlice) Filter(filters ...Int64PtrFilter) *Int64PtrSlice {

	res := MakeInt64PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *int64) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Int64PtrSlice) FilterVal(filters ...Int64PtrValFilter) *Int64PtrSlice {
	return this.Filter(Int64PtrValFilterBatchToInt64PtrFilter(filters...)...)
}

func (this *Int64PtrSlice) Map(mappers ...Int64PtrMapper) *Int64PtrSlice {

	res := MakeInt64PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *int64) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Int64PtrSlice) MapVal(mappers ...Int64PtrValMapper) *Int64PtrSlice {
	return this.Map(Int64PtrValMapperBatchToInt64PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Int64PtrSlice) First() *int64 {
	return this.At(0)
}

func (this *Int64PtrSlice) Last() *int64 {
	return this.At(this.Len() - 1)
}

func (this *Int64PtrSlice) Ats(is ...int) *Int64PtrSlice {
	res := MakeInt64PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Int64PtrSlice) Slice(left, right int) *Int64PtrSlice {
	return MakeInt64PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Int64PtrSlice) At(i int) *int64 {
	return this.GetRaw()[i]
}

func (this *Int64PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Int64PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *int32  @aliasPreix: Int32Ptr

type Int32PtrSlice struct {
	raw []*int32
}

func MakeInt32PtrSlice(param ...int) *Int32PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Int32PtrSlice{
		raw: make([]*int32, p1, p2),
	}
}

func MakeInt32PtrSliceFromRaw(raw []*int32) *Int32PtrSlice {
	if raw == nil {
		raw = make([]*int32, 0)
	}
	return &Int32PtrSlice{
		raw: raw,
	}
}

func (this *Int32PtrSlice) Append(vals ...*int32) {
	this.raw = append(this.raw, vals...)
}

func (this *Int32PtrSlice) AppendSlice(slice *Int32PtrSlice) {
	slice.RangeVal(func(val *int32) {
		this.Append(val)
	})
}

func (this *Int32PtrSlice) AppendIf(filter Int32PtrFilter, vals ...*int32) {
	this.AppendSlice(MakeInt32PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Int32PtrSlice) Range(rangers ...Int32PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Int32PtrSlice) RangeVal(rangers ...Int32PtrValRanger) {
	this.Range(Int32PtrValRangerBatchToInt32PtrRanger(rangers...)...)
}

func (this *Int32PtrSlice) GetRaw() []*int32 {
	return this.raw
}

func (this *Int32PtrSlice) Filter(filters ...Int32PtrFilter) *Int32PtrSlice {

	res := MakeInt32PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *int32) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Int32PtrSlice) FilterVal(filters ...Int32PtrValFilter) *Int32PtrSlice {
	return this.Filter(Int32PtrValFilterBatchToInt32PtrFilter(filters...)...)
}

func (this *Int32PtrSlice) Map(mappers ...Int32PtrMapper) *Int32PtrSlice {

	res := MakeInt32PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *int32) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Int32PtrSlice) MapVal(mappers ...Int32PtrValMapper) *Int32PtrSlice {
	return this.Map(Int32PtrValMapperBatchToInt32PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Int32PtrSlice) First() *int32 {
	return this.At(0)
}

func (this *Int32PtrSlice) Last() *int32 {
	return this.At(this.Len() - 1)
}

func (this *Int32PtrSlice) Ats(is ...int) *Int32PtrSlice {
	res := MakeInt32PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Int32PtrSlice) Slice(left, right int) *Int32PtrSlice {
	return MakeInt32PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Int32PtrSlice) At(i int) *int32 {
	return this.GetRaw()[i]
}

func (this *Int32PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Int32PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *int16  @aliasPreix: Int16Ptr

type Int16PtrSlice struct {
	raw []*int16
}

func MakeInt16PtrSlice(param ...int) *Int16PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Int16PtrSlice{
		raw: make([]*int16, p1, p2),
	}
}

func MakeInt16PtrSliceFromRaw(raw []*int16) *Int16PtrSlice {
	if raw == nil {
		raw = make([]*int16, 0)
	}
	return &Int16PtrSlice{
		raw: raw,
	}
}

func (this *Int16PtrSlice) Append(vals ...*int16) {
	this.raw = append(this.raw, vals...)
}

func (this *Int16PtrSlice) AppendSlice(slice *Int16PtrSlice) {
	slice.RangeVal(func(val *int16) {
		this.Append(val)
	})
}

func (this *Int16PtrSlice) AppendIf(filter Int16PtrFilter, vals ...*int16) {
	this.AppendSlice(MakeInt16PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Int16PtrSlice) Range(rangers ...Int16PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Int16PtrSlice) RangeVal(rangers ...Int16PtrValRanger) {
	this.Range(Int16PtrValRangerBatchToInt16PtrRanger(rangers...)...)
}

func (this *Int16PtrSlice) GetRaw() []*int16 {
	return this.raw
}

func (this *Int16PtrSlice) Filter(filters ...Int16PtrFilter) *Int16PtrSlice {

	res := MakeInt16PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *int16) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Int16PtrSlice) FilterVal(filters ...Int16PtrValFilter) *Int16PtrSlice {
	return this.Filter(Int16PtrValFilterBatchToInt16PtrFilter(filters...)...)
}

func (this *Int16PtrSlice) Map(mappers ...Int16PtrMapper) *Int16PtrSlice {

	res := MakeInt16PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *int16) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Int16PtrSlice) MapVal(mappers ...Int16PtrValMapper) *Int16PtrSlice {
	return this.Map(Int16PtrValMapperBatchToInt16PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Int16PtrSlice) First() *int16 {
	return this.At(0)
}

func (this *Int16PtrSlice) Last() *int16 {
	return this.At(this.Len() - 1)
}

func (this *Int16PtrSlice) Ats(is ...int) *Int16PtrSlice {
	res := MakeInt16PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Int16PtrSlice) Slice(left, right int) *Int16PtrSlice {
	return MakeInt16PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Int16PtrSlice) At(i int) *int16 {
	return this.GetRaw()[i]
}

func (this *Int16PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Int16PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *int8  @aliasPreix: Int8Ptr

type Int8PtrSlice struct {
	raw []*int8
}

func MakeInt8PtrSlice(param ...int) *Int8PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Int8PtrSlice{
		raw: make([]*int8, p1, p2),
	}
}

func MakeInt8PtrSliceFromRaw(raw []*int8) *Int8PtrSlice {
	if raw == nil {
		raw = make([]*int8, 0)
	}
	return &Int8PtrSlice{
		raw: raw,
	}
}

func (this *Int8PtrSlice) Append(vals ...*int8) {
	this.raw = append(this.raw, vals...)
}

func (this *Int8PtrSlice) AppendSlice(slice *Int8PtrSlice) {
	slice.RangeVal(func(val *int8) {
		this.Append(val)
	})
}

func (this *Int8PtrSlice) AppendIf(filter Int8PtrFilter, vals ...*int8) {
	this.AppendSlice(MakeInt8PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Int8PtrSlice) Range(rangers ...Int8PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Int8PtrSlice) RangeVal(rangers ...Int8PtrValRanger) {
	this.Range(Int8PtrValRangerBatchToInt8PtrRanger(rangers...)...)
}

func (this *Int8PtrSlice) GetRaw() []*int8 {
	return this.raw
}

func (this *Int8PtrSlice) Filter(filters ...Int8PtrFilter) *Int8PtrSlice {

	res := MakeInt8PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *int8) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Int8PtrSlice) FilterVal(filters ...Int8PtrValFilter) *Int8PtrSlice {
	return this.Filter(Int8PtrValFilterBatchToInt8PtrFilter(filters...)...)
}

func (this *Int8PtrSlice) Map(mappers ...Int8PtrMapper) *Int8PtrSlice {

	res := MakeInt8PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *int8) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Int8PtrSlice) MapVal(mappers ...Int8PtrValMapper) *Int8PtrSlice {
	return this.Map(Int8PtrValMapperBatchToInt8PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Int8PtrSlice) First() *int8 {
	return this.At(0)
}

func (this *Int8PtrSlice) Last() *int8 {
	return this.At(this.Len() - 1)
}

func (this *Int8PtrSlice) Ats(is ...int) *Int8PtrSlice {
	res := MakeInt8PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Int8PtrSlice) Slice(left, right int) *Int8PtrSlice {
	return MakeInt8PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Int8PtrSlice) At(i int) *int8 {
	return this.GetRaw()[i]
}

func (this *Int8PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Int8PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *uint  @aliasPreix: UintPtr

type UintPtrSlice struct {
	raw []*uint
}

func MakeUintPtrSlice(param ...int) *UintPtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &UintPtrSlice{
		raw: make([]*uint, p1, p2),
	}
}

func MakeUintPtrSliceFromRaw(raw []*uint) *UintPtrSlice {
	if raw == nil {
		raw = make([]*uint, 0)
	}
	return &UintPtrSlice{
		raw: raw,
	}
}

func (this *UintPtrSlice) Append(vals ...*uint) {
	this.raw = append(this.raw, vals...)
}

func (this *UintPtrSlice) AppendSlice(slice *UintPtrSlice) {
	slice.RangeVal(func(val *uint) {
		this.Append(val)
	})
}

func (this *UintPtrSlice) AppendIf(filter UintPtrFilter, vals ...*uint) {
	this.AppendSlice(MakeUintPtrSliceFromRaw(vals).Filter(filter))
}

func (this *UintPtrSlice) Range(rangers ...UintPtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *UintPtrSlice) RangeVal(rangers ...UintPtrValRanger) {
	this.Range(UintPtrValRangerBatchToUintPtrRanger(rangers...)...)
}

func (this *UintPtrSlice) GetRaw() []*uint {
	return this.raw
}

func (this *UintPtrSlice) Filter(filters ...UintPtrFilter) *UintPtrSlice {

	res := MakeUintPtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *uint) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *UintPtrSlice) FilterVal(filters ...UintPtrValFilter) *UintPtrSlice {
	return this.Filter(UintPtrValFilterBatchToUintPtrFilter(filters...)...)
}

func (this *UintPtrSlice) Map(mappers ...UintPtrMapper) *UintPtrSlice {

	res := MakeUintPtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *uint) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *UintPtrSlice) MapVal(mappers ...UintPtrValMapper) *UintPtrSlice {
	return this.Map(UintPtrValMapperBatchToUintPtrMapper(mappers...)...)
}

// let it crash apis
func (this *UintPtrSlice) First() *uint {
	return this.At(0)
}

func (this *UintPtrSlice) Last() *uint {
	return this.At(this.Len() - 1)
}

func (this *UintPtrSlice) Ats(is ...int) *UintPtrSlice {
	res := MakeUintPtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *UintPtrSlice) Slice(left, right int) *UintPtrSlice {
	return MakeUintPtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *UintPtrSlice) At(i int) *uint {
	return this.GetRaw()[i]
}

func (this *UintPtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *UintPtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *uint8  @aliasPreix: Uint8Ptr

type Uint8PtrSlice struct {
	raw []*uint8
}

func MakeUint8PtrSlice(param ...int) *Uint8PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Uint8PtrSlice{
		raw: make([]*uint8, p1, p2),
	}
}

func MakeUint8PtrSliceFromRaw(raw []*uint8) *Uint8PtrSlice {
	if raw == nil {
		raw = make([]*uint8, 0)
	}
	return &Uint8PtrSlice{
		raw: raw,
	}
}

func (this *Uint8PtrSlice) Append(vals ...*uint8) {
	this.raw = append(this.raw, vals...)
}

func (this *Uint8PtrSlice) AppendSlice(slice *Uint8PtrSlice) {
	slice.RangeVal(func(val *uint8) {
		this.Append(val)
	})
}

func (this *Uint8PtrSlice) AppendIf(filter Uint8PtrFilter, vals ...*uint8) {
	this.AppendSlice(MakeUint8PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Uint8PtrSlice) Range(rangers ...Uint8PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Uint8PtrSlice) RangeVal(rangers ...Uint8PtrValRanger) {
	this.Range(Uint8PtrValRangerBatchToUint8PtrRanger(rangers...)...)
}

func (this *Uint8PtrSlice) GetRaw() []*uint8 {
	return this.raw
}

func (this *Uint8PtrSlice) Filter(filters ...Uint8PtrFilter) *Uint8PtrSlice {

	res := MakeUint8PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *uint8) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Uint8PtrSlice) FilterVal(filters ...Uint8PtrValFilter) *Uint8PtrSlice {
	return this.Filter(Uint8PtrValFilterBatchToUint8PtrFilter(filters...)...)
}

func (this *Uint8PtrSlice) Map(mappers ...Uint8PtrMapper) *Uint8PtrSlice {

	res := MakeUint8PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *uint8) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Uint8PtrSlice) MapVal(mappers ...Uint8PtrValMapper) *Uint8PtrSlice {
	return this.Map(Uint8PtrValMapperBatchToUint8PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Uint8PtrSlice) First() *uint8 {
	return this.At(0)
}

func (this *Uint8PtrSlice) Last() *uint8 {
	return this.At(this.Len() - 1)
}

func (this *Uint8PtrSlice) Ats(is ...int) *Uint8PtrSlice {
	res := MakeUint8PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Uint8PtrSlice) Slice(left, right int) *Uint8PtrSlice {
	return MakeUint8PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Uint8PtrSlice) At(i int) *uint8 {
	return this.GetRaw()[i]
}

func (this *Uint8PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Uint8PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *uint16  @aliasPreix: Uint16Ptr

type Uint16PtrSlice struct {
	raw []*uint16
}

func MakeUint16PtrSlice(param ...int) *Uint16PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Uint16PtrSlice{
		raw: make([]*uint16, p1, p2),
	}
}

func MakeUint16PtrSliceFromRaw(raw []*uint16) *Uint16PtrSlice {
	if raw == nil {
		raw = make([]*uint16, 0)
	}
	return &Uint16PtrSlice{
		raw: raw,
	}
}

func (this *Uint16PtrSlice) Append(vals ...*uint16) {
	this.raw = append(this.raw, vals...)
}

func (this *Uint16PtrSlice) AppendSlice(slice *Uint16PtrSlice) {
	slice.RangeVal(func(val *uint16) {
		this.Append(val)
	})
}

func (this *Uint16PtrSlice) AppendIf(filter Uint16PtrFilter, vals ...*uint16) {
	this.AppendSlice(MakeUint16PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Uint16PtrSlice) Range(rangers ...Uint16PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Uint16PtrSlice) RangeVal(rangers ...Uint16PtrValRanger) {
	this.Range(Uint16PtrValRangerBatchToUint16PtrRanger(rangers...)...)
}

func (this *Uint16PtrSlice) GetRaw() []*uint16 {
	return this.raw
}

func (this *Uint16PtrSlice) Filter(filters ...Uint16PtrFilter) *Uint16PtrSlice {

	res := MakeUint16PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *uint16) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Uint16PtrSlice) FilterVal(filters ...Uint16PtrValFilter) *Uint16PtrSlice {
	return this.Filter(Uint16PtrValFilterBatchToUint16PtrFilter(filters...)...)
}

func (this *Uint16PtrSlice) Map(mappers ...Uint16PtrMapper) *Uint16PtrSlice {

	res := MakeUint16PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *uint16) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Uint16PtrSlice) MapVal(mappers ...Uint16PtrValMapper) *Uint16PtrSlice {
	return this.Map(Uint16PtrValMapperBatchToUint16PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Uint16PtrSlice) First() *uint16 {
	return this.At(0)
}

func (this *Uint16PtrSlice) Last() *uint16 {
	return this.At(this.Len() - 1)
}

func (this *Uint16PtrSlice) Ats(is ...int) *Uint16PtrSlice {
	res := MakeUint16PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Uint16PtrSlice) Slice(left, right int) *Uint16PtrSlice {
	return MakeUint16PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Uint16PtrSlice) At(i int) *uint16 {
	return this.GetRaw()[i]
}

func (this *Uint16PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Uint16PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *uint32  @aliasPreix: Uint32Ptr

type Uint32PtrSlice struct {
	raw []*uint32
}

func MakeUint32PtrSlice(param ...int) *Uint32PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Uint32PtrSlice{
		raw: make([]*uint32, p1, p2),
	}
}

func MakeUint32PtrSliceFromRaw(raw []*uint32) *Uint32PtrSlice {
	if raw == nil {
		raw = make([]*uint32, 0)
	}
	return &Uint32PtrSlice{
		raw: raw,
	}
}

func (this *Uint32PtrSlice) Append(vals ...*uint32) {
	this.raw = append(this.raw, vals...)
}

func (this *Uint32PtrSlice) AppendSlice(slice *Uint32PtrSlice) {
	slice.RangeVal(func(val *uint32) {
		this.Append(val)
	})
}

func (this *Uint32PtrSlice) AppendIf(filter Uint32PtrFilter, vals ...*uint32) {
	this.AppendSlice(MakeUint32PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Uint32PtrSlice) Range(rangers ...Uint32PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Uint32PtrSlice) RangeVal(rangers ...Uint32PtrValRanger) {
	this.Range(Uint32PtrValRangerBatchToUint32PtrRanger(rangers...)...)
}

func (this *Uint32PtrSlice) GetRaw() []*uint32 {
	return this.raw
}

func (this *Uint32PtrSlice) Filter(filters ...Uint32PtrFilter) *Uint32PtrSlice {

	res := MakeUint32PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *uint32) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Uint32PtrSlice) FilterVal(filters ...Uint32PtrValFilter) *Uint32PtrSlice {
	return this.Filter(Uint32PtrValFilterBatchToUint32PtrFilter(filters...)...)
}

func (this *Uint32PtrSlice) Map(mappers ...Uint32PtrMapper) *Uint32PtrSlice {

	res := MakeUint32PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *uint32) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Uint32PtrSlice) MapVal(mappers ...Uint32PtrValMapper) *Uint32PtrSlice {
	return this.Map(Uint32PtrValMapperBatchToUint32PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Uint32PtrSlice) First() *uint32 {
	return this.At(0)
}

func (this *Uint32PtrSlice) Last() *uint32 {
	return this.At(this.Len() - 1)
}

func (this *Uint32PtrSlice) Ats(is ...int) *Uint32PtrSlice {
	res := MakeUint32PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Uint32PtrSlice) Slice(left, right int) *Uint32PtrSlice {
	return MakeUint32PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Uint32PtrSlice) At(i int) *uint32 {
	return this.GetRaw()[i]
}

func (this *Uint32PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Uint32PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *uint64  @aliasPreix: Uint64Ptr

type Uint64PtrSlice struct {
	raw []*uint64
}

func MakeUint64PtrSlice(param ...int) *Uint64PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Uint64PtrSlice{
		raw: make([]*uint64, p1, p2),
	}
}

func MakeUint64PtrSliceFromRaw(raw []*uint64) *Uint64PtrSlice {
	if raw == nil {
		raw = make([]*uint64, 0)
	}
	return &Uint64PtrSlice{
		raw: raw,
	}
}

func (this *Uint64PtrSlice) Append(vals ...*uint64) {
	this.raw = append(this.raw, vals...)
}

func (this *Uint64PtrSlice) AppendSlice(slice *Uint64PtrSlice) {
	slice.RangeVal(func(val *uint64) {
		this.Append(val)
	})
}

func (this *Uint64PtrSlice) AppendIf(filter Uint64PtrFilter, vals ...*uint64) {
	this.AppendSlice(MakeUint64PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Uint64PtrSlice) Range(rangers ...Uint64PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Uint64PtrSlice) RangeVal(rangers ...Uint64PtrValRanger) {
	this.Range(Uint64PtrValRangerBatchToUint64PtrRanger(rangers...)...)
}

func (this *Uint64PtrSlice) GetRaw() []*uint64 {
	return this.raw
}

func (this *Uint64PtrSlice) Filter(filters ...Uint64PtrFilter) *Uint64PtrSlice {

	res := MakeUint64PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *uint64) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Uint64PtrSlice) FilterVal(filters ...Uint64PtrValFilter) *Uint64PtrSlice {
	return this.Filter(Uint64PtrValFilterBatchToUint64PtrFilter(filters...)...)
}

func (this *Uint64PtrSlice) Map(mappers ...Uint64PtrMapper) *Uint64PtrSlice {

	res := MakeUint64PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *uint64) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Uint64PtrSlice) MapVal(mappers ...Uint64PtrValMapper) *Uint64PtrSlice {
	return this.Map(Uint64PtrValMapperBatchToUint64PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Uint64PtrSlice) First() *uint64 {
	return this.At(0)
}

func (this *Uint64PtrSlice) Last() *uint64 {
	return this.At(this.Len() - 1)
}

func (this *Uint64PtrSlice) Ats(is ...int) *Uint64PtrSlice {
	res := MakeUint64PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Uint64PtrSlice) Slice(left, right int) *Uint64PtrSlice {
	return MakeUint64PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Uint64PtrSlice) At(i int) *uint64 {
	return this.GetRaw()[i]
}

func (this *Uint64PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Uint64PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *float32  @aliasPreix: Float32Ptr

type Float32PtrSlice struct {
	raw []*float32
}

func MakeFloat32PtrSlice(param ...int) *Float32PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Float32PtrSlice{
		raw: make([]*float32, p1, p2),
	}
}

func MakeFloat32PtrSliceFromRaw(raw []*float32) *Float32PtrSlice {
	if raw == nil {
		raw = make([]*float32, 0)
	}
	return &Float32PtrSlice{
		raw: raw,
	}
}

func (this *Float32PtrSlice) Append(vals ...*float32) {
	this.raw = append(this.raw, vals...)
}

func (this *Float32PtrSlice) AppendSlice(slice *Float32PtrSlice) {
	slice.RangeVal(func(val *float32) {
		this.Append(val)
	})
}

func (this *Float32PtrSlice) AppendIf(filter Float32PtrFilter, vals ...*float32) {
	this.AppendSlice(MakeFloat32PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Float32PtrSlice) Range(rangers ...Float32PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Float32PtrSlice) RangeVal(rangers ...Float32PtrValRanger) {
	this.Range(Float32PtrValRangerBatchToFloat32PtrRanger(rangers...)...)
}

func (this *Float32PtrSlice) GetRaw() []*float32 {
	return this.raw
}

func (this *Float32PtrSlice) Filter(filters ...Float32PtrFilter) *Float32PtrSlice {

	res := MakeFloat32PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *float32) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Float32PtrSlice) FilterVal(filters ...Float32PtrValFilter) *Float32PtrSlice {
	return this.Filter(Float32PtrValFilterBatchToFloat32PtrFilter(filters...)...)
}

func (this *Float32PtrSlice) Map(mappers ...Float32PtrMapper) *Float32PtrSlice {

	res := MakeFloat32PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *float32) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Float32PtrSlice) MapVal(mappers ...Float32PtrValMapper) *Float32PtrSlice {
	return this.Map(Float32PtrValMapperBatchToFloat32PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Float32PtrSlice) First() *float32 {
	return this.At(0)
}

func (this *Float32PtrSlice) Last() *float32 {
	return this.At(this.Len() - 1)
}

func (this *Float32PtrSlice) Ats(is ...int) *Float32PtrSlice {
	res := MakeFloat32PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Float32PtrSlice) Slice(left, right int) *Float32PtrSlice {
	return MakeFloat32PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Float32PtrSlice) At(i int) *float32 {
	return this.GetRaw()[i]
}

func (this *Float32PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Float32PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *float64  @aliasPreix: Float64Ptr

type Float64PtrSlice struct {
	raw []*float64
}

func MakeFloat64PtrSlice(param ...int) *Float64PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Float64PtrSlice{
		raw: make([]*float64, p1, p2),
	}
}

func MakeFloat64PtrSliceFromRaw(raw []*float64) *Float64PtrSlice {
	if raw == nil {
		raw = make([]*float64, 0)
	}
	return &Float64PtrSlice{
		raw: raw,
	}
}

func (this *Float64PtrSlice) Append(vals ...*float64) {
	this.raw = append(this.raw, vals...)
}

func (this *Float64PtrSlice) AppendSlice(slice *Float64PtrSlice) {
	slice.RangeVal(func(val *float64) {
		this.Append(val)
	})
}

func (this *Float64PtrSlice) AppendIf(filter Float64PtrFilter, vals ...*float64) {
	this.AppendSlice(MakeFloat64PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Float64PtrSlice) Range(rangers ...Float64PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Float64PtrSlice) RangeVal(rangers ...Float64PtrValRanger) {
	this.Range(Float64PtrValRangerBatchToFloat64PtrRanger(rangers...)...)
}

func (this *Float64PtrSlice) GetRaw() []*float64 {
	return this.raw
}

func (this *Float64PtrSlice) Filter(filters ...Float64PtrFilter) *Float64PtrSlice {

	res := MakeFloat64PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *float64) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Float64PtrSlice) FilterVal(filters ...Float64PtrValFilter) *Float64PtrSlice {
	return this.Filter(Float64PtrValFilterBatchToFloat64PtrFilter(filters...)...)
}

func (this *Float64PtrSlice) Map(mappers ...Float64PtrMapper) *Float64PtrSlice {

	res := MakeFloat64PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *float64) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Float64PtrSlice) MapVal(mappers ...Float64PtrValMapper) *Float64PtrSlice {
	return this.Map(Float64PtrValMapperBatchToFloat64PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Float64PtrSlice) First() *float64 {
	return this.At(0)
}

func (this *Float64PtrSlice) Last() *float64 {
	return this.At(this.Len() - 1)
}

func (this *Float64PtrSlice) Ats(is ...int) *Float64PtrSlice {
	res := MakeFloat64PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Float64PtrSlice) Slice(left, right int) *Float64PtrSlice {
	return MakeFloat64PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Float64PtrSlice) At(i int) *float64 {
	return this.GetRaw()[i]
}

func (this *Float64PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Float64PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *bool  @aliasPreix: BoolPtr

type BoolPtrSlice struct {
	raw []*bool
}

func MakeBoolPtrSlice(param ...int) *BoolPtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &BoolPtrSlice{
		raw: make([]*bool, p1, p2),
	}
}

func MakeBoolPtrSliceFromRaw(raw []*bool) *BoolPtrSlice {
	if raw == nil {
		raw = make([]*bool, 0)
	}
	return &BoolPtrSlice{
		raw: raw,
	}
}

func (this *BoolPtrSlice) Append(vals ...*bool) {
	this.raw = append(this.raw, vals...)
}

func (this *BoolPtrSlice) AppendSlice(slice *BoolPtrSlice) {
	slice.RangeVal(func(val *bool) {
		this.Append(val)
	})
}

func (this *BoolPtrSlice) AppendIf(filter BoolPtrFilter, vals ...*bool) {
	this.AppendSlice(MakeBoolPtrSliceFromRaw(vals).Filter(filter))
}

func (this *BoolPtrSlice) Range(rangers ...BoolPtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *BoolPtrSlice) RangeVal(rangers ...BoolPtrValRanger) {
	this.Range(BoolPtrValRangerBatchToBoolPtrRanger(rangers...)...)
}

func (this *BoolPtrSlice) GetRaw() []*bool {
	return this.raw
}

func (this *BoolPtrSlice) Filter(filters ...BoolPtrFilter) *BoolPtrSlice {

	res := MakeBoolPtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *bool) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *BoolPtrSlice) FilterVal(filters ...BoolPtrValFilter) *BoolPtrSlice {
	return this.Filter(BoolPtrValFilterBatchToBoolPtrFilter(filters...)...)
}

func (this *BoolPtrSlice) Map(mappers ...BoolPtrMapper) *BoolPtrSlice {

	res := MakeBoolPtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *bool) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *BoolPtrSlice) MapVal(mappers ...BoolPtrValMapper) *BoolPtrSlice {
	return this.Map(BoolPtrValMapperBatchToBoolPtrMapper(mappers...)...)
}

// let it crash apis
func (this *BoolPtrSlice) First() *bool {
	return this.At(0)
}

func (this *BoolPtrSlice) Last() *bool {
	return this.At(this.Len() - 1)
}

func (this *BoolPtrSlice) Ats(is ...int) *BoolPtrSlice {
	res := MakeBoolPtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *BoolPtrSlice) Slice(left, right int) *BoolPtrSlice {
	return MakeBoolPtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *BoolPtrSlice) At(i int) *bool {
	return this.GetRaw()[i]
}

func (this *BoolPtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *BoolPtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *rune  @aliasPreix: RunePtr

type RunePtrSlice struct {
	raw []*rune
}

func MakeRunePtrSlice(param ...int) *RunePtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &RunePtrSlice{
		raw: make([]*rune, p1, p2),
	}
}

func MakeRunePtrSliceFromRaw(raw []*rune) *RunePtrSlice {
	if raw == nil {
		raw = make([]*rune, 0)
	}
	return &RunePtrSlice{
		raw: raw,
	}
}

func (this *RunePtrSlice) Append(vals ...*rune) {
	this.raw = append(this.raw, vals...)
}

func (this *RunePtrSlice) AppendSlice(slice *RunePtrSlice) {
	slice.RangeVal(func(val *rune) {
		this.Append(val)
	})
}

func (this *RunePtrSlice) AppendIf(filter RunePtrFilter, vals ...*rune) {
	this.AppendSlice(MakeRunePtrSliceFromRaw(vals).Filter(filter))
}

func (this *RunePtrSlice) Range(rangers ...RunePtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *RunePtrSlice) RangeVal(rangers ...RunePtrValRanger) {
	this.Range(RunePtrValRangerBatchToRunePtrRanger(rangers...)...)
}

func (this *RunePtrSlice) GetRaw() []*rune {
	return this.raw
}

func (this *RunePtrSlice) Filter(filters ...RunePtrFilter) *RunePtrSlice {

	res := MakeRunePtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *rune) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *RunePtrSlice) FilterVal(filters ...RunePtrValFilter) *RunePtrSlice {
	return this.Filter(RunePtrValFilterBatchToRunePtrFilter(filters...)...)
}

func (this *RunePtrSlice) Map(mappers ...RunePtrMapper) *RunePtrSlice {

	res := MakeRunePtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *rune) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *RunePtrSlice) MapVal(mappers ...RunePtrValMapper) *RunePtrSlice {
	return this.Map(RunePtrValMapperBatchToRunePtrMapper(mappers...)...)
}

// let it crash apis
func (this *RunePtrSlice) First() *rune {
	return this.At(0)
}

func (this *RunePtrSlice) Last() *rune {
	return this.At(this.Len() - 1)
}

func (this *RunePtrSlice) Ats(is ...int) *RunePtrSlice {
	res := MakeRunePtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *RunePtrSlice) Slice(left, right int) *RunePtrSlice {
	return MakeRunePtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *RunePtrSlice) At(i int) *rune {
	return this.GetRaw()[i]
}

func (this *RunePtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *RunePtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *complex64  @aliasPreix: Complex64Ptr

type Complex64PtrSlice struct {
	raw []*complex64
}

func MakeComplex64PtrSlice(param ...int) *Complex64PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Complex64PtrSlice{
		raw: make([]*complex64, p1, p2),
	}
}

func MakeComplex64PtrSliceFromRaw(raw []*complex64) *Complex64PtrSlice {
	if raw == nil {
		raw = make([]*complex64, 0)
	}
	return &Complex64PtrSlice{
		raw: raw,
	}
}

func (this *Complex64PtrSlice) Append(vals ...*complex64) {
	this.raw = append(this.raw, vals...)
}

func (this *Complex64PtrSlice) AppendSlice(slice *Complex64PtrSlice) {
	slice.RangeVal(func(val *complex64) {
		this.Append(val)
	})
}

func (this *Complex64PtrSlice) AppendIf(filter Complex64PtrFilter, vals ...*complex64) {
	this.AppendSlice(MakeComplex64PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Complex64PtrSlice) Range(rangers ...Complex64PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Complex64PtrSlice) RangeVal(rangers ...Complex64PtrValRanger) {
	this.Range(Complex64PtrValRangerBatchToComplex64PtrRanger(rangers...)...)
}

func (this *Complex64PtrSlice) GetRaw() []*complex64 {
	return this.raw
}

func (this *Complex64PtrSlice) Filter(filters ...Complex64PtrFilter) *Complex64PtrSlice {

	res := MakeComplex64PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *complex64) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Complex64PtrSlice) FilterVal(filters ...Complex64PtrValFilter) *Complex64PtrSlice {
	return this.Filter(Complex64PtrValFilterBatchToComplex64PtrFilter(filters...)...)
}

func (this *Complex64PtrSlice) Map(mappers ...Complex64PtrMapper) *Complex64PtrSlice {

	res := MakeComplex64PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *complex64) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Complex64PtrSlice) MapVal(mappers ...Complex64PtrValMapper) *Complex64PtrSlice {
	return this.Map(Complex64PtrValMapperBatchToComplex64PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Complex64PtrSlice) First() *complex64 {
	return this.At(0)
}

func (this *Complex64PtrSlice) Last() *complex64 {
	return this.At(this.Len() - 1)
}

func (this *Complex64PtrSlice) Ats(is ...int) *Complex64PtrSlice {
	res := MakeComplex64PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Complex64PtrSlice) Slice(left, right int) *Complex64PtrSlice {
	return MakeComplex64PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Complex64PtrSlice) At(i int) *complex64 {
	return this.GetRaw()[i]
}

func (this *Complex64PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Complex64PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// autogen @type: *complex32  @aliasPreix: Complex32Ptr

type Complex32PtrSlice struct {
	raw []*complex32
}

func MakeComplex32PtrSlice(param ...int) *Complex32PtrSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &Complex32PtrSlice{
		raw: make([]*complex32, p1, p2),
	}
}

func MakeComplex32PtrSliceFromRaw(raw []*complex32) *Complex32PtrSlice {
	if raw == nil {
		raw = make([]*complex32, 0)
	}
	return &Complex32PtrSlice{
		raw: raw,
	}
}

func (this *Complex32PtrSlice) Append(vals ...*complex32) {
	this.raw = append(this.raw, vals...)
}

func (this *Complex32PtrSlice) AppendSlice(slice *Complex32PtrSlice) {
	slice.RangeVal(func(val *complex32) {
		this.Append(val)
	})
}

func (this *Complex32PtrSlice) AppendIf(filter Complex32PtrFilter, vals ...*complex32) {
	this.AppendSlice(MakeComplex32PtrSliceFromRaw(vals).Filter(filter))
}

func (this *Complex32PtrSlice) Range(rangers ...Complex32PtrRanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *Complex32PtrSlice) RangeVal(rangers ...Complex32PtrValRanger) {
	this.Range(Complex32PtrValRangerBatchToComplex32PtrRanger(rangers...)...)
}

func (this *Complex32PtrSlice) GetRaw() []*complex32 {
	return this.raw
}

func (this *Complex32PtrSlice) Filter(filters ...Complex32PtrFilter) *Complex32PtrSlice {

	res := MakeComplex32PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *complex32) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *Complex32PtrSlice) FilterVal(filters ...Complex32PtrValFilter) *Complex32PtrSlice {
	return this.Filter(Complex32PtrValFilterBatchToComplex32PtrFilter(filters...)...)
}

func (this *Complex32PtrSlice) Map(mappers ...Complex32PtrMapper) *Complex32PtrSlice {

	res := MakeComplex32PtrSlice(0, len(this.GetRaw()))

	this.Range(func(i int, val *complex32) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *Complex32PtrSlice) MapVal(mappers ...Complex32PtrValMapper) *Complex32PtrSlice {
	return this.Map(Complex32PtrValMapperBatchToComplex32PtrMapper(mappers...)...)
}

// let it crash apis
func (this *Complex32PtrSlice) First() *complex32 {
	return this.At(0)
}

func (this *Complex32PtrSlice) Last() *complex32 {
	return this.At(this.Len() - 1)
}

func (this *Complex32PtrSlice) Ats(is ...int) *Complex32PtrSlice {
	res := MakeComplex32PtrSlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *Complex32PtrSlice) Slice(left, right int) *Complex32PtrSlice {
	return MakeComplex32PtrSliceFromRaw(this.GetRaw()[left:right])
}

func (this *Complex32PtrSlice) At(i int) *complex32 {
	return this.GetRaw()[i]
}

func (this *Complex32PtrSlice) Len() int {
	return len(this.GetRaw())
}

func (this *Complex32PtrSlice) Cap() int {
	return cap(this.GetRaw())
}

// generate sliceFunc

// common area
type BoolFunc func() bool

// autogen @type: string  @aliasPreix: String

// ===  gen area =====
type StringFilter func(int, string) bool
type StringMapper func(int, string) string
type StringRanger func(int, string)

func (filter StringFilter) Something() StringFilter {
	if filter == nil {
		return func(int, string) bool {
			return false
		}
	}
	return filter
}

func (mapper StringMapper) Something() StringMapper {
	if mapper == nil {
		return func(i int, val string) string {
			return val
		}
	}
	return mapper
}

func (ranger StringRanger) Something() StringRanger {
	if ranger == nil {
		return func(int, string) {}
	}
	return ranger
}

type StringValFilter func(string) bool
type StringValMapper func(string) string
type StringValRanger func(string)

func (filter StringValFilter) Something() StringValFilter {
	if filter == nil {
		return func(string) bool {
			return false
		}
	}
	return filter
}

func (mapper StringValMapper) Something() StringValMapper {
	if mapper == nil {
		return func(val string) string {
			return val
		}
	}
	return mapper
}

func (ranger StringValRanger) Something() StringValRanger {
	if ranger == nil {
		return func(string) {}
	}
	return ranger
}

// ==== util func

func StringValRangerToStringRanger(ranger StringValRanger) StringRanger {
	return func(i int, val string) {
		ranger.Something()(val)
	}
}

func StringValRangerBatchToStringRanger(rangers ...StringValRanger) []StringRanger {
	res := make([]StringRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, StringValRangerToStringRanger(ranger))
	}
	return res
}

func StringValMapperToStringMapper(mapper StringValMapper) StringMapper {
	return func(i int, val string) string {
		return mapper.Something()(val)
	}
}

func StringValMapperBatchToStringMapper(mappers ...StringValMapper) []StringMapper {
	res := make([]StringMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, StringValMapperToStringMapper(mapper))
	}
	return res
}

func StringValFilterToStringFilter(filter StringValFilter) StringFilter {
	return func(i int, val string) bool {
		return filter.Something()(val)
	}
}

func StringValFilterBatchToStringFilter(filters ...StringValFilter) []StringFilter {
	res := make([]StringFilter, len(filters))
	for _, filter := range filters {
		res = append(res, StringValFilterToStringFilter(filter))
	}
	return res
}

func StringIf(f BoolFunc, a, b string) string {
	if f() {
		return a
	}
	return b
}

// autogen @type: int  @aliasPreix: Int

// ===  gen area =====
type IntFilter func(int, int) bool
type IntMapper func(int, int) int
type IntRanger func(int, int)

func (filter IntFilter) Something() IntFilter {
	if filter == nil {
		return func(int, int) bool {
			return false
		}
	}
	return filter
}

func (mapper IntMapper) Something() IntMapper {
	if mapper == nil {
		return func(i int, val int) int {
			return val
		}
	}
	return mapper
}

func (ranger IntRanger) Something() IntRanger {
	if ranger == nil {
		return func(int, int) {}
	}
	return ranger
}

type IntValFilter func(int) bool
type IntValMapper func(int) int
type IntValRanger func(int)

func (filter IntValFilter) Something() IntValFilter {
	if filter == nil {
		return func(int) bool {
			return false
		}
	}
	return filter
}

func (mapper IntValMapper) Something() IntValMapper {
	if mapper == nil {
		return func(val int) int {
			return val
		}
	}
	return mapper
}

func (ranger IntValRanger) Something() IntValRanger {
	if ranger == nil {
		return func(int) {}
	}
	return ranger
}

// ==== util func

func IntValRangerToIntRanger(ranger IntValRanger) IntRanger {
	return func(i int, val int) {
		ranger.Something()(val)
	}
}

func IntValRangerBatchToIntRanger(rangers ...IntValRanger) []IntRanger {
	res := make([]IntRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, IntValRangerToIntRanger(ranger))
	}
	return res
}

func IntValMapperToIntMapper(mapper IntValMapper) IntMapper {
	return func(i int, val int) int {
		return mapper.Something()(val)
	}
}

func IntValMapperBatchToIntMapper(mappers ...IntValMapper) []IntMapper {
	res := make([]IntMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, IntValMapperToIntMapper(mapper))
	}
	return res
}

func IntValFilterToIntFilter(filter IntValFilter) IntFilter {
	return func(i int, val int) bool {
		return filter.Something()(val)
	}
}

func IntValFilterBatchToIntFilter(filters ...IntValFilter) []IntFilter {
	res := make([]IntFilter, len(filters))
	for _, filter := range filters {
		res = append(res, IntValFilterToIntFilter(filter))
	}
	return res
}

func IntIf(f BoolFunc, a, b int) int {
	if f() {
		return a
	}
	return b
}

// autogen @type: int64  @aliasPreix: Int64

// ===  gen area =====
type Int64Filter func(int, int64) bool
type Int64Mapper func(int, int64) int64
type Int64Ranger func(int, int64)

func (filter Int64Filter) Something() Int64Filter {
	if filter == nil {
		return func(int, int64) bool {
			return false
		}
	}
	return filter
}

func (mapper Int64Mapper) Something() Int64Mapper {
	if mapper == nil {
		return func(i int, val int64) int64 {
			return val
		}
	}
	return mapper
}

func (ranger Int64Ranger) Something() Int64Ranger {
	if ranger == nil {
		return func(int, int64) {}
	}
	return ranger
}

type Int64ValFilter func(int64) bool
type Int64ValMapper func(int64) int64
type Int64ValRanger func(int64)

func (filter Int64ValFilter) Something() Int64ValFilter {
	if filter == nil {
		return func(int64) bool {
			return false
		}
	}
	return filter
}

func (mapper Int64ValMapper) Something() Int64ValMapper {
	if mapper == nil {
		return func(val int64) int64 {
			return val
		}
	}
	return mapper
}

func (ranger Int64ValRanger) Something() Int64ValRanger {
	if ranger == nil {
		return func(int64) {}
	}
	return ranger
}

// ==== util func

func Int64ValRangerToInt64Ranger(ranger Int64ValRanger) Int64Ranger {
	return func(i int, val int64) {
		ranger.Something()(val)
	}
}

func Int64ValRangerBatchToInt64Ranger(rangers ...Int64ValRanger) []Int64Ranger {
	res := make([]Int64Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Int64ValRangerToInt64Ranger(ranger))
	}
	return res
}

func Int64ValMapperToInt64Mapper(mapper Int64ValMapper) Int64Mapper {
	return func(i int, val int64) int64 {
		return mapper.Something()(val)
	}
}

func Int64ValMapperBatchToInt64Mapper(mappers ...Int64ValMapper) []Int64Mapper {
	res := make([]Int64Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Int64ValMapperToInt64Mapper(mapper))
	}
	return res
}

func Int64ValFilterToInt64Filter(filter Int64ValFilter) Int64Filter {
	return func(i int, val int64) bool {
		return filter.Something()(val)
	}
}

func Int64ValFilterBatchToInt64Filter(filters ...Int64ValFilter) []Int64Filter {
	res := make([]Int64Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Int64ValFilterToInt64Filter(filter))
	}
	return res
}

func Int64If(f BoolFunc, a, b int64) int64 {
	if f() {
		return a
	}
	return b
}

// autogen @type: int32  @aliasPreix: Int32

// ===  gen area =====
type Int32Filter func(int, int32) bool
type Int32Mapper func(int, int32) int32
type Int32Ranger func(int, int32)

func (filter Int32Filter) Something() Int32Filter {
	if filter == nil {
		return func(int, int32) bool {
			return false
		}
	}
	return filter
}

func (mapper Int32Mapper) Something() Int32Mapper {
	if mapper == nil {
		return func(i int, val int32) int32 {
			return val
		}
	}
	return mapper
}

func (ranger Int32Ranger) Something() Int32Ranger {
	if ranger == nil {
		return func(int, int32) {}
	}
	return ranger
}

type Int32ValFilter func(int32) bool
type Int32ValMapper func(int32) int32
type Int32ValRanger func(int32)

func (filter Int32ValFilter) Something() Int32ValFilter {
	if filter == nil {
		return func(int32) bool {
			return false
		}
	}
	return filter
}

func (mapper Int32ValMapper) Something() Int32ValMapper {
	if mapper == nil {
		return func(val int32) int32 {
			return val
		}
	}
	return mapper
}

func (ranger Int32ValRanger) Something() Int32ValRanger {
	if ranger == nil {
		return func(int32) {}
	}
	return ranger
}

// ==== util func

func Int32ValRangerToInt32Ranger(ranger Int32ValRanger) Int32Ranger {
	return func(i int, val int32) {
		ranger.Something()(val)
	}
}

func Int32ValRangerBatchToInt32Ranger(rangers ...Int32ValRanger) []Int32Ranger {
	res := make([]Int32Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Int32ValRangerToInt32Ranger(ranger))
	}
	return res
}

func Int32ValMapperToInt32Mapper(mapper Int32ValMapper) Int32Mapper {
	return func(i int, val int32) int32 {
		return mapper.Something()(val)
	}
}

func Int32ValMapperBatchToInt32Mapper(mappers ...Int32ValMapper) []Int32Mapper {
	res := make([]Int32Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Int32ValMapperToInt32Mapper(mapper))
	}
	return res
}

func Int32ValFilterToInt32Filter(filter Int32ValFilter) Int32Filter {
	return func(i int, val int32) bool {
		return filter.Something()(val)
	}
}

func Int32ValFilterBatchToInt32Filter(filters ...Int32ValFilter) []Int32Filter {
	res := make([]Int32Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Int32ValFilterToInt32Filter(filter))
	}
	return res
}

func Int32If(f BoolFunc, a, b int32) int32 {
	if f() {
		return a
	}
	return b
}

// autogen @type: int16  @aliasPreix: Int16

// ===  gen area =====
type Int16Filter func(int, int16) bool
type Int16Mapper func(int, int16) int16
type Int16Ranger func(int, int16)

func (filter Int16Filter) Something() Int16Filter {
	if filter == nil {
		return func(int, int16) bool {
			return false
		}
	}
	return filter
}

func (mapper Int16Mapper) Something() Int16Mapper {
	if mapper == nil {
		return func(i int, val int16) int16 {
			return val
		}
	}
	return mapper
}

func (ranger Int16Ranger) Something() Int16Ranger {
	if ranger == nil {
		return func(int, int16) {}
	}
	return ranger
}

type Int16ValFilter func(int16) bool
type Int16ValMapper func(int16) int16
type Int16ValRanger func(int16)

func (filter Int16ValFilter) Something() Int16ValFilter {
	if filter == nil {
		return func(int16) bool {
			return false
		}
	}
	return filter
}

func (mapper Int16ValMapper) Something() Int16ValMapper {
	if mapper == nil {
		return func(val int16) int16 {
			return val
		}
	}
	return mapper
}

func (ranger Int16ValRanger) Something() Int16ValRanger {
	if ranger == nil {
		return func(int16) {}
	}
	return ranger
}

// ==== util func

func Int16ValRangerToInt16Ranger(ranger Int16ValRanger) Int16Ranger {
	return func(i int, val int16) {
		ranger.Something()(val)
	}
}

func Int16ValRangerBatchToInt16Ranger(rangers ...Int16ValRanger) []Int16Ranger {
	res := make([]Int16Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Int16ValRangerToInt16Ranger(ranger))
	}
	return res
}

func Int16ValMapperToInt16Mapper(mapper Int16ValMapper) Int16Mapper {
	return func(i int, val int16) int16 {
		return mapper.Something()(val)
	}
}

func Int16ValMapperBatchToInt16Mapper(mappers ...Int16ValMapper) []Int16Mapper {
	res := make([]Int16Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Int16ValMapperToInt16Mapper(mapper))
	}
	return res
}

func Int16ValFilterToInt16Filter(filter Int16ValFilter) Int16Filter {
	return func(i int, val int16) bool {
		return filter.Something()(val)
	}
}

func Int16ValFilterBatchToInt16Filter(filters ...Int16ValFilter) []Int16Filter {
	res := make([]Int16Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Int16ValFilterToInt16Filter(filter))
	}
	return res
}

func Int16If(f BoolFunc, a, b int16) int16 {
	if f() {
		return a
	}
	return b
}

// autogen @type: int8  @aliasPreix: Int8

// ===  gen area =====
type Int8Filter func(int, int8) bool
type Int8Mapper func(int, int8) int8
type Int8Ranger func(int, int8)

func (filter Int8Filter) Something() Int8Filter {
	if filter == nil {
		return func(int, int8) bool {
			return false
		}
	}
	return filter
}

func (mapper Int8Mapper) Something() Int8Mapper {
	if mapper == nil {
		return func(i int, val int8) int8 {
			return val
		}
	}
	return mapper
}

func (ranger Int8Ranger) Something() Int8Ranger {
	if ranger == nil {
		return func(int, int8) {}
	}
	return ranger
}

type Int8ValFilter func(int8) bool
type Int8ValMapper func(int8) int8
type Int8ValRanger func(int8)

func (filter Int8ValFilter) Something() Int8ValFilter {
	if filter == nil {
		return func(int8) bool {
			return false
		}
	}
	return filter
}

func (mapper Int8ValMapper) Something() Int8ValMapper {
	if mapper == nil {
		return func(val int8) int8 {
			return val
		}
	}
	return mapper
}

func (ranger Int8ValRanger) Something() Int8ValRanger {
	if ranger == nil {
		return func(int8) {}
	}
	return ranger
}

// ==== util func

func Int8ValRangerToInt8Ranger(ranger Int8ValRanger) Int8Ranger {
	return func(i int, val int8) {
		ranger.Something()(val)
	}
}

func Int8ValRangerBatchToInt8Ranger(rangers ...Int8ValRanger) []Int8Ranger {
	res := make([]Int8Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Int8ValRangerToInt8Ranger(ranger))
	}
	return res
}

func Int8ValMapperToInt8Mapper(mapper Int8ValMapper) Int8Mapper {
	return func(i int, val int8) int8 {
		return mapper.Something()(val)
	}
}

func Int8ValMapperBatchToInt8Mapper(mappers ...Int8ValMapper) []Int8Mapper {
	res := make([]Int8Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Int8ValMapperToInt8Mapper(mapper))
	}
	return res
}

func Int8ValFilterToInt8Filter(filter Int8ValFilter) Int8Filter {
	return func(i int, val int8) bool {
		return filter.Something()(val)
	}
}

func Int8ValFilterBatchToInt8Filter(filters ...Int8ValFilter) []Int8Filter {
	res := make([]Int8Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Int8ValFilterToInt8Filter(filter))
	}
	return res
}

func Int8If(f BoolFunc, a, b int8) int8 {
	if f() {
		return a
	}
	return b
}

// autogen @type: uint  @aliasPreix: Uint

// ===  gen area =====
type UintFilter func(int, uint) bool
type UintMapper func(int, uint) uint
type UintRanger func(int, uint)

func (filter UintFilter) Something() UintFilter {
	if filter == nil {
		return func(int, uint) bool {
			return false
		}
	}
	return filter
}

func (mapper UintMapper) Something() UintMapper {
	if mapper == nil {
		return func(i int, val uint) uint {
			return val
		}
	}
	return mapper
}

func (ranger UintRanger) Something() UintRanger {
	if ranger == nil {
		return func(int, uint) {}
	}
	return ranger
}

type UintValFilter func(uint) bool
type UintValMapper func(uint) uint
type UintValRanger func(uint)

func (filter UintValFilter) Something() UintValFilter {
	if filter == nil {
		return func(uint) bool {
			return false
		}
	}
	return filter
}

func (mapper UintValMapper) Something() UintValMapper {
	if mapper == nil {
		return func(val uint) uint {
			return val
		}
	}
	return mapper
}

func (ranger UintValRanger) Something() UintValRanger {
	if ranger == nil {
		return func(uint) {}
	}
	return ranger
}

// ==== util func

func UintValRangerToUintRanger(ranger UintValRanger) UintRanger {
	return func(i int, val uint) {
		ranger.Something()(val)
	}
}

func UintValRangerBatchToUintRanger(rangers ...UintValRanger) []UintRanger {
	res := make([]UintRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, UintValRangerToUintRanger(ranger))
	}
	return res
}

func UintValMapperToUintMapper(mapper UintValMapper) UintMapper {
	return func(i int, val uint) uint {
		return mapper.Something()(val)
	}
}

func UintValMapperBatchToUintMapper(mappers ...UintValMapper) []UintMapper {
	res := make([]UintMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, UintValMapperToUintMapper(mapper))
	}
	return res
}

func UintValFilterToUintFilter(filter UintValFilter) UintFilter {
	return func(i int, val uint) bool {
		return filter.Something()(val)
	}
}

func UintValFilterBatchToUintFilter(filters ...UintValFilter) []UintFilter {
	res := make([]UintFilter, len(filters))
	for _, filter := range filters {
		res = append(res, UintValFilterToUintFilter(filter))
	}
	return res
}

func UintIf(f BoolFunc, a, b uint) uint {
	if f() {
		return a
	}
	return b
}

// autogen @type: uint8  @aliasPreix: Uint8

// ===  gen area =====
type Uint8Filter func(int, uint8) bool
type Uint8Mapper func(int, uint8) uint8
type Uint8Ranger func(int, uint8)

func (filter Uint8Filter) Something() Uint8Filter {
	if filter == nil {
		return func(int, uint8) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint8Mapper) Something() Uint8Mapper {
	if mapper == nil {
		return func(i int, val uint8) uint8 {
			return val
		}
	}
	return mapper
}

func (ranger Uint8Ranger) Something() Uint8Ranger {
	if ranger == nil {
		return func(int, uint8) {}
	}
	return ranger
}

type Uint8ValFilter func(uint8) bool
type Uint8ValMapper func(uint8) uint8
type Uint8ValRanger func(uint8)

func (filter Uint8ValFilter) Something() Uint8ValFilter {
	if filter == nil {
		return func(uint8) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint8ValMapper) Something() Uint8ValMapper {
	if mapper == nil {
		return func(val uint8) uint8 {
			return val
		}
	}
	return mapper
}

func (ranger Uint8ValRanger) Something() Uint8ValRanger {
	if ranger == nil {
		return func(uint8) {}
	}
	return ranger
}

// ==== util func

func Uint8ValRangerToUint8Ranger(ranger Uint8ValRanger) Uint8Ranger {
	return func(i int, val uint8) {
		ranger.Something()(val)
	}
}

func Uint8ValRangerBatchToUint8Ranger(rangers ...Uint8ValRanger) []Uint8Ranger {
	res := make([]Uint8Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Uint8ValRangerToUint8Ranger(ranger))
	}
	return res
}

func Uint8ValMapperToUint8Mapper(mapper Uint8ValMapper) Uint8Mapper {
	return func(i int, val uint8) uint8 {
		return mapper.Something()(val)
	}
}

func Uint8ValMapperBatchToUint8Mapper(mappers ...Uint8ValMapper) []Uint8Mapper {
	res := make([]Uint8Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Uint8ValMapperToUint8Mapper(mapper))
	}
	return res
}

func Uint8ValFilterToUint8Filter(filter Uint8ValFilter) Uint8Filter {
	return func(i int, val uint8) bool {
		return filter.Something()(val)
	}
}

func Uint8ValFilterBatchToUint8Filter(filters ...Uint8ValFilter) []Uint8Filter {
	res := make([]Uint8Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Uint8ValFilterToUint8Filter(filter))
	}
	return res
}

func Uint8If(f BoolFunc, a, b uint8) uint8 {
	if f() {
		return a
	}
	return b
}

// autogen @type: uint16  @aliasPreix: Uint16

// ===  gen area =====
type Uint16Filter func(int, uint16) bool
type Uint16Mapper func(int, uint16) uint16
type Uint16Ranger func(int, uint16)

func (filter Uint16Filter) Something() Uint16Filter {
	if filter == nil {
		return func(int, uint16) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint16Mapper) Something() Uint16Mapper {
	if mapper == nil {
		return func(i int, val uint16) uint16 {
			return val
		}
	}
	return mapper
}

func (ranger Uint16Ranger) Something() Uint16Ranger {
	if ranger == nil {
		return func(int, uint16) {}
	}
	return ranger
}

type Uint16ValFilter func(uint16) bool
type Uint16ValMapper func(uint16) uint16
type Uint16ValRanger func(uint16)

func (filter Uint16ValFilter) Something() Uint16ValFilter {
	if filter == nil {
		return func(uint16) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint16ValMapper) Something() Uint16ValMapper {
	if mapper == nil {
		return func(val uint16) uint16 {
			return val
		}
	}
	return mapper
}

func (ranger Uint16ValRanger) Something() Uint16ValRanger {
	if ranger == nil {
		return func(uint16) {}
	}
	return ranger
}

// ==== util func

func Uint16ValRangerToUint16Ranger(ranger Uint16ValRanger) Uint16Ranger {
	return func(i int, val uint16) {
		ranger.Something()(val)
	}
}

func Uint16ValRangerBatchToUint16Ranger(rangers ...Uint16ValRanger) []Uint16Ranger {
	res := make([]Uint16Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Uint16ValRangerToUint16Ranger(ranger))
	}
	return res
}

func Uint16ValMapperToUint16Mapper(mapper Uint16ValMapper) Uint16Mapper {
	return func(i int, val uint16) uint16 {
		return mapper.Something()(val)
	}
}

func Uint16ValMapperBatchToUint16Mapper(mappers ...Uint16ValMapper) []Uint16Mapper {
	res := make([]Uint16Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Uint16ValMapperToUint16Mapper(mapper))
	}
	return res
}

func Uint16ValFilterToUint16Filter(filter Uint16ValFilter) Uint16Filter {
	return func(i int, val uint16) bool {
		return filter.Something()(val)
	}
}

func Uint16ValFilterBatchToUint16Filter(filters ...Uint16ValFilter) []Uint16Filter {
	res := make([]Uint16Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Uint16ValFilterToUint16Filter(filter))
	}
	return res
}

func Uint16If(f BoolFunc, a, b uint16) uint16 {
	if f() {
		return a
	}
	return b
}

// autogen @type: uint32  @aliasPreix: Uint32

// ===  gen area =====
type Uint32Filter func(int, uint32) bool
type Uint32Mapper func(int, uint32) uint32
type Uint32Ranger func(int, uint32)

func (filter Uint32Filter) Something() Uint32Filter {
	if filter == nil {
		return func(int, uint32) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint32Mapper) Something() Uint32Mapper {
	if mapper == nil {
		return func(i int, val uint32) uint32 {
			return val
		}
	}
	return mapper
}

func (ranger Uint32Ranger) Something() Uint32Ranger {
	if ranger == nil {
		return func(int, uint32) {}
	}
	return ranger
}

type Uint32ValFilter func(uint32) bool
type Uint32ValMapper func(uint32) uint32
type Uint32ValRanger func(uint32)

func (filter Uint32ValFilter) Something() Uint32ValFilter {
	if filter == nil {
		return func(uint32) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint32ValMapper) Something() Uint32ValMapper {
	if mapper == nil {
		return func(val uint32) uint32 {
			return val
		}
	}
	return mapper
}

func (ranger Uint32ValRanger) Something() Uint32ValRanger {
	if ranger == nil {
		return func(uint32) {}
	}
	return ranger
}

// ==== util func

func Uint32ValRangerToUint32Ranger(ranger Uint32ValRanger) Uint32Ranger {
	return func(i int, val uint32) {
		ranger.Something()(val)
	}
}

func Uint32ValRangerBatchToUint32Ranger(rangers ...Uint32ValRanger) []Uint32Ranger {
	res := make([]Uint32Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Uint32ValRangerToUint32Ranger(ranger))
	}
	return res
}

func Uint32ValMapperToUint32Mapper(mapper Uint32ValMapper) Uint32Mapper {
	return func(i int, val uint32) uint32 {
		return mapper.Something()(val)
	}
}

func Uint32ValMapperBatchToUint32Mapper(mappers ...Uint32ValMapper) []Uint32Mapper {
	res := make([]Uint32Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Uint32ValMapperToUint32Mapper(mapper))
	}
	return res
}

func Uint32ValFilterToUint32Filter(filter Uint32ValFilter) Uint32Filter {
	return func(i int, val uint32) bool {
		return filter.Something()(val)
	}
}

func Uint32ValFilterBatchToUint32Filter(filters ...Uint32ValFilter) []Uint32Filter {
	res := make([]Uint32Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Uint32ValFilterToUint32Filter(filter))
	}
	return res
}

func Uint32If(f BoolFunc, a, b uint32) uint32 {
	if f() {
		return a
	}
	return b
}

// autogen @type: uint64  @aliasPreix: Uint64

// ===  gen area =====
type Uint64Filter func(int, uint64) bool
type Uint64Mapper func(int, uint64) uint64
type Uint64Ranger func(int, uint64)

func (filter Uint64Filter) Something() Uint64Filter {
	if filter == nil {
		return func(int, uint64) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint64Mapper) Something() Uint64Mapper {
	if mapper == nil {
		return func(i int, val uint64) uint64 {
			return val
		}
	}
	return mapper
}

func (ranger Uint64Ranger) Something() Uint64Ranger {
	if ranger == nil {
		return func(int, uint64) {}
	}
	return ranger
}

type Uint64ValFilter func(uint64) bool
type Uint64ValMapper func(uint64) uint64
type Uint64ValRanger func(uint64)

func (filter Uint64ValFilter) Something() Uint64ValFilter {
	if filter == nil {
		return func(uint64) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint64ValMapper) Something() Uint64ValMapper {
	if mapper == nil {
		return func(val uint64) uint64 {
			return val
		}
	}
	return mapper
}

func (ranger Uint64ValRanger) Something() Uint64ValRanger {
	if ranger == nil {
		return func(uint64) {}
	}
	return ranger
}

// ==== util func

func Uint64ValRangerToUint64Ranger(ranger Uint64ValRanger) Uint64Ranger {
	return func(i int, val uint64) {
		ranger.Something()(val)
	}
}

func Uint64ValRangerBatchToUint64Ranger(rangers ...Uint64ValRanger) []Uint64Ranger {
	res := make([]Uint64Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Uint64ValRangerToUint64Ranger(ranger))
	}
	return res
}

func Uint64ValMapperToUint64Mapper(mapper Uint64ValMapper) Uint64Mapper {
	return func(i int, val uint64) uint64 {
		return mapper.Something()(val)
	}
}

func Uint64ValMapperBatchToUint64Mapper(mappers ...Uint64ValMapper) []Uint64Mapper {
	res := make([]Uint64Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Uint64ValMapperToUint64Mapper(mapper))
	}
	return res
}

func Uint64ValFilterToUint64Filter(filter Uint64ValFilter) Uint64Filter {
	return func(i int, val uint64) bool {
		return filter.Something()(val)
	}
}

func Uint64ValFilterBatchToUint64Filter(filters ...Uint64ValFilter) []Uint64Filter {
	res := make([]Uint64Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Uint64ValFilterToUint64Filter(filter))
	}
	return res
}

func Uint64If(f BoolFunc, a, b uint64) uint64 {
	if f() {
		return a
	}
	return b
}

// autogen @type: float32  @aliasPreix: Float32

// ===  gen area =====
type Float32Filter func(int, float32) bool
type Float32Mapper func(int, float32) float32
type Float32Ranger func(int, float32)

func (filter Float32Filter) Something() Float32Filter {
	if filter == nil {
		return func(int, float32) bool {
			return false
		}
	}
	return filter
}

func (mapper Float32Mapper) Something() Float32Mapper {
	if mapper == nil {
		return func(i int, val float32) float32 {
			return val
		}
	}
	return mapper
}

func (ranger Float32Ranger) Something() Float32Ranger {
	if ranger == nil {
		return func(int, float32) {}
	}
	return ranger
}

type Float32ValFilter func(float32) bool
type Float32ValMapper func(float32) float32
type Float32ValRanger func(float32)

func (filter Float32ValFilter) Something() Float32ValFilter {
	if filter == nil {
		return func(float32) bool {
			return false
		}
	}
	return filter
}

func (mapper Float32ValMapper) Something() Float32ValMapper {
	if mapper == nil {
		return func(val float32) float32 {
			return val
		}
	}
	return mapper
}

func (ranger Float32ValRanger) Something() Float32ValRanger {
	if ranger == nil {
		return func(float32) {}
	}
	return ranger
}

// ==== util func

func Float32ValRangerToFloat32Ranger(ranger Float32ValRanger) Float32Ranger {
	return func(i int, val float32) {
		ranger.Something()(val)
	}
}

func Float32ValRangerBatchToFloat32Ranger(rangers ...Float32ValRanger) []Float32Ranger {
	res := make([]Float32Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Float32ValRangerToFloat32Ranger(ranger))
	}
	return res
}

func Float32ValMapperToFloat32Mapper(mapper Float32ValMapper) Float32Mapper {
	return func(i int, val float32) float32 {
		return mapper.Something()(val)
	}
}

func Float32ValMapperBatchToFloat32Mapper(mappers ...Float32ValMapper) []Float32Mapper {
	res := make([]Float32Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Float32ValMapperToFloat32Mapper(mapper))
	}
	return res
}

func Float32ValFilterToFloat32Filter(filter Float32ValFilter) Float32Filter {
	return func(i int, val float32) bool {
		return filter.Something()(val)
	}
}

func Float32ValFilterBatchToFloat32Filter(filters ...Float32ValFilter) []Float32Filter {
	res := make([]Float32Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Float32ValFilterToFloat32Filter(filter))
	}
	return res
}

func Float32If(f BoolFunc, a, b float32) float32 {
	if f() {
		return a
	}
	return b
}

// autogen @type: float64  @aliasPreix: Float64

// ===  gen area =====
type Float64Filter func(int, float64) bool
type Float64Mapper func(int, float64) float64
type Float64Ranger func(int, float64)

func (filter Float64Filter) Something() Float64Filter {
	if filter == nil {
		return func(int, float64) bool {
			return false
		}
	}
	return filter
}

func (mapper Float64Mapper) Something() Float64Mapper {
	if mapper == nil {
		return func(i int, val float64) float64 {
			return val
		}
	}
	return mapper
}

func (ranger Float64Ranger) Something() Float64Ranger {
	if ranger == nil {
		return func(int, float64) {}
	}
	return ranger
}

type Float64ValFilter func(float64) bool
type Float64ValMapper func(float64) float64
type Float64ValRanger func(float64)

func (filter Float64ValFilter) Something() Float64ValFilter {
	if filter == nil {
		return func(float64) bool {
			return false
		}
	}
	return filter
}

func (mapper Float64ValMapper) Something() Float64ValMapper {
	if mapper == nil {
		return func(val float64) float64 {
			return val
		}
	}
	return mapper
}

func (ranger Float64ValRanger) Something() Float64ValRanger {
	if ranger == nil {
		return func(float64) {}
	}
	return ranger
}

// ==== util func

func Float64ValRangerToFloat64Ranger(ranger Float64ValRanger) Float64Ranger {
	return func(i int, val float64) {
		ranger.Something()(val)
	}
}

func Float64ValRangerBatchToFloat64Ranger(rangers ...Float64ValRanger) []Float64Ranger {
	res := make([]Float64Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Float64ValRangerToFloat64Ranger(ranger))
	}
	return res
}

func Float64ValMapperToFloat64Mapper(mapper Float64ValMapper) Float64Mapper {
	return func(i int, val float64) float64 {
		return mapper.Something()(val)
	}
}

func Float64ValMapperBatchToFloat64Mapper(mappers ...Float64ValMapper) []Float64Mapper {
	res := make([]Float64Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Float64ValMapperToFloat64Mapper(mapper))
	}
	return res
}

func Float64ValFilterToFloat64Filter(filter Float64ValFilter) Float64Filter {
	return func(i int, val float64) bool {
		return filter.Something()(val)
	}
}

func Float64ValFilterBatchToFloat64Filter(filters ...Float64ValFilter) []Float64Filter {
	res := make([]Float64Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Float64ValFilterToFloat64Filter(filter))
	}
	return res
}

func Float64If(f BoolFunc, a, b float64) float64 {
	if f() {
		return a
	}
	return b
}

// autogen @type: bool  @aliasPreix: Bool

// ===  gen area =====
type BoolFilter func(int, bool) bool
type BoolMapper func(int, bool) bool
type BoolRanger func(int, bool)

func (filter BoolFilter) Something() BoolFilter {
	if filter == nil {
		return func(int, bool) bool {
			return false
		}
	}
	return filter
}

func (mapper BoolMapper) Something() BoolMapper {
	if mapper == nil {
		return func(i int, val bool) bool {
			return val
		}
	}
	return mapper
}

func (ranger BoolRanger) Something() BoolRanger {
	if ranger == nil {
		return func(int, bool) {}
	}
	return ranger
}

type BoolValFilter func(bool) bool
type BoolValMapper func(bool) bool
type BoolValRanger func(bool)

func (filter BoolValFilter) Something() BoolValFilter {
	if filter == nil {
		return func(bool) bool {
			return false
		}
	}
	return filter
}

func (mapper BoolValMapper) Something() BoolValMapper {
	if mapper == nil {
		return func(val bool) bool {
			return val
		}
	}
	return mapper
}

func (ranger BoolValRanger) Something() BoolValRanger {
	if ranger == nil {
		return func(bool) {}
	}
	return ranger
}

// ==== util func

func BoolValRangerToBoolRanger(ranger BoolValRanger) BoolRanger {
	return func(i int, val bool) {
		ranger.Something()(val)
	}
}

func BoolValRangerBatchToBoolRanger(rangers ...BoolValRanger) []BoolRanger {
	res := make([]BoolRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, BoolValRangerToBoolRanger(ranger))
	}
	return res
}

func BoolValMapperToBoolMapper(mapper BoolValMapper) BoolMapper {
	return func(i int, val bool) bool {
		return mapper.Something()(val)
	}
}

func BoolValMapperBatchToBoolMapper(mappers ...BoolValMapper) []BoolMapper {
	res := make([]BoolMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, BoolValMapperToBoolMapper(mapper))
	}
	return res
}

func BoolValFilterToBoolFilter(filter BoolValFilter) BoolFilter {
	return func(i int, val bool) bool {
		return filter.Something()(val)
	}
}

func BoolValFilterBatchToBoolFilter(filters ...BoolValFilter) []BoolFilter {
	res := make([]BoolFilter, len(filters))
	for _, filter := range filters {
		res = append(res, BoolValFilterToBoolFilter(filter))
	}
	return res
}

func BoolIf(f BoolFunc, a, b bool) bool {
	if f() {
		return a
	}
	return b
}

// autogen @type: rune  @aliasPreix: Rune

// ===  gen area =====
type RuneFilter func(int, rune) bool
type RuneMapper func(int, rune) rune
type RuneRanger func(int, rune)

func (filter RuneFilter) Something() RuneFilter {
	if filter == nil {
		return func(int, rune) bool {
			return false
		}
	}
	return filter
}

func (mapper RuneMapper) Something() RuneMapper {
	if mapper == nil {
		return func(i int, val rune) rune {
			return val
		}
	}
	return mapper
}

func (ranger RuneRanger) Something() RuneRanger {
	if ranger == nil {
		return func(int, rune) {}
	}
	return ranger
}

type RuneValFilter func(rune) bool
type RuneValMapper func(rune) rune
type RuneValRanger func(rune)

func (filter RuneValFilter) Something() RuneValFilter {
	if filter == nil {
		return func(rune) bool {
			return false
		}
	}
	return filter
}

func (mapper RuneValMapper) Something() RuneValMapper {
	if mapper == nil {
		return func(val rune) rune {
			return val
		}
	}
	return mapper
}

func (ranger RuneValRanger) Something() RuneValRanger {
	if ranger == nil {
		return func(rune) {}
	}
	return ranger
}

// ==== util func

func RuneValRangerToRuneRanger(ranger RuneValRanger) RuneRanger {
	return func(i int, val rune) {
		ranger.Something()(val)
	}
}

func RuneValRangerBatchToRuneRanger(rangers ...RuneValRanger) []RuneRanger {
	res := make([]RuneRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, RuneValRangerToRuneRanger(ranger))
	}
	return res
}

func RuneValMapperToRuneMapper(mapper RuneValMapper) RuneMapper {
	return func(i int, val rune) rune {
		return mapper.Something()(val)
	}
}

func RuneValMapperBatchToRuneMapper(mappers ...RuneValMapper) []RuneMapper {
	res := make([]RuneMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, RuneValMapperToRuneMapper(mapper))
	}
	return res
}

func RuneValFilterToRuneFilter(filter RuneValFilter) RuneFilter {
	return func(i int, val rune) bool {
		return filter.Something()(val)
	}
}

func RuneValFilterBatchToRuneFilter(filters ...RuneValFilter) []RuneFilter {
	res := make([]RuneFilter, len(filters))
	for _, filter := range filters {
		res = append(res, RuneValFilterToRuneFilter(filter))
	}
	return res
}

func RuneIf(f BoolFunc, a, b rune) rune {
	if f() {
		return a
	}
	return b
}

// autogen @type: complex64  @aliasPreix: Complex64

// ===  gen area =====
type Complex64Filter func(int, complex64) bool
type Complex64Mapper func(int, complex64) complex64
type Complex64Ranger func(int, complex64)

func (filter Complex64Filter) Something() Complex64Filter {
	if filter == nil {
		return func(int, complex64) bool {
			return false
		}
	}
	return filter
}

func (mapper Complex64Mapper) Something() Complex64Mapper {
	if mapper == nil {
		return func(i int, val complex64) complex64 {
			return val
		}
	}
	return mapper
}

func (ranger Complex64Ranger) Something() Complex64Ranger {
	if ranger == nil {
		return func(int, complex64) {}
	}
	return ranger
}

type Complex64ValFilter func(complex64) bool
type Complex64ValMapper func(complex64) complex64
type Complex64ValRanger func(complex64)

func (filter Complex64ValFilter) Something() Complex64ValFilter {
	if filter == nil {
		return func(complex64) bool {
			return false
		}
	}
	return filter
}

func (mapper Complex64ValMapper) Something() Complex64ValMapper {
	if mapper == nil {
		return func(val complex64) complex64 {
			return val
		}
	}
	return mapper
}

func (ranger Complex64ValRanger) Something() Complex64ValRanger {
	if ranger == nil {
		return func(complex64) {}
	}
	return ranger
}

// ==== util func

func Complex64ValRangerToComplex64Ranger(ranger Complex64ValRanger) Complex64Ranger {
	return func(i int, val complex64) {
		ranger.Something()(val)
	}
}

func Complex64ValRangerBatchToComplex64Ranger(rangers ...Complex64ValRanger) []Complex64Ranger {
	res := make([]Complex64Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Complex64ValRangerToComplex64Ranger(ranger))
	}
	return res
}

func Complex64ValMapperToComplex64Mapper(mapper Complex64ValMapper) Complex64Mapper {
	return func(i int, val complex64) complex64 {
		return mapper.Something()(val)
	}
}

func Complex64ValMapperBatchToComplex64Mapper(mappers ...Complex64ValMapper) []Complex64Mapper {
	res := make([]Complex64Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Complex64ValMapperToComplex64Mapper(mapper))
	}
	return res
}

func Complex64ValFilterToComplex64Filter(filter Complex64ValFilter) Complex64Filter {
	return func(i int, val complex64) bool {
		return filter.Something()(val)
	}
}

func Complex64ValFilterBatchToComplex64Filter(filters ...Complex64ValFilter) []Complex64Filter {
	res := make([]Complex64Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Complex64ValFilterToComplex64Filter(filter))
	}
	return res
}

func Complex64If(f BoolFunc, a, b complex64) complex64 {
	if f() {
		return a
	}
	return b
}

// autogen @type: complex32  @aliasPreix: Complex32

// ===  gen area =====
type Complex32Filter func(int, complex32) bool
type Complex32Mapper func(int, complex32) complex32
type Complex32Ranger func(int, complex32)

func (filter Complex32Filter) Something() Complex32Filter {
	if filter == nil {
		return func(int, complex32) bool {
			return false
		}
	}
	return filter
}

func (mapper Complex32Mapper) Something() Complex32Mapper {
	if mapper == nil {
		return func(i int, val complex32) complex32 {
			return val
		}
	}
	return mapper
}

func (ranger Complex32Ranger) Something() Complex32Ranger {
	if ranger == nil {
		return func(int, complex32) {}
	}
	return ranger
}

type Complex32ValFilter func(complex32) bool
type Complex32ValMapper func(complex32) complex32
type Complex32ValRanger func(complex32)

func (filter Complex32ValFilter) Something() Complex32ValFilter {
	if filter == nil {
		return func(complex32) bool {
			return false
		}
	}
	return filter
}

func (mapper Complex32ValMapper) Something() Complex32ValMapper {
	if mapper == nil {
		return func(val complex32) complex32 {
			return val
		}
	}
	return mapper
}

func (ranger Complex32ValRanger) Something() Complex32ValRanger {
	if ranger == nil {
		return func(complex32) {}
	}
	return ranger
}

// ==== util func

func Complex32ValRangerToComplex32Ranger(ranger Complex32ValRanger) Complex32Ranger {
	return func(i int, val complex32) {
		ranger.Something()(val)
	}
}

func Complex32ValRangerBatchToComplex32Ranger(rangers ...Complex32ValRanger) []Complex32Ranger {
	res := make([]Complex32Ranger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Complex32ValRangerToComplex32Ranger(ranger))
	}
	return res
}

func Complex32ValMapperToComplex32Mapper(mapper Complex32ValMapper) Complex32Mapper {
	return func(i int, val complex32) complex32 {
		return mapper.Something()(val)
	}
}

func Complex32ValMapperBatchToComplex32Mapper(mappers ...Complex32ValMapper) []Complex32Mapper {
	res := make([]Complex32Mapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Complex32ValMapperToComplex32Mapper(mapper))
	}
	return res
}

func Complex32ValFilterToComplex32Filter(filter Complex32ValFilter) Complex32Filter {
	return func(i int, val complex32) bool {
		return filter.Something()(val)
	}
}

func Complex32ValFilterBatchToComplex32Filter(filters ...Complex32ValFilter) []Complex32Filter {
	res := make([]Complex32Filter, len(filters))
	for _, filter := range filters {
		res = append(res, Complex32ValFilterToComplex32Filter(filter))
	}
	return res
}

func Complex32If(f BoolFunc, a, b complex32) complex32 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *string  @aliasPreix: StringPtr

// ===  gen area =====
type StringPtrFilter func(int, *string) bool
type StringPtrMapper func(int, *string) *string
type StringPtrRanger func(int, *string)

func (filter StringPtrFilter) Something() StringPtrFilter {
	if filter == nil {
		return func(int, *string) bool {
			return false
		}
	}
	return filter
}

func (mapper StringPtrMapper) Something() StringPtrMapper {
	if mapper == nil {
		return func(i int, val *string) *string {
			return val
		}
	}
	return mapper
}

func (ranger StringPtrRanger) Something() StringPtrRanger {
	if ranger == nil {
		return func(int, *string) {}
	}
	return ranger
}

type StringPtrValFilter func(*string) bool
type StringPtrValMapper func(*string) *string
type StringPtrValRanger func(*string)

func (filter StringPtrValFilter) Something() StringPtrValFilter {
	if filter == nil {
		return func(*string) bool {
			return false
		}
	}
	return filter
}

func (mapper StringPtrValMapper) Something() StringPtrValMapper {
	if mapper == nil {
		return func(val *string) *string {
			return val
		}
	}
	return mapper
}

func (ranger StringPtrValRanger) Something() StringPtrValRanger {
	if ranger == nil {
		return func(*string) {}
	}
	return ranger
}

// ==== util func

func StringPtrValRangerToStringPtrRanger(ranger StringPtrValRanger) StringPtrRanger {
	return func(i int, val *string) {
		ranger.Something()(val)
	}
}

func StringPtrValRangerBatchToStringPtrRanger(rangers ...StringPtrValRanger) []StringPtrRanger {
	res := make([]StringPtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, StringPtrValRangerToStringPtrRanger(ranger))
	}
	return res
}

func StringPtrValMapperToStringPtrMapper(mapper StringPtrValMapper) StringPtrMapper {
	return func(i int, val *string) *string {
		return mapper.Something()(val)
	}
}

func StringPtrValMapperBatchToStringPtrMapper(mappers ...StringPtrValMapper) []StringPtrMapper {
	res := make([]StringPtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, StringPtrValMapperToStringPtrMapper(mapper))
	}
	return res
}

func StringPtrValFilterToStringPtrFilter(filter StringPtrValFilter) StringPtrFilter {
	return func(i int, val *string) bool {
		return filter.Something()(val)
	}
}

func StringPtrValFilterBatchToStringPtrFilter(filters ...StringPtrValFilter) []StringPtrFilter {
	res := make([]StringPtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, StringPtrValFilterToStringPtrFilter(filter))
	}
	return res
}

func StringPtrIf(f BoolFunc, a, b *string) *string {
	if f() {
		return a
	}
	return b
}

// autogen @type: *int  @aliasPreix: IntPtr

// ===  gen area =====
type IntPtrFilter func(int, *int) bool
type IntPtrMapper func(int, *int) *int
type IntPtrRanger func(int, *int)

func (filter IntPtrFilter) Something() IntPtrFilter {
	if filter == nil {
		return func(int, *int) bool {
			return false
		}
	}
	return filter
}

func (mapper IntPtrMapper) Something() IntPtrMapper {
	if mapper == nil {
		return func(i int, val *int) *int {
			return val
		}
	}
	return mapper
}

func (ranger IntPtrRanger) Something() IntPtrRanger {
	if ranger == nil {
		return func(int, *int) {}
	}
	return ranger
}

type IntPtrValFilter func(*int) bool
type IntPtrValMapper func(*int) *int
type IntPtrValRanger func(*int)

func (filter IntPtrValFilter) Something() IntPtrValFilter {
	if filter == nil {
		return func(*int) bool {
			return false
		}
	}
	return filter
}

func (mapper IntPtrValMapper) Something() IntPtrValMapper {
	if mapper == nil {
		return func(val *int) *int {
			return val
		}
	}
	return mapper
}

func (ranger IntPtrValRanger) Something() IntPtrValRanger {
	if ranger == nil {
		return func(*int) {}
	}
	return ranger
}

// ==== util func

func IntPtrValRangerToIntPtrRanger(ranger IntPtrValRanger) IntPtrRanger {
	return func(i int, val *int) {
		ranger.Something()(val)
	}
}

func IntPtrValRangerBatchToIntPtrRanger(rangers ...IntPtrValRanger) []IntPtrRanger {
	res := make([]IntPtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, IntPtrValRangerToIntPtrRanger(ranger))
	}
	return res
}

func IntPtrValMapperToIntPtrMapper(mapper IntPtrValMapper) IntPtrMapper {
	return func(i int, val *int) *int {
		return mapper.Something()(val)
	}
}

func IntPtrValMapperBatchToIntPtrMapper(mappers ...IntPtrValMapper) []IntPtrMapper {
	res := make([]IntPtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, IntPtrValMapperToIntPtrMapper(mapper))
	}
	return res
}

func IntPtrValFilterToIntPtrFilter(filter IntPtrValFilter) IntPtrFilter {
	return func(i int, val *int) bool {
		return filter.Something()(val)
	}
}

func IntPtrValFilterBatchToIntPtrFilter(filters ...IntPtrValFilter) []IntPtrFilter {
	res := make([]IntPtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, IntPtrValFilterToIntPtrFilter(filter))
	}
	return res
}

func IntPtrIf(f BoolFunc, a, b *int) *int {
	if f() {
		return a
	}
	return b
}

// autogen @type: *int64  @aliasPreix: Int64Ptr

// ===  gen area =====
type Int64PtrFilter func(int, *int64) bool
type Int64PtrMapper func(int, *int64) *int64
type Int64PtrRanger func(int, *int64)

func (filter Int64PtrFilter) Something() Int64PtrFilter {
	if filter == nil {
		return func(int, *int64) bool {
			return false
		}
	}
	return filter
}

func (mapper Int64PtrMapper) Something() Int64PtrMapper {
	if mapper == nil {
		return func(i int, val *int64) *int64 {
			return val
		}
	}
	return mapper
}

func (ranger Int64PtrRanger) Something() Int64PtrRanger {
	if ranger == nil {
		return func(int, *int64) {}
	}
	return ranger
}

type Int64PtrValFilter func(*int64) bool
type Int64PtrValMapper func(*int64) *int64
type Int64PtrValRanger func(*int64)

func (filter Int64PtrValFilter) Something() Int64PtrValFilter {
	if filter == nil {
		return func(*int64) bool {
			return false
		}
	}
	return filter
}

func (mapper Int64PtrValMapper) Something() Int64PtrValMapper {
	if mapper == nil {
		return func(val *int64) *int64 {
			return val
		}
	}
	return mapper
}

func (ranger Int64PtrValRanger) Something() Int64PtrValRanger {
	if ranger == nil {
		return func(*int64) {}
	}
	return ranger
}

// ==== util func

func Int64PtrValRangerToInt64PtrRanger(ranger Int64PtrValRanger) Int64PtrRanger {
	return func(i int, val *int64) {
		ranger.Something()(val)
	}
}

func Int64PtrValRangerBatchToInt64PtrRanger(rangers ...Int64PtrValRanger) []Int64PtrRanger {
	res := make([]Int64PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Int64PtrValRangerToInt64PtrRanger(ranger))
	}
	return res
}

func Int64PtrValMapperToInt64PtrMapper(mapper Int64PtrValMapper) Int64PtrMapper {
	return func(i int, val *int64) *int64 {
		return mapper.Something()(val)
	}
}

func Int64PtrValMapperBatchToInt64PtrMapper(mappers ...Int64PtrValMapper) []Int64PtrMapper {
	res := make([]Int64PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Int64PtrValMapperToInt64PtrMapper(mapper))
	}
	return res
}

func Int64PtrValFilterToInt64PtrFilter(filter Int64PtrValFilter) Int64PtrFilter {
	return func(i int, val *int64) bool {
		return filter.Something()(val)
	}
}

func Int64PtrValFilterBatchToInt64PtrFilter(filters ...Int64PtrValFilter) []Int64PtrFilter {
	res := make([]Int64PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Int64PtrValFilterToInt64PtrFilter(filter))
	}
	return res
}

func Int64PtrIf(f BoolFunc, a, b *int64) *int64 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *int32  @aliasPreix: Int32Ptr

// ===  gen area =====
type Int32PtrFilter func(int, *int32) bool
type Int32PtrMapper func(int, *int32) *int32
type Int32PtrRanger func(int, *int32)

func (filter Int32PtrFilter) Something() Int32PtrFilter {
	if filter == nil {
		return func(int, *int32) bool {
			return false
		}
	}
	return filter
}

func (mapper Int32PtrMapper) Something() Int32PtrMapper {
	if mapper == nil {
		return func(i int, val *int32) *int32 {
			return val
		}
	}
	return mapper
}

func (ranger Int32PtrRanger) Something() Int32PtrRanger {
	if ranger == nil {
		return func(int, *int32) {}
	}
	return ranger
}

type Int32PtrValFilter func(*int32) bool
type Int32PtrValMapper func(*int32) *int32
type Int32PtrValRanger func(*int32)

func (filter Int32PtrValFilter) Something() Int32PtrValFilter {
	if filter == nil {
		return func(*int32) bool {
			return false
		}
	}
	return filter
}

func (mapper Int32PtrValMapper) Something() Int32PtrValMapper {
	if mapper == nil {
		return func(val *int32) *int32 {
			return val
		}
	}
	return mapper
}

func (ranger Int32PtrValRanger) Something() Int32PtrValRanger {
	if ranger == nil {
		return func(*int32) {}
	}
	return ranger
}

// ==== util func

func Int32PtrValRangerToInt32PtrRanger(ranger Int32PtrValRanger) Int32PtrRanger {
	return func(i int, val *int32) {
		ranger.Something()(val)
	}
}

func Int32PtrValRangerBatchToInt32PtrRanger(rangers ...Int32PtrValRanger) []Int32PtrRanger {
	res := make([]Int32PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Int32PtrValRangerToInt32PtrRanger(ranger))
	}
	return res
}

func Int32PtrValMapperToInt32PtrMapper(mapper Int32PtrValMapper) Int32PtrMapper {
	return func(i int, val *int32) *int32 {
		return mapper.Something()(val)
	}
}

func Int32PtrValMapperBatchToInt32PtrMapper(mappers ...Int32PtrValMapper) []Int32PtrMapper {
	res := make([]Int32PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Int32PtrValMapperToInt32PtrMapper(mapper))
	}
	return res
}

func Int32PtrValFilterToInt32PtrFilter(filter Int32PtrValFilter) Int32PtrFilter {
	return func(i int, val *int32) bool {
		return filter.Something()(val)
	}
}

func Int32PtrValFilterBatchToInt32PtrFilter(filters ...Int32PtrValFilter) []Int32PtrFilter {
	res := make([]Int32PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Int32PtrValFilterToInt32PtrFilter(filter))
	}
	return res
}

func Int32PtrIf(f BoolFunc, a, b *int32) *int32 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *int16  @aliasPreix: Int16Ptr

// ===  gen area =====
type Int16PtrFilter func(int, *int16) bool
type Int16PtrMapper func(int, *int16) *int16
type Int16PtrRanger func(int, *int16)

func (filter Int16PtrFilter) Something() Int16PtrFilter {
	if filter == nil {
		return func(int, *int16) bool {
			return false
		}
	}
	return filter
}

func (mapper Int16PtrMapper) Something() Int16PtrMapper {
	if mapper == nil {
		return func(i int, val *int16) *int16 {
			return val
		}
	}
	return mapper
}

func (ranger Int16PtrRanger) Something() Int16PtrRanger {
	if ranger == nil {
		return func(int, *int16) {}
	}
	return ranger
}

type Int16PtrValFilter func(*int16) bool
type Int16PtrValMapper func(*int16) *int16
type Int16PtrValRanger func(*int16)

func (filter Int16PtrValFilter) Something() Int16PtrValFilter {
	if filter == nil {
		return func(*int16) bool {
			return false
		}
	}
	return filter
}

func (mapper Int16PtrValMapper) Something() Int16PtrValMapper {
	if mapper == nil {
		return func(val *int16) *int16 {
			return val
		}
	}
	return mapper
}

func (ranger Int16PtrValRanger) Something() Int16PtrValRanger {
	if ranger == nil {
		return func(*int16) {}
	}
	return ranger
}

// ==== util func

func Int16PtrValRangerToInt16PtrRanger(ranger Int16PtrValRanger) Int16PtrRanger {
	return func(i int, val *int16) {
		ranger.Something()(val)
	}
}

func Int16PtrValRangerBatchToInt16PtrRanger(rangers ...Int16PtrValRanger) []Int16PtrRanger {
	res := make([]Int16PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Int16PtrValRangerToInt16PtrRanger(ranger))
	}
	return res
}

func Int16PtrValMapperToInt16PtrMapper(mapper Int16PtrValMapper) Int16PtrMapper {
	return func(i int, val *int16) *int16 {
		return mapper.Something()(val)
	}
}

func Int16PtrValMapperBatchToInt16PtrMapper(mappers ...Int16PtrValMapper) []Int16PtrMapper {
	res := make([]Int16PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Int16PtrValMapperToInt16PtrMapper(mapper))
	}
	return res
}

func Int16PtrValFilterToInt16PtrFilter(filter Int16PtrValFilter) Int16PtrFilter {
	return func(i int, val *int16) bool {
		return filter.Something()(val)
	}
}

func Int16PtrValFilterBatchToInt16PtrFilter(filters ...Int16PtrValFilter) []Int16PtrFilter {
	res := make([]Int16PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Int16PtrValFilterToInt16PtrFilter(filter))
	}
	return res
}

func Int16PtrIf(f BoolFunc, a, b *int16) *int16 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *int8  @aliasPreix: Int8Ptr

// ===  gen area =====
type Int8PtrFilter func(int, *int8) bool
type Int8PtrMapper func(int, *int8) *int8
type Int8PtrRanger func(int, *int8)

func (filter Int8PtrFilter) Something() Int8PtrFilter {
	if filter == nil {
		return func(int, *int8) bool {
			return false
		}
	}
	return filter
}

func (mapper Int8PtrMapper) Something() Int8PtrMapper {
	if mapper == nil {
		return func(i int, val *int8) *int8 {
			return val
		}
	}
	return mapper
}

func (ranger Int8PtrRanger) Something() Int8PtrRanger {
	if ranger == nil {
		return func(int, *int8) {}
	}
	return ranger
}

type Int8PtrValFilter func(*int8) bool
type Int8PtrValMapper func(*int8) *int8
type Int8PtrValRanger func(*int8)

func (filter Int8PtrValFilter) Something() Int8PtrValFilter {
	if filter == nil {
		return func(*int8) bool {
			return false
		}
	}
	return filter
}

func (mapper Int8PtrValMapper) Something() Int8PtrValMapper {
	if mapper == nil {
		return func(val *int8) *int8 {
			return val
		}
	}
	return mapper
}

func (ranger Int8PtrValRanger) Something() Int8PtrValRanger {
	if ranger == nil {
		return func(*int8) {}
	}
	return ranger
}

// ==== util func

func Int8PtrValRangerToInt8PtrRanger(ranger Int8PtrValRanger) Int8PtrRanger {
	return func(i int, val *int8) {
		ranger.Something()(val)
	}
}

func Int8PtrValRangerBatchToInt8PtrRanger(rangers ...Int8PtrValRanger) []Int8PtrRanger {
	res := make([]Int8PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Int8PtrValRangerToInt8PtrRanger(ranger))
	}
	return res
}

func Int8PtrValMapperToInt8PtrMapper(mapper Int8PtrValMapper) Int8PtrMapper {
	return func(i int, val *int8) *int8 {
		return mapper.Something()(val)
	}
}

func Int8PtrValMapperBatchToInt8PtrMapper(mappers ...Int8PtrValMapper) []Int8PtrMapper {
	res := make([]Int8PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Int8PtrValMapperToInt8PtrMapper(mapper))
	}
	return res
}

func Int8PtrValFilterToInt8PtrFilter(filter Int8PtrValFilter) Int8PtrFilter {
	return func(i int, val *int8) bool {
		return filter.Something()(val)
	}
}

func Int8PtrValFilterBatchToInt8PtrFilter(filters ...Int8PtrValFilter) []Int8PtrFilter {
	res := make([]Int8PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Int8PtrValFilterToInt8PtrFilter(filter))
	}
	return res
}

func Int8PtrIf(f BoolFunc, a, b *int8) *int8 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *uint  @aliasPreix: UintPtr

// ===  gen area =====
type UintPtrFilter func(int, *uint) bool
type UintPtrMapper func(int, *uint) *uint
type UintPtrRanger func(int, *uint)

func (filter UintPtrFilter) Something() UintPtrFilter {
	if filter == nil {
		return func(int, *uint) bool {
			return false
		}
	}
	return filter
}

func (mapper UintPtrMapper) Something() UintPtrMapper {
	if mapper == nil {
		return func(i int, val *uint) *uint {
			return val
		}
	}
	return mapper
}

func (ranger UintPtrRanger) Something() UintPtrRanger {
	if ranger == nil {
		return func(int, *uint) {}
	}
	return ranger
}

type UintPtrValFilter func(*uint) bool
type UintPtrValMapper func(*uint) *uint
type UintPtrValRanger func(*uint)

func (filter UintPtrValFilter) Something() UintPtrValFilter {
	if filter == nil {
		return func(*uint) bool {
			return false
		}
	}
	return filter
}

func (mapper UintPtrValMapper) Something() UintPtrValMapper {
	if mapper == nil {
		return func(val *uint) *uint {
			return val
		}
	}
	return mapper
}

func (ranger UintPtrValRanger) Something() UintPtrValRanger {
	if ranger == nil {
		return func(*uint) {}
	}
	return ranger
}

// ==== util func

func UintPtrValRangerToUintPtrRanger(ranger UintPtrValRanger) UintPtrRanger {
	return func(i int, val *uint) {
		ranger.Something()(val)
	}
}

func UintPtrValRangerBatchToUintPtrRanger(rangers ...UintPtrValRanger) []UintPtrRanger {
	res := make([]UintPtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, UintPtrValRangerToUintPtrRanger(ranger))
	}
	return res
}

func UintPtrValMapperToUintPtrMapper(mapper UintPtrValMapper) UintPtrMapper {
	return func(i int, val *uint) *uint {
		return mapper.Something()(val)
	}
}

func UintPtrValMapperBatchToUintPtrMapper(mappers ...UintPtrValMapper) []UintPtrMapper {
	res := make([]UintPtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, UintPtrValMapperToUintPtrMapper(mapper))
	}
	return res
}

func UintPtrValFilterToUintPtrFilter(filter UintPtrValFilter) UintPtrFilter {
	return func(i int, val *uint) bool {
		return filter.Something()(val)
	}
}

func UintPtrValFilterBatchToUintPtrFilter(filters ...UintPtrValFilter) []UintPtrFilter {
	res := make([]UintPtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, UintPtrValFilterToUintPtrFilter(filter))
	}
	return res
}

func UintPtrIf(f BoolFunc, a, b *uint) *uint {
	if f() {
		return a
	}
	return b
}

// autogen @type: *uint8  @aliasPreix: Uint8Ptr

// ===  gen area =====
type Uint8PtrFilter func(int, *uint8) bool
type Uint8PtrMapper func(int, *uint8) *uint8
type Uint8PtrRanger func(int, *uint8)

func (filter Uint8PtrFilter) Something() Uint8PtrFilter {
	if filter == nil {
		return func(int, *uint8) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint8PtrMapper) Something() Uint8PtrMapper {
	if mapper == nil {
		return func(i int, val *uint8) *uint8 {
			return val
		}
	}
	return mapper
}

func (ranger Uint8PtrRanger) Something() Uint8PtrRanger {
	if ranger == nil {
		return func(int, *uint8) {}
	}
	return ranger
}

type Uint8PtrValFilter func(*uint8) bool
type Uint8PtrValMapper func(*uint8) *uint8
type Uint8PtrValRanger func(*uint8)

func (filter Uint8PtrValFilter) Something() Uint8PtrValFilter {
	if filter == nil {
		return func(*uint8) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint8PtrValMapper) Something() Uint8PtrValMapper {
	if mapper == nil {
		return func(val *uint8) *uint8 {
			return val
		}
	}
	return mapper
}

func (ranger Uint8PtrValRanger) Something() Uint8PtrValRanger {
	if ranger == nil {
		return func(*uint8) {}
	}
	return ranger
}

// ==== util func

func Uint8PtrValRangerToUint8PtrRanger(ranger Uint8PtrValRanger) Uint8PtrRanger {
	return func(i int, val *uint8) {
		ranger.Something()(val)
	}
}

func Uint8PtrValRangerBatchToUint8PtrRanger(rangers ...Uint8PtrValRanger) []Uint8PtrRanger {
	res := make([]Uint8PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Uint8PtrValRangerToUint8PtrRanger(ranger))
	}
	return res
}

func Uint8PtrValMapperToUint8PtrMapper(mapper Uint8PtrValMapper) Uint8PtrMapper {
	return func(i int, val *uint8) *uint8 {
		return mapper.Something()(val)
	}
}

func Uint8PtrValMapperBatchToUint8PtrMapper(mappers ...Uint8PtrValMapper) []Uint8PtrMapper {
	res := make([]Uint8PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Uint8PtrValMapperToUint8PtrMapper(mapper))
	}
	return res
}

func Uint8PtrValFilterToUint8PtrFilter(filter Uint8PtrValFilter) Uint8PtrFilter {
	return func(i int, val *uint8) bool {
		return filter.Something()(val)
	}
}

func Uint8PtrValFilterBatchToUint8PtrFilter(filters ...Uint8PtrValFilter) []Uint8PtrFilter {
	res := make([]Uint8PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Uint8PtrValFilterToUint8PtrFilter(filter))
	}
	return res
}

func Uint8PtrIf(f BoolFunc, a, b *uint8) *uint8 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *uint16  @aliasPreix: Uint16Ptr

// ===  gen area =====
type Uint16PtrFilter func(int, *uint16) bool
type Uint16PtrMapper func(int, *uint16) *uint16
type Uint16PtrRanger func(int, *uint16)

func (filter Uint16PtrFilter) Something() Uint16PtrFilter {
	if filter == nil {
		return func(int, *uint16) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint16PtrMapper) Something() Uint16PtrMapper {
	if mapper == nil {
		return func(i int, val *uint16) *uint16 {
			return val
		}
	}
	return mapper
}

func (ranger Uint16PtrRanger) Something() Uint16PtrRanger {
	if ranger == nil {
		return func(int, *uint16) {}
	}
	return ranger
}

type Uint16PtrValFilter func(*uint16) bool
type Uint16PtrValMapper func(*uint16) *uint16
type Uint16PtrValRanger func(*uint16)

func (filter Uint16PtrValFilter) Something() Uint16PtrValFilter {
	if filter == nil {
		return func(*uint16) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint16PtrValMapper) Something() Uint16PtrValMapper {
	if mapper == nil {
		return func(val *uint16) *uint16 {
			return val
		}
	}
	return mapper
}

func (ranger Uint16PtrValRanger) Something() Uint16PtrValRanger {
	if ranger == nil {
		return func(*uint16) {}
	}
	return ranger
}

// ==== util func

func Uint16PtrValRangerToUint16PtrRanger(ranger Uint16PtrValRanger) Uint16PtrRanger {
	return func(i int, val *uint16) {
		ranger.Something()(val)
	}
}

func Uint16PtrValRangerBatchToUint16PtrRanger(rangers ...Uint16PtrValRanger) []Uint16PtrRanger {
	res := make([]Uint16PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Uint16PtrValRangerToUint16PtrRanger(ranger))
	}
	return res
}

func Uint16PtrValMapperToUint16PtrMapper(mapper Uint16PtrValMapper) Uint16PtrMapper {
	return func(i int, val *uint16) *uint16 {
		return mapper.Something()(val)
	}
}

func Uint16PtrValMapperBatchToUint16PtrMapper(mappers ...Uint16PtrValMapper) []Uint16PtrMapper {
	res := make([]Uint16PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Uint16PtrValMapperToUint16PtrMapper(mapper))
	}
	return res
}

func Uint16PtrValFilterToUint16PtrFilter(filter Uint16PtrValFilter) Uint16PtrFilter {
	return func(i int, val *uint16) bool {
		return filter.Something()(val)
	}
}

func Uint16PtrValFilterBatchToUint16PtrFilter(filters ...Uint16PtrValFilter) []Uint16PtrFilter {
	res := make([]Uint16PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Uint16PtrValFilterToUint16PtrFilter(filter))
	}
	return res
}

func Uint16PtrIf(f BoolFunc, a, b *uint16) *uint16 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *uint32  @aliasPreix: Uint32Ptr

// ===  gen area =====
type Uint32PtrFilter func(int, *uint32) bool
type Uint32PtrMapper func(int, *uint32) *uint32
type Uint32PtrRanger func(int, *uint32)

func (filter Uint32PtrFilter) Something() Uint32PtrFilter {
	if filter == nil {
		return func(int, *uint32) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint32PtrMapper) Something() Uint32PtrMapper {
	if mapper == nil {
		return func(i int, val *uint32) *uint32 {
			return val
		}
	}
	return mapper
}

func (ranger Uint32PtrRanger) Something() Uint32PtrRanger {
	if ranger == nil {
		return func(int, *uint32) {}
	}
	return ranger
}

type Uint32PtrValFilter func(*uint32) bool
type Uint32PtrValMapper func(*uint32) *uint32
type Uint32PtrValRanger func(*uint32)

func (filter Uint32PtrValFilter) Something() Uint32PtrValFilter {
	if filter == nil {
		return func(*uint32) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint32PtrValMapper) Something() Uint32PtrValMapper {
	if mapper == nil {
		return func(val *uint32) *uint32 {
			return val
		}
	}
	return mapper
}

func (ranger Uint32PtrValRanger) Something() Uint32PtrValRanger {
	if ranger == nil {
		return func(*uint32) {}
	}
	return ranger
}

// ==== util func

func Uint32PtrValRangerToUint32PtrRanger(ranger Uint32PtrValRanger) Uint32PtrRanger {
	return func(i int, val *uint32) {
		ranger.Something()(val)
	}
}

func Uint32PtrValRangerBatchToUint32PtrRanger(rangers ...Uint32PtrValRanger) []Uint32PtrRanger {
	res := make([]Uint32PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Uint32PtrValRangerToUint32PtrRanger(ranger))
	}
	return res
}

func Uint32PtrValMapperToUint32PtrMapper(mapper Uint32PtrValMapper) Uint32PtrMapper {
	return func(i int, val *uint32) *uint32 {
		return mapper.Something()(val)
	}
}

func Uint32PtrValMapperBatchToUint32PtrMapper(mappers ...Uint32PtrValMapper) []Uint32PtrMapper {
	res := make([]Uint32PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Uint32PtrValMapperToUint32PtrMapper(mapper))
	}
	return res
}

func Uint32PtrValFilterToUint32PtrFilter(filter Uint32PtrValFilter) Uint32PtrFilter {
	return func(i int, val *uint32) bool {
		return filter.Something()(val)
	}
}

func Uint32PtrValFilterBatchToUint32PtrFilter(filters ...Uint32PtrValFilter) []Uint32PtrFilter {
	res := make([]Uint32PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Uint32PtrValFilterToUint32PtrFilter(filter))
	}
	return res
}

func Uint32PtrIf(f BoolFunc, a, b *uint32) *uint32 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *uint64  @aliasPreix: Uint64Ptr

// ===  gen area =====
type Uint64PtrFilter func(int, *uint64) bool
type Uint64PtrMapper func(int, *uint64) *uint64
type Uint64PtrRanger func(int, *uint64)

func (filter Uint64PtrFilter) Something() Uint64PtrFilter {
	if filter == nil {
		return func(int, *uint64) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint64PtrMapper) Something() Uint64PtrMapper {
	if mapper == nil {
		return func(i int, val *uint64) *uint64 {
			return val
		}
	}
	return mapper
}

func (ranger Uint64PtrRanger) Something() Uint64PtrRanger {
	if ranger == nil {
		return func(int, *uint64) {}
	}
	return ranger
}

type Uint64PtrValFilter func(*uint64) bool
type Uint64PtrValMapper func(*uint64) *uint64
type Uint64PtrValRanger func(*uint64)

func (filter Uint64PtrValFilter) Something() Uint64PtrValFilter {
	if filter == nil {
		return func(*uint64) bool {
			return false
		}
	}
	return filter
}

func (mapper Uint64PtrValMapper) Something() Uint64PtrValMapper {
	if mapper == nil {
		return func(val *uint64) *uint64 {
			return val
		}
	}
	return mapper
}

func (ranger Uint64PtrValRanger) Something() Uint64PtrValRanger {
	if ranger == nil {
		return func(*uint64) {}
	}
	return ranger
}

// ==== util func

func Uint64PtrValRangerToUint64PtrRanger(ranger Uint64PtrValRanger) Uint64PtrRanger {
	return func(i int, val *uint64) {
		ranger.Something()(val)
	}
}

func Uint64PtrValRangerBatchToUint64PtrRanger(rangers ...Uint64PtrValRanger) []Uint64PtrRanger {
	res := make([]Uint64PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Uint64PtrValRangerToUint64PtrRanger(ranger))
	}
	return res
}

func Uint64PtrValMapperToUint64PtrMapper(mapper Uint64PtrValMapper) Uint64PtrMapper {
	return func(i int, val *uint64) *uint64 {
		return mapper.Something()(val)
	}
}

func Uint64PtrValMapperBatchToUint64PtrMapper(mappers ...Uint64PtrValMapper) []Uint64PtrMapper {
	res := make([]Uint64PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Uint64PtrValMapperToUint64PtrMapper(mapper))
	}
	return res
}

func Uint64PtrValFilterToUint64PtrFilter(filter Uint64PtrValFilter) Uint64PtrFilter {
	return func(i int, val *uint64) bool {
		return filter.Something()(val)
	}
}

func Uint64PtrValFilterBatchToUint64PtrFilter(filters ...Uint64PtrValFilter) []Uint64PtrFilter {
	res := make([]Uint64PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Uint64PtrValFilterToUint64PtrFilter(filter))
	}
	return res
}

func Uint64PtrIf(f BoolFunc, a, b *uint64) *uint64 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *float32  @aliasPreix: Float32Ptr

// ===  gen area =====
type Float32PtrFilter func(int, *float32) bool
type Float32PtrMapper func(int, *float32) *float32
type Float32PtrRanger func(int, *float32)

func (filter Float32PtrFilter) Something() Float32PtrFilter {
	if filter == nil {
		return func(int, *float32) bool {
			return false
		}
	}
	return filter
}

func (mapper Float32PtrMapper) Something() Float32PtrMapper {
	if mapper == nil {
		return func(i int, val *float32) *float32 {
			return val
		}
	}
	return mapper
}

func (ranger Float32PtrRanger) Something() Float32PtrRanger {
	if ranger == nil {
		return func(int, *float32) {}
	}
	return ranger
}

type Float32PtrValFilter func(*float32) bool
type Float32PtrValMapper func(*float32) *float32
type Float32PtrValRanger func(*float32)

func (filter Float32PtrValFilter) Something() Float32PtrValFilter {
	if filter == nil {
		return func(*float32) bool {
			return false
		}
	}
	return filter
}

func (mapper Float32PtrValMapper) Something() Float32PtrValMapper {
	if mapper == nil {
		return func(val *float32) *float32 {
			return val
		}
	}
	return mapper
}

func (ranger Float32PtrValRanger) Something() Float32PtrValRanger {
	if ranger == nil {
		return func(*float32) {}
	}
	return ranger
}

// ==== util func

func Float32PtrValRangerToFloat32PtrRanger(ranger Float32PtrValRanger) Float32PtrRanger {
	return func(i int, val *float32) {
		ranger.Something()(val)
	}
}

func Float32PtrValRangerBatchToFloat32PtrRanger(rangers ...Float32PtrValRanger) []Float32PtrRanger {
	res := make([]Float32PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Float32PtrValRangerToFloat32PtrRanger(ranger))
	}
	return res
}

func Float32PtrValMapperToFloat32PtrMapper(mapper Float32PtrValMapper) Float32PtrMapper {
	return func(i int, val *float32) *float32 {
		return mapper.Something()(val)
	}
}

func Float32PtrValMapperBatchToFloat32PtrMapper(mappers ...Float32PtrValMapper) []Float32PtrMapper {
	res := make([]Float32PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Float32PtrValMapperToFloat32PtrMapper(mapper))
	}
	return res
}

func Float32PtrValFilterToFloat32PtrFilter(filter Float32PtrValFilter) Float32PtrFilter {
	return func(i int, val *float32) bool {
		return filter.Something()(val)
	}
}

func Float32PtrValFilterBatchToFloat32PtrFilter(filters ...Float32PtrValFilter) []Float32PtrFilter {
	res := make([]Float32PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Float32PtrValFilterToFloat32PtrFilter(filter))
	}
	return res
}

func Float32PtrIf(f BoolFunc, a, b *float32) *float32 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *float64  @aliasPreix: Float64Ptr

// ===  gen area =====
type Float64PtrFilter func(int, *float64) bool
type Float64PtrMapper func(int, *float64) *float64
type Float64PtrRanger func(int, *float64)

func (filter Float64PtrFilter) Something() Float64PtrFilter {
	if filter == nil {
		return func(int, *float64) bool {
			return false
		}
	}
	return filter
}

func (mapper Float64PtrMapper) Something() Float64PtrMapper {
	if mapper == nil {
		return func(i int, val *float64) *float64 {
			return val
		}
	}
	return mapper
}

func (ranger Float64PtrRanger) Something() Float64PtrRanger {
	if ranger == nil {
		return func(int, *float64) {}
	}
	return ranger
}

type Float64PtrValFilter func(*float64) bool
type Float64PtrValMapper func(*float64) *float64
type Float64PtrValRanger func(*float64)

func (filter Float64PtrValFilter) Something() Float64PtrValFilter {
	if filter == nil {
		return func(*float64) bool {
			return false
		}
	}
	return filter
}

func (mapper Float64PtrValMapper) Something() Float64PtrValMapper {
	if mapper == nil {
		return func(val *float64) *float64 {
			return val
		}
	}
	return mapper
}

func (ranger Float64PtrValRanger) Something() Float64PtrValRanger {
	if ranger == nil {
		return func(*float64) {}
	}
	return ranger
}

// ==== util func

func Float64PtrValRangerToFloat64PtrRanger(ranger Float64PtrValRanger) Float64PtrRanger {
	return func(i int, val *float64) {
		ranger.Something()(val)
	}
}

func Float64PtrValRangerBatchToFloat64PtrRanger(rangers ...Float64PtrValRanger) []Float64PtrRanger {
	res := make([]Float64PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Float64PtrValRangerToFloat64PtrRanger(ranger))
	}
	return res
}

func Float64PtrValMapperToFloat64PtrMapper(mapper Float64PtrValMapper) Float64PtrMapper {
	return func(i int, val *float64) *float64 {
		return mapper.Something()(val)
	}
}

func Float64PtrValMapperBatchToFloat64PtrMapper(mappers ...Float64PtrValMapper) []Float64PtrMapper {
	res := make([]Float64PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Float64PtrValMapperToFloat64PtrMapper(mapper))
	}
	return res
}

func Float64PtrValFilterToFloat64PtrFilter(filter Float64PtrValFilter) Float64PtrFilter {
	return func(i int, val *float64) bool {
		return filter.Something()(val)
	}
}

func Float64PtrValFilterBatchToFloat64PtrFilter(filters ...Float64PtrValFilter) []Float64PtrFilter {
	res := make([]Float64PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Float64PtrValFilterToFloat64PtrFilter(filter))
	}
	return res
}

func Float64PtrIf(f BoolFunc, a, b *float64) *float64 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *bool  @aliasPreix: BoolPtr

// ===  gen area =====
type BoolPtrFilter func(int, *bool) bool
type BoolPtrMapper func(int, *bool) *bool
type BoolPtrRanger func(int, *bool)

func (filter BoolPtrFilter) Something() BoolPtrFilter {
	if filter == nil {
		return func(int, *bool) bool {
			return false
		}
	}
	return filter
}

func (mapper BoolPtrMapper) Something() BoolPtrMapper {
	if mapper == nil {
		return func(i int, val *bool) *bool {
			return val
		}
	}
	return mapper
}

func (ranger BoolPtrRanger) Something() BoolPtrRanger {
	if ranger == nil {
		return func(int, *bool) {}
	}
	return ranger
}

type BoolPtrValFilter func(*bool) bool
type BoolPtrValMapper func(*bool) *bool
type BoolPtrValRanger func(*bool)

func (filter BoolPtrValFilter) Something() BoolPtrValFilter {
	if filter == nil {
		return func(*bool) bool {
			return false
		}
	}
	return filter
}

func (mapper BoolPtrValMapper) Something() BoolPtrValMapper {
	if mapper == nil {
		return func(val *bool) *bool {
			return val
		}
	}
	return mapper
}

func (ranger BoolPtrValRanger) Something() BoolPtrValRanger {
	if ranger == nil {
		return func(*bool) {}
	}
	return ranger
}

// ==== util func

func BoolPtrValRangerToBoolPtrRanger(ranger BoolPtrValRanger) BoolPtrRanger {
	return func(i int, val *bool) {
		ranger.Something()(val)
	}
}

func BoolPtrValRangerBatchToBoolPtrRanger(rangers ...BoolPtrValRanger) []BoolPtrRanger {
	res := make([]BoolPtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, BoolPtrValRangerToBoolPtrRanger(ranger))
	}
	return res
}

func BoolPtrValMapperToBoolPtrMapper(mapper BoolPtrValMapper) BoolPtrMapper {
	return func(i int, val *bool) *bool {
		return mapper.Something()(val)
	}
}

func BoolPtrValMapperBatchToBoolPtrMapper(mappers ...BoolPtrValMapper) []BoolPtrMapper {
	res := make([]BoolPtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, BoolPtrValMapperToBoolPtrMapper(mapper))
	}
	return res
}

func BoolPtrValFilterToBoolPtrFilter(filter BoolPtrValFilter) BoolPtrFilter {
	return func(i int, val *bool) bool {
		return filter.Something()(val)
	}
}

func BoolPtrValFilterBatchToBoolPtrFilter(filters ...BoolPtrValFilter) []BoolPtrFilter {
	res := make([]BoolPtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, BoolPtrValFilterToBoolPtrFilter(filter))
	}
	return res
}

func BoolPtrIf(f BoolFunc, a, b *bool) *bool {
	if f() {
		return a
	}
	return b
}

// autogen @type: *rune  @aliasPreix: RunePtr

// ===  gen area =====
type RunePtrFilter func(int, *rune) bool
type RunePtrMapper func(int, *rune) *rune
type RunePtrRanger func(int, *rune)

func (filter RunePtrFilter) Something() RunePtrFilter {
	if filter == nil {
		return func(int, *rune) bool {
			return false
		}
	}
	return filter
}

func (mapper RunePtrMapper) Something() RunePtrMapper {
	if mapper == nil {
		return func(i int, val *rune) *rune {
			return val
		}
	}
	return mapper
}

func (ranger RunePtrRanger) Something() RunePtrRanger {
	if ranger == nil {
		return func(int, *rune) {}
	}
	return ranger
}

type RunePtrValFilter func(*rune) bool
type RunePtrValMapper func(*rune) *rune
type RunePtrValRanger func(*rune)

func (filter RunePtrValFilter) Something() RunePtrValFilter {
	if filter == nil {
		return func(*rune) bool {
			return false
		}
	}
	return filter
}

func (mapper RunePtrValMapper) Something() RunePtrValMapper {
	if mapper == nil {
		return func(val *rune) *rune {
			return val
		}
	}
	return mapper
}

func (ranger RunePtrValRanger) Something() RunePtrValRanger {
	if ranger == nil {
		return func(*rune) {}
	}
	return ranger
}

// ==== util func

func RunePtrValRangerToRunePtrRanger(ranger RunePtrValRanger) RunePtrRanger {
	return func(i int, val *rune) {
		ranger.Something()(val)
	}
}

func RunePtrValRangerBatchToRunePtrRanger(rangers ...RunePtrValRanger) []RunePtrRanger {
	res := make([]RunePtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, RunePtrValRangerToRunePtrRanger(ranger))
	}
	return res
}

func RunePtrValMapperToRunePtrMapper(mapper RunePtrValMapper) RunePtrMapper {
	return func(i int, val *rune) *rune {
		return mapper.Something()(val)
	}
}

func RunePtrValMapperBatchToRunePtrMapper(mappers ...RunePtrValMapper) []RunePtrMapper {
	res := make([]RunePtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, RunePtrValMapperToRunePtrMapper(mapper))
	}
	return res
}

func RunePtrValFilterToRunePtrFilter(filter RunePtrValFilter) RunePtrFilter {
	return func(i int, val *rune) bool {
		return filter.Something()(val)
	}
}

func RunePtrValFilterBatchToRunePtrFilter(filters ...RunePtrValFilter) []RunePtrFilter {
	res := make([]RunePtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, RunePtrValFilterToRunePtrFilter(filter))
	}
	return res
}

func RunePtrIf(f BoolFunc, a, b *rune) *rune {
	if f() {
		return a
	}
	return b
}

// autogen @type: *complex64  @aliasPreix: Complex64Ptr

// ===  gen area =====
type Complex64PtrFilter func(int, *complex64) bool
type Complex64PtrMapper func(int, *complex64) *complex64
type Complex64PtrRanger func(int, *complex64)

func (filter Complex64PtrFilter) Something() Complex64PtrFilter {
	if filter == nil {
		return func(int, *complex64) bool {
			return false
		}
	}
	return filter
}

func (mapper Complex64PtrMapper) Something() Complex64PtrMapper {
	if mapper == nil {
		return func(i int, val *complex64) *complex64 {
			return val
		}
	}
	return mapper
}

func (ranger Complex64PtrRanger) Something() Complex64PtrRanger {
	if ranger == nil {
		return func(int, *complex64) {}
	}
	return ranger
}

type Complex64PtrValFilter func(*complex64) bool
type Complex64PtrValMapper func(*complex64) *complex64
type Complex64PtrValRanger func(*complex64)

func (filter Complex64PtrValFilter) Something() Complex64PtrValFilter {
	if filter == nil {
		return func(*complex64) bool {
			return false
		}
	}
	return filter
}

func (mapper Complex64PtrValMapper) Something() Complex64PtrValMapper {
	if mapper == nil {
		return func(val *complex64) *complex64 {
			return val
		}
	}
	return mapper
}

func (ranger Complex64PtrValRanger) Something() Complex64PtrValRanger {
	if ranger == nil {
		return func(*complex64) {}
	}
	return ranger
}

// ==== util func

func Complex64PtrValRangerToComplex64PtrRanger(ranger Complex64PtrValRanger) Complex64PtrRanger {
	return func(i int, val *complex64) {
		ranger.Something()(val)
	}
}

func Complex64PtrValRangerBatchToComplex64PtrRanger(rangers ...Complex64PtrValRanger) []Complex64PtrRanger {
	res := make([]Complex64PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Complex64PtrValRangerToComplex64PtrRanger(ranger))
	}
	return res
}

func Complex64PtrValMapperToComplex64PtrMapper(mapper Complex64PtrValMapper) Complex64PtrMapper {
	return func(i int, val *complex64) *complex64 {
		return mapper.Something()(val)
	}
}

func Complex64PtrValMapperBatchToComplex64PtrMapper(mappers ...Complex64PtrValMapper) []Complex64PtrMapper {
	res := make([]Complex64PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Complex64PtrValMapperToComplex64PtrMapper(mapper))
	}
	return res
}

func Complex64PtrValFilterToComplex64PtrFilter(filter Complex64PtrValFilter) Complex64PtrFilter {
	return func(i int, val *complex64) bool {
		return filter.Something()(val)
	}
}

func Complex64PtrValFilterBatchToComplex64PtrFilter(filters ...Complex64PtrValFilter) []Complex64PtrFilter {
	res := make([]Complex64PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Complex64PtrValFilterToComplex64PtrFilter(filter))
	}
	return res
}

func Complex64PtrIf(f BoolFunc, a, b *complex64) *complex64 {
	if f() {
		return a
	}
	return b
}

// autogen @type: *complex32  @aliasPreix: Complex32Ptr

// ===  gen area =====
type Complex32PtrFilter func(int, *complex32) bool
type Complex32PtrMapper func(int, *complex32) *complex32
type Complex32PtrRanger func(int, *complex32)

func (filter Complex32PtrFilter) Something() Complex32PtrFilter {
	if filter == nil {
		return func(int, *complex32) bool {
			return false
		}
	}
	return filter
}

func (mapper Complex32PtrMapper) Something() Complex32PtrMapper {
	if mapper == nil {
		return func(i int, val *complex32) *complex32 {
			return val
		}
	}
	return mapper
}

func (ranger Complex32PtrRanger) Something() Complex32PtrRanger {
	if ranger == nil {
		return func(int, *complex32) {}
	}
	return ranger
}

type Complex32PtrValFilter func(*complex32) bool
type Complex32PtrValMapper func(*complex32) *complex32
type Complex32PtrValRanger func(*complex32)

func (filter Complex32PtrValFilter) Something() Complex32PtrValFilter {
	if filter == nil {
		return func(*complex32) bool {
			return false
		}
	}
	return filter
}

func (mapper Complex32PtrValMapper) Something() Complex32PtrValMapper {
	if mapper == nil {
		return func(val *complex32) *complex32 {
			return val
		}
	}
	return mapper
}

func (ranger Complex32PtrValRanger) Something() Complex32PtrValRanger {
	if ranger == nil {
		return func(*complex32) {}
	}
	return ranger
}

// ==== util func

func Complex32PtrValRangerToComplex32PtrRanger(ranger Complex32PtrValRanger) Complex32PtrRanger {
	return func(i int, val *complex32) {
		ranger.Something()(val)
	}
}

func Complex32PtrValRangerBatchToComplex32PtrRanger(rangers ...Complex32PtrValRanger) []Complex32PtrRanger {
	res := make([]Complex32PtrRanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, Complex32PtrValRangerToComplex32PtrRanger(ranger))
	}
	return res
}

func Complex32PtrValMapperToComplex32PtrMapper(mapper Complex32PtrValMapper) Complex32PtrMapper {
	return func(i int, val *complex32) *complex32 {
		return mapper.Something()(val)
	}
}

func Complex32PtrValMapperBatchToComplex32PtrMapper(mappers ...Complex32PtrValMapper) []Complex32PtrMapper {
	res := make([]Complex32PtrMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, Complex32PtrValMapperToComplex32PtrMapper(mapper))
	}
	return res
}

func Complex32PtrValFilterToComplex32PtrFilter(filter Complex32PtrValFilter) Complex32PtrFilter {
	return func(i int, val *complex32) bool {
		return filter.Something()(val)
	}
}

func Complex32PtrValFilterBatchToComplex32PtrFilter(filters ...Complex32PtrValFilter) []Complex32PtrFilter {
	res := make([]Complex32PtrFilter, len(filters))
	for _, filter := range filters {
		res = append(res, Complex32PtrValFilterToComplex32PtrFilter(filter))
	}
	return res
}

func Complex32PtrIf(f BoolFunc, a, b *complex32) *complex32 {
	if f() {
		return a
	}
	return b
}
