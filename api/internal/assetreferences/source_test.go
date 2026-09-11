package assetreferences

import (
	"reflect"
	"testing"
)

func TestMediaKeysFollowRenderedUses(t *testing.T) {
	a, b := "34kRPmE5SHCmWUG0tPXBi", "34kRPmFQqIl8gb7G3IKcO"
	content := "![a](/media/" + a + "?format=webp&preset=inline&v=1)\n![again][picture]\n\n[picture]: /media/" + a + "?v=1\n\n<img src=\"https://docs.test/media/" + b + "?v=1\">\n\n`![code](/media/ignored)`\n\n```md\n![code](/media/ignored)\n```\n![external](https://other.test/media/ignored)"
	if got := MediaKeys(content, "https://docs.test"); !reflect.DeepEqual(got, []string{a, b}) {
		t.Fatalf("keys=%v", got)
	}
}
