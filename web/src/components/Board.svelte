<script lang="ts">
    import { onMount } from 'svelte';
    import { game, connect, sendAction } from '../stores/game';
    import Card from './Card.svelte';

    onMount(() => {
        // Auto-connect for MVP demo. Real app would have a Join form.
        connect("lobby1", "Player" + Math.floor(Math.random()*1000));
    });

    function startGame() {
        sendAction("START_GAME");
    }
</script>

<div class="flex flex-col items-center justify-center min-h-screen bg-gray-900 text-white p-4">
    <header class="mb-8 text-center">
        <h1 class="text-4xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-red-500 to-purple-600">
            EL IMPOSTOR
        </h1>
        <p class="text-gray-400 mt-2">Status: {$game.status}</p>
    </header>

    <main class="w-full max-w-4xl flex flex-col md:flex-row items-center gap-8">
        
        <!-- Player List -->
        <div class="w-full md:w-1/3 bg-gray-800 rounded-lg p-6 shadow-lg border border-gray-700">
            <h3 class="text-xl font-bold mb-4 border-b border-gray-700 pb-2">Players</h3>
            <ul class="space-y-2">
                {#each $game.players as player}
                    <li class="flex items-center justify-between p-2 rounded bg-gray-700">
                        <span>{player.name}</span>
                        {#if player.id === $game.me.id}
                            <span class="text-xs bg-blue-500 px-2 py-1 rounded">YOU</span>
                        {/if}
                    </li>
                {/each}
            </ul>
        </div>

        <!-- Game Area -->
        <div class="w-full md:w-2/3 flex flex-col items-center">
            
            {#if $game.status === 'WAITING'}
                <div class="text-center py-10">
                    <p class="text-xl mb-6">Waiting for host to start...</p>
                    {#if $game.me.isLeader} 
                        <button 
                            class="bg-green-600 hover:bg-green-500 text-white font-bold py-3 px-8 rounded-full shadow-lg transform transition hover:scale-105"
                            on:click={startGame}
                        >
                            START GAME
                        </button>
                    {/if}
                </div>
            {:else if $game.status === 'PLAYING'}
                <Card />
            {/if}

        </div>
    </main>
</div>
