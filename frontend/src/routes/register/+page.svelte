<script>
    import { auth } from '$lib/stores/auth.js';
    import { toast } from '$lib/stores/toast.js';
    import { goto } from '$app/navigation';

    let loading = $state(false);
    let form = $state({
        student_id: '', email: '', password: '', confirm_password: '',
        first_name: '', last_name: '', phone: '',
        faculty: 'คณะเทคโนโลยีสารสนเทศ', department: '',
    });

    const departments = [
        'เทคโนโลยีสารสนเทศ', 'วิทยาการข้อมูลและการวิเคราะห์เชิงธุรกิจ',
        'เทคโนโลยีสารสนเทศทางธุรกิจ',
    ];

    async function handleRegister(e) {
        e.preventDefault();
        if (form.password !== form.confirm_password) { toast.error('รหัสผ่านไม่ตรงกัน'); return; }
        if (form.password.length < 6) { toast.error('รหัสผ่านต้องมีอย่างน้อย 6 ตัวอักษร'); return; }
        loading = true;
        try {
            await auth.register({ student_id: form.student_id, email: form.email, password: form.password, first_name: form.first_name, last_name: form.last_name, phone: form.phone, faculty: form.faculty, department: form.department });
            toast.success('สมัครสมาชิกสำเร็จ! ยินดีต้อนรับ 🎉');
            goto('/home');
        } catch (err) { toast.error(err.message); }
        finally { loading = false; }
    }
</script>

<svelte:head>
    <title>สมัครสมาชิก — IT KMITL Workshop Portal</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center py-20 px-4 relative bg-gradient-to-br from-white via-it-50/30 to-white">
    <div class="absolute top-1/4 -right-32 w-80 h-80 bg-it-100/50 rounded-full blur-[100px]"></div>
    <div class="absolute bottom-1/4 -left-32 w-80 h-80 bg-it-50/50 rounded-full blur-[100px]"></div>

    <div class="relative z-10 w-full max-w-lg">
        <div class="text-center mb-8">
            <img src="/logo.png" alt="IT KMITL Logo" class="w-14 h-14 mx-auto mb-4 rounded-2xl" />
            <h1 class="text-3xl font-bold text-text-primary mb-2">สมัครสมาชิก</h1>
            <p class="text-text-secondary">สร้างบัญชีเพื่อเข้าร่วมกิจกรรม IT KMITL</p>
        </div>

        <form onsubmit={handleRegister} class="bg-white border border-gray-100 shadow-xl rounded-2xl p-8 space-y-5" id="register-form">
            <div class="grid grid-cols-2 gap-4">
                <div>
                    <label for="first_name" class="form-label">ชื่อ <span class="text-red-500">*</span></label>
                    <input type="text" id="first_name" bind:value={form.first_name} class="form-input" placeholder="ชื่อ" required />
                </div>
                <div>
                    <label for="last_name" class="form-label">นามสกุล <span class="text-red-500">*</span></label>
                    <input type="text" id="last_name" bind:value={form.last_name} class="form-input" placeholder="นามสกุล" required />
                </div>
            </div>
            <div>
                <label for="student_id" class="form-label">รหัสนักศึกษา <span class="text-red-500">*</span></label>
                <input type="text" id="student_id" bind:value={form.student_id} class="form-input" placeholder="เช่น 66010XXX" required />
            </div>
            <div>
                <label for="email" class="form-label">อีเมล <span class="text-red-500">*</span></label>
                <input type="email" id="email" bind:value={form.email} class="form-input" placeholder="your.email@kmitl.ac.th" required />
            </div>
            <div>
                <label for="phone" class="form-label">เบอร์โทรศัพท์</label>
                <input type="tel" id="phone" bind:value={form.phone} class="form-input" placeholder="0XX-XXX-XXXX" />
            </div>
            <div>
                <label for="department" class="form-label">สาขา</label>
                <select id="department" bind:value={form.department} class="form-input">
                    <option value="">เลือกสาขา</option>
                    {#each departments as dept}<option value={dept}>{dept}</option>{/each}
                </select>
            </div>
            <div>
                <label for="password" class="form-label">รหัสผ่าน <span class="text-red-500">*</span></label>
                <input type="password" id="password" bind:value={form.password} class="form-input" placeholder="อย่างน้อย 6 ตัวอักษร" required minlength="6" />
            </div>
            <div>
                <label for="confirm_password" class="form-label">ยืนยันรหัสผ่าน <span class="text-red-500">*</span></label>
                <input type="password" id="confirm_password" bind:value={form.confirm_password} class="form-input" placeholder="กรอกรหัสผ่านอีกครั้ง" required />
            </div>
            <button type="submit" class="btn-primary w-full !py-3 text-base" disabled={loading} id="register-submit-btn">
                {#if loading}
                    <svg class="animate-spin w-5 h-5 mr-2" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path></svg>
                    กำลังสมัคร...
                {:else}สมัครสมาชิก{/if}
            </button>
            <p class="text-center text-sm text-text-secondary">มีบัญชีอยู่แล้ว? <a href="/login" class="text-it-600 hover:text-it-500 font-medium transition-colors">เข้าสู่ระบบ</a></p>
        </form>
    </div>
</div>
