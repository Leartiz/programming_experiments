package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func referenceFullJustify(words []string, maxWidth int) []string {
	var result []string
	i := 0
	for i < len(words) {
		j := i
		lineLen := 0
		for j < len(words) {
			add := len(words[j])
			if j > i {
				add++
			}
			if lineLen+add > maxWidth {
				break
			}
			lineLen += add
			j++
		}
		lineWords := words[i:j]
		i = j

		if i >= len(words) {
			result = append(result, referenceLastLine(lineWords, maxWidth))
		} else {
			result = append(result, referenceJustifyLine(lineWords, maxWidth))
		}
	}
	return result
}

func referenceJustifyLine(words []string, maxWidth int) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", maxWidth-len(words[0]))
	}

	chars := 0
	for _, w := range words {
		chars += len(w)
	}
	spaces := maxWidth - chars
	gaps := len(words) - 1
	base := spaces / gaps
	extra := spaces % gaps

	var b strings.Builder
	for i, w := range words {
		b.WriteString(w)
		if i == len(words)-1 {
			break
		}
		b.WriteString(strings.Repeat(" ", base))
		if extra > 0 {
			b.WriteByte(' ')
			extra--
		}
	}
	return b.String()
}

func referenceLastLine(words []string, maxWidth int) string {
	line := strings.Join(words, " ")
	return line + strings.Repeat(" ", maxWidth-len(line))
}

func assertFullJustify(t *testing.T, words []string, maxWidth int, want []string) {
	t.Helper()

	got := fullJustify(words, maxWidth)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("fullJustify(%q, %d)\ngot:  %#v\nwant: %#v", words, maxWidth, got, want)
	}

	if err := validateOutput(words, maxWidth, got); err != nil {
		t.Errorf("validateOutput: %v", err)
	}

	ref := referenceFullJustify(words, maxWidth)
	if !reflect.DeepEqual(got, ref) {
		t.Errorf("mismatch with reference for %q, maxWidth=%d\ngot:  %#v\nref:  %#v", words, maxWidth, got, ref)
	}
}

func validateOutput(words []string, maxWidth int, lines []string) error {
	if len(lines) == 0 {
		return fmt.Errorf("empty result")
	}

	var recovered []string
	for i, line := range lines {
		if len(line) != maxWidth {
			return fmt.Errorf("line %d length=%d, want %d: %q", i, len(line), maxWidth, line)
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			return fmt.Errorf("line %d is blank: %q", i, line)
		}

		parts := strings.Fields(line)
		recovered = append(recovered, parts...)
	}

	if !reflect.DeepEqual(recovered, words) {
		return fmt.Errorf("words mismatch: got %v, want %v", recovered, words)
	}

	last := lines[len(lines)-1]
	if strings.HasSuffix(strings.TrimRight(last, " "), "  ") {
		// last line must be left-justified: no double spaces between words
		inner := strings.TrimRight(last, " ")
		if strings.Contains(inner, "  ") {
			return fmt.Errorf("last line has extra internal spaces: %q", last)
		}
	}

	return nil
}

func TestFullJustify_LeetCodeExamples(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		{
			name:     "example1",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			name:     "example2",
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			name:     "example3",
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFullJustify(t, tt.words, tt.maxWidth, tt.want)
		})
	}
}

func TestFullJustify_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		{
			name:     "single_word",
			words:    []string{"hello"},
			maxWidth: 10,
			want:     []string{"hello     "},
		},
		{
			name:     "single_line_two_words",
			words:    []string{"hi", "there"},
			maxWidth: 8,
			want:     []string{"hi there"},
		},
		{
			name:     "exact_fit_one_space_per_gap",
			words:    []string{"ab", "cd"},
			maxWidth: 5,
			want:     []string{"ab cd"},
		},
		{
			name:     "one_word_per_line_except_last",
			words:    []string{"aa", "bb", "cc", "dd"},
			maxWidth: 4,
			want: []string{
				"aa  ",
				"bb  ",
				"cc  ",
				"dd  ",
			},
		},
		{
			name:     "single_word_middle_line",
			words:    []string{"aa", "bbbb", "cc"},
			maxWidth: 4,
			want: []string{
				"aa  ",
				"bbbb",
				"cc  ",
			},
		},
		{
			name:     "max_width_equals_word_length",
			words:    []string{"word", "word", "word"},
			maxWidth: 4,
			want: []string{
				"word",
				"word",
				"word",
			},
		},
		{
			name:     "all_words_on_last_line",
			words:    []string{"a", "b", "c"},
			maxWidth: 7,
			want:     []string{"a b c  "},
		},
		{
			name:     "two_words_per_line_when_fit",
			words:    []string{"a", "b", "c", "d", "e"},
			maxWidth: 3,
			want: []string{
				"a b",
				"c d",
				"e  ",
			},
		},
		{
			name:     "two_words_last_line_padding",
			words:    []string{"go", "to", "a", "party"},
			maxWidth: 5,
			want: []string{
				"go to",
				"a    ",
				"party",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFullJustify(t, tt.words, tt.maxWidth, tt.want)
		})
	}
}

// Adversarial cases aimed at nextLineWords, uneven space distribution,
// single-word non-last lines, and exact-fit packing.
func TestFullJustify_Adversarial(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		{
			name:     "uneven_remainder_three_gaps",
			words:    []string{"a", "b", "c", "d"},
			maxWidth: 7,
			want:     []string{"a b c d"},
		},
		{
			name:     "many_single_char_words",
			words:    []string{"a", "b", "c", "d", "e", "f", "g", "h"},
			maxWidth: 5,
			want: []string{
				"a b c",
				"d e f",
				"g h  ",
			},
		},
		{
			name:     "long_word_forces_single_word_line",
			words:    []string{"short", "verylongwordhere", "x"},
			maxWidth: 16,
			want: []string{
				"short           ",
				"verylongwordhere",
				"x               ",
			},
		},
		{
			name:     "exact_fit_full_line_non_last",
			words:    []string{"abc", "def", "ghi", "j"},
			maxWidth: 11,
			want: []string{
				"abc def ghi",
				"j          ",
			},
		},
		{
			name:     "max_width_100_stress",
			words:    []string{strings.Repeat("a", 20), strings.Repeat("b", 20), strings.Repeat("c", 20), strings.Repeat("d", 20), strings.Repeat("e", 20)},
			maxWidth: 100,
			want: []string{
				strings.Repeat("a", 20) + strings.Repeat(" ", 7) + strings.Repeat("b", 20) + strings.Repeat(" ", 7) + strings.Repeat("c", 20) + strings.Repeat(" ", 6) + strings.Repeat("d", 20),
				strings.Repeat("e", 20) + strings.Repeat(" ", 80),
			},
		},
		{
			name:     "space_between_zero_extra_one",
			words:    []string{"ab", "cd", "ef", "g"},
			maxWidth: 9,
			want: []string{
				"ab  cd ef",
				"g        ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertFullJustify(t, tt.words, tt.maxWidth, tt.want)
		})
	}
}

func TestNextLineWords(t *testing.T) {
	tests := []struct {
		name      string
		words     []string
		wordIndex int
		maxWidth  int
		wantWords []string
		wantIndex int
	}{
		{
			name:      "packs_until_overflow",
			words:     []string{"This", "is", "an", "example"},
			wordIndex: 0,
			maxWidth:  16,
			wantWords: []string{"This", "is", "an"},
			wantIndex: 3,
		},
		{
			name:      "exact_fit",
			words:     []string{"ab", "cd"},
			wordIndex: 0,
			maxWidth:  5,
			wantWords: []string{"ab", "cd"},
			wantIndex: 2,
		},
		{
			name:      "single_word_only",
			words:     []string{"acknowledgment", "shall"},
			wordIndex: 0,
			maxWidth:  16,
			wantWords: []string{"acknowledgment"},
			wantIndex: 1,
		},
		{
			name:      "from_middle",
			words:     []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			wordIndex: 3,
			maxWidth:  16,
			wantWords: []string{"acknowledgment"},
			wantIndex: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWords, gotIndex := nextLineWords(tt.words, tt.wordIndex, tt.maxWidth)
			if !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("words: got %v, want %v", gotWords, tt.wantWords)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("index: got %d, want %d", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestLineJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     string
	}{
		{
			name:     "three_words",
			words:    []string{"This", "is", "an"},
			maxWidth: 16,
			want:     "This    is    an",
		},
		{
			name:     "single_word_non_last",
			words:    []string{"acknowledgment"},
			maxWidth: 16,
			want:     "acknowledgment  ",
		},
		{
			name:     "uneven_spaces",
			words:    []string{"example", "of", "text"},
			maxWidth: 16,
			want:     "example  of text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lineJustify(tt.words, tt.maxWidth)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			if len(got) != tt.maxWidth {
				t.Errorf("line length=%d, want %d", len(got), tt.maxWidth)
			}
		})
	}
}

func TestLastLineJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     string
	}{
		{
			name:     "single_word",
			words:    []string{"justification."},
			maxWidth: 16,
			want:     "justification.  ",
		},
		{
			name:     "two_words",
			words:    []string{"shall", "be"},
			maxWidth: 16,
			want:     "shall be        ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lastLineJustify(tt.words, tt.maxWidth)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			if len(got) != tt.maxWidth {
				t.Errorf("line length=%d, want %d", len(got), tt.maxWidth)
			}
		})
	}
}

func TestFullJustify_Fuzz(t *testing.T) {
	rng := newRand(1)

	for trial := 0; trial < 50000; trial++ {
		maxWidth := rng.Intn(100) + 1
		n := rng.Intn(20) + 1
		words := make([]string, n)
		for i := range words {
			wl := rng.Intn(20) + 1
			if wl > maxWidth {
				wl = maxWidth
			}
			words[i] = randomWord(rng, wl)
		}

		got := fullJustify(words, maxWidth)
		want := referenceFullJustify(words, maxWidth)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("trial=%d maxWidth=%d words=%v\ngot:  %#v\nwant: %#v", trial, maxWidth, words, got, want)
		}
		if err := validateOutput(words, maxWidth, got); err != nil {
			t.Fatalf("trial=%d maxWidth=%d words=%v: %v", trial, maxWidth, words, err)
		}
	}
}

type testRand struct {
	state uint64
}

func newRand(seed uint64) *testRand {
	return &testRand{state: seed}
}

func (r *testRand) Intn(n int) int {
	r.state = r.state*6364136223846793005 + 1
	return int(r.state % uint64(n))
}

func randomWord(r *testRand, length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func TestFullJustify_ExhaustiveSmallCases(t *testing.T) {
	alphabet := []string{"a", "b", "c", "d", "ab", "bc", "cd", "abc"}

	for maxWidth := 1; maxWidth <= 8; maxWidth++ {
		for size := 1; size <= 4; size++ {
			var words []string
			var build func(depth int)
			build = func(depth int) {
				if depth == size {
					if len(words) == 0 {
						return
					}
					for _, w := range words {
						if len(w) > maxWidth {
							return
						}
					}

					got := fullJustify(words, maxWidth)
					ref := referenceFullJustify(words, maxWidth)
					if !reflect.DeepEqual(got, ref) {
						t.Errorf("maxWidth=%d words=%v\ngot: %v\nref: %v", maxWidth, words, got, ref)
					}
					if err := validateOutput(words, maxWidth, got); err != nil {
						t.Errorf("maxWidth=%d words=%v: %v", maxWidth, words, err)
					}
					return
				}
				for _, w := range alphabet {
					words = append(words, w)
					build(depth + 1)
					words = words[:len(words)-1]
				}
			}
			build(0)
		}
	}
}
