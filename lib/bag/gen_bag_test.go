package bag

import "testing"

func deepEqual(lhs, rhs []string) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, val := range lhs {
		if val != rhs[i] {
			return false
		}
	}
	return true
}

func TestStringSliceAppend(t *testing.T) {
	input := []string{"a", "b"}
	expect := []string{"a", "b", "c"}
	testStringSliceAppend(t, MakeStringSliceFromRaw(input), expect, "c")
}

func TestStringSliceAppendIf(t *testing.T) {
	input := []string{"a", "b"}
	expect := []string{"a", "b", "c"}
	testStringSliceAppendIf(t, MakeStringSliceFromRaw(input), expect, func(i int, val string) bool {
		return val == ""
	}, "c", "", "")
}

func TestStringSliceAppendSlice(t *testing.T) {
	input := []string{"a", "b"}
	appended := []string{"c", "d"}
	expect := []string{"a", "b", "c", "d"}
	testStringSliceAppendSlice(t, MakeStringSliceFromRaw(input), expect, MakeStringSliceFromRaw(appended))
}

func TestStringSliceAt(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	testStringSliceAt(t, MakeStringSliceFromRaw(input), 2, "c")
}

func TestStringSliceAts(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	expect := []string{"a", "d"}
	testStringSliceAts(t, MakeStringSliceFromRaw(input), expect, 0, 3)
}

func TestStringSliceCap(t *testing.T) {
	expect := 100
	input := make([]string, 0, expect)
	if expect != MakeStringSliceFromRaw(input).Cap() {
		t.Fatalf("fail")
	}
}

func TestStringSliceFilter(t *testing.T) {
	input := []string{"a", "b", "c", "d", "", "", "!", "!"}
	expect := []string{"a", "b", "c", "d"}
	testStringFilter(t, MakeStringSliceFromRaw(input), expect,
		func(i int, val string) bool { return val == "" },
		func(i int, val string) bool { return val == "!" })
}

func TestStringSliceFilterVal(t *testing.T) {
	input := []string{"a", "b", "c", "d", "", "", "!", "!"}
	expect := []string{"a", "b", "c", "d"}
	testStringFilterVal(t, MakeStringSliceFromRaw(input), expect,
		func(val string) bool { return val == "" },
		func(val string) bool { return val == "!" })
}

func TestStringSliceFirst(t *testing.T) {
	input := []string{"a", "b", "c", "d", "", "", "!", "!"}
	expect := "a"
	testStringSliceFirst(t, MakeStringSliceFromRaw(input), expect)
}

func TestStringSliceGetRaw(t *testing.T) {

}

func TestStringSliceLast(t *testing.T) {

}

func TestStringSliceLen(t *testing.T) {

}

func TestStringSliceMap(t *testing.T) {
	input := []string{"a", "b", "c", "d", "", "", "!", "!"}
	expect := []string{"aaa", "aba", "aca", "ada", "aa", "aa", "a!a", "a!a"}
	testStringSliceMap(t, MakeStringSliceFromRaw(input), expect,
		func(i int, val string) string { return "a" + val },
		func(i int, val string) string { return val + "a" })
}

func TestStringSliceMapVal(t *testing.T) {

}

func TestStringSliceRange(t *testing.T) {

}

func TestStringSliceRangeVal(t *testing.T) {

}

func TestStringSliceSlice(t *testing.T) {
	input := []string{"a", "b", "c", "d", "", "", "!", "!"}
	expect := []string{"a", "b", "c", "d"}
	testStringSliceSlice(t, MakeStringSliceFromRaw(input), expect, 0, 4)
}

// internal test
func testStringSliceAppendIf(t *testing.T, input *StringSlice, expect []string, f StringFilter, append ...string) {
	input.AppendIf(f, append...)
	testCheck(t, input.GetRaw(), expect)
}

func testStringSliceAppend(t *testing.T, input *StringSlice, expect []string, append ...string) {
	input.Append(append...)
	testCheck(t, input.GetRaw(), expect)
}

func testStringSliceAppendSlice(t *testing.T, input *StringSlice, expect []string, append *StringSlice) {
	input.AppendSlice(append)
	testCheck(t, input.GetRaw(), expect)
}

func testStringSliceAts(t *testing.T, input *StringSlice, expect []string, ats ...int) {
	testCheck(t, input.Ats(ats...).GetRaw(), expect)
}

func testStringSliceAt(t *testing.T, input *StringSlice, at int, expect string) {
	testCheck(t, []string{input.At(at)}, []string{expect})
}

func testStringFilter(t *testing.T, input *StringSlice, expect []string, filter ...StringFilter) {
	input = input.Filter(filter...)
	testCheck(t, input.GetRaw(), expect)
}

func testStringFilterVal(t *testing.T, input *StringSlice, expect []string, filter ...StringValFilter) {
	input = input.FilterVal(filter...)
	testCheck(t, input.GetRaw(), expect)
}

func testStringSliceFirst(t *testing.T, input *StringSlice, expect string) {
	testCheck(t, []string{input.First()}, []string{expect})
}

func testStringSliceMap(t *testing.T, input *StringSlice, expect []string, mapper ...StringMapper) {
	input = input.Map(mapper...)
	testCheck(t, input.GetRaw(), expect)
}

func testStringSliceSlice(t *testing.T, input *StringSlice, expect []string, left, right int) {
	input = input.Slice(left, right)
	testCheck(t, input.GetRaw(), expect)
}

func testCheck(t *testing.T, input []string, expect []string) {
	if !deepEqual(input, expect) {
		t.Logf("\n%+v\n%+v", input, expect)
		t.Fatalf("fail")
	}
}
