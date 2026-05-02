<script>
    import { onMount } from 'svelte';
    import { page } from '$app/state';
    import { goto } from '$app/navigation';
    import { auth } from '$lib/stores/auth.js';
    import { toast } from '$lib/stores/toast.js';
    import { api } from '$lib/utils/api.js';

    let activity = $state(null);
    let loading = $state(true);
    let isRegistered = $state(false);
    let registrationId = $state(null);
    let registering = $state(false);
    let note = $state('');

    const activityId = $derived(page.params.id);

    function fmtDate(d) { return new Date(d).toLocaleDateString('th-TH', { weekday:'long', day:'numeric', month:'long', year:'numeric', hour:'2-digit', minute:'2-digit' }); }
    function fmtShort(d) { return new Date(d).toLocaleDateString('th-TH', { day:'numeric', month:'short', year:'numeric' }); }
    function getStatus(s) { const m = { open:{c:'badge-open',t:'เปิดรับสมัคร'}, full:{c:'badge-full',t:'เต็มแล้ว'}, closed:{c:'badge-closed',t:'ปิดรับสมัคร'}, cancelled:{c:'badge-closed',t:'ยกเลิก'} }; return m[s]||{c:'badge-open',t:s}; }
    function catIcon(c) { const m = {Workshop:'🛠️',Competition:'🏆',Camp:'🏕️',Seminar:'🎤',Talk:'💬'}; return m[c]||'📌'; }

    onMount(async () => {
        if (!$auth.authenticated && !$auth.loading) { goto('/login'); return; }
        try {
            const [aRes, cRes] = await Promise.all([api.getActivity(activityId), api.checkRegistration(activityId)]);
            activity = aRes.activity;
            isRegistered = cRes.registered;
            registrationId = cRes.registration_id || null;
        } catch (err) { toast.error('ไม่สามารถโหลดข้อมูลกิจกรรม'); }
        finally { loading = false; }
    });

    async function handleRegister() {
        registering = true;
        try {
            await api.registerActivity({ activity_id: activityId, note });
            isRegistered = true; activity.current_participants += 1;
            toast.success('สมัครกิจกรรมสำเร็จ! 🎉');
        } catch (err) { toast.error(err.message); }
        finally { registering = false; }
    }

    async function handleCancel() {
        if (!registrationId) return;
        try {
            await api.cancelRegistration(registrationId);
            isRegistered = false; registrationId = null;
            activity.current_participants = Math.max(0, activity.current_participants - 1);
            toast.success('ยกเลิกการสมัครสำเร็จ');
        } catch (err) { toast.error(err.message); }
    }
</script>

<svelte:head><title>{activity?.title || 'กิจกรรม'} — IT KMITL Workshop Portal</title></svelte:head>

<div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8 bg-white min-h-screen">
    <a href="/home" class="inline-flex items-center gap-2 text-sm text-text-secondary hover:text-it-600 transition-colors mb-6">
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" /></svg>
        กลับไปหน้ากิจกรรม
    </a>

    {#if loading}
        <div class="animate-pulse space-y-6"><div class="h-64 bg-gray-100 rounded-2xl"></div><div class="h-8 bg-gray-100 rounded w-2/3"></div><div class="h-4 bg-gray-100 rounded w-full"></div></div>
    {:else if activity}
        
        <div class="relative h-64 sm:h-80 rounded-2xl overflow-hidden mb-8 bg-it-50">
            {#if activity.image_url}
                <img src={activity.image_url} alt={activity.title} class="w-full h-full object-cover" />
            {:else}
                <div class="absolute inset-0 bg-gradient-to-br from-it-100 to-it-200 flex items-center justify-center"><span class="text-7xl">{catIcon(activity.category)}</span></div>
            {/if}
            <div class="absolute inset-0 bg-gradient-to-t from-white via-transparent to-transparent"></div>
            <div class="absolute bottom-6 left-6 right-6 z-20">
                <div class="flex items-center gap-3 mb-3">
                    <span class="text-3xl">{catIcon(activity.category)}</span>
                    <span class="badge badge-category">{activity.category}</span>
                    <span class="badge {getStatus(activity.status).c}">{getStatus(activity.status).t}</span>
                </div>
                <h1 class="text-2xl sm:text-3xl font-bold text-it-900">{activity.title}</h1>
            </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
            <div class="lg:col-span-2 space-y-6">
                <div class="bg-gray-50 border border-gray-100 rounded-2xl p-6 sm:p-8">
                    <h2 class="text-xl font-bold text-text-primary mb-4">รายละเอียด</h2>
                    <p class="text-text-secondary leading-relaxed whitespace-pre-line">{activity.description}</p>
                </div>

                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div class="card p-5">
                        <div class="flex items-center gap-3">
                            <div class="w-10 h-10 rounded-xl bg-it-50 flex items-center justify-center"><svg class="w-5 h-5 text-it-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" /></svg></div>
                            <div><p class="text-xs text-text-muted">วันที่จัด</p><p class="text-sm font-medium text-text-primary">{fmtShort(activity.start_date)}</p></div>
                        </div>
                    </div>
                    <div class="card p-5">
                        <div class="flex items-center gap-3">
                            <div class="w-10 h-10 rounded-xl bg-orange-50 flex items-center justify-center"><svg class="w-5 h-5 text-flamingo-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" /></svg></div>
                            <div><p class="text-xs text-text-muted">ปิดรับสมัคร</p><p class="text-sm font-medium text-text-primary">{fmtShort(activity.register_deadline)}</p></div>
                        </div>
                    </div>
                    <div class="card p-5 sm:col-span-2">
                        <div class="flex items-center gap-3">
                            <div class="w-10 h-10 rounded-xl bg-green-50 flex items-center justify-center"><svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" /></svg></div>
                            <div><p class="text-xs text-text-muted">สถานที่</p><p class="text-sm font-medium text-text-primary">{activity.location}</p></div>
                        </div>
                    </div>
                </div>
            </div>

            
            <div class="lg:col-span-1">
                <div class="bg-white border border-gray-200 shadow-lg rounded-2xl p-6 sticky top-24">
                    <h3 class="font-bold text-text-primary mb-4">สมัครเข้าร่วม</h3>
                    <div class="mb-6">
                        <div class="flex justify-between mb-2"><span class="text-sm text-text-secondary">ผู้สมัคร</span><span class="text-sm font-bold text-text-primary">{activity.current_participants}/{activity.max_participants}</span></div>
                        <div class="w-full h-2 bg-gray-100 rounded-full overflow-hidden"><div class="h-full rounded-full bg-gradient-to-r from-it-500 to-it-400 transition-all" style="width: {(activity.current_participants / activity.max_participants) * 100}%"></div></div>
                        <p class="text-xs text-text-muted mt-1.5">เหลือ {activity.max_participants - activity.current_participants} ที่นั่ง</p>
                    </div>

                    {#if isRegistered}
                        <div class="p-4 rounded-xl bg-green-50 border border-green-200 mb-4">
                            <div class="flex items-center gap-2 text-green-700"><svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" /></svg><span class="font-semibold text-sm">คุณสมัครกิจกรรมนี้แล้ว</span></div>
                        </div>
                        <button onclick={handleCancel} class="btn-danger w-full !py-3">ยกเลิกการสมัคร</button>
                    {:else if activity.status === 'open'}
                        <div class="mb-4"><label for="note" class="form-label">หมายเหตุ (ไม่จำเป็น)</label><textarea id="note" bind:value={note} class="form-input !h-20 resize-none" placeholder="เช่น ต้องการข้อมูลเพิ่มเติม..."></textarea></div>
                        <button onclick={handleRegister} disabled={registering} class="btn-primary w-full !py-3">{#if registering}กำลังสมัคร...{:else}สมัครเข้าร่วมกิจกรรม{/if}</button>
                    {:else}
                        <div class="p-4 rounded-xl bg-red-50 border border-red-200"><p class="text-sm text-red-600 font-medium">กิจกรรมนี้{activity.status === 'full' ? 'เต็มแล้ว' : 'ปิดรับสมัครแล้ว'}</p></div>
                    {/if}
                </div>
            </div>
        </div>
    {:else}
        <div class="text-center py-20"><div class="text-5xl mb-4">😕</div><h2 class="text-xl font-bold text-text-primary mb-2">ไม่พบกิจกรรม</h2><a href="/home" class="btn-primary mt-4">กลับหน้ากิจกรรม</a></div>
    {/if}
</div>
