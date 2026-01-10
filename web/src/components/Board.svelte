<script lang="ts">
    import { game } from '../stores/game';
    import Landing from './Landing.svelte';
    import LobbyRoom from './LobbyRoom.svelte';
    import GameView from './GameView.svelte';
</script>

<div class="flex flex-col items-center justify-center min-h-screen bg-gray-900 text-white p-4">
    <header class="mb-8 text-center">
        <h1 class="text-4xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-red-500 to-purple-600">
            EL IMPOSTOR
        </h1>
        {#if $game.status !== 'CONNECTING'}
           <p class="text-gray-500 text-xs mt-1 uppercase tracking-widest">{$game.status}</p>
        {/if}
    </header>

    <main class="w-full flex justify-center">
        {#if $game.status === 'CONNECTING'}
            <Landing />
        {:else if $game.status === 'WAITING'}
            <LobbyRoom lobbyId={$game.lobbyId || '???'} />
        {:else if $game.status === 'PLAYING'}
            <GameView />
        {/if}
    </main>
</div>
