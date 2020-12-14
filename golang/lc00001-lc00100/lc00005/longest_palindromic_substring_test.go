package lc00005

import (
	"testing"
)

type caseType struct {
	s        string
	excepted string
}

var cases = []caseType{
	{"babad", "bab"},
	{"cbbd", "bb"},
	{"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth", "ranynar"},
	{"aacabdkacaa", "aca"},
	{"edllaaaaallf", "llaaaaall"},
}

func TestLongestPalindrome(t *testing.T) {
	for _, c := range cases {
		result := longestPalindrome(c.s)
		if result != c.excepted {
			t.Fatalf("longestPalindrome function failed, s: %s, execpted: %s, result: %s", c.s, c.excepted, result)
		}
	}
}

func TestCenterSpread(t *testing.T) {
	for _, c := range cases {
		result := centerSpread(c.s)
		if result != c.excepted {
			t.Fatalf("centerSpread function failed, s: %s, execpted: %s, result: %s", c.s, c.excepted, result)
		}
	}
}

func TestManacher(t *testing.T) {
	for _, c := range cases {
		result := manacher(c.s)
		if result != c.excepted {
			t.Fatalf("manacher function failed, s: %s, execpted: %s, result: %s", c.s, c.excepted, result)
		}
	}
}
