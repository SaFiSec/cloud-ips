package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderVersionOutput(t *testing.T) {
	var remoteFilepathTests = []struct {
		comment string
		want    string
		params  cmdVersion
	}{
		{
			"case 1",
			`Version  v1.0.0
Date     2018-01-01T00:00:00+11:00
API      v1`,
			cmdVersion{
				APICompatibility: 1,
				BuildVersion: "v1.0.0",
				BuildDate: "2018-01-01T00:00:00+11:00",
			},
		},
		{
			"case 2",
			`Version  v1.0.0-b715353
Date     2018-01-12T16:42:06+11:00
API      v2`,
			cmdVersion{
				APICompatibility: 2,
				BuildVersion: "v1.0.0-b715353",
				BuildDate: "2018-01-12T16:42:06+11:00",
			},
		},
	}

	for _, testCase := range remoteFilepathTests {
		actual := RenderVersionOutput(&testCase.params)
		assert.Equal(t, testCase.want, actual, testCase.comment)
	}
}