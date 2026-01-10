<script lang="ts">
    import { game, sendAction } from '../stores/game';
    import { language } from '../stores/language';
    import { t } from '../lib/i18n';

    $: lang = $language;

    let votedTargetId: string | null = null;

    function vote(targetId: string) {
        // Toggle logic handled in UI state, always send action
        sendAction("CAST_VOTE", { target_id: targetId });
    }

    // Computed: Count votes per player (if we had full data)
    // For MVP transparency, backend needs to send vote updates.
    // Assuming game.votes is a map { voterId: targetId }
</script>

<div class="bg-gray-800 rounded-xl p-6 shadow-xl border border-gray-700 w-full max-w-md">
    <h3 class="text-xl font-bold mb-6 text-white text-center border-b border-gray-600 pb-2">
        {t('vote.title', lang)}
    </h3>
    
    <div class="space-y-4">
        {#each $game.players as player}
            {#if player.id !== $game.me.id}
                <button 
                    class="w-full flex items-center justify-between p-4 rounded-lg transition-all border
                           {votedTargetId === player.id 
                             ? 'bg-red-900/50 border-red-500 shadow-red-500/20 shadow-lg' 
                             : 'bg-gray-700 hover:bg-gray-600 border-gray-600'}"
                    on:click={() => {
                        if (votedTargetId === player.id) {
                            votedTargetId = null; // Toggle off locally
                        } else {
                            votedTargetId = player.id;
                        }
                        vote(player.id);
                    }}
                >
                    <span class="font-bold text-lg text-gray-200">{player.name}</span>
                    {#if votedTargetId === player.id}
                        <span class="text-red-400 font-bold text-sm tracking-wider">{t('vote.eliminate', lang)}</span>
                    {:else}
                         <span class="text-gray-500 text-xs uppercase tracking-widest">{t('vote.tapToVote', lang)}</span>
                    {/if}
                </button>
            {/if}
        {/each}
    </div>
    
    <p class="text-center text-gray-500 text-sm mt-6">
        {votedTargetId ? t('vote.waiting', lang) : t('vote.selectImpostor', lang)}
    </p>
</div>
