<script lang="ts">
    import { onMount } from 'svelte';
    import { connect } from '../stores/game';
    import { offline } from '../stores/offline';
    import { language } from '../stores/language';
    import { fade } from 'svelte/transition';
    import { t } from '../lib/i18n';

    $: lang = $language;

    let playerName = '';
    let lobbyCode = '';

    onMount(() => {
        // Auto-fill lobby from URL query param
        const params = new URLSearchParams(window.location.search);
        const lobbyParam = params.get('lobby');
        if (lobbyParam) {
            lobbyCode = lobbyParam;
        }
    });

    async function handleJoin() {
        if (!playerName) return alert("Enter your name!");
        
        let finalLobby = lobbyCode;
        
        if (!finalLobby) {
            // Create New Game via API
            try {
                const res = await fetch('http://localhost:8080/api/lobby', { method: 'POST' });
                const data = await res.json();
                finalLobby = data.lobby_id;
            } catch (e) {
                console.error(e);
                return alert("Error creating lobby");
            }
        }
        
        connect(finalLobby, playerName);
    }
</script>

<div class="flex flex-col items-center justify-center w-full max-w-md p-8 bg-white/5 backdrop-blur-xl rounded-3xl shadow-2xl border border-white/10 relative overflow-hidden group">
    <!-- Glow effect -->
    <div class="absolute -top-10 -right-10 w-40 h-40 bg-blue-500/30 rounded-full blur-3xl group-hover:bg-blue-500/50 transition duration-1000"></div>
    <div class="absolute -bottom-10 -left-10 w-40 h-40 bg-purple-500/30 rounded-full blur-3xl group-hover:bg-purple-500/50 transition duration-1000"></div>

    

    <h2 class="text-3xl font-bold mb-2 text-white relative z-10">{t('landing.title', lang)}</h2>
<div class="flex justify-center mt-2">
  <button
    class="relative bg-black/60 rounded-full p-1 flex items-center w-24 border border-white/20 shadow-lg group cursor-pointer backdrop-blur-sm"
    on:click={() => language.update(l => l === 'en' ? 'es' : 'en')}
  >
    <div class="absolute w-1/2 h-full top-0 left-0 bg-gradient-to-r from-amber-700 to-amber-600 rounded-full transition-transform duration-300 shadow-md {$language === 'es' ? 'translate-x-[100%]' : 'translate-x-0'}"></div>
    <span class="relative z-10 w-1/2 text-center text-xs font-bold transition-colors duration-300 py-1 {$language === 'es' ? 'text-gray-400' : 'text-white'}">EN</span>
    <span class="relative z-10 w-1/2 text-center text-xs font-bold transition-colors duration-300 py-1 {$language === 'es' ? 'text-white' : 'text-gray-400'}">ES</span>
  </button>
</div>
    <p class="text-gray-400 mb-8 text-sm text-center relative z-10">{t('landing.enterCredentials', lang)}</p>
    
    <div class="w-full space-y-5 relative z-10">
        <div class="space-y-1">
            <label for="name" class="block text-xs font-bold text-gray-300 uppercase tracking-widest ml-1">{t('landing.nameLabel', lang)}</label>
            <input 
                id="name"
                bind:value={playerName}
                type="text" 
                class="w-full px-4 py-3 bg-black/40 border border-white/10 rounded-xl text-white placeholder-gray-500 focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all duration-300 hover:bg-black/50"
                placeholder="e.g. Maverick"
            />
        </div>

        <div class="space-y-1">
            <label for="lobby" class="block text-xs font-bold text-gray-300 uppercase tracking-widest ml-1">{t('landing.lobbyLabel', lang)}</label>
            <input 
                id="lobby"
                bind:value={lobbyCode}
                type="text" 
                class="w-full px-4 py-3 bg-black/40 border border-white/10 rounded-xl text-white placeholder-gray-500 focus:ring-2 focus:ring-purple-500 focus:border-transparent outline-none transition-all duration-300 hover:bg-black/50"
                placeholder="Leave blank to create new"
            />
        </div>

        <button 
            on:click={handleJoin}
            class="w-full py-4 mt-6 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-bold rounded-xl shadow-lg hover:shadow-blue-500/30 hover:from-blue-500 hover:to-purple-500 transform transition-all duration-300 hover:scale-[1.02] active:scale-95 uppercase tracking-wider flex items-center justify-center gap-2"
        >
            <span>{t('landing.joinButton', lang)}</span>
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z" clip-rule="evenodd" />
            </svg>
        </button>

        {#if !lobbyCode}
        <button 
            transition:fade
            on:click={() => offline.startSetup()}
            class="w-full py-4 bg-gradient-to-r from-gray-700 to-gray-600 text-white font-bold rounded-xl shadow-lg hover:shadow-gray-500/30 hover:from-gray-600 hover:to-gray-500 transform transition-all duration-300 hover:scale-[1.02] active:scale-95 uppercase tracking-wider flex items-center justify-center gap-2"
        >
            <span>{t('landing.playOffline', lang)}</span>
        </button>
        {/if}
    </div>
</div>
