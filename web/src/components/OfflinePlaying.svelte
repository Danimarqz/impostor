<script lang="ts">
    import { offline } from '../stores/offline';
    import { t } from '../lib/i18n';

    $: lang = $offline.config.lang as 'en' | 'es';
    
    function finishGame() {
        offline.finishGame();
    }
</script>

<div class="w-full max-w-md flex flex-col items-center text-center">
    <div class="mb-8 relative">
         <div class="absolute -inset-4 bg-blue-500/20 rounded-full blur-xl animate-pulse"></div>
         <span class="text-8xl relative z-10">🕵️‍♂️</span>
    </div>

    <h2 class="text-3xl font-bold text-white mb-4">{t('offline.missionProgress', lang)}</h2>
    <p class="text-gray-300 mb-8 max-w-xs">
        {t('offline.askQuestions', lang)}
    </p>

    <div class="bg-white/5 border border-white/10 rounded-2xl p-6 w-full mb-8 backdrop-blur-sm">
        <p class="text-xs font-bold text-gray-500 uppercase tracking-widest mb-2">{t('offline.category', lang)}</p>
        <p class="text-2xl font-bold text-white mb-4">{$offline.config.category}</p>

        {#if $offline.startingPlayerId}
            <div class="pt-4 border-t border-white/10">
                <p class="text-lg text-yellow-400 font-bold animate-pulse">
                    {t('offline.player', lang)} {$offline.startingPlayerId} {t('offline.starts', lang)}
                </p>
            </div>
        {/if}
    </div>

    <button 
        on:click={finishGame}
        class="w-full py-4 bg-red-600/90 text-white font-bold rounded-xl shadow-lg border border-red-500 hover:bg-red-500 active:scale-95 transition-all uppercase tracking-wider"
    >
        {t('offline.identifyImpostor', lang)}
    </button>
</div>
