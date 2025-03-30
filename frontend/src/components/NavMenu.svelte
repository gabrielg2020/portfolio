<script lang="ts">
  import Button from "./Button.svelte";
  import { X } from "@lucide/svelte";

  let closed: boolean = false;

  const closeMenu = () => {
    closed = true;
  };

  export let titleLinkMap: object = {
    "00 Home": "home",
    "01 Projects": "projects",
    "02 Experiences": "experiences",
    "03 Contact": "contact",
  };
</script>

{#if closed == false}
  <div class="nav-menu">
    <div class="header">
      <h1 class="h5">Portfolio</h1>
      <!-- svelte-ignore a11y_no_static_element_interactions -->
      <!-- svelte-ignore a11y_click_events_have_key_events -->
      <span class="icon action" on:click={closeMenu}><X size="32" /></span>
    </div>
    <div class="sections">
      {#each Object.entries(titleLinkMap) as [title, link]}
        <a href="#{link}" class="h6">{title}</a>
      {/each}
    </div>
    <div class="footer">
      <Button text="Website Docs" />
      <Button text="GitHub Repo" />
    </div>
  </div>

  <div class="blur-background"></div>
{/if}

<style>
  .nav-menu {
    /* Positoning */
    position: fixed;
    top: 0;
    left: 0;
    z-index: 100;

    /* Sizing */
    box-sizing: border-box;
    width: 250px;
    height: 100vh;
    padding: 16px;
    
    /* Children's Positioning */
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  
    /* Styling */
    background-color: var(--background);
    box-shadow: -2px 0 10px rgba(0, 0, 0, 0.1);
    box-sizing: border-box;
  }

  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .sections {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .footer {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .blur-background {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 99;
    width: 100vw;
    height: 100vh;
    background-color: var(--background-transparent);
    backdrop-filter: blur(2px);
    -webkit-backdrop-filter: blur(2px); 
  }
</style>
