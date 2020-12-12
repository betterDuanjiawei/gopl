package word

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	//var s = []struct{
	//	{"aaa", true},
	//	{"aaa", true},
	//	{"aaa", true},
	//}
	var tests = []struct{
		input string
		want bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); test.want !=  got {
			t.Errorf("IsPalindrome(%v) = %t", test.input, got)
		}
	}
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ { // n+1 不是 n-1
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}

	return string(runes)
}

func TestRandomPalindrome(t *testing.T)  {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		s := randomPalindrome(rng)
		if !IsPalindrome(s) {
			t.Errorf("IsPalindrome(%q) = false", s)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//IsPalindrome("A man, a plan, a canal: Panama")
		//IsPalindrome2("A man, a plan, a canal: Panama")
		IsPalindrome3("A man, a plan, a canal: Panama")
	}
}