<script lang="ts">
    import { offline } from '../stores/offline';
    import { fade, scale } from 'svelte/transition';
    import { onMount, afterUpdate } from 'svelte';

    let revealed = false;
    let nextReady = false;

    import { t } from '../lib/i18n';
    
    $: currentPlayer = $offline.players[$offline.currentPlayerIndex];
    $: isImpostor = currentPlayer.role === 'IMPOSTOR';
    $: lang = $offline.config.lang as 'en' | 'es';

    // Reset revealed state when player index changes
    $: if ($offline.currentPlayerIndex !== undefined) {
        revealed = false;
    }

    function toggleReveal() {
        revealed = !revealed;
        if (revealed) {
            // Short delay to prevent accidental double clicks closing it immediately? 
            // Or just let user close it.
        }
    }

    function handleNext() {
        // Just call next player, the key block in HTML will reset the component state
        offline.nextPlayer();
    }
</script>

<div class="w-full max-w-md flex flex-col items-center">
    <h2 class="text-2xl font-bold text-white mb-8 tracking-widest uppercase">{t('offline.passTheDevice', lang)}</h2>

    {#key currentPlayer.id}
    <div 
        class="relative w-full aspect-[3/4] max-h-[500px] perspective-1000 group cursor-pointer outline-none focus:ring-4 focus:ring-purple-500/50 rounded-3xl" 
        on:click={toggleReveal}
        on:keydown={(e) => (e.key === 'Enter' || e.key === ' ') && toggleReveal()}
        role="button"
        tabindex="0"
    >
        <div class="relative w-full h-full duration-500 preserve-3d transition-transform {revealed ? 'rotate-y-180' : ''}">
            
            <!-- Front (Hidden) -->
            <div class="absolute inset-0 backface-hidden bg-gradient-to-br from-indigo-900 to-purple-900 rounded-3xl shadow-2xl border border-white/10 flex flex-col items-center justify-center p-8">
                <div class="text-6xl mb-4">🤫</div>
                <h3 class="text-3xl font-bold text-white mb-2">{t('offline.player', lang)} {currentPlayer.id}</h3>
                <p class="text-blue-200 text-center text-sm">{t('offline.tapToSee', lang)}</p>
                <p class="mt-8 text-white/30 text-xs uppercase tracking-widest">{t('offline.keepSecret', lang)}</p>
            </div>

            <!-- Back (Revealed) -->
            <div class="absolute inset-0 backface-hidden rotate-y-180 bg-white rounded-3xl shadow-2xl flex flex-col items-center justify-center p-8 border-4 border-purple-500">
                <p class="text-gray-500 text-sm font-bold uppercase tracking-widest mb-2">{t('offline.yourRole', lang)}</p>
                 <h2 class="text-4xl font-black mb-6 {isImpostor ? 'text-red-600' : 'text-blue-600'}">
                    {isImpostor ? t('game.impostor', lang) : t('game.civilians', lang)}
                </h2>

                <div class="w-full h-px bg-gray-200 my-4"></div>

                <p class="text-gray-500 text-sm font-bold uppercase tracking-widest mb-2">{t('offline.secretWord', lang)}</p>
                <p class="text-3xl font-bold text-gray-800 text-center break-words">
                    {currentPlayer.word}
                </p>
                
                {#if isImpostor && $offline.config.difficulty === 'NORMAL'}
                     <p class="mt-4 text-xs text-red-500 text-center font-bold">{t('offline.blendIn', lang)}</p>
                {/if}
            </div>
        </div>
    </div>
    {/key}

    {#if revealed}
        <button 
            in:scale
            on:click|stopPropagation={handleNext}
            class="mt-8 px-12 py-4 bg-white text-purple-900 font-bold rounded-full shadow-lg hover:bg-gray-100 active:scale-95 transition-all uppercase tracking-wider"
        >
            {t('offline.nextPlayer', lang)}
        </button>
    {/if}
</div>

<style>
    .perspective-1000 { perspective: 1000px; }
    .preserve-3d { transform-style: preserve-3d; }
    .backface-hidden { backface-visibility: hidden; }
    .rotate-y-180 { transform: rotateY(180deg); }
</style>
