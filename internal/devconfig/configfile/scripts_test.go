package configfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const config = `{
	"shell": {
		"scripts": {
			// Comment before the script name
			"before-name": "echo before-name",

			"before-vaule": /* Comment before the script value */ "echo before-vaule",

			// Comment before the script name
			"both": /* Comment before the script value too */ "echo both",

			//Comment without a space after the marker
			"without-space": "echo without-space",

			// First comment line
			// Second comment line
			//   Third comment line, with two spaces of indentation
			"multiline": "echo multiline",

			"without-comment": "echo without-comment"
		}
	}
}
`

func TestScriptsCommentExtraction(t *testing.T) {
	cfg, err := LoadBytes([]byte(config))
	require.NoError(t, err)

	comments := make(map[string]string)
	for name, script := range cfg.Scripts() {
		comments[name] = script.Comments
	}

	assert.Equal(t, map[string]string{
		"before-name":     "Comment before the script name",
		"before-vaule":    "Comment before the script value",
		"both":            "Comment before the script name\nComment before the script value too",
		"without-space":   "Comment without a space after the marker",
		"multiline":       "First comment line\nSecond comment line\n  Third comment line, with two spaces of indentation",
		"without-comment": "",
	}, comments)
}
