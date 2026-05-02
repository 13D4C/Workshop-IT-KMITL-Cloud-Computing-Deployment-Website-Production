<script>
    import { toast } from '$lib/stores/toast.js';
</script>

{#if $toast.length > 0}
    <div class="fixed bottom-6 right-6 z-[100] flex flex-col gap-3 max-w-sm">
        {#each $toast as t (t.id)}
            <div
                class="animate-slide-up rounded-xl px-5 py-4 shadow-lg border flex items-start gap-3
                    {t.type === 'success' ? 'bg-green-50 border-green-200 text-green-700' : ''}
                    {t.type === 'error' ? 'bg-red-50 border-red-200 text-red-700' : ''}
                    {t.type === 'info' ? 'bg-blue-50 border-blue-200 text-blue-700' : ''}
                    {t.type === 'warning' ? 'bg-yellow-50 border-yellow-200 text-yellow-700' : ''}"
            >
                <div class="mt-0.5">
                    {#if t.type === 'success'}
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg>
                    {:else if t.type === 'error'}
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" /></svg>
                    {:else if t.type === 'warning'}
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>
                    {:else}
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
                    {/if}
                </div>
                <p class="text-sm font-medium flex-1">{t.message}</p>
                <button onclick={() => toast.remove(t.id)} class="ml-2 opacity-60 hover:opacity-100 transition-opacity" aria-label="ปิด">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" /></svg>
                </button>
            </div>
        {/each}
    </div>
{/if}
