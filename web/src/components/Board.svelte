<script lang="ts">
    import { game } from '../stores/game';
    import Landing from './Landing.svelte';
    import LobbyRoom from './LobbyRoom.svelte';
    import GameView from './GameView.svelte';
    import Chat from './Chat.svelte';
    import { fade, fly } from 'svelte/transition';
    import { language } from '../stores/language';
    import { t } from '../lib/i18n';

    $: lang = $language;

    let showMobileChat = false;
    let unreadCount = 0;
    let lastMsgCount = 0;

    // React to new messages for unread count
    $: {
        if ($game.messages.length > lastMsgCount) {
             if (!showMobileChat) {
                 unreadCount += ($game.messages.length - lastMsgCount);
             }
             lastMsgCount = $game.messages.length;
        }
    }

    // Reset unread when opening
    $: if (showMobileChat) unreadCount = 0;
</script>

<div class="flex min-h-screen w-full relative overflow-hidden">
    <!-- Main Content Area -->
    <div class="flex-1 flex flex-col items-center p-4 relative z-10 overflow-y-auto h-screen">
        <header class="mb-12 text-center animate-fade-in-down mt-8">
            <h1 class="text-6xl md:text-8xl font-black text-transparent bg-clip-text bg-gradient-to-r from-blue-400 via-purple-500 to-pink-500 drop-shadow-lg tracking-tighter">
                IMPOSTOR
            </h1>
            {#if $game.status !== 'CONNECTING'}
            <div class="mt-4 px-4 py-1 bg-white/10 backdrop-blur-md rounded-full inline-block border border-white/10">
                    <p class="text-blue-200 text-xs font-bold uppercase tracking-[0.2em] shadow-sm">
                        {$game.status} {t('header.phase', lang)}
                    </p>
            </div>
            {/if}
        </header>

        <main class="w-full flex justify-center perspective-1000 pb-20">
            {#if $game.status === 'CONNECTING'}
                <Landing />
            {:else if $game.status === 'WAITING'}
                <LobbyRoom lobbyId={$game.lobbyId || '???'} />
            {:else if ['PLAYING', 'VOTING', 'FINISHED'].includes($game.status)}
                <GameView />
            {/if}
        </main>
    </div>

    <!-- Fixed Sidebar Chat (Desktop) -->
    {#if $game.status !== 'CONNECTING'}
        <aside class="hidden md:flex w-80 h-screen sticky top-0 border-l border-white/10 bg-black/20 backdrop-blur-md flex-col p-4 shadow-2xl z-20">
            <h3 class="text-white/50 text-xs font-bold uppercase tracking-widest mb-4">{t('chat.title', lang)}</h3>
            <div class="flex-1 overflow-hidden flex flex-col">
                <Chat />
            </div>
        </aside>

        <!-- Mobile Floating Action Button (FAB) -->
        <button 
            class="md:hidden fixed bottom-6 right-6 w-14 h-14 bg-blue-600 rounded-full shadow-2xl z-50 flex items-center justify-center border border-white/20 active:scale-95 transition-transform"
            on:click={() => showMobileChat = !showMobileChat}
            aria-label="Toggle Chat"
        >
            {#if unreadCount > 0}
                <span class="absolute -top-1 -right-1 bg-red-500 text-white text-[10px] font-bold px-2 py-0.5 rounded-full border border-gray-900 animate-pulse">
                    {unreadCount}
                </span>
            {/if}
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
            </svg>
        </button>

        <!-- Mobile Chat Overlay -->
        {#if showMobileChat}
            <!-- Backdrop -->
            <button 
                class="md:hidden fixed inset-0 bg-black/60 backdrop-blur-sm z-40"
                transition:fade={{ duration: 200 }}
                on:click={() => showMobileChat = false}
                aria-label="Close Chat"
            ></button>
            
            <!-- Chat Drawer -->
            <div 
                class="md:hidden fixed right-0 top-0 bottom-0 w-80 bg-gray-900 shadow-2xl border-l border-white/10 z-50 flex flex-col p-4"
                transition:fly={{ x: 300, duration: 300 }}
            >
                <div class="flex justify-between items-center mb-4 border-b border-white/10 pb-4">
                    <h3 class="text-white/50 text-xs font-bold uppercase tracking-widest">{t('chat.title', lang)}</h3>
                    <button on:click={() => showMobileChat = false} class="text-gray-400 hover:text-white" aria-label="Close Chat">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                        </svg>
                    </button>
                </div>
                <div class="flex-1 overflow-hidden flex flex-col">
                    <Chat />
                </div>
            </div>
        {/if}
    {/if}
</div>

<style>
    .animate-fade-in-down {
        animation: fadeInDown 0.8s ease-out;
    }
    @keyframes fadeInDown {
        from { opacity: 0; transform: translateY(-20px); }
        to { opacity: 1; transform: translateY(0); }
    }
</style>
