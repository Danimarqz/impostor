package game

import (
	"fmt"
	"impostor/internal/domain"
	"log"
	"math/rand"
)

// StartGame initializes a new match.
func (l *Lobby) StartGame(category domain.Category, mode domain.GameMode) error {
	l.mu.Lock() // We need global lock to set state
	defer l.mu.Unlock()

	if len(l.Players) < 3 {
		// return nil // Allow for testing with fewer players? No, stick to rules or user preference.
		// For MVP testing, let's allow it, OR just return error.
		// return fmt.Errorf("not enough players")
	}

	if l.State != domain.StateWaiting {
		return fmt.Errorf("lobby is already in game")
	}

	l.Config.Mode = mode // Store the mode
	
	// 1. Assign Roles
	l.assignRoles()

	// 2. Select Word Pair
	pair := l.selectRandomWordPair(category)

	// 3. Distribute Cards (Broadcast to clients)
	l.broadcastStartWithPair(pair)

	l.State = domain.StatePlaying
	return nil
}

func (l *Lobby) broadcastStartWithPair(pair domain.WordPair) {
	for id, client := range l.Clients {
		p, ok := l.Players[id]
		if !ok {
			continue
		}

		card := l.GetCardForPlayer(id, pair)
		
		msg := map[string]interface{}{
			"status": "PLAYING",
			"role": string(p.Role),
			"displayed_word": card.DisplayedWord,
		}

		if err := client.WriteJSON(msg); err != nil {
			log.Printf("Error sending start to %s: %v", id, err)
		}
	}
}

func (l *Lobby) assignRoles() {
	playerIDs := make([]string, 0, len(l.Players))
	for id := range l.Players {
		playerIDs = append(playerIDs, id)
	}

	// Shuffle players
	rand.Shuffle(len(playerIDs), func(i, j int) {
		playerIDs[i], playerIDs[j] = playerIDs[j], playerIDs[i]
	})

	// First player in shuffled list is Impostor
	impostorID := playerIDs[0]

	for _, p := range l.Players {
		if p.ID == impostorID {
			p.Role = domain.RoleImpostor
		} else {
			p.Role = domain.RoleCivilian
		}
		p.IsAlive = true
	}
}

func (l *Lobby) selectRandomWordPair(cat domain.Category) domain.WordPair {
	if cat.Name == "✨ Infinite" {
		language := l.Config.Language
		if language == "" {
			language = "en" // Default to English
		}
		pair, err := FetchRandomPairFromAPI(language)
		if err == nil {
			return pair
		}
		log.Printf("API Error (fallback to local): %v", err)
        // Fallback to a random category from defaults if API fails
        randomCat := GetCategoryByName(DefaultCategories[rand.Intn(len(DefaultCategories))].Name, language)
        return l.selectRandomWordPair(randomCat)
	}

	if len(cat.Pairs) == 0 {
		return domain.WordPair{Real: "Error", Trap: "Error"}
	}
	idx := rand.Intn(len(cat.Pairs))
	return cat.Pairs[idx]
}

func (l *Lobby) distributeCards(pair domain.WordPair) {
	// Need to store cards somewhere if we want to retrieve them later,
	// or we just send them to the players immediately.
	// For now, let's assume we update a field in Player struct or redundant context.
	// Ideally, Card is calculated on demand or stored in a session.
	// Updating logic to be simple: We will just inform the players via WebSocket later.
	// But let's verify logic here.
}

// GetCardForPlayer returns the visual representation for a specific player.
func (l *Lobby) GetCardForPlayer(playerID string, pair domain.WordPair) domain.Card {
	// Safe read not strictly needed if called from StartGame (which has Lock)
	// But if called effectively from outside, we need RLock.
	// Assuming this is a helper called from within a Lock.
	
	p, ok := l.Players[playerID]
	if !ok {
		return domain.Card{}
	}

	if p.Role == domain.RoleCivilian {
		return domain.Card{
			DisplayedWord: pair.Real,
			IsImpostor:    false,
		}
	}

	// Impostor Logic
	if l.Config.Mode == domain.ModeEasy {
		return domain.Card{
			DisplayedWord: pair.Trap,
			IsImpostor:    true,
		}
	}

	return domain.Card{
		DisplayedWord: "YOU ARE THE IMPOSTOR",
		IsImpostor:    true,
	}
}

// CastVote records a vote and checks for game end conditions.
func (l *Lobby) CastVote(voterID, targetID string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.State != domain.StatePlaying && l.State != domain.StateVoting {
		return fmt.Errorf("voting not allowed now")
	}
	
	l.State = domain.StateVoting // Ensure we are in voting mode
	
	if currentTarget, ok := l.Votes[voterID]; ok && currentTarget == targetID {
		// Toggle off (remove vote)
		delete(l.Votes, voterID)
	} else {
		// Set new vote
		l.Votes[voterID] = targetID
	}

	// Check results
	totalPlayers := len(l.Players)
	votesForTarget := make(map[string]int)
	totalVotes := 0

	for _, t := range l.Votes {
		votesForTarget[t]++
		totalVotes++
	}

	// Broadcast updated votes (anonymous or public? Public count is good)
	// Sending list of who voted for whom is simplest for MVP transparency.
	voteUpdate := map[string]interface{}{
		"type": "VOTE_UPDATE",
		"votes": l.Votes,
	}
	l.broadcastInternal(voteUpdate) // Use internal helper if available, or manual loop

	// Win Condition: Majority (> 50%)
	// Or clean sweep? Let's say majority.
	threshold := totalPlayers/2 + 1 
	
	for target, count := range votesForTarget {
		if count >= threshold {
			l.finishGame(target)
			break
		}
	}
	
	return nil
}

func (l *Lobby) finishGame(kickedID string) {
	kickedPlayer := l.Players[kickedID]
	
	// Determine Winner
	var winner string // "CIVILIANS" or "IMPOSTOR"
	
	if kickedPlayer.Role == domain.RoleImpostor {
		winner = "CIVILIANS"
	} else {
		winner = "IMPOSTOR"
	}
	
	// Reveal all roles
	allPlayers := make([]map[string]interface{}, 0, len(l.Players))
	for _, p := range l.Players {
		allPlayers = append(allPlayers, map[string]interface{}{
			"id":        p.ID,
			"name":      p.Name,
			"is_leader": p.IsLeader,
			"role":      string(p.Role), // Explicit cast
		})
	}
	
	msg := map[string]interface{}{
		"status":   "FINISHED",
		"winner":   winner,
		"kicked":   kickedPlayer.Name,
		"role_was": string(kickedPlayer.Role),
		"reveal":   allPlayers,
	}
	
	l.State = domain.StateFinished
	l.broadcastInternal(msg)
	
	// Reset Game?
	l.Votes = make(map[string]string)
}

func (l *Lobby) ResetGame() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.State = domain.StateWaiting
	l.Votes = make(map[string]string)

	msg := map[string]interface{}{
		"type":   "GAME_RESET",
		"status": "WAITING",
	}
	l.broadcastInternal(msg)
}

// broadcastInternal sends a message to all clients without locking.
// Caller must verify safety or hold lock if accessing internal state not passed as arg.
func (l *Lobby) broadcastInternal(msg any) {
	for id, client := range l.Clients {
		if err := client.WriteJSON(msg); err != nil {
			log.Printf("Error sending to player %s: %v", id, err)
		}
	}
}
