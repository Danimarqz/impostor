import { writable } from 'svelte/store';

// Define the shape of our frontend state (mirroring Go structs)
export interface GameState {
  status: 'CONNECTING' | 'WAITING' | 'PLAYING' | 'VOTING' | 'FINISHED';
  me: {
    id: string;
    name: string;
    isLeader: boolean;
    role?: 'CIVILIAN' | 'IMPOSTOR';
    word?: string;
  };
  players: Array<{ id: string; name: string }>;
}

const initialState: GameState = {
  status: 'CONNECTING',
  me: { id: '', name: '', isLeader: false },
  players: []
};

export const game = writable<GameState>(initialState);

export const updateGame = (newData: Partial<GameState>) => {
  game.update(current => ({ ...current, ...newData }));
};

let socket: WebSocket;

export const connect = (lobbyId: string, playerName: string) => {
  // Generate random ID for demo purposes if not persisted
  const playerId = Math.random().toString(36).substring(7);

  // Connect to Backend (adjust port if needed, assuming 8080)
  socket = new WebSocket(`ws://localhost:8080/ws/${lobbyId}?playerId=${playerId}&playerName=${playerName}`);

  socket.onopen = () => {
    console.log("Connected to WS");
    updateGame({
      status: 'WAITING',
      me: { id: playerId, name: playerName, isLeader: false }
    });
  };

  socket.onmessage = (event) => {
    try {
      const data = JSON.parse(event.data);
      console.log("WS Data:", data);

      // Handle specific messages
      // For MVP: If data has 'state', update it
      if (data.type === 'PLAYER_JOINED') {
        const joinedPlayer = data.player;
        game.update(g => {
          const isMe = g.me.id === joinedPlayer.id;
          return {
            ...g,
            me: isMe ? { ...g.me, isLeader: joinedPlayer.is_leader } : g.me,
            players: [...g.players, joinedPlayer]
          };
        });
      }
      if (data.players) {
        game.update(g => ({ ...g, players: data.players }));
      }
      if (data.status) { // Assuming backend sends full state or partial updates
        game.update(g => ({ ...g, status: data.status }));
      }
      if (data.role) {
        game.update(g => ({
          ...g,
          me: { ...g.me, role: data.role, word: data.displayed_word }
        }));
      }
      // Add other handlers (player list updates etc)
    } catch (e) {
      console.error("Parse error", e);
    }
  };

  socket.onclose = () => {
    console.log("Disconnected");
    updateGame({ status: 'CONNECTING' });
  };
};

export const sendAction = (action: string, payload?: any) => {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify({ action, ...payload }));
  }
};
