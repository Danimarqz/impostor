package game

import (
	"fmt"
	"impostor/internal/domain"
	"log"
	"math/rand"
	"time"
)

// StartGame initiates the match if conditions are met.
func (l *Lobby) StartGame(category domain.Category) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.Players) < 3 {
		return fmt.Errorf("not enough players (minimum 3)")
	}

	if l.State != domain.StateWaiting {
		return fmt.Errorf("lobby is already in game")
	}

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
	rand.Seed(time.Now().UnixNano())
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

	// Hard Mode
	return domain.Card{
		DisplayedWord: "YOU ARE THE IMPOSTOR",
		IsImpostor:    true,
	}
}
