<script>
    import { auth } from '$lib/stores/auth.js';
    import { toast } from '$lib/stores/toast.js';
    import { goto } from '$app/navigation';

    let loading = $state(false);
    let email = $state('');
    let password = $state('');

    async function handleLogin(e) {
        e.preventDefault();
        if (!email || !password) { toast.error('กรุณากรอกอีเมลและรหัสผ่าน'); return; }
        loading = true;
        try {
            await auth.login(email, password);
            toast.success('เข้าสู่ระบบสำเร็จ! 🎉');
            goto('/home');
        } catch (err) { toast.error(err.message); }
        finally { loading = false; }
    }
</script>

<svelte:head>
    <title>เข้าสู่ระบบ — IT KMITL Workshop Portal</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center py-20 px-4 relative bg-gradient-to-br from-white via-it-50/30 to-white">
    <div class="absolute top-1/3 -left-32 w-80 h-80 bg-it-100/50 rounded-full blur-[100px]"></div>
    <div class="absolute bottom-1/3 -right-32 w-80 h-80 bg-it-50/50 rounded-full blur-[100px]"></div>

    <div class="relative z-10 w-full max-w-md">
        <div class="text-center mb-8">
            <img src="/logo.png" alt="IT KMITL Logo" class="w-14 h-14 mx-auto mb-4 rounded-2xl" />
            <h1 class="text-3xl font-bold text-text-primary mb-2">เข้าสู่ระบบ</h1>
            <p class="text-text-secondary">ยินดีต้อนรับกลับ! เข้าสู่ระบบเพื่อดูกิจกรรม</p>
        </div>

        <form onsubmit={handleLogin} class="bg-white border border-gray-100 shadow-xl rounded-2xl p-8 space-y-6" id="login-form">
            <div>
                <label for="login-email" class="form-label">อีเมล</label>
                <input type="email" id="login-email" bind:value={email} class="form-input" placeholder="your.email@kmitl.ac.th" required />
            </div>
            <div>
                <label for="login-password" class="form-label">รหัสผ่าน</label>
                <input type="password" id="login-password" bind:value={password} class="form-input" placeholder="รหัสผ่าน" required />
            </div>
            <button type="submit" class="btn-primary w-full !py-3 text-base" disabled={loading} id="login-submit-btn">
                {#if loading}
                    <svg class="animate-spin w-5 h-5 mr-2" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path></svg>
                    กำลังเข้าสู่ระบบ...
                {:else}เข้าสู่ระบบ{/if}
            </button>
            <div class="flex items-center gap-4">
                <div class="flex-1 h-px bg-gray-200"></div>
                <span class="text-xs text-text-muted">หรือ</span>
                <div class="flex-1 h-px bg-gray-200"></div>
            </div>
            <p class="text-center text-sm text-text-secondary">ยังไม่มีบัญชี? <a href="/register" class="text-it-600 hover:text-it-500 font-medium transition-colors">สมัครสมาชิก</a></p>
        </form>
    </div>
</div>
