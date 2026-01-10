<script lang="ts">
    import { game, sendAction } from '../stores/game';
    import { language } from '../stores/language';
    import { t } from '../lib/i18n';

    import { onMount } from 'svelte';

    export let lobbyId: string;

    let copied = false;
    let isEasyMode = false;
    let categories: string[] = [];
    let selectedCategory = 'General';

    $: lang = $language;

    // Fetch categories when language changes
    $: {
        fetchCategories();
    }

    async function fetchCategories() {
        try {
            const res = await fetch(`/api/categories?lang=${$language}`);
            const data = await res.json();
            categories = data.categories || ['General'];
            selectedCategory = categories[0];
        } catch (e) {
            console.error("Failed to fetch categories", e);
            categories = ['General'];
        }
    }

    onMount(() => {
        fetchCategories();
    });

    function copyLink() {
        const url = `${window.location.origin}/?lobby=${lobbyId}`;
        navigator.clipboard.writeText(url);
        copied = true;
        setTimeout(() => copied = false, 2000);
    }

    function startGame() {
        sendAction("START_GAME", { 
            mode: isEasyMode ? 'easy' : 'hard',
            category: selectedCategory,
            language: $language
        });
    }
</script>

<div class="w-full max-w-5xl flex flex-col items-center animate-fade-in">
    
    <!-- Header info -->
    <div class="bg-white/5 backdrop-blur-md rounded-2xl px-8 py-4 mb-8 flex flex-col md:flex-row items-center gap-4 shadow-xl border border-white/10 w-full justify-between">
        <div class="flex items-center gap-2">
            <span class="text-blue-300 text-xs font-bold uppercase tracking-widest">{t('lobby.operationId', lang)}</span>
            <span class="text-white/20">|</span>
            <span class="text-xl font-mono text-white tracking-widest font-bold">{lobbyId}</span>
        </div>
        <button 
            on:click={copyLink}
            class="group flex items-center gap-2 text-xs font-bold bg-white/10 hover:bg-white/20 px-4 py-2 rounded-lg text-white transition-all border border-white/5"
        >
            <span>{copied ? t('lobby.copied', lang) : t('lobby.copyCode', lang)}</span>
            {#if !copied}
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-blue-300 group-hover:text-white transition-colors" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
            {/if}
        </button>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 w-full">
        <!-- Player List -->
        <div class="col-span-1 lg:col-span-1 flex flex-col gap-4 order-2 md:order-1">
            <div class="bg-black/20 backdrop-blur-md rounded-2xl p-6 shadow-xl border border-white/10 flex-1">
                <h3 class="text-lg font-bold mb-4 border-b border-white/10 pb-3 text-white flex justify-between items-center">
                    <span>{t('lobby.squadMembers', lang)}</span>
                    <span class="bg-blue-600/20 text-blue-300 px-2 py-0.5 rounded text-xs">{$game.players.length}</span>
                </h3>
                <ul class="space-y-3">
                    {#each $game.players as player}
                        <li class="flex items-center justify-between p-3 rounded-xl bg-white/5 border border-white/5 hover:bg-white/10 transition">
                            <div class="flex items-center gap-3">
                                <div class="w-8 h-8 rounded-full bg-gradient-to-br from-gray-700 to-gray-800 flex items-center justify-center text-xs font-bold ring-2 ring-white/5">
                                    {player.name.substring(0,2).toUpperCase()}
                                </div>
                                <span class="font-medium text-gray-200">{player.name}</span>
                            </div>
                            <div class="flex items-center gap-2">
                                {#if player.is_leader}
                                    <span class="text-[10px] bg-yellow-500/20 text-yellow-300 px-2 py-0.5 rounded border border-yellow-500/20 font-bold tracking-wider">{t('lobby.host', lang)}</span>
                                {/if}
                                {#if player.id === $game.me.id}
                                    <span class="text-[10px] bg-blue-500/20 text-blue-300 px-2 py-0.5 rounded border border-blue-500/20 font-bold tracking-wider">{t('lobby.you', lang)}</span>
                                {/if}
                            </div>
                        </li>
                    {/each}
                </ul>
            </div>
            <!-- Chat Removed -->
        </div>

        <!-- Waiting / Start Area -->
        <div class="col-span-1 md:col-span-1 lg:col-span-2 order-1 md:order-2">
            <div class="bg-gradient-to-br from-gray-800/50 to-gray-900/50 backdrop-blur-md rounded-2xl p-8 border border-white/10 h-full flex flex-col items-center justify-center relative overflow-hidden">
                 <!-- Decorative grid kind of thing -->
                 <div class="absolute inset-0 bg-[url('/grid.svg')] opacity-10"></div>

                {#if $game.me.isLeader}
                    <div class="w-full max-w-lg space-y-8 relative z-10">
                        <!-- Language Selection -->
                        <div class="space-y-4 flex flex-col items-center">
                            <div class="flex items-center space-x-2">
                                <span class="text-amber-300 font-bold uppercase text-xs tracking-[0.2em]">{t('lobby.language', lang)}</span>
                                <button
                                    class="relative bg-black/40 rounded-full p-0.5 flex items-center w-20 border border-white/10 shadow-inner group cursor-pointer"
                                    on:click={() => language.update(l => l === 'en' ? 'es' : 'en')}
                                >
                                    <div class="absolute w-1/2 h-full top-0 left-0 bg-gradient-to-r from-amber-700 to-amber-600 rounded-full transition-transform duration-300 shadow-md {$language === 'es' ? 'translate-x-[100%]' : 'translate-x-0'}"></div>
                                    <span class="relative z-10 w-1/2 text-center text-xs font-bold transition-colors duration-300 {$language === 'es' ? 'text-gray-500' : 'text-white'}">EN</span>
                                    <span class="relative z-10 w-1/2 text-center text-xs font-bold transition-colors duration-300 {$language === 'es' ? 'text-white' : 'text-gray-500'}">ES</span>
                                </button>
                            </div>
                        </div>

                        <!-- Category Selection -->
                        <div class="space-y-4 text-center">
                            <span class="text-blue-300 font-bold uppercase text-xs tracking-[0.2em]">{t('lobby.selectSector', lang)}</span>
                            <div class="flex flex-wrap gap-2 justify-center">
                                {#each categories as cat}
                                    <button 
                                        type="button"
                                        class="px-5 py-2 rounded-xl text-sm font-bold border transition-all duration-300 cursor-pointer relative z-10 hover:-translate-y-1
                                               {selectedCategory === cat 
                                                 ? (cat === '✨ Infinite' ? 'bg-gradient-to-r from-purple-600 to-pink-600 border-purple-500 text-white shadow-lg shadow-purple-500/20 animate-pulse' : 'bg-blue-600 border-blue-500 text-white shadow-lg shadow-blue-500/20')
                                                 : (cat === '✨ Infinite' ? 'bg-gray-800/80 border-purple-500/50 text-purple-300 hover:bg-purple-900/20 hover:border-purple-400' : 'bg-gray-800/80 border-gray-700 text-gray-400 hover:bg-gray-700 hover:border-gray-500')}"
                                        on:click={() => selectedCategory = cat}
                                    >
                                        {cat}
                                    </button>
                                {/each}
                            </div>
                        </div>

                        <!-- Difficulty Selection -->
                        <div class="space-y-4 flex flex-col items-center">
                            <span class="text-purple-300 font-bold uppercase text-xs tracking-[0.2em]">{t('lobby.difficulty', lang)}</span>
                            <button 
                                class="relative bg-black/40 rounded-full p-1.5 flex items-center w-64 border border-white/10 shadow-inner group cursor-pointer"
                                on:click={() => isEasyMode = !isEasyMode}
                            >
                                <div class="absolute w-1/2 h-full top-0 left-0 bg-gradient-to-r from-gray-700 to-gray-600 rounded-full transition-transform duration-300 shadow-md {isEasyMode ? 'translate-x-[100%]' : 'translate-x-0'}"></div>
                                
                                <span class="relative z-10 w-1/2 text-center text-xs font-bold transition-colors duration-300 flex flex-col py-1 {isEasyMode ? 'text-gray-500' : 'text-white'}">
                                    <span class="tracking-widest">{t('lobby.hard', lang)}</span>
                                </span>
                                
                                <span class="relative z-10 w-1/2 text-center text-xs font-bold transition-colors duration-300 flex flex-col py-1 {isEasyMode ? 'text-white' : 'text-gray-500'}">
                                    <span class="tracking-widest">{t('lobby.easy', lang)}</span>
                                </span>
                            </button>
                            <p class="text-xs text-gray-400 font-mono">
                                {isEasyMode ? t('lobby.trapDetected', lang) : t('lobby.blindMode', lang)}
                            </p>
                        </div>
                    </div>

                    <div class="mt-12 flex flex-col items-center gap-4 relative z-10">
                        <p class="text-gray-300 text-sm">{t('lobby.systemsReady', lang)}</p>
                        <button 
                            class="group relative bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-500 hover:to-emerald-500 text-white font-black py-4 px-12 rounded-2xl shadow-xl hover:shadow-green-500/20 transform transition-all duration-300 hover:scale-[1.02] active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none"
                            on:click={startGame}
                            disabled={$game.players.length < 3}
                        >
                            <span class="relative z-10 flex items-center gap-3">
                                <span>{t('lobby.initiateLaunch', lang)}</span>
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 group-hover:rotate-90 transition-transform duration-500" viewBox="0 0 20 20" fill="currentColor">
                                    <path fill-rule="evenodd" d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z" clip-rule="evenodd" />
                                </svg>
                            </span>
                        </button>
                        {#if $game.players.length < 3}
                            <p class="text-red-400 bg-red-900/20 px-4 py-2 rounded-lg text-xs border border-red-500/20 flex items-center gap-2">
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" /></svg>
                                {t('lobby.minPlayers', lang)}
                            </p>
                        {/if}
                    </div>

                {:else}
                    <div class="flex flex-col items-center justify-center h-full">
                        <div class="relative w-24 h-24 mb-8">
                             <div class="absolute inset-0 border-4 border-blue-500/30 rounded-full animate-ping"></div>
                             <div class="absolute inset-0 border-4 border-t-blue-500 border-r-transparent border-b-transparent border-l-transparent rounded-full animate-spin"></div>
                             <div class="absolute inset-2 bg-blue-500/10 rounded-full flex items-center justify-center">
                                 <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                                 </svg>
                             </div>
                        </div>
                        <h2 class="text-2xl font-bold text-white mb-2 tracking-widest">{t('lobby.standby', lang)}</h2>
                        <p class="text-gray-400">{t('lobby.waitingHost', lang)}</p>
                    </div>
                {/if}
            </div>
        </div>
    </div>
</div>
