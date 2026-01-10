<script lang="ts">
    import Card from './Card.svelte';
    import VotingPanel from './VotingPanel.svelte';
    import { game } from '../stores/game';
    import { language } from '../stores/language';
    import { t } from '../lib/i18n';
    
    $: lang = $language;
</script>

<div class="flex flex-col items-center justify-center w-full animate-fade-in p-4">
    {#if $game.status === 'FINISHED'}
        <div class="text-center animate-bounce-in mb-8 p-6 bg-gray-800 rounded-xl border border-gray-600 shadow-2xl z-50">
            <h1 class="text-5xl font-black text-transparent bg-clip-text bg-gradient-to-r from-yellow-400 to-red-600 mb-4">
                {t('game.gameOver', lang)}
            </h1>
            <h2 class="text-3xl font-bold text-white mb-2">
                {t('game.winner', lang)}: <span class="{ $game.winner === 'CIVILIANS' ? 'text-blue-400' : 'text-red-500' }">{$game.winner === 'CIVILIANS' ? t('game.civilians', lang) : t('game.impostor', lang)}</span>
            </h2>
            <p class="text-gray-400 mb-6">
                {$game.kicked} {t('game.wasThe', lang)} <span class="font-bold text-white">{$game.role_was === 'IMPOSTOR' ? t('game.impostor', lang) : 'CIVILIAN'}</span>
            </p>
            
            <div class="mt-4 border-t border-gray-600 pt-4">
                <h3 class="text-sm uppercase text-gray-500 mb-2">Identities Revealed</h3>
                <div class="grid grid-cols-2 gap-2">
                    {#each $game.players as p}
                        <div class="flex items-center justify-between bg-gray-700 px-3 py-1 rounded">
                            <span class="text-gray-200">{p.name}</span>
                            <span class="text-xs font-bold {p.role === 'IMPOSTOR' ? 'text-red-400' : 'text-blue-400'}">
                                {p.role || '???'}
                            </span>
                        </div>
                    {/each}
                </div>
            </div>
            
             <button 
                class="mt-6 bg-white text-gray-900 font-bold py-2 px-6 rounded-full hover:bg-gray-200 transition"
                on:click={() => {
                    // Send Reset Command
                    import('../stores/game').then(({ sendAction }) => {
                        sendAction("RESET_GAME", {});
                    });
                }}
            >
                {t('game.playAgain', lang)}
            </button>
        </div>
    {/if}

    {#if $game.status !== 'FINISHED'}
    <div class="w-full max-w-6xl flex flex-col items-center gap-8 animate-fade-in">
        <!-- Top Row: Card and Voting -->
        <div class="flex flex-col md:flex-row gap-8 w-full justify-center items-start">
             <!-- Card Area -->
             <div class="flex-1 flex flex-col items-center">
                <div class="bg-white/5 p-4 rounded-2xl backdrop-blur-sm border border-white/5">
                    <Card />
                </div>
                <div class="mt-4 text-center text-gray-500 text-xs uppercase tracking-widest">
                    <p>{t('game.tapToReveal', lang)}</p>
                </div>
             </div>

             <!-- Voting Area -->
             <div class="flex-1 flex justify-center w-full">
                 <VotingPanel />
             </div>
        </div>
    </div>
    {/if}
</div>

<style>
    .animate-bounce-in {
        animation: bounceIn 1s ease-out forwards;
    }
    @keyframes bounceIn {
        0% { opacity: 0; transform: scale(0.3) translateY(-100px); }
        50% { opacity: 1; transform: scale(1.05) translateY(10px); }
        70% { transform: scale(0.9) translateY(-5px); }
        100% { transform: scale(1) translateY(0); }
    }
</style>
