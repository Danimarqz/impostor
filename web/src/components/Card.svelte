<script lang="ts">
    import { game } from '../stores/game';
    import { language } from '../stores/language';
    import { t } from '../lib/i18n';
    
    $: lang = $language;
    let isFlipped = false;

    // Reactive statement: Calculate card color based on role
    $: borderColor = $game.me.role === 'IMPOSTOR' ? 'border-red-500' : 'border-blue-500';

    function toggleFlip() {
        if ($game.status === 'PLAYING') {
            isFlipped = !isFlipped;
        }
    }
</script>

<!-- Using Tailwind for 3D Transform primitives (requires custom utilities or arbitrary values usually, 
     but here we use standard classes for layout and verify logic first) -->
<!-- For valid 3D flip in Tailwind, we usually need 'perspective' on container and 'preserve-3d' on children -->

<button 
    type="button" 
    class="group perspective-1000 w-64 h-96 cursor-pointer bg-transparent border-none p-0 focus:outline-none" 
    on:click={toggleFlip}
    on:keydown={(e) => e.key === 'Enter' && toggleFlip()}
    disabled={$game.status !== 'PLAYING' && $game.status !== 'FINISHED'}
>
    <!-- Card Inner Container -->
    <div class="relative w-full h-full text-center transition-transform duration-700 transform-style-3d {isFlipped ? 'rotate-y-180' : ''}">
        
        <!-- Front (Secret) -->
        <div class="absolute w-full h-full backface-hidden bg-gradient-to-br from-gray-800 to-gray-900 border-2 border-gray-600 rounded-xl shadow-2xl flex flex-col items-center justify-center p-4">
            <div class="w-16 h-16 mb-4 rounded-full bg-gray-700 flex items-center justify-center border border-gray-500">
                <span class="text-3xl">🤫</span>
            </div>
            <h2 class="text-3xl font-bold tracking-widest text-white mb-2">{t('card.secret', lang)}</h2>
            <p class="text-sm text-gray-400">{t('card.tapToReveal', lang)}</p>
        </div>

        <!-- Back (Revealed) -->
        <div class="absolute w-full h-full backface-hidden rotate-y-180 bg-white rounded-xl shadow-2xl flex flex-col items-center justify-center border-4 ${borderColor} p-6 overflow-hidden">
             <!-- Background Texture/Icon -->
             <div class="absolute inset-0 opacity-10 flex items-center justify-center pointer-events-none">
                 <span class="text-9xl">{$game.me.role === 'IMPOSTOR' ? '🔪' : '👥'}</span>
             </div>

            <h2 class="relative text-sm font-bold mb-4 uppercase text-gray-500 tracking-wider">
                {$game.me.role === 'IMPOSTOR' ? t('card.yourMission', lang) : t('card.secretWord', lang)}
            </h2>
            
            <p class="relative text-3xl font-extrabold text-center text-gray-900 break-words leading-tight">
                {$game.me.word}
            </p>
            
            {#if $game.me.role === 'IMPOSTOR'}
                <div class="relative mt-8 px-4 py-1 bg-red-100 text-red-600 rounded-full text-xs font-bold uppercase tracking-widest animate-pulse border border-red-200">
                    {t('card.impostor', lang)}
                </div>
            {:else}
                 <div class="relative mt-8 px-4 py-1 bg-blue-100 text-blue-600 rounded-full text-xs font-bold uppercase tracking-widest border border-blue-200">
                    {t('card.civilian', lang)}
                </div>
            {/if}
        </div>
    </div>
</button>

<style>
    .perspective-1000 { perspective: 1000px; }
    .transform-style-3d { transform-style: preserve-3d; }
    .backface-hidden { backface-visibility: hidden; }
    .rotate-y-180 { transform: rotateY(180deg); }
</style>
