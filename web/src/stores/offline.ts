import { writable, get } from 'svelte/store';
import { getCookie, setCookie } from '../lib/cookie';

export type OfflinePhase = 'IDLE' | 'SETUP' | 'PASSING' | 'PLAYING' | 'FINISHED';

export interface OfflineConfig {
  playerCount: number;
  category: string;
  difficulty: 'NORMAL' | 'EASY';
  lang: string;
}

export interface Player {
  id: number;
  role: 'CIVILIAN' | 'IMPOSTOR';
  word: string; // Real word for civs, Trap (or nothing) for Impostor depending on diff
  isRevealed: boolean;
}

export interface OfflineGameState {
  status: OfflinePhase;
  config: OfflineConfig;
  players: Player[];
  currentPlayerIndex: number;
  wordPair: { real: string; trap: string };
  impostorIndex: number;
  startingPlayerId?: number;
  winner?: 'CIVILIAN' | 'IMPOSTOR';
}

const DEFAULT_CONFIG: OfflineConfig = {
  playerCount: 3,
  category: 'General',
  difficulty: 'NORMAL',
  lang: 'en'
};

const initialState: OfflineGameState = {
  status: 'IDLE',
  config: DEFAULT_CONFIG,
  players: [],
  currentPlayerIndex: 0,
  wordPair: { real: '', trap: '' },
  impostorIndex: -1
};

function createOfflineStore() {
  // Try to load from cookie
  let startState = initialState;
  const saved = getCookie('impostor_offline_game');
  if (saved) {
    try {
      startState = JSON.parse(saved);
    } catch (e) {
      console.error("Failed to parse offline cookie", e);
    }
  }

  const { subscribe, set, update } = writable<OfflineGameState>(startState);

  return {
    subscribe,

    // Setup & Start
    startGame: async (config: OfflineConfig) => {
      // Fetch word from API
      try {
        const res = await fetch(`http://localhost:8080/api/word?category=${encodeURIComponent(config.category)}&lang=${config.lang}`);
        if (!res.ok) throw new Error("Failed to fetch word");
        const wordPair = await res.json();

        const impostorIdx = Math.floor(Math.random() * config.playerCount);
        const startingPlayerId = Math.floor(Math.random() * config.playerCount) + 1;

        const players: Player[] = Array.from({ length: config.playerCount }, (_, i) => {
          const isImpostor = i === impostorIdx;
          let displayedWord = wordPair.real;

          if (isImpostor) {
            // In Easy mode, Impostor sees the Trap word (Hint).
            // In Normal mode, Impostor sees "Impostor" (handled by UI, but here we can set word null or special)
            displayedWord = config.difficulty === 'EASY' ? wordPair.trap : 'IMPOSTOR';
          }

          return {
            id: i + 1,
            role: isImpostor ? 'IMPOSTOR' : 'CIVILIAN',
            word: displayedWord,
            isRevealed: false
          };
        });

        const newState: OfflineGameState = {
          status: 'PASSING',
          config,
          players,
          currentPlayerIndex: 0,
          wordPair,
          impostorIndex: impostorIdx,
          startingPlayerId
        };

        set(newState);
        saveState(newState);

      } catch (e) {
        console.error(e);
        alert("Could not start game. Server unreachable?");
      }
    },

    // Phase Transitions
    nextPlayer: () => {
      update(s => {
        const nextIdx = s.currentPlayerIndex + 1;
        if (nextIdx >= s.players.length) {
          const nextState = { ...s, status: 'PLAYING' as OfflinePhase };
          saveState(nextState);
          return nextState;
        }
        const nextState = { ...s, currentPlayerIndex: nextIdx };
        saveState(nextState);
        return nextState;
      });
    },

    finishGame: () => {
      update(s => {
        const nextState = { ...s, status: 'FINISHED' as OfflinePhase };
        saveState(nextState);
        return nextState;
      });
    },

    restartGame: () => {
      // Restart with same config
      update(s => {
        // We basically need to call startGame again, but stores can't easily await async in update.
        // We'll reset to SETUP but keep config, and UI triggers start?
        // Or better: The UI calls `offline.startGame(currentConfig)` again.
        // Here we just provide a reset helper?
        return s; // No-op, handled by UI action
      });
    },

    reset: () => {
      set(initialState);
      setCookie('impostor_offline_game', '', -1);
    },

    startSetup: () => {
      update(s => ({ ...s, status: 'SETUP' }));
    },


    updateConfig: (cfg: Partial<OfflineConfig>) => {
      update(s => ({ ...s, config: { ...s.config, ...cfg } }));
    }
  };
}

// Helper to save state
function saveState(state: OfflineGameState) {
  setCookie('impostor_offline_game', JSON.stringify(state));
}

export const offline = createOfflineStore();
