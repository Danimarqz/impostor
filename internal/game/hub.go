package game

import (
	"impostor/internal/domain"
	"sync"
)

// Hub manages the global state of all active lobbies.
// Uses RWMutex to allow multiple concurrent reads (e.g., checking lobby status)
// but block writes (e.g., creating/deleting a lobby).
type Hub struct {
	lobbies map[string]*Lobby
	mu      sync.RWMutex // Protects the lobbies map
}

// NewHub creates a new Hub instance.
func NewHub() *Hub {
	return &Hub{
		lobbies: make(map[string]*Lobby),
	}
}

// CreateLobby initializes a new lobby safely.
func (h *Hub) CreateLobby(id string, host *domain.Player) *Lobby {
	h.mu.Lock() // WRITE LOCK: No one else can read or write to the map
	defer h.mu.Unlock()

	l := NewLobby(id, host)
	h.lobbies[id] = l
	return l
}

// GetLobby retrieves a lobby by ID safely.
func (h *Hub) GetLobby(id string) (*Lobby, bool) {
	h.mu.RLock() // READ LOCK: Others can read, but no one can write (delete/create)
	defer h.mu.RUnlock()

	l, ok := h.lobbies[id]
	return l, ok
}

// DeleteLobby removes a lobby when empty or finished.
func (h *Hub) DeleteLobby(id string) {
	h.mu.Lock() // WRITE LOCK
	defer h.mu.Unlock()

	delete(h.lobbies, id)
}

// Lobby represents a specific match and its players.
// It also needs its own concurrency control to protect its internal state.
type Lobby struct {
	ID      string
	Players map[string]*domain.Player
	Clients map[string]NetworkClient // Map playerID -> Connection
	Config  domain.LobbyConfig
	State   domain.LobbyState
	
	// Mutex to protect the Lobby's internal state (separate from Hub)
	// This allows actions in Lobby A not to block Lobby B.
	mu sync.RWMutex
}

func NewLobby(id string, host *domain.Player) *Lobby {
	players := make(map[string]*domain.Player)
	if host != nil {
		players[host.ID] = host
	}
	return &Lobby{
		ID:      id,
		Players: players,
		State:   domain.StateWaiting,
	}
}

// AddPlayer adds a player to the lobby in a thread-safe manner.
func (l *Lobby) AddPlayer(p *domain.Player) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Players[p.ID] = p
}
