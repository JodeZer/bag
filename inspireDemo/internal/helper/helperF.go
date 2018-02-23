package helper

type StringFilter func(string) bool
type StringMender func(string) string
type StringRanger func(string)

func (filter StringFilter) Something() StringFilter {
	if filter == nil {
		return func(string) bool {
			return false
		}
	}
	return filter
}

func (mender StringMender) Something() StringMender {
	if mender == nil {
		return func(str string) string {
			return str
		}
	}
	return mender
}

func (ranger StringRanger) Something() StringRanger {
	if ranger == nil {
		return func(string) {}
	}
	return ranger
}