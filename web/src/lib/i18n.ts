// Simple i18n system for bilingual support (EN/ES)

export type Language = 'en' | 'es';

interface Translations {
  [key: string]: {
    en: string;
    es: string;
  };
}

const translations: Translations = {
  // Landing Page
  'landing.title': { en: 'Welcome, Agent', es: 'Bienvenido, Agente' },
  'landing.nameLabel': { en: 'Your Code Name', es: 'Tu Nombre en Clave' },
  'landing.lobbyLabel': { en: 'Mission Code (Optional)', es: 'Código de Misión (Opcional)' },
  'landing.joinButton': { en: 'JOIN THE MISSION', es: 'UNIRSE A LA MISIÓN' },

  // Lobby Room
  'lobby.operationId': { en: 'Operation ID', es: 'ID de Operación' },
  'lobby.copyCode': { en: 'COPY ACCESS CODE', es: 'COPIAR CÓDIGO DE ACCESO' },
  'lobby.copied': { en: 'COPIED TO CLIPBOARD', es: 'COPIADO AL PORTAPAPELES' },
  'lobby.squadMembers': { en: 'Squad Members', es: 'Miembros del Escuadrón' },
  'lobby.host': { en: 'HOST', es: 'ANFITRIÓN' },
  'lobby.you': { en: 'YOU', es: 'TÚ' },
  'lobby.selectSector': { en: 'Select Target Sector', es: 'Seleccionar Sector Objetivo' },
  'lobby.difficulty': { en: 'Mission Difficulty', es: 'Dificultad de Misión' },
  'lobby.hard': { en: 'HARD', es: 'DIFÍCIL' },
  'lobby.easy': { en: 'EASY', es: 'FÁCIL' },
  'lobby.trapDetected': { en: '>> TRAP WORD DETECTED FOR IMPOSTOR', es: '>> PALABRA TRAMPA DETECTADA PARA IMPOSTOR' },
  'lobby.blindMode': { en: '>> BLIND MODE ACTIVE', es: '>> MODO CIEGO ACTIVO' },
  'lobby.systemsReady': { en: 'All systems nominal. Awaiting command.', es: 'Todos los sistemas nominales. Esperando comando.' },
  'lobby.initiateLaunch': { en: 'INITIATE LAUNCH', es: 'INICIAR LANZAMIENTO' },
  'lobby.minPlayers': { en: 'Minimum 3 operatives required', es: 'Se requieren mínimo 3 operativos' },
  'lobby.standby': { en: 'STANDBY', es: 'EN ESPERA' },
  'lobby.waitingHost': { en: 'Waiting for mission commander to finalize parameters...', es: 'Esperando que el comandante finalice los parámetros...' },
  'lobby.language': { en: 'Language', es: 'Idioma' },

  // Game View
  'game.tapToReveal': { en: 'Tap card to hide/reveal', es: 'Toca la carta para ocultar/revelar' },
  'game.gameOver': { en: 'GAME OVER', es: 'FIN DEL JUEGO' },
  'game.winner': { en: 'WINNER', es: 'GANADOR' },
  'game.civilians': { en: 'CIVILIANS', es: 'CIVILES' },
  'game.impostor': { en: 'IMPOSTOR', es: 'IMPOSTOR' },
  'game.eliminated': { en: 'was eliminated', es: 'fue eliminado' },
  'game.wasThe': { en: 'They were the', es: 'Era' },
  'game.playAgain': { en: 'PLAY AGAIN (KEEP LOBBY)', es: 'JUGAR DE NUEVO (MANTENER LOBBY)' },

  // Voting Panel
  'vote.title': { en: 'VOTE TO ELIMINATE', es: 'VOTA PARA ELIMINAR' },
  'vote.eliminate': { en: 'ELIMINATE', es: 'ELIMINAR' },
  'vote.tapToVote': { en: 'Vote', es: 'Votar' },
  'vote.waiting': { en: 'Waiting for others...', es: 'Esperando a los demás...' },
  'vote.selectImpostor': { en: 'Select the Impostor carefully!', es: '¡Selecciona al Impostor con cuidado!' },

  // Board/Header
  'header.phase': { en: 'PHASE', es: 'FASE' },
  'header.connecting': { en: 'CONNECTING', es: 'CONECTANDO' },
  'header.waiting': { en: 'WAITING', es: 'EN ESPERA' },
  'header.playing': { en: 'PLAYING', es: 'DE JUEGO' },
  'header.voting': { en: 'VOTING', es: 'DE VOTACIÓN' },
  'header.finished': { en: 'FINISHED', es: 'FINALIZADA' },

  // Card Component
  'card.secret': { en: 'SECRET', es: 'SECRETO' },
  'card.tapToReveal': { en: 'Tap to reveal role', es: 'Toca para revelar rol' },
  'card.yourMission': { en: 'YOUR MISSION', es: 'TU MISIÓN' },
  'card.secretWord': { en: 'SECRET WORD', es: 'PALABRA SECRETA' },
  'card.impostor': { en: 'Impostor', es: 'Impostor' },
  'card.civilian': { en: 'Civilian', es: 'Civil' },

  // Chat Component
  'chat.title': { en: 'Mission Comms', es: 'Comunicaciones de Misión' },
  'chat.startConversation': { en: 'Start the conversation...', es: 'Inicia la conversación...' },
  'chat.typeMessage': { en: 'Type a message...', es: 'Escribe un mensaje...' },
  'chat.send': { en: 'SEND', es: 'ENVIAR' },
};

export function t(key: string, lang: Language = 'en'): string {
  const translation = translations[key];
  if (!translation) {
    console.warn(`Translation missing for key: ${key}`);
    return key;
  }
  return translation[lang];
}

// Helper to get all translation keys
export function getAllKeys(): string[] {
  return Object.keys(translations);
}
