<script lang="ts">
    import { game } from '../stores/game';
    
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
    class="perspective-1000 w-64 h-96 cursor-pointer bg-transparent border-none p-0" 
    on:click={toggleFlip}
    on:keydown={(e) => e.key === 'Enter' && toggleFlip()}
>
    <div class="relative w-full h-full text-center transition-transform duration-700 transform-style-3d {isFlipped ? 'rotate-y-180' : ''}">
        
        <!-- Front -->
        <div class="absolute w-full h-full backface-hidden bg-gray-800 text-white flex items-center justify-center rounded-xl shadow-xl border-4 border-gray-600">
            <h2 class="text-3xl font-bold tracking-widest">SECRET</h2>
            <p class="absolute bottom-4 text-sm text-gray-400">Tap to reveal</p>
        </div>

        <!-- Back -->
        <div class="absolute w-full h-full backface-hidden rotate-y-180 bg-white text-gray-900 flex flex-col items-center justify-center rounded-xl shadow-xl border-8 ${borderColor}">
            <h2 class="text-xl font-bold mb-4 uppercase text-gray-500">Your Word</h2>
            <p class="text-4xl font-extrabold text-center px-2">{$game.me.word}</p>
            
            {#if $game.me.role === 'IMPOSTOR'}
                <div class="mt-8 px-4 py-2 bg-red-100 text-red-600 rounded-full font-bold animate-pulse">
                    IMPOSTOR
                </div>
            {/if}
        </div>
    </div>
</button>

<style>
    /* Custom utility classes if Tailwind plugin not present for 3d */
    .perspective-1000 { perspective: 1000px; }
    .transform-style-3d { transform-style: preserve-3d; }
    .backface-hidden { backface-visibility: hidden; }
    .rotate-y-180 { transform: rotateY(180deg); }
</style>
