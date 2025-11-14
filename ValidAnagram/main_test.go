package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "範例1：字母異位詞",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "範例2：不是字母異位詞",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "空字串",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "單一字元_相同",
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			name: "單一字元_不同",
			s:    "a",
			t:    "b",
			want: false,
		},
		{
			name: "長度不同",
			s:    "ab",
			t:    "abc",
			want: false,
		},
		{
			name: "完全相同的字串",
			s:    "hello",
			t:    "hello",
			want: true,
		},
		{
			name: "重複字母",
			s:    "aabbcc",
			t:    "abcabc",
			want: true,
		},
		{
			name: "重複字母_數量不同",
			s:    "aab",
			t:    "abb",
			want: false,
		},
		{
			name: "所有字母",
			s:    "abcdefghijklmnopqrstuvwxyz",
			t:    "zyxwvutsrqponmlkjihgfedcba",
			want: true,
		},
		{
			name: "長字串_字母異位詞",
			s:    "listen",
			t:    "silent",
			want: true,
		},
		{
			name: "長字串_不是字母異位詞",
			s:    "listen",
			t:    "enlist",
			want: true,
		},
		{
			name: "只有一個字母不同",
			s:    "abc",
			t:    "abd",
			want: false,
		},
		{
			name: "Unicode字元（超出a-z範圍）",
			s:    "你好",
			t:    "好你",
			want: true, // 對於陣列法可能失敗，但雜湊表法可以處理
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 對於 Unicode 測試，只使用支援的方法
			if tt.name == "Unicode字元（超出a-z範圍）" {
				got := isAnagramHashMap(tt.s, tt.t)
				if got != tt.want {
					t.Errorf("isAnagramHashMap() = %v, want %v", got, tt.want)
				}
			} else {
				got := isAnagram(tt.s, tt.t)
				if got != tt.want {
					t.Errorf("isAnagram() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestIsAnagramHashMap(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "範例1",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "範例2",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "Unicode字元",
			s:    "你好世界",
			t:    "界世好你",
			want: true,
		},
		{
			name: "混合字元",
			s:    "Hello123",
			t:    "321olleH",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagramHashMap(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isAnagramHashMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAnagramSorting(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "範例1",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "範例2",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "空字串",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "長字串",
			s:    "listen",
			t:    "silent",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagramSorting(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isAnagramSorting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAnagramTwoMaps(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "範例1",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "範例2",
			s:    "rat",
			t:    "car",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagramTwoMaps(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isAnagramTwoMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基準測試：比較不同解法的效能
func BenchmarkIsAnagram(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
		t    string
	}{
		{
			name: "短字串_7個字元",
			s:    "anagram",
			t:    "nagaram",
		},
		{
			name: "中字串_26個字元（所有字母）",
			s:    "abcdefghijklmnopqrstuvwxyz",
			t:    "zyxwvutsrqponmlkjihgfedcba",
		},
		{
			name: "長字串_1000個字元",
			s:    string(make([]byte, 1000)),
			t:    string(make([]byte, 1000)),
		},
		{
			name: "不是字母異位詞_提前退出",
			s:    "abc",
			t:    "def",
		},
		{
			name: "長度不同_提前退出",
			s:    "abc",
			t:    "abcd",
		},
	}

	// 初始化長字串測試案例
	for i := range benchmarks {
		if benchmarks[i].name == "長字串_1000個字元" {
			sBytes := make([]byte, 1000)
			tBytes := make([]byte, 1000)
			for j := 0; j < 1000; j++ {
				sBytes[j] = byte('a' + (j % 26))
				tBytes[j] = byte('a' + ((1000 - j - 1) % 26))
			}
			benchmarks[i].s = string(sBytes)
			benchmarks[i].t = string(tBytes)
		}
	}

	for _, bm := range benchmarks {
		b.Run("陣列計數_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isAnagram(bm.s, bm.t)
			}
		})

		b.Run("雜湊表_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isAnagramHashMap(bm.s, bm.t)
			}
		})

		b.Run("排序_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isAnagramSorting(bm.s, bm.t)
			}
		})

		b.Run("兩個Map_"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isAnagramTwoMaps(bm.s, bm.t)
			}
		})
	}
}
