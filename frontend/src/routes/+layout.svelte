<script>
    import '../app.css';
    import Navbar from '$lib/components/Navbar.svelte';
    import Footer from '$lib/components/Footer.svelte';
    import Toast from '$lib/components/Toast.svelte';
    import { auth } from '$lib/stores/auth.js';
    import { onMount } from 'svelte';

    let { children } = $props();

    onMount(() => {
        auth.init();
    });
</script>

<svelte:head>
    <title>IT KMITL Workshop Portal</title>
</svelte:head>

<div class="min-h-screen flex flex-col">
    <Navbar />
    <main class="flex-1 pt-16">
        {@render children()}
    </main>
    <Footer />
</div>

<Toast />
