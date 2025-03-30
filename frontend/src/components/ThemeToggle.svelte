<script lang="ts">
  import { theme } from "../store/theme";
  import { Moon, Sun } from "@lucide/svelte";
</script>

<button on:click={theme.toggle} aria-label="Toggle theme">
  {#if $theme === 'dark'}
    <Moon />
  {:else}
    <Sun />
  {/if}
</button>

<style>
  button {
    display: flex;
    align-items: center;
    justify-content: center;
    background: none;
    border: none;
    cursor: pointer;
    font-size: 1.5rem;
    padding: 0.5rem;
    border-radius: 50%;
    background-color: var(--background-transparent);
    color: var(--text);
  }
  
  button:hover {
    background-color: var(--text-transparent);
    color: var(--background);
  }
</style>