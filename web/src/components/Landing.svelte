<script lang="ts">
    import { onMount } from 'svelte';
    import { connect } from '../stores/game';

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

    function handleJoin() {
        if (!playerName) return alert("Enter your name!");
        // If no lobby code, generate random one (Create Game)
        const finalLobby = lobbyCode || Math.random().toString(36).substring(7);
        connect(finalLobby, playerName);
    }
</script>

<div class="flex flex-col items-center justify-center w-full max-w-md p-8 bg-gray-800 rounded-xl shadow-2xl border border-gray-700">
    <h2 class="text-3xl font-bold mb-8 text-white">Join the Game</h2>
    
    <div class="w-full space-y-4">
        <div>
            <label for="name" class="block text-sm font-medium text-gray-400 mb-1">Nickname</label>
            <input 
                id="name"
                bind:value={playerName}
                type="text" 
                class="w-full px-4 py-3 bg-gray-900 border border-gray-600 rounded-lg text-white focus:ring-2 focus:ring-blue-500 outline-none transition"
                placeholder="Ex: Impostor99"
            />
        </div>

        <div>
            <label for="lobby" class="block text-sm font-medium text-gray-400 mb-1">Lobby Code (Optional)</label>
            <input 
                id="lobby"
                bind:value={lobbyCode}
                type="text" 
                class="w-full px-4 py-3 bg-gray-900 border border-gray-600 rounded-lg text-white focus:ring-2 focus:ring-purple-500 outline-none transition"
                placeholder="Leave empty to create new"
            />
        </div>

        <button 
            on:click={handleJoin}
            class="w-full py-4 mt-4 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-bold rounded-lg shadow-lg hover:from-blue-500 hover:to-purple-500 transform transition hover:scale-[1.02]"
        >
            {lobbyCode ? 'JOIN LOBBY' : 'CREATE NEW GAME'}
        </button>
    </div>
</div>
