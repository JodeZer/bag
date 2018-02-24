package helper

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
