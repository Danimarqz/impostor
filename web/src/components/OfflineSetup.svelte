<script lang="ts">
    import { onMount } from 'svelte';
    import { offline } from '../stores/offline';
    import { language } from '../stores/language';
    import { t } from '../lib/i18n';

    $: lang = $language;

    let categories: string[] = ['General'];
    let selectedCategory = 'General';
    let playerCount = 3;
    let difficulty: 'NORMAL' | 'EASY' = 'NORMAL';
    let loading = false;

    // Reactive statement to fetch categories whenever lang changes
    $: if (lang) {
        fetchCategories();
    }

    async function fetchCategories() {
        try {
            const res = await fetch(`http://localhost:8080/api/categories?lang=${lang}`);
            if (res.ok) {
                const data = await res.json();
                categories = data.categories;
                // Reset selection if current category is not in new list, or just keep it if possible? 
                // For safety, defaulting to first one or keeping 'General' if exists is better.
                if (!categories.includes(selectedCategory)) {
                     selectedCategory = categories[0] || 'General';
                }
            }
        } catch (e) {
            console.error("Failed to load categories", e);
        }
    }

    async function handleStart() {
        loading = true;
        await offline.startGame({
            playerCount,
            category: selectedCategory,
            difficulty,
            lang: lang
        });
        loading = false;
    }
</script>

<div class="w-full max-w-md p-6 md:p-8 bg-white/5 backdrop-blur-xl rounded-3xl shadow-2xl border border-white/10 flex flex-col items-center">
    <h2 class="text-3xl font-bold mb-6 text-white text-center"> {t('offline.setup', lang)} </h2>
    
    <div class="w-full space-y-6">
        <!-- Category -->
        <div class="space-y-2">
             <label for="category-select" class="block text-xs font-bold text-gray-300 uppercase tracking-widest ml-1">{t('offline.category', lang)}</label>
             <select id="category-select" bind:value={selectedCategory} class="w-full px-4 py-3 bg-black/40 border border-white/10 rounded-xl text-white outline-none focus:ring-2 focus:ring-purple-500">
                {#each categories as cat}
                    <option value={cat}>{cat}</option>
                {/each}
             </select>
        </div>

        <!-- Player Count -->
         <div class="space-y-2">
             <label for="player-count" class="block text-xs font-bold text-gray-300 uppercase tracking-widest ml-1">{t('offline.players', lang)}: {playerCount}</label>
             <input id="player-count" type="range" min="3" max="12" bind:value={playerCount} class="w-full h-2 bg-gray-700 rounded-lg appearance-none cursor-pointer accent-purple-500"/>
             <div class="flex justify-between text-xs text-gray-400 px-1">
                <span>3</span>
                <span>12</span>
             </div>
        </div>

        <!-- Difficulty -->
        <div class="space-y-2" role="group" aria-labelledby="difficulty-label">
             <span id="difficulty-label" class="block text-xs font-bold text-gray-300 uppercase tracking-widest ml-1">{t('offline.difficulty', lang)}</span>
             <div class="flex bg-black/40 p-1 rounded-xl">
                <button 
                    class="flex-1 py-2 rounded-lg text-sm font-bold transition-all {difficulty === 'NORMAL' ? 'bg-purple-600 text-white shadow-lg' : 'text-gray-400 hover:text-white'}"
                    on:click={() => difficulty = 'NORMAL'}
                >
                    {t('offline.normal', lang)}
                </button>
                <button 
                    class="flex-1 py-2 rounded-lg text-sm font-bold transition-all {difficulty === 'EASY' ? 'bg-green-600 text-white shadow-lg' : 'text-gray-400 hover:text-white'}"
                    on:click={() => difficulty = 'EASY'}
                >
                    {t('offline.easy', lang)}
                </button>
             </div>
             <p class="text-xs text-gray-500 text-center mt-1">
                {difficulty === 'NORMAL' ? t('offline.diffNormalDesc', lang) : t('offline.diffEasyDesc', lang)}
             </p>
        </div>

        <button 
            on:click={handleStart}
            disabled={loading}
            class="w-full py-4 mt-6 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-bold rounded-xl shadow-lg hover:shadow-blue-500/30 hover:scale-[1.02] active:scale-95 transition-all uppercase tracking-wider disabled:opacity-50 disabled:cursor-not-allowed"
        >
            {loading ? 'INITIALIZING...' : t('offline.startGame', lang)}
        </button>
    </div>
</div>
