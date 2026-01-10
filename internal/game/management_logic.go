package game

import "log"

// RemovePlayer removes a player and handles leader reassignment.
// Returns true if the lobby is empty and should be deleted.
func (l *Lobby) RemovePlayer(playerID string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	p, exists := l.Players[playerID]
	if !exists {
		return len(l.Players) == 0
	}

	wasLeader := p.IsLeader
	delete(l.Players, playerID)
	delete(l.Clients, playerID)

	if len(l.Players) == 0 {
		return true // Lobby is empty
	}

	if wasLeader {
		// Assign new leader (just pick one, map order is random but acceptable for MVP)
		for _, remainingPlayer := range l.Players {
			remainingPlayer.IsLeader = true
			log.Printf("Player %s is the new leader", remainingPlayer.Name)
			
			// Broadcast new leader info
			// We can send a generic UPDATE_PLAYER or PLAYER_LEFT with new leader info
			// Ideally we broadcast the whole list again or a specific event.
			// Let's rely on the caller to broadcast "PLAYER_LEFT" and updated list?
			// No, we should do it here if possible, but we need to handle the socket writes carefully (we are locked).
			// Let's just return, and let the caller broadcast.
			// But wait, the caller doesn't know WHO new leader is.
			break // Only assign one
		}
	}

	return false
}

func (l *Lobby) BroadcastPlayerList() {
	l.mu.RLock()
	log.Printf("Broadcasting player list for lobby %s. Count: %d", l.ID, len(l.Players))
	// Gather data first to avoid holding lock during network I/O or recursive locking
	players := make([]map[string]interface{}, 0, len(l.Players))
	for _, p := range l.Players {
		players = append(players, map[string]interface{}{
			"id":        p.ID,
			"name":      p.Name,
			"is_leader": p.IsLeader,
		})
	}
	l.mu.RUnlock()

	msg := map[string]any{
		"type":    "PLAYER_LIST", // Frontend should handle this to replace the list
		"players": players,
	}

	l.Broadcast(msg)
}
