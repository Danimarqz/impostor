import { writable } from 'svelte/store';

// Define the shape of our frontend state (mirroring Go structs)
export interface GameState {
  status: 'CONNECTING' | 'WAITING' | 'PLAYING' | 'VOTING' | 'FINISHED';
  lobbyId?: string;
  me: {
    id: string;
    name: string;
    isLeader: boolean;
    role?: 'CIVILIAN' | 'IMPOSTOR';
    word?: string;
  };
  players: Array<{ id: string; name: string; is_leader?: boolean; role?: string }>;
  messages: Array<{ from: string; text: string }>;
  winner?: string;
  kicked?: string;
  role_was?: string;
}

const initialState: GameState = {
  status: 'CONNECTING',
  lobbyId: '',
  me: { id: '', name: '', isLeader: false },
  players: [],
  messages: []
};

export const game = writable<GameState>(initialState);

export const updateGame = (newData: Partial<GameState>) => {
  game.update(current => ({ ...current, ...newData }));
};

let socket: WebSocket;

export const connect = (lobbyId: string, playerName: string) => {
  // Generate random ID for demo purposes if not persisted
  const playerId = Math.random().toString(36).substring(7);

  // Connect to Backend (dynamically determine host)
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
  const host = window.location.host;
  socket = new WebSocket(`${protocol}//${host}/ws/${lobbyId}?playerId=${playerId}&playerName=${playerName}`);

  socket.onopen = () => {
    console.log("Connected to WS");
    updateGame({
      status: 'WAITING',
      lobbyId: lobbyId,
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
      if (data.type === 'CHAT_MESSAGE') {
        game.update(g => ({ ...g, messages: [...g.messages, { from: data.from, text: data.text }] }));
      }
      if (data.type === 'PLAYER_LIST') {
        game.update(g => {
          // Check if I am now the leader
          const me = data.players.find((p: any) => p.id === g.me.id);
          return {
            ...g,
            players: data.players,
            me: me ? { ...g.me, isLeader: me.is_leader } : g.me
          };
        });
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
      if (data.type === 'VOTE_UPDATE') {
        // Store votes if we want to show counts
        // For now, let's just log it or update a 'votes' field in state
        console.log("Votes updated:", data.votes);
      }
      if (data.status === 'FINISHED') {
        game.update(g => ({
          ...g,
          status: 'FINISHED',
          winner: data.winner,
          kicked: data.kicked,
          role_was: data.role_was,
          players: data.reveal || g.players // Update players with roles if provided
        }));
      }
      if (data.type === 'GAME_RESET' || (data.status === 'WAITING' && !data.type)) {
        game.update(g => ({ ...g, status: 'WAITING', winner: undefined, kicked: undefined, role_was: undefined }));
      }
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
