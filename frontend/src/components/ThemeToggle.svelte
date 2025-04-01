<script lang="ts">
  import { Moon, Sun } from '@lucide/svelte';
  import { darkMode, toggleDarkMode } from '../scripts/store/darkMode';

  // Props
  export let size: number = 20;
  export let showLabel: boolean = false;
</script>

<button
  on:click={toggleDarkMode}
  class="flex items-center gap-2 text-gray-500 hover:text-blue-600 dark:text-gray-400 dark:hover:text-blue-400 transition-colors p-1 rounded-md"
  aria-label="Toggle dark mode"
  title="Toggle dark mode"
>
  {#if $darkMode}
    <Sun size={size} />
    {#if showLabel}
      <span>Light Mode</span>
    {/if}
  {:else}
    <Moon size={size} />
    {#if showLabel}
      <span>Dark Mode</span>
    {/if}
  {/if}
</button>