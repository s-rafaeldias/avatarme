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
