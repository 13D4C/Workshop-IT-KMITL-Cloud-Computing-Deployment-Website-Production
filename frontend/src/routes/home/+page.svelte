<script>
    import { onMount } from 'svelte';
    import { auth } from '$lib/stores/auth.js';
    import { toast } from '$lib/stores/toast.js';
    import { goto } from '$app/navigation';
    import { api } from '$lib/utils/api.js';
    import ActivityCard from '$lib/components/ActivityCard.svelte';

    let activities = $state([]);
    let myRegistrations = $state([]);
    let loading = $state(true);
    let searchQuery = $state('');
    let selectedCategory = $state('');
    let activeTab = $state('all');
    const categories = ['', 'Workshop', 'Competition', 'Camp', 'Seminar', 'Talk'];

    const filteredActivities = $derived(() => {
        let f = activities;
        if (searchQuery) { const q = searchQuery.toLowerCase(); f = f.filter(a => a.title.toLowerCase().includes(q) || a.description.toLowerCase().includes(q)); }
        if (selectedCategory) f = f.filter(a => a.category === selectedCategory);
        return f;
    });

    onMount(async () => {
        if (!$auth.authenticated && !$auth.loading) { goto('/login'); return; }
        try {
            const [aRes, rRes] = await Promise.all([api.getActivities(), api.getMyRegistrations()]);
            activities = aRes.activities || [];
            myRegistrations = rRes.registrations || [];
        } catch (err) { toast.error('ไม่สามารถโหลดข้อมูลได้'); }
        finally { loading = false; }
    });

    async function cancelReg(regId) {
        try {
            await api.cancelRegistration(regId);
            myRegistrations = myRegistrations.filter(r => r.id !== regId);
            const res = await api.getActivities();
            activities = res.activities || [];
            toast.success('ยกเลิกการสมัครสำเร็จ');
        } catch (err) { toast.error(err.message); }
    }

    function fmtDate(d) { return new Date(d).toLocaleDateString('th-TH', { day:'numeric', month:'short', year:'numeric', hour:'2-digit', minute:'2-digit' }); }
</script>

<svelte:head><title>กิจกรรม — IT KMITL Workshop Portal</title></svelte:head>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 bg-white min-h-screen">
    <div class="mb-8">
        <h1 class="text-3xl sm:text-4xl font-bold text-text-primary mb-2">
            {#if $auth.user}สวัสดี, {$auth.user.first_name}! 👋{:else}กิจกรรมทั้งหมด{/if}
        </h1>
        <p class="text-text-secondary">ค้นหาและสมัครกิจกรรมที่คุณสนใจ</p>
    </div>

    <div class="flex gap-2 mb-6">
        <button class="px-5 py-2.5 rounded-xl text-sm font-medium transition-all {activeTab === 'all' ? 'bg-it-50 text-it-600 border border-it-200' : 'text-text-secondary hover:bg-gray-50 border border-transparent'}" onclick={() => activeTab = 'all'}>กิจกรรมทั้งหมด</button>
        <button class="px-5 py-2.5 rounded-xl text-sm font-medium transition-all {activeTab === 'my' ? 'bg-it-50 text-it-600 border border-it-200' : 'text-text-secondary hover:bg-gray-50 border border-transparent'}" onclick={() => activeTab = 'my'}>
            กิจกรรมที่สมัครแล้ว
            {#if myRegistrations.length > 0}<span class="ml-1.5 px-2 py-0.5 rounded-full bg-flamingo-500/10 text-flamingo-500 text-xs font-bold">{myRegistrations.length}</span>{/if}
        </button>
    </div>

    {#if activeTab === 'all'}
        <div class="flex flex-col sm:flex-row gap-4 mb-8">
            <div class="flex-1 relative">
                <svg class="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-text-muted" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
                <input type="text" bind:value={searchQuery} class="form-input !pl-10" placeholder="ค้นหากิจกรรม..." />
            </div>
            <div class="flex gap-2 flex-wrap">
                {#each categories as cat}
                    <button class="px-4 py-2 rounded-lg text-sm font-medium transition-all {selectedCategory === cat ? 'bg-it-50 text-it-600 border border-it-200' : 'text-text-secondary bg-gray-50 hover:bg-gray-100 border border-transparent'}" onclick={() => selectedCategory = cat}>{cat || 'ทั้งหมด'}</button>
                {/each}
            </div>
        </div>

        {#if loading}
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {#each [1,2,3] as _}<div class="card animate-pulse"><div class="h-48 bg-gray-100"></div><div class="p-5 space-y-3"><div class="h-4 bg-gray-100 rounded w-1/3"></div><div class="h-6 bg-gray-100 rounded w-3/4"></div></div></div>{/each}
            </div>
        {:else if filteredActivities().length === 0}
            <div class="text-center py-20"><div class="text-5xl mb-4">🔍</div><h3 class="text-xl font-semibold text-text-primary mb-2">ไม่พบกิจกรรม</h3><p class="text-text-secondary">ลองเปลี่ยนคำค้นหาหรือตัวกรอง</p></div>
        {:else}
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {#each filteredActivities() as activity (activity.id)}<ActivityCard {activity} />{/each}
            </div>
        {/if}
    {:else}
        {#if myRegistrations.length === 0}
            <div class="text-center py-20"><div class="text-5xl mb-4">📋</div><h3 class="text-xl font-semibold text-text-primary mb-2">ยังไม่ได้สมัครกิจกรรม</h3><button onclick={() => activeTab = 'all'} class="btn-primary mt-4">ดูกิจกรรมทั้งหมด</button></div>
        {:else}
            <div class="space-y-4">
                {#each myRegistrations as reg (reg.id)}
                    <div class="card p-6 flex flex-col sm:flex-row gap-4 items-start sm:items-center">
                        <div class="w-12 h-12 rounded-2xl bg-it-50 flex items-center justify-center shrink-0">
                            <svg class="w-6 h-6 text-it-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4" /></svg>
                        </div>
                        <div class="flex-1 min-w-0">
                            <a href="/activity/{reg.activity_id}" class="text-lg font-bold text-text-primary hover:text-it-600 transition-colors">{reg.activity?.title || 'กิจกรรม'}</a>
                            <p class="text-sm text-text-secondary mt-1">{fmtDate(reg.activity?.start_date)} · สมัครเมื่อ {fmtDate(reg.registered_at)}</p>
                        </div>
                        <button onclick={() => cancelReg(reg.id)} class="btn-danger shrink-0">ยกเลิก</button>
                    </div>
                {/each}
            </div>
        {/if}
    {/if}
</div>
