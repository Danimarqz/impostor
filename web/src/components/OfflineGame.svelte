<script lang="ts">
    import { offline } from '../stores/offline';
    import OfflineSetup from './OfflineSetup.svelte';
    import OfflinePassAndPlay from './OfflinePassAndPlay.svelte';
    import OfflinePlaying from './OfflinePlaying.svelte';
    import OfflineResult from './OfflineResult.svelte';
    import { fly } from 'svelte/transition';
</script>

<div class="w-full flex justify-center items-center" in:fly={{ y: 20, duration: 500 }}>
    {#if $offline.status === 'SETUP'}
        <OfflineSetup />
    {:else if $offline.status === 'PASSING'}
        <OfflinePassAndPlay />
    {:else if $offline.status === 'PLAYING'}
        <OfflinePlaying />
    {:else if $offline.status === 'FINISHED'}
        <OfflineResult />
    {/if}
</div>
