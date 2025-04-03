package WildcardMatching

import "testing"

func TestIsMatch(t *testing.T) {
	testCases := []struct {
		s        string
		p        string
		expected bool
		desc     string
	}{
		{"aa", "a", false, "基本不匹配情況"},
		{"aa", "*", true, "星號匹配任何內容"},
		{"cb", "?a", false, "問號匹配一個字符"},
		{"adceb", "*a*b", true, "多個星號與字符"},
		{"acdcb", "a*c?b", false, "星號和問號的組合"},
		{"", "", true, "空字串和模式"},
		{"", "*", true, "空字串和星號模式"},
		{"abc", "", false, "非空字串和空模式"},
		{"abc", "abc", true, "精確匹配"},
		{"abcde", "a*e", true, "中間有星號"},
		{"mississippi", "m*issip*", true, "多個星號"},
		{"mississippi", "m*issip*i", true, "多個星號以字符結尾"},
		{"mississippi", "m*issip*?", true, "多個星號以問號結尾"},
		{"aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba", "a*******b", false, "複雜模式與多個星號"},
		{"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb", false, "非常長的字串和模式"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := isMatch(tc.s, tc.p)
			if got != tc.expected {
				t.Errorf("isMatch(%q, %q) = %v; 期望 %v", tc.s, tc.p, got, tc.expected)
			}
		})
	}
}
