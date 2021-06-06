package prettier

import (
	"strconv"
	"strings"
	"testing"
)

func assertPretty(t *testing.T, doc Doc, maxLineWidth int, expected string) {
	var builder strings.Builder
	Prettier(&builder, doc, maxLineWidth, "    ")
	actual := builder.String()

	if actual != expected {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestJSON(t *testing.T) {

	array := func(elements []Doc) Doc {
		return WrapBrackets(
			Join(
				Concat{
					Text(","),
					Line{},
				},
				elements...,
			),
			SoftLine{},
		)
	}

	const innerElementCount = 5
	innerArrayElements := make([]Doc, innerElementCount)
	for i := 0; i < innerElementCount; i++ {
		innerArrayElements[i] = Text(strconv.Itoa(i + 1))
	}
	innerArray := array(innerArrayElements)
	outerArray := array([]Doc{innerArray, innerArray, innerArray})

	assertPretty(t, outerArray, 80, `[[1, 2, 3, 4, 5], [1, 2, 3, 4, 5], [1, 2, 3, 4, 5]]`)

	assertPretty(t, outerArray, 20, `[
    [1, 2, 3, 4, 5],
    [1, 2, 3, 4, 5],
    [1, 2, 3, 4, 5]
]`)

	assertPretty(t, outerArray, 10, `[
    [
        1,
        2,
        3,
        4,
        5
    ],
    [
        1,
        2,
        3,
        4,
        5
    ],
    [
        1,
        2,
        3,
        4,
        5
    ]
]`)
}

func TestHardline(t *testing.T) {

	doc := Join(HardLine{}, Text("first"), Text("second"))

	assertPretty(t, doc, 100, "first\nsecond")
}
