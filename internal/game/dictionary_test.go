package game

import (
	"testing"
)

func TestGetCategoryByName(t *testing.T) {
	tests := []struct {
		name     string
		catName  string
		lang     string
		wantName string
	}{
		{"English existing", "Animals", "en", "Animals"},
		{"Spanish existing", "Animales", "es", "Animales"},
		{"Default fallback", "NonExistent", "en", "General"},
		{"Infinite category", "✨ Infinite", "en", "✨ Infinite"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCategoryByName(tt.catName, tt.lang)
			if got.Name != tt.wantName {
				t.Errorf("GetCategoryByName() = %v, want %v", got.Name, tt.wantName)
			}
		})
	}
}

func TestGetRandomWord(t *testing.T) {
	// Simple smoke test
	word, err := GetRandomWord("General", "en")
	if err != nil {
		t.Errorf("GetRandomWord() error = %v", err)
	}
	if word.Real == "" || word.Trap == "" {
		t.Errorf("GetRandomWord() returned empty word components: %v", word)
	}

	// Test Spanish
	wordES, err := GetRandomWord("General", "es")
	if err != nil {
		t.Errorf("GetRandomWord(es) error = %v", err)
	}
	if wordES.Real == "" {
		t.Errorf("GetRandomWord(es) returned empty word")
	}
}
