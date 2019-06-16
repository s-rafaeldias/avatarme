package patch

import (
	"github.com/s-rafaeldias/avatarme/util"
	"testing"
)

func TestGetHexColor(t *testing.T) {

	tables := []struct {
		text  string
		color string
	}{
		{"color1 for test", "#7D1ADC"},
		{"color2 for test", "#006F24"},
		{"color3 for test", "#6A4EE8"},
		{"color4 for test", "#88B463"},
	}

	for _, table := range tables {
		text := util.GetMD5(table.text)
		color := GetHexColor(text[:])

		if color != table.color {
			t.Errorf("Color for text '%s' was incorrect\ngot: '%s', expected: '%s'",
				table.text, color, table.color)
		}
	}

}

func TestGetPatch(t *testing.T) {
	cases := []struct {
		text   string
		corner byte
		center byte
		side   byte
	}{
		{"text1 for test", byte(10), byte(11), byte(12)},
		{"text2 for test", byte(10), byte(11), byte(12)},
		{"text3 for test", byte(10), byte(11), byte(12)},
		{"text4 for test", byte(10), byte(11), byte(12)},
	}

	for _, testCase := range cases {
		text := util.GetMD5(testCase.text)
		corner, side, center := GetPatch(text[:])
		if corner != testCase.corner {
			t.Errorf("Corner value incorrect\ngot %d, expected %d", corner, testCase.corner)
		}
		if side != testCase.side {
			t.Errorf("Side value incorrect\ngot %d, expected %d", side, testCase.side)
		}
		if center != testCase.center {
			t.Errorf("Center value incorrect\ngot %d, expected %d", center, testCase.center)
		}
	}
}
