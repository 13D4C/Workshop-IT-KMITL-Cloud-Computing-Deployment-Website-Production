<script>
    import { auth } from '$lib/stores/auth.js';
    import { goto } from '$app/navigation';
    import { page } from '$app/state';

    let mobileMenuOpen = $state(false);

    function toggleMenu() {
        mobileMenuOpen = !mobileMenuOpen;
    }

    async function handleLogout() {
        await auth.logout();
        goto('/');
    }
</script>

<nav class="fixed top-0 left-0 right-0 z-50 bg-white/80 backdrop-blur-xl border-b border-gray-100">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
            
            <a href="/" class="flex items-center gap-3 group">
                <img src="/logo.png" alt="IT KMITL Logo" class="w-9 h-9 rounded-lg transition-transform group-hover:scale-110" />
                <div class="hidden sm:block">
                    <span class="font-bold text-it-800 text-sm">IT KMITL</span>
                    <span class="text-text-muted text-xs block -mt-0.5">Workshop Portal</span>
                </div>
            </a>

            
            <div class="hidden md:flex items-center gap-1">
                <a href="/" class="px-4 py-2 rounded-lg text-sm font-medium transition-all duration-300 {page.url.pathname === '/' ? 'text-it-600 bg-it-50' : 'text-text-secondary hover:text-text-primary hover:bg-gray-50'}">
                    หน้าแรก
                </a>
                {#if $auth.authenticated}
                    <a href="/home" class="px-4 py-2 rounded-lg text-sm font-medium transition-all duration-300 {page.url.pathname === '/home' ? 'text-it-600 bg-it-50' : 'text-text-secondary hover:text-text-primary hover:bg-gray-50'}">
                        กิจกรรม
                    </a>
                {/if}
            </div>

            
            <div class="hidden md:flex items-center gap-3">
                {#if $auth.authenticated}
                    <div class="flex items-center gap-3">
                        <div class="text-right">
                            <p class="text-sm font-medium text-text-primary">{$auth.user?.first_name} {$auth.user?.last_name}</p>
                            <p class="text-xs text-text-muted">{$auth.user?.student_id}</p>
                        </div>
                        <div class="w-9 h-9 rounded-full bg-gradient-to-br from-it-500 to-it-700 flex items-center justify-center text-white font-semibold text-sm">
                            {$auth.user?.first_name?.[0] || 'U'}
                        </div>
                        <button onclick={handleLogout} class="btn-secondary !py-2 !px-4 !text-xs">
                            ออกจากระบบ
                        </button>
                    </div>
                {:else}
                    <a href="/login" class="btn-secondary !py-2 !px-4 !text-sm">เข้าสู่ระบบ</a>
                    <a href="/register" class="btn-primary !py-2 !px-4 !text-sm">สมัครสมาชิก</a>
                {/if}
            </div>

            
            <button onclick={toggleMenu} class="md:hidden p-2 rounded-lg text-text-secondary hover:text-text-primary hover:bg-gray-50 transition-colors" id="mobile-menu-btn" aria-label="เมนู">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    {#if mobileMenuOpen}
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    {:else}
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                    {/if}
                </svg>
            </button>
        </div>

        
        {#if mobileMenuOpen}
            <div class="md:hidden py-4 border-t border-gray-100 animate-slide-up">
                <div class="flex flex-col gap-2">
                    <a href="/" onclick={toggleMenu} class="px-4 py-2 rounded-lg text-sm font-medium text-text-secondary hover:text-text-primary hover:bg-gray-50 transition-all">
                        หน้าแรก
                    </a>
                    {#if $auth.authenticated}
                        <a href="/home" onclick={toggleMenu} class="px-4 py-2 rounded-lg text-sm font-medium text-text-secondary hover:text-text-primary hover:bg-gray-50 transition-all">
                            กิจกรรม
                        </a>
                        <div class="pt-2 mt-2 border-t border-gray-100">
                            <p class="px-4 text-sm text-text-primary font-medium">{$auth.user?.first_name} {$auth.user?.last_name}</p>
                            <button onclick={() => { handleLogout(); toggleMenu(); }} class="mt-2 w-full btn-secondary !text-sm !py-2">
                                ออกจากระบบ
                            </button>
                        </div>
                    {:else}
                        <div class="pt-2 mt-2 border-t border-gray-100 flex gap-2">
                            <a href="/login" onclick={toggleMenu} class="flex-1 btn-secondary !text-sm !py-2 text-center">เข้าสู่ระบบ</a>
                            <a href="/register" onclick={toggleMenu} class="flex-1 btn-primary !text-sm !py-2 text-center">สมัครสมาชิก</a>
                        </div>
                    {/if}
                </div>
            </div>
        {/if}
    </div>
</nav>
