package cleanstring

import "testing"

func TestCleanSentence(t *testing.T) {
  got := Clean("The quick brown fox jumps over the lazy dog.")
  want := "the_quick_brown_fox_jumps_over_the_lazy_dog"
  if got != want {
    t.Errorf("got: %s, want: %s", got, want)
  }
}

func TestUnbrokenNonAlphanum(t *testing.T) {
  got := Clean("Acme & Co.")
  want := "acme_co"
  if got != want {
    t.Errorf("got: %s, want: %s", got, want)
  }
}

func TestUnbrokenNonAlphanumHead(t *testing.T) {
  got := Clean("... hello? Is anybody there?")
  want := "hello_is_anybody_there"
  if got != want {
    t.Errorf("got: %s, want: %s", got, want)
  }
}

func TestUnbrokenNonAlphanumTail(t *testing.T) {
  got := Clean("What did you just say?!?!")
  want := "what_did_you_just_say"
  if got != want {
    t.Errorf("got: %s, want: %s", got, want)
  }
}
