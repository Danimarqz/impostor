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
  'landing.enterCredentials': { en: 'Enter your credentials to access the system.', es: 'Introduce tus credenciales para acceder al sistema.' },
  'landing.playOffline': { en: 'PLAY OFFLINE', es: 'JUGAR OFFLINE' },

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

  // Offline Mode
  'offline.passTheDevice': { en: 'Pass the Device', es: 'Pasa el Dispositivo' },
  'offline.player': { en: 'Player', es: 'Jugador' },
  'offline.tapToSee': { en: 'Tap card to see your secret word', es: 'Toca para ver tu palabra secreta' },
  'offline.keepSecret': { en: 'Keep it secret!', es: '¡Mantenlo en secreto!' },
  'offline.yourRole': { en: 'Your Role', es: 'Tu Rol' },
  'offline.secretWord': { en: 'Secret Word', es: 'Palabra Secreta' },
  'offline.nextPlayer': { en: 'Next Player', es: 'Siguiente Jugador' },
  'offline.missionProgress': { en: 'Mission in Progress', es: 'Misión en Progreso' },
  'offline.askQuestions': { en: "Ask questions. Find the Impostor. Don't reveal your word.", es: "Haz preguntas. Encuentra al Impostor. No reveles tu palabra." },
  'offline.category': { en: 'Category', es: 'Categoría' },
  'offline.identifyImpostor': { en: 'Identify Impostor (Finish)', es: 'Identificar Impostor (Finalizar)' },
  'offline.missionReport': { en: 'Mission Report', es: 'Reporte de Misión' },
  'offline.impostorWas': { en: 'The Impostor Was', es: 'El Impostor Era' },
  'offline.civilianWord': { en: 'Civilian Word', es: 'Palabra Civil' },
  'offline.impostorTrap': { en: 'Impostor Trap', es: 'Trampa de Impostor' },
  'offline.playAgain': { en: 'Play Again', es: 'Jugar de Nuevo' },
  'offline.backToMenu': { en: 'Back to Menu', es: 'Volver al Menú' },
  'offline.blendIn': { en: "Blend in. Don't get caught.", es: "Disimula. Que no te atrapen." },
  'offline.startGame': { en: 'START GAME', es: 'INICIAR JUEGO' },
  'offline.setup': { en: 'Offline Setup', es: 'Configuración Offline' },
  'offline.players': { en: 'Players', es: 'Jugadores' },
  'offline.difficulty': { en: 'Difficulty', es: 'Dificultad' },
  'offline.normal': { en: 'NORMAL', es: 'NORMAL' },
  'offline.easy': { en: 'EASY', es: 'FÁCIL' },
  'offline.diffNormalDesc': { en: 'Impostor only knows they are the Impostor.', es: 'El Impostor solo sabe que es el Impostor.' },
  'offline.diffEasyDesc': { en: 'Impostor gets a hint (the trap word).', es: 'El Impostor recibe una pista (la palabra trampa).' },
  'offline.starts': { en: 'starts the discussion!', es: '¡empieza la discusión!' },
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
