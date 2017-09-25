package htmlsafe

import (
	"bytes"
	"testing"
)

func testHelper(t *testing.T, raw, expected string) {
	t.Helper()

	r := bytes.NewReader([]byte(raw))
	b, err := Run(r)

	if err != nil {
		t.Errorf("expected nil err, was: %v", err)
	} else if actual := b.String(); actual != expected {
		t.Errorf("got `%s`, expected `%s`", actual, expected)
	}
}

func TestSanitizeScript(t *testing.T) {
	var raw string

	raw = "<div><script><div>Hello</div></script>what</div>\n<script name=\"abc\">Link</script>"
	testHelper(t, raw, "<div>what</div>\n")

	raw = "Style is eaten: <style><xtag></style>"
	testHelper(t, raw, "Style is eaten: ")
}

func TestTagClose(t *testing.T) {
	testHelper(t, "<a>", "<a></a>")
	testHelper(t, "<a><b></c><c>", "<a><b><c></c></b></a>")
}

func TestVoidElement(t *testing.T) {
	testHelper(t, "<link>", "<link>")
	testHelper(t, "<p><link></link>", "<p><link></p>")
	testHelper(t, "<p><link></p></link>", "<p><link></p>")
}
