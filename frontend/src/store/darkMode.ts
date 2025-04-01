import { writable } from 'svelte/store';

// Check if localStorage is available (client-side)
const isClient = typeof window !== 'undefined';

// Try to get the user's preference from localStorage
const getInitialTheme = (): boolean => {
  if (!isClient) return false;
  
  // Check localStorage first
  const storedTheme = localStorage.getItem('darkMode');
  if (storedTheme !== null) {
    return storedTheme === 'true';
  }
  
  // If not in localStorage, check user's system preference
  if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
    return true;
  }
  
  // Default to light mode
  return false;
};

// Create writable store with initial value
export const darkMode = writable<boolean>(false);

// Initialize store value when in browser environment
if (isClient) {
  darkMode.set(getInitialTheme());
  
  // Update localStorage when the value changes
  darkMode.subscribe(value => {
    localStorage.setItem('darkMode', String(value));
    
    // Apply theme to document
    if (value) {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  });
}

// Toggle function
export function toggleDarkMode(): void {
  darkMode.update(value => !value);
}