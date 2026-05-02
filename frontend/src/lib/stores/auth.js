import { writable } from 'svelte/store';
import { api } from '$lib/utils/api.js';

function createAuthStore() {
    const { subscribe, set, update } = writable({
        user: null,
        loading: true,
        authenticated: false,
    });

    return {
        subscribe,
        
        init: async () => {
            try {
                const data = await api.me();
                set({ user: data.user, loading: false, authenticated: true });
            } catch {
                set({ user: null, loading: false, authenticated: false });
            }
        },

        login: async (email, password) => {
            const data = await api.login({ email, password });
            set({ user: data.user, loading: false, authenticated: true });
            return data;
        },

        register: async (userData) => {
            const data = await api.register(userData);
            set({ user: data.user, loading: false, authenticated: true });
            return data;
        },

        logout: async () => {
            try {
                await api.logout();
            } catch {
                // ignore
            }
            set({ user: null, loading: false, authenticated: false });
        },

        set,
    };
}

export const auth = createAuthStore();
