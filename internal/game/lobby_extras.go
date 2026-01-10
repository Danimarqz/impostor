package game

import (
	"log"
)

// Lobby methods expanded

// Broadcast sends a message to all players in the lobby.
// StartGame and other actions will use this.
// For now, we don't have the WebSocket connection in the Player struct domain.
// We need a way to link Player ID -> Websocket Connection.
// This is missing in the current design.
// Real-world solution: The 'Player' struct in domain shouldn't hold the connection.
// The 'Lobby' in 'game' package should hold a map of connections.

// Let's safe update Lobby struct in hub.go (or move it here) to include connections.
// But wait, `hub.go` has the struct definition. I should update it there or moved it.
// I will keep it in `hub.go` for now but we need to solve the connection mapping.

// Solution: Add a `Clients` map to Lobby struct in `hub.go`.
// map[string]*websocket.Conn (Wait, `game` package shouldn't depend on `fiber/websocket` if we want clean arch)
// Interface? `NetworkClient` interface with `WriteJSON`.

type NetworkClient interface {
	WriteJSON(v any) error
}

// Update Lobby struct in `hub.go` to include `Clients map[string]NetworkClient`
// But I can't easily edit `hub.go` struct definition without replacing the file or multi-replace.
// Let's assume I'll do that in `hub.go`.

// Helper function to be called from handler
func (l *Lobby) RegisterClient(playerID string, client NetworkClient) {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	// We might need a separate map for clients if we don't want to pollute domain.Player
	if l.Clients == nil {
		l.Clients = make(map[string]NetworkClient)
	}
	l.Clients[playerID] = client
}

func (l *Lobby) Broadcast(msg any) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for id, client := range l.Clients {
		if err := client.WriteJSON(msg); err != nil {
			log.Printf("Error sending to player %s: %v", id, err)
			// Maybe unregister?
		}
	}
}
