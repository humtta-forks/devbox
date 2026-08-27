package configfile

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
