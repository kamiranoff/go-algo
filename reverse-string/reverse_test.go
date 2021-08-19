package reverse_string

import "testing"

func TestReverse(t *testing.T) {
	reversed := reverse("abcd")
	if reversed != "dcba" {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", reversed, "dcba")
	}

	reversed = reverse("  adcd")

	if reversed != "dcba  " {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", reversed, "dcba  ")
	}
}
