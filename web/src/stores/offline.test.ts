// @vitest-environment happy-dom
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { get } from 'svelte/store';
import { offline } from './offline';

// Mock dependencies
global.fetch = vi.fn();
// Mock document.cookie
Object.defineProperty(global.document, 'cookie', {
  value: '',
  writable: true,
});

describe('Offline Store', () => {
  beforeEach(() => {
    offline.reset();
    vi.clearAllMocks();
    document.cookie = '';
  });

  it('should initialize in IDLE status', () => {
    const state = get(offline);
    expect(state.status).toBe('IDLE');
  });

  it('should start game correctly', async () => {
    const mockWord = { real: 'Cat', trap: 'Dog' };
    (global.fetch as any).mockResolvedValue({
      ok: true,
      json: async () => mockWord
    });

    await offline.startGame({
      playerCount: 3,
      category: 'Animals',
      difficulty: 'NORMAL',
      lang: 'en'
    });

    const state = get(offline);
    expect(state.status).toBe('PASSING');
    expect(state.players).toHaveLength(3);
    expect(state.wordPair).toEqual(mockWord);
    expect(state.impostorIndex).toBeGreaterThanOrEqual(0);
    expect(state.impostorIndex).toBeGreaterThanOrEqual(0);
    expect(state.impostorIndex).toBeLessThan(3);

    // Check starting player
    expect(state.startingPlayerId).toBeDefined();
    expect(state.startingPlayerId).toBeGreaterThanOrEqual(1);
    expect(state.startingPlayerId).toBeLessThanOrEqual(3);

    // check roles
    const impostorCount = state.players.filter(p => p.role === 'IMPOSTOR').length;
    expect(impostorCount).toBe(1);
  });

  it('should advance players', async () => {
    // Setup game manually or mock start
    const mockWord = { real: 'Cat', trap: 'Dog' };
    (global.fetch as any).mockResolvedValue({
      ok: true,
      json: async () => mockWord
    });
    await offline.startGame({ playerCount: 3, category: 'A', difficulty: 'NORMAL', lang: 'en' });

    // Player 1
    let state = get(offline);
    expect(state.currentPlayerIndex).toBe(0);

    // Next -> Player 2
    offline.nextPlayer();
    state = get(offline);
    expect(state.currentPlayerIndex).toBe(1);

    // Next -> Player 3
    offline.nextPlayer();
    state = get(offline);
    expect(state.currentPlayerIndex).toBe(2);

    // Next -> Playing Phase
    offline.nextPlayer();
    state = get(offline);
    expect(state.status).toBe('PLAYING');
  });

  it('should finish game', () => {
    offline.finishGame();
    expect(get(offline).status).toBe('FINISHED');
  });
});
