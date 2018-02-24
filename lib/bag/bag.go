package bag

import (
	"github.com/xxx/pkg"
)

// auto generate by bag

// generate slice

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

// autogen @type: pkg.StructA  @aliasPreix: PkgStructA

type PkgStructASlice struct {
	raw []pkg.StructA
}

func MakePkgStructASlice(param ...int) *PkgStructASlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &PkgStructASlice{
		raw: make([]pkg.StructA, p1, p2),
	}
}

func MakePkgStructASliceFromRaw(raw []pkg.StructA) *PkgStructASlice {
	if raw == nil {
		raw = make([]pkg.StructA, 0)
	}
	return &PkgStructASlice{
		raw: raw,
	}
}

func (this *PkgStructASlice) Append(vals ...pkg.StructA) {
	this.raw = append(this.raw, vals...)
}

func (this *PkgStructASlice) AppendSlice(slice *PkgStructASlice) {
	slice.RangeVal(func(val pkg.StructA) {
		this.Append(val)
	})
}

func (this *PkgStructASlice) AppendIf(filter PkgStructAFilter, vals ...pkg.StructA) {
	this.AppendSlice(MakePkgStructASliceFromRaw(vals).Filter(filter))
}

func (this *PkgStructASlice) Range(rangers ...PkgStructARanger) {
	for _, ranger := range rangers {
		for i, val := range this.GetRaw() {
			ranger.Something()(i, val)
		}
	}
}

func (this *PkgStructASlice) RangeVal(rangers ...PkgStructAValRanger) {
	this.Range(PkgStructAValRangerBatchToPkgStructARanger(rangers...)...)
}

func (this *PkgStructASlice) GetRaw() []pkg.StructA {
	return this.raw
}

func (this *PkgStructASlice) Filter(filters ...PkgStructAFilter) *PkgStructASlice {

	res := MakePkgStructASlice(0, len(this.GetRaw()))

	this.Range(func(i int, val pkg.StructA) {
		for _, f := range filters {
			if f.Something()(i, val) {
				return
			}
		}
		res.Append(val)
	})

	return res
}

func (this *PkgStructASlice) FilterVal(filters ...PkgStructAValFilter) *PkgStructASlice {
	return this.Filter(PkgStructAValFilterBatchToPkgStructAFilter(filters...)...)
}

func (this *PkgStructASlice) Map(mappers ...PkgStructAMapper) *PkgStructASlice {

	res := MakePkgStructASlice(0, len(this.GetRaw()))

	this.Range(func(i int, val pkg.StructA) {
		for _, f := range mappers {
			val = f.Something()(i, val)
		}
		res.Append(val)
	})

	return res
}

func (this *PkgStructASlice) MapVal(mappers ...PkgStructAValMapper) *PkgStructASlice {
	return this.Map(PkgStructAValMapperBatchToPkgStructAMapper(mappers...)...)
}

// let it crash apis
func (this *PkgStructASlice) First() pkg.StructA {
	return this.At(0)
}

func (this *PkgStructASlice) Last() pkg.StructA {
	return this.At(this.Len() - 1)
}

func (this *PkgStructASlice) Ats(is ...int) *PkgStructASlice {
	res := MakePkgStructASlice(0, len(is))
	for _, i := range is {
		res.Append(this.At(i))
	}
	return res
}

func (this *PkgStructASlice) Slice(left, right int) *PkgStructASlice {
	return MakePkgStructASliceFromRaw(this.GetRaw()[left:right])
}

func (this *PkgStructASlice) At(i int) pkg.StructA {
	return this.GetRaw()[i]
}

func (this *PkgStructASlice) Len() int {
	return len(this.GetRaw())
}

func (this *PkgStructASlice) Cap() int {
	return cap(this.GetRaw())
}

// generate sliceFunc

// autogen @type: int  @aliasPreix: Int

// common area
type BoolFunc func() bool

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

// autogen @type: string  @aliasPreix: String

// common area
type BoolFunc func() bool

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

// autogen @type: pkg.StructA  @aliasPreix: PkgStructA

// common area
type BoolFunc func() bool

// ===  gen area =====
type PkgStructAFilter func(int, pkg.StructA) bool
type PkgStructAMapper func(int, pkg.StructA) pkg.StructA
type PkgStructARanger func(int, pkg.StructA)

func (filter PkgStructAFilter) Something() PkgStructAFilter {
	if filter == nil {
		return func(int, pkg.StructA) bool {
			return false
		}
	}
	return filter
}

func (mapper PkgStructAMapper) Something() PkgStructAMapper {
	if mapper == nil {
		return func(i int, val pkg.StructA) pkg.StructA {
			return val
		}
	}
	return mapper
}

func (ranger PkgStructARanger) Something() PkgStructARanger {
	if ranger == nil {
		return func(int, pkg.StructA) {}
	}
	return ranger
}

type PkgStructAValFilter func(pkg.StructA) bool
type PkgStructAValMapper func(pkg.StructA) pkg.StructA
type PkgStructAValRanger func(pkg.StructA)

func (filter PkgStructAValFilter) Something() PkgStructAValFilter {
	if filter == nil {
		return func(pkg.StructA) bool {
			return false
		}
	}
	return filter
}

func (mapper PkgStructAValMapper) Something() PkgStructAValMapper {
	if mapper == nil {
		return func(val pkg.StructA) pkg.StructA {
			return val
		}
	}
	return mapper
}

func (ranger PkgStructAValRanger) Something() PkgStructAValRanger {
	if ranger == nil {
		return func(pkg.StructA) {}
	}
	return ranger
}

// ==== util func

func PkgStructAValRangerToPkgStructARanger(ranger PkgStructAValRanger) PkgStructARanger {
	return func(i int, val pkg.StructA) {
		ranger.Something()(val)
	}
}

func PkgStructAValRangerBatchToPkgStructARanger(rangers ...PkgStructAValRanger) []PkgStructARanger {
	res := make([]PkgStructARanger, len(rangers))
	for _, ranger := range rangers {
		res = append(res, PkgStructAValRangerToPkgStructARanger(ranger))
	}
	return res
}

func PkgStructAValMapperToPkgStructAMapper(mapper PkgStructAValMapper) PkgStructAMapper {
	return func(i int, val pkg.StructA) pkg.StructA {
		return mapper.Something()(val)
	}
}

func PkgStructAValMapperBatchToPkgStructAMapper(mappers ...PkgStructAValMapper) []PkgStructAMapper {
	res := make([]PkgStructAMapper, len(mappers))
	for _, mapper := range mappers {
		res = append(res, PkgStructAValMapperToPkgStructAMapper(mapper))
	}
	return res
}

func PkgStructAValFilterToPkgStructAFilter(filter PkgStructAValFilter) PkgStructAFilter {
	return func(i int, val pkg.StructA) bool {
		return filter.Something()(val)
	}
}

func PkgStructAValFilterBatchToPkgStructAFilter(filters ...PkgStructAValFilter) []PkgStructAFilter {
	res := make([]PkgStructAFilter, len(filters))
	for _, filter := range filters {
		res = append(res, PkgStructAValFilterToPkgStructAFilter(filter))
	}
	return res
}

func PkgStructAIf(f BoolFunc, a, b pkg.StructA) pkg.StructA {
	if f() {
		return a
	}
	return b
}
