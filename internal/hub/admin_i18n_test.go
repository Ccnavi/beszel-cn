package hub

import (
	"strings"
	"testing"
)

func TestInjectAdminI18nScript(t *testing.T) {
	html := "<html><body><div id=\"app\"></div></body></html>"
	injected := injectAdminI18nScript(html)
	if !strings.Contains(injected, adminI18nScriptPath) {
		t.Fatalf("expected injected HTML to contain %q", adminI18nScriptPath)
	}
	if strings.Count(injectAdminI18nScript(injected), adminI18nScriptPath) != 1 {
		t.Fatal("expected admin i18n script to be injected only once")
	}
}

func TestInjectAdminI18nScriptWithoutBody(t *testing.T) {
	html := "<div id=\"app\"></div>"
	injected := injectAdminI18nScript(html)
	if !strings.HasSuffix(injected, `</script>`) {
		t.Fatalf("expected script to be appended, got %q", injected)
	}
}
