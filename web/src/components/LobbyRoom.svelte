<script lang="ts">
    import { game, sendAction } from '../stores/game';

    export let lobbyId: string;

    let copied = false;

    function copyLink() {
        const url = `${window.location.origin}/?lobby=${lobbyId}`;
        navigator.clipboard.writeText(url);
        copied = true;
        setTimeout(() => copied = false, 2000);
    }

    function startGame() {
        sendAction("START_GAME");
    }
</script>

<div class="w-full max-w-4xl flex flex-col items-center">
    
    <!-- Header info -->
    <div class="bg-gray-800 rounded-full px-8 py-3 mb-8 flex items-center gap-4 shadow-lg border border-gray-600">
        <span class="text-gray-400">LOBBY CODE:</span>
        <span class="text-2xl font-mono text-white tracking-widest">{lobbyId}</span>
        <button 
            on:click={copyLink}
            class="ml-4 text-xs bg-gray-700 hover:bg-gray-600 px-3 py-1 rounded text-blue-300 transition"
        >
            {copied ? 'COPIED!' : 'COPY LINK'}
        </button>
    </div>

    <div class="flex flex-col md:flex-row w-full gap-8">
        <!-- Player List -->
        <div class="flex-1 bg-gray-800 rounded-xl p-6 shadow-xl border border-gray-700">
            <h3 class="text-xl font-bold mb-4 border-b border-gray-700 pb-2 text-white">
                Players ({$game.players.length})
            </h3>
            <ul class="space-y-3">
                {#each $game.players as player}
                    <li class="flex items-center justify-between p-3 rounded bg-gray-700/50 border border-gray-700">
                        <span class="font-medium text-gray-200">{player.name}</span>
                        <div class="flex items-center gap-2">
                            {#if player.is_leader}
                                <span class="text-xs bg-yellow-500/20 text-yellow-400 px-2 py-0.5 rounded border border-yellow-500/30">HOST</span>
                            {/if}
                            {#if player.id === $game.me.id}
                                <span class="text-xs bg-blue-500/20 text-blue-400 px-2 py-0.5 rounded border border-blue-500/30">YOU</span>
                            {/if}
                        </div>
                    </li>
                {/each}
            </ul>
        </div>

        <!-- Waiting / Start Area -->
        <div class="flex-1 flex flex-col items-center justify-center bg-gray-800/50 rounded-xl p-6 border border-gray-700 border-dashed">
            {#if $game.me.isLeader}
                <p class="text-gray-300 mb-6 text-center">You are the host.<br>Start when everyone is ready.</p>
                <button 
                    class="bg-green-600 hover:bg-green-500 text-white font-bold py-4 px-10 rounded-full shadow-lg transform transition hover:scale-105"
                    on:click={startGame}
                    disabled={$game.players.length < 3}
                >
                    START GAME
                </button>
                {#if $game.players.length < 3}
                    <p class="text-red-400 text-sm mt-4">Need at least 3 players</p>
                {/if}
            {:else}
                <div class="animate-pulse flex flex-col items-center">
                    <div class="w-16 h-16 border-4 border-t-blue-500 border-gray-600 rounded-full animate-spin mb-4"></div>
                    <p class="text-xl text-gray-300">Waiting for host...</p>
                </div>
            {/if}
        </div>
    </div>
</div>
