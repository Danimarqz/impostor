package game

import (
	"encoding/json"
	"fmt"
	"impostor/internal/domain"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// DatamuseWord represents the JSON response from Datamuse API
type DatamuseWord struct {
	Word  string   `json:"word"`
	Score int      `json:"score"`
	Tags  []string `json:"tags,omitempty"`
}

var seedWordsEN = []string{
	"fire", "water", "earth", "wind", "magic", "science", "space", "time",
	"love", "war", "peace", "king", "queen", "apple", "banana", "car",
	"house", "dog", "cat", "bird", "fish", "book", "computer", "phone",
	"music", "art", "happy", "sad", "fast", "slow", "red", "blue",
	"city", "forest", "mountain", "ocean", "river", "desert", "snow", "rain",
	"adventure", "mystery", "history", "future", "robot", "alien", "ghost", "vampire",
	"coffee", "tea", "beer", "wine", "pizza", "burger", "sushi", "cake",
    "doctor", "teacher", "police", "thief", "judge", "lawyer", "actor", "singer",
}

var seedWordsES = []string{
	"fuego", "agua", "tierra", "viento", "magia", "ciencia", "espacio", "tiempo",
	"amor", "guerra", "paz", "rey", "reina", "manzana", "plátano", "coche",
	"casa", "perro", "gato", "pájaro", "pez", "libro", "ordenador", "teléfono",
	"música", "arte", "feliz", "triste", "rápido", "lento", "rojo", "azul",
	"ciudad", "bosque", "montaña", "océano", "río", "desierto", "nieve", "lluvia",
	"aventura", "misterio", "historia", "futuro", "robot", "alien", "fantasma", "vampiro",
	"café", "té", "cerveza", "vino", "pizza", "hamburguesa", "sushi", "pastel",
    "doctor", "profesor", "policía", "ladrón", "juez", "abogado", "actor", "cantante",
}

// FetchRandomPairFromAPI attempts to get a related word pair using Datamuse.
// It uses a random seed word to find related words.
// language should be "en" or "es"
func FetchRandomPairFromAPI(language string) (domain.WordPair, error) {
	// Select appropriate seed words based on language
	seedWords := seedWordsEN
	if language == "es" {
		seedWords = seedWordsES
	}
	
	// Try up to 3 different seeds in case one fails
	maxRetries := 3
	for attempt := range maxRetries {
		// 1. Pick a random seed word
		seed := seedWords[rand.Intn(len(seedWords))]
		
		// 2. Fetch related words (Triggers/Associations are often good "Impostor" alternatives)
		url := fmt.Sprintf("https://api.datamuse.com/words?rel_trg=%s&max=20", seed)
		if rand.Float32() > 0.5 {
			// 50% chance to look for synonyms instead of associations (Harder)
			url = fmt.Sprintf("https://api.datamuse.com/words?rel_syn=%s&max=20", seed)
		}
		
		// Add vocabulary parameter for Spanish
		if language == "es" {
			url += "&v=es"
		}

		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get(url)
		if err != nil {
			if attempt == maxRetries-1 {
				return domain.WordPair{}, err
			}
			continue // Try another seed
		}
		defer resp.Body.Close()

		var results []DatamuseWord
		if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
			if attempt == maxRetries-1 {
				return domain.WordPair{}, err
			}
			continue // Try another seed
		}

		if len(results) == 0 {
			if attempt == maxRetries-1 {
				return domain.WordPair{}, fmt.Errorf("no results after %d attempts", maxRetries)
			}
			continue // Try another seed
		}

		// 3. Pick a random related word as the Trap
		trap := results[rand.Intn(len(results))].Word

		// Ensure they are not identical (API sometimes returns the word itself)
		if strings.EqualFold(seed, trap) {
			if len(results) > 1 {
				trap = results[(rand.Intn(len(results)-1) + 1)%len(results)].Word
			} else {
				continue // Try another seed
			}
		}

		// Capitalize for display
		return domain.WordPair{
			Real: strings.Title(seed),
			Trap: strings.Title(trap),
		}, nil
	}
	
	return domain.WordPair{}, fmt.Errorf("failed after %d attempts", maxRetries)
}
