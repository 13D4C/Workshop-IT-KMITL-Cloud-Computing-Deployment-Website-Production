<script>
    let { activity } = $props();

    function formatDate(dateStr) {
        const date = new Date(dateStr);
        return date.toLocaleDateString('th-TH', { day: 'numeric', month: 'short', year: 'numeric' });
    }

    function getStatusBadge(status) {
        switch(status) {
            case 'open': return { class: 'badge-open', text: 'เปิดรับสมัคร' };
            case 'full': return { class: 'badge-full', text: 'เต็มแล้ว' };
            case 'closed': return { class: 'badge-closed', text: 'ปิดรับสมัคร' };
            case 'cancelled': return { class: 'badge-closed', text: 'ยกเลิก' };
            default: return { class: 'badge-open', text: status };
        }
    }

    function getCategoryIcon(category) {
        switch(category) {
            case 'Workshop': return '🛠️';
            case 'Competition': return '🏆';
            case 'Camp': return '🏕️';
            case 'Seminar': return '🎤';
            case 'Talk': return '💬';
            default: return '📌';
        }
    }

    const badge = $derived(getStatusBadge(activity.status));
    const spotsLeft = $derived(activity.max_participants - activity.current_participants);
    const spotsPercent = $derived((activity.current_participants / activity.max_participants) * 100);
</script>

<a href="/activity/{activity.id}" class="card block group" id="activity-card-{activity.id}">
    
    <div class="relative h-48 overflow-hidden bg-it-50">
        {#if activity.image_url}
            <img
                src={activity.image_url}
                alt={activity.title}
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
            />
        {:else}
            <div class="absolute inset-0 bg-gradient-to-br from-it-100 to-it-200 flex items-center justify-center">
                <span class="text-5xl">{getCategoryIcon(activity.category)}</span>
            </div>
        {/if}
        <div class="absolute inset-0 bg-gradient-to-t from-white/60 via-transparent to-transparent"></div>

        
        <div class="absolute top-3 right-3 z-10">
            <span class="badge {badge.class}">{badge.text}</span>
        </div>
    </div>

    
    <div class="p-5">
        <div class="flex items-center gap-2 mb-2">
            <span class="badge badge-category text-xs">{getCategoryIcon(activity.category)} {activity.category}</span>
        </div>

        <h3 class="font-bold text-lg text-text-primary mb-2 group-hover:text-it-600 transition-colors line-clamp-2">
            {activity.title}
        </h3>

        <p class="text-sm text-text-secondary line-clamp-2 mb-4">
            {activity.description}
        </p>

        
        <div class="space-y-2 text-xs text-text-muted">
            <div class="flex items-center gap-2">
                <svg class="w-4 h-4 text-it-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                <span>{formatDate(activity.start_date)}</span>
            </div>
            <div class="flex items-center gap-2">
                <svg class="w-4 h-4 text-it-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                <span class="truncate">{activity.location}</span>
            </div>
        </div>

        
        <div class="mt-4">
            <div class="flex justify-between items-center mb-1.5">
                <span class="text-xs text-text-muted">ผู้สมัคร</span>
                <span class="text-xs font-semibold {spotsLeft <= 5 ? 'text-flamingo-500' : 'text-text-secondary'}">
                    {activity.current_participants}/{activity.max_participants}
                </span>
            </div>
            <div class="w-full h-1.5 bg-gray-100 rounded-full overflow-hidden">
                <div
                    class="h-full rounded-full transition-all duration-700 {spotsPercent >= 90 ? 'bg-gradient-to-r from-flamingo-500 to-red-500' : 'bg-gradient-to-r from-it-500 to-it-400'}"
                    style="width: {spotsPercent}%"
                ></div>
            </div>
        </div>
    </div>
</a>

<style>
    .line-clamp-2 {
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
        overflow: hidden;
    }
</style>
