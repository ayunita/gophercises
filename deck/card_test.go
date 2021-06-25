package deck

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Three, Suit: Spade})
	fmt.Println(Card{Rank: King, Suit: Club})
	fmt.Println(Card{Rank: Six, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Three of Spades
	// King of Clubs
	// Six of Diamonds
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	want := 13 * 4
	got := len(cards)
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestAbsRank(t *testing.T) {
	testCases := []struct {
		Suit    Suit
		Rank    Rank
		absRank int
	}{
		{Spade, Five, 5},
		{Diamond, Three, 16},
		{Club, King, 39},
		{Heart, Jack, 50},
	}
	for _, c := range testCases {
		want := c.absRank
		got := absRank(Card{Rank: c.Rank, Suit: c.Suit})
		if want != got {
			t.Errorf("%s %s: want %d, got %d", c.Rank.String(), c.Suit.String(), want, got)
		}
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	want := Card{Rank: Ace, Suit: Spade}
	got := cards[0]
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	want := Card{Rank: Ace, Suit: Spade}
	got := cards[0]
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	want := 3
	got := count
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Suit == Spade || card.Suit == Heart || card.Suit == Diamond
	}

	cards := New(Filter(filter))
	want := Club
	for _, c := range cards {
		got := c.Suit
		if want != got {
			t.Errorf("want: %s, got: %s", want, got)
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	want := 13 * 4 * 3
	got := len(cards)
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
func TestShuffle(t *testing.T) {
	before := New()
	after := Shuffle(before)
	if reflect.DeepEqual(before, after) {
		t.Error("Before and after must be different")
	}
}

func TestFYShuffle(t *testing.T) {
	cards := New()
	before := append([]Card(nil), cards...)
	after := FYShuffle(cards)
	if reflect.DeepEqual(before, after) {
		t.Error("Before and after must be different")
	}
}
