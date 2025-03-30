import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

interface ThemeStore extends Writable<string> {
  toggle: () => void;
}

// Initialize theme from localStorage (if available) or system preference
function createThemeStore(): ThemeStore {
  // Get the initial value from localStorage or system preference
  const getInitialTheme = (): string => {
    // Check if we're in a browser environment
    if (typeof window !== 'undefined') {
      // Check localStorage first
      const storedTheme = localStorage.getItem('theme');
      if (storedTheme) {
        return storedTheme;
      }
      
      // If no stored preference, check system preference
      if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
        return 'dark';
      }
    }
    
    // Default to light if not in browser or no preference
    return 'light';
  };

  const theme = writable<string>(getInitialTheme());
  
  // Apply theme effect
  if (typeof window !== 'undefined') {
    // Initialize theme on page load
    const currentTheme = getInitialTheme();
    if (currentTheme === 'dark') {
      document.documentElement.classList.add('dark-theme');
    }
    
    theme.subscribe(value => {
      // Update localStorage
      localStorage.setItem('theme', value);
      
      // Apply/remove class on html element
      if (value === 'dark') {
        document.documentElement.classList.add('dark-theme');
      } else {
        document.documentElement.classList.remove('dark-theme');
      }
    });
    
    // Listen for system preference changes
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', e => {
      if (!localStorage.getItem('theme')) { // Only auto-switch if user hasn't set a preference
        theme.set(e.matches ? 'dark' : 'light');
      }
    });
  }

  // Return writable store with custom toggle method
  return {
    ...theme,
    toggle: () => {
      theme.update(current => current === 'dark' ? 'light' : 'dark');
    }
  };
}

export const theme = createThemeStore();


