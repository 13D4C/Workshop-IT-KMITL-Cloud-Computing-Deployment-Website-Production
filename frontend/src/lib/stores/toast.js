import { writable } from 'svelte/store';

function createToastStore() {
    const { subscribe, update } = writable([]);

    let id = 0;

    function add(message, type = 'info', duration = 4000) {
        const toastId = id++;
        update(toasts => [...toasts, { id: toastId, message, type }]);
        
        setTimeout(() => {
            remove(toastId);
        }, duration);
    }

    function remove(toastId) {
        update(toasts => toasts.filter(t => t.id !== toastId));
    }

    return {
        subscribe,
        success: (msg) => add(msg, 'success'),
        error: (msg) => add(msg, 'error'),
        info: (msg) => add(msg, 'info'),
        warning: (msg) => add(msg, 'warning'),
        remove,
    };
}

export const toast = createToastStore();
