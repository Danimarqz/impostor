import { writable } from 'svelte/store';
import type { Language } from '../lib/i18n';

// Initialize from localStorage or default to 'en'
const storedLanguage = (typeof window !== 'undefined' ? localStorage.getItem('language') : null) as Language | null;
const initialLanguage: Language = storedLanguage || 'en';

// Language store with localStorage sync
export const language = writable<Language>(initialLanguage);

// Subscribe to changes and save to localStorage
if (typeof window !== 'undefined') {
  language.subscribe(value => {
    localStorage.setItem('language', value);
  });
}

// Helper to toggle language
export function toggleLanguage() {
  language.update(lang => lang === 'en' ? 'es' : 'en');
}
