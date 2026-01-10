<script lang="ts">
    import { game, sendAction } from '../stores/game';
    import { afterUpdate } from 'svelte';
    import { language } from '../stores/language';
    import { t } from '../lib/i18n';

    $: lang = $language;

    let message = '';
    let chatContainer: HTMLElement;

    function sendMessage() {
        if (!message.trim()) return;
        sendAction('CHAT_MESSAGE', { message });
        message = '';
    }

    afterUpdate(() => {
        if (chatContainer) chatContainer.scrollTop = chatContainer.scrollHeight;
    });
</script>

<div class="flex flex-col h-full w-full bg-black/40 rounded-xl border border-white/10 overflow-hidden shadow-inner">
    <!-- Messages Area -->
    <div class="flex-1 overflow-y-auto p-4 space-y-3 scrollbar-thin scrollbar-thumb-white/20" bind:this={chatContainer}>
        {#each $game.messages as msg}
            <div class="text-sm">
                <span class="font-bold text-blue-400">{msg.from}:</span>
                <span class="text-gray-200">{msg.text}</span>
            </div>
        {/each}
        {#if $game.messages.length === 0}
            <p class="text-gray-500 text-xs italic text-center mt-4">{t('chat.startConversation', lang)}</p>
        {/if}
    </div>

    <!-- Input Area -->
    <div class="p-2 border-t border-gray-700 flex gap-2">
        <input 
            type="text" 
            bind:value={message}
            on:keydown={(e) => e.key === 'Enter' && sendMessage()}
            class="flex-1 bg-gray-800 text-white text-sm rounded px-3 py-2 outline-none focus:ring-1 focus:ring-blue-500"
            placeholder={t('chat.typeMessage', lang)}
        />
        <button 
            on:click={sendMessage}
            class="bg-blue-600 hover:bg-blue-500 text-white font-bold px-3 py-2 rounded transition flex items-center justify-center"
            aria-label={t('chat.send', lang)}
        >
            <!-- Send Icon -->
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11a1 1 0 112 0v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z" />
            </svg>
        </button>
    </div>
</div>
