package cleanstring

import "testing"

func TestCleanSentence(t *testing.T) {
  got := Clean("The quick brown fox jumps over the lazy dog.")
  want := "the_quick_brown_fox_jumps_over_the_lazy_dog"
  if got != want {
    t.Errorf("got: %s, want: %s", got, want)
  }
}
