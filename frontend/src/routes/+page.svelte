<script>
    import { auth } from '$lib/stores/auth.js';
</script>

<svelte:head>
    <title>IT KMITL Workshop Portal — ศูนย์รวมกิจกรรมและ Workshop คณะ IT</title>
    <meta name="description" content="สมัครเข้าร่วม Workshop, Competition, Camp และกิจกรรมสุดพิเศษจากคณะเทคโนโลยีสารสนเทศ สถาบันเทคโนโลยีพระจอมเกล้าเจ้าคุณทหารลาดกระบัง" />
</svelte:head>


<section class="relative min-h-screen flex items-center justify-center overflow-hidden bg-gradient-to-br from-white via-it-50/50 to-white">
    
    <div class="absolute inset-0">
        <div class="absolute top-1/4 -left-32 w-96 h-96 bg-it-200/40 rounded-full blur-[120px] animate-float"></div>
        <div class="absolute bottom-1/4 -right-32 w-96 h-96 bg-it-100/50 rounded-full blur-[120px] animate-float" style="animation-delay: 3s;"></div>
        <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-flamingo-400/5 rounded-full blur-[150px]"></div>
        <div class="absolute inset-0 opacity-[0.03]" style="background-image: linear-gradient(rgba(0,0,0,0.05) 1px, transparent 1px), linear-gradient(90deg, rgba(0,0,0,0.05) 1px, transparent 1px); background-size: 60px 60px;"></div>
    </div>

    <div class="relative z-10 max-w-5xl mx-auto px-4 text-center">
        
        <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-it-50 border border-it-100 mb-8 animate-fade-in">
            <div class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></div>
            <span class="text-sm text-it-700 font-medium">เปิดรับสมัครกิจกรรมใหม่แล้ว</span>
        </div>

        
        <h1 class="text-5xl sm:text-6xl lg:text-7xl font-black mb-6 tracking-tight animate-slide-up">
            <span class="text-it-900">IT KMITL</span><br />
            <span class="gradient-text">Workshop Portal</span>
        </h1>

        
        <p class="text-lg sm:text-xl text-text-secondary max-w-2xl mx-auto mb-10 leading-relaxed animate-slide-up delay-200" style="opacity: 0;">
            ศูนย์รวมกิจกรรม Workshop, Competition, Camp และ Tech Talk
            จากคณะเทคโนโลยีสารสนเทศ สถาบันเทคโนโลยีพระจอมเกล้าเจ้าคุณทหารลาดกระบัง
        </p>

        
        <div class="flex flex-col sm:flex-row gap-4 justify-center animate-slide-up delay-300" style="opacity: 0;">
            {#if $auth.authenticated}
                <a href="/home" class="btn-primary text-lg px-8 py-4">
                    <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" /></svg>
                    ดูกิจกรรมทั้งหมด
                </a>
            {:else}
                <a href="/register" class="btn-primary text-lg px-8 py-4" id="hero-register-btn">
                    <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z" /></svg>
                    สมัครสมาชิกเลย
                </a>
                <a href="/login" class="btn-secondary text-lg px-8 py-4" id="hero-login-btn">เข้าสู่ระบบ</a>
            {/if}
        </div>

        
        <div class="grid grid-cols-3 gap-6 mt-20 max-w-lg mx-auto animate-slide-up delay-400" style="opacity: 0;">
            <div class="text-center">
                <div class="text-3xl font-black text-it-600">6+</div>
                <div class="text-xs text-text-muted mt-1">กิจกรรม</div>
            </div>
            <div class="text-center">
                <div class="text-3xl font-black text-it-600">425</div>
                <div class="text-xs text-text-muted mt-1">ที่นั่งรวม</div>
            </div>
            <div class="text-center">
                <div class="text-3xl font-black text-flamingo-500">IT</div>
                <div class="text-xs text-text-muted mt-1">KMITL</div>
            </div>
        </div>
    </div>

    
    <div class="absolute bottom-8 left-1/2 -translate-x-1/2 animate-bounce">
        <svg class="w-6 h-6 text-text-muted" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" /></svg>
    </div>
</section>


<section class="py-24 relative bg-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
            <h2 class="text-3xl sm:text-4xl font-bold text-text-primary mb-4">ทำไมต้องเข้าร่วมกับเรา?</h2>
            <p class="text-text-secondary max-w-2xl mx-auto">กิจกรรมคุณภาพที่คัดสรรมาเพื่อนักศึกษาคณะ IT โดยเฉพาะ</p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div class="card p-8 text-center group">
                <div class="w-16 h-16 mx-auto mb-6 rounded-2xl bg-it-50 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
                    <svg class="w-8 h-8 text-it-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" /></svg>
                </div>
                <h3 class="text-xl font-bold text-text-primary mb-3">เรียนรู้จากผู้เชี่ยวชาญ</h3>
                <p class="text-sm text-text-secondary leading-relaxed">Workshop ที่สอนโดยอาจารย์และผู้เชี่ยวชาญในสายงาน พร้อม Hands-on Lab จริง</p>
            </div>

            <div class="card p-8 text-center group">
                <div class="w-16 h-16 mx-auto mb-6 rounded-2xl bg-blue-50 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
                    <svg class="w-8 h-8 text-it-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" /></svg>
                </div>
                <h3 class="text-xl font-bold text-text-primary mb-3">Networking</h3>
                <p class="text-sm text-text-secondary leading-relaxed">พบปะเพื่อนร่วมสาย สร้าง connection กับ Senior Engineers จากบริษัทชั้นนำ</p>
            </div>

            <div class="card p-8 text-center group">
                <div class="w-16 h-16 mx-auto mb-6 rounded-2xl bg-green-50 flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
                    <svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" /></svg>
                </div>
                <h3 class="text-xl font-bold text-text-primary mb-3">ใบ Certificate</h3>
                <p class="text-sm text-text-secondary leading-relaxed">รับ Certificate เมื่อเข้าร่วมกิจกรรมครบ เพิ่มประสบการณ์ใน Portfolio ของคุณ</p>
            </div>
        </div>
    </div>
</section>


<section class="py-24 relative border-t border-gray-100 bg-gray-50/50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
            <h2 class="text-3xl sm:text-4xl font-bold text-text-primary mb-4">สร้างด้วยเทคโนโลยีชั้นนำ</h2>
            <p class="text-text-secondary">พร้อม Deploy บน AWS ด้วย High Availability Architecture</p>
        </div>
        <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-6 gap-4">
            {#each [
                { name: 'Svelte', icon: '⚡' },
                { name: 'Go Fiber', icon: '🚀' },
                { name: 'PostgreSQL', icon: '🐘' },
                { name: 'Docker', icon: '🐳' },
                { name: 'TailwindCSS', icon: '🎨' },
                { name: 'AWS', icon: '☁️' },
            ] as tech}
                <div class="card p-4 text-center group cursor-default">
                    <div class="text-3xl mb-2 group-hover:scale-125 transition-transform duration-300">{tech.icon}</div>
                    <div class="text-xs font-semibold text-text-secondary">{tech.name}</div>
                </div>
            {/each}
        </div>
    </div>
</section>


<section class="py-24 relative bg-white">
    <div class="max-w-4xl mx-auto px-4 text-center">
        <div class="rounded-3xl p-12 sm:p-16 relative overflow-hidden bg-gradient-to-br from-it-600 to-it-800">
            <div class="absolute top-0 right-0 w-64 h-64 bg-white/5 rounded-full blur-[80px]"></div>
            <div class="absolute bottom-0 left-0 w-64 h-64 bg-flamingo-500/10 rounded-full blur-[80px]"></div>
            <div class="relative z-10">
                <h2 class="text-3xl sm:text-4xl font-bold text-white mb-4">พร้อมเริ่มต้นแล้วหรือยัง?</h2>
                <p class="text-it-100 mb-8 max-w-lg mx-auto">สมัครสมาชิกวันนี้เพื่อไม่พลาดกิจกรรมสุดพิเศษจากคณะ IT KMITL</p>
                {#if !$auth.authenticated}
                    <a href="/register" class="inline-flex items-center justify-center px-10 py-4 bg-white text-it-700 rounded-xl font-bold text-lg hover:bg-it-50 transition-all hover:-translate-y-1 shadow-lg">สมัครสมาชิกฟรี</a>
                {:else}
                    <a href="/home" class="inline-flex items-center justify-center px-10 py-4 bg-white text-it-700 rounded-xl font-bold text-lg hover:bg-it-50 transition-all hover:-translate-y-1 shadow-lg">ดูกิจกรรมทั้งหมด</a>
                {/if}
            </div>
        </div>
    </div>
</section>
