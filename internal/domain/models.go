package domain

// Role represents the player's role in the game.
type Role string

const (
	RoleCivilian Role = "CIVILIAN"
	RoleImpostor Role = "IMPOSTOR"
)

// GameMode defines the difficulty of the match.
type GameMode string

const (
	ModeHard GameMode = "HARD" // Impostor sees "YOU ARE THE IMPOSTOR"
	ModeEasy GameMode = "EASY" // Impostor sees a trap word
)

// Player represents a connected user.
type Player struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsLeader bool   `json:"is_leader"`
	Role     Role   `json:"role,omitempty"` // Omitted if not revealed
	IsAlive  bool   `json:"is_alive"`
}

// Card contains the information displayed to each player.
type Card struct {
	DisplayedWord string `json:"displayed_word"`
	IsImpostor    bool   `json:"is_impostor"`
	// In Easy mode, if IsImpostor is true, DisplayedWord is the trap word.
	// In Hard mode, if IsImpostor is true, DisplayedWord is "YOU ARE THE IMPOSTOR".
}

// Category groups words by theme.
type Category struct {
	Name string    `json:"name"`
	Pairs []WordPair `json:"pairs"`
}

// WordPair links a real word with its "trap" version for Easy mode.
type WordPair struct {
	Real string `json:"real"`
	Trap string `json:"trap"`
}

// LobbyConfig defines the rules of the match.
type LobbyConfig struct {
	Mode           GameMode `json:"mode"`
	Language       string   `json:"language"` // "en" or "es"
	Categories     []string `json:"categories"`
	Rounds         int      `json:"rounds"`
	VotingTimeSecs int      `json:"voting_time_secs"`
}

// LobbyState defines the current phase of the match.
type LobbyState string

const (
	StateWaiting  LobbyState = "WAITING"
	StatePlaying  LobbyState = "PLAYING"
	StateVoting   LobbyState = "VOTING"
	StateFinished LobbyState = "FINISHED"
)
