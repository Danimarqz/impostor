<script lang="ts">
    import { offline } from '../stores/offline';
    import { language } from '../stores/language';
    import { t } from '../lib/i18n';

    $: impostor = $offline.players[$offline.impostorIndex];
    $: wordPair = $offline.wordPair;
    $: lang = $offline.config.lang as 'en' | 'es';

    function handlePlayAgain() {
        // Re-use same settings, just fetch new word and assign roles
        offline.startGame($offline.config);
    }

    function handleMenu() {
        offline.reset();
    }
</script>

<div class="w-full max-w-md flex flex-col items-center text-center">
    <h2 class="text-4xl font-black text-white mb-2 uppercase tracking-tight">{t('offline.missionReport', lang)}</h2>
    
    <div class="my-8 w-full bg-white/5 border border-white/10 rounded-3xl p-8 backdrop-blur-xl">
        <p class="text-xs font-bold text-gray-500 uppercase tracking-widest mb-4">{t('offline.impostorWas', lang)}</p>
        
        <div class="flex flex-col items-center">
             <div class="w-20 h-20 bg-red-600 rounded-full flex items-center justify-center text-3xl font-bold text-white mb-4 shadow-xl border-4 border-red-400">
                {impostor.id}
             </div>
             <h3 class="text-3xl font-bold text-white">{t('offline.player', lang)} {impostor.id}</h3>
        </div>

        <div class="mt-8 grid grid-cols-2 gap-4">
            <div class="bg-black/30 rounded-xl p-4">
                <p class="text-[10px] font-bold text-gray-400 uppercase tracking-widest mb-1">{t('offline.civilianWord', lang)}</p>
                <p class="text-lg font-bold text-blue-400">{wordPair.real}</p>
            </div>
            <div class="bg-black/30 rounded-xl p-4">
                <p class="text-[10px] font-bold text-gray-400 uppercase tracking-widest mb-1">{t('offline.impostorTrap', lang)}</p>
                <p class="text-lg font-bold text-red-400">{wordPair.trap}</p>
            </div>
        </div>
    </div>

    <div class="w-full space-y-4">
        <button 
            on:click={handlePlayAgain}
            class="w-full py-4 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-bold rounded-xl shadow-lg hover:shadow-blue-500/30 active:scale-95 transition-all uppercase tracking-wider"
        >
            {t('offline.playAgain', lang)}
        </button>

        <button 
            on:click={handleMenu}
            class="w-full py-4 bg-transparent border border-white/20 text-white/70 font-bold rounded-xl hover:bg-white/5 hover:text-white active:scale-95 transition-all uppercase tracking-wider"
        >
            {t('offline.backToMenu', lang)}
        </button>
    </div>
</div>
