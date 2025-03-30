<script lang="ts">
  import Button from "./Button.svelte";
  import { Menu } from "@lucide/svelte";
  import ThemeToggle from "./ThemeToggle.svelte";
  import NavMenu from "./NavMenu.svelte";
  import { onMount } from "svelte";

  let isMenuOpen = false;
  let isMobile = false;

  const toggleMenu = () => {
    isMenuOpen = !isMenuOpen;
  };

  const titleLinkMap = {
    "00 Home": "home",
    "01 Projects": "projects",
    "02 Experiences": "experiences",
    "03 Contact": "contact",
  };
  
  // Update mobile state based on screen size
  function updateLayout() {
    isMobile = window.innerWidth <= 1023;
    
    // Close the menu when transitioning from mobile to desktop
    if (!isMobile) {
      isMenuOpen = false;
    }
  }
  
  onMount(() => {
    updateLayout();
    window.addEventListener('resize', updateLayout);
    
    return () => {
      window.removeEventListener('resize', updateLayout);
    };
  });
</script>

<nav class:mobile={isMobile}>
  <div class="nav-container">
    <!-- Navigation Links (visible on desktop) -->
    {#if !isMobile}
      <div class="nav-links">
        {#each Object.entries(titleLinkMap) as [title, link]}
          <a href="#{link}" class="sub-h2">{title}</a>
        {/each}
      </div>
    {/if}
    
    <!-- Right Side Content -->
    <div class="right-side">
      {#if isMobile}
        <!-- Mobile Menu Button -->
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <div 
          class="menu-button" 
          on:click={toggleMenu}
          role="button"
          tabindex="0"
          aria-label="Toggle menu"
        >
          <Menu />
        </div>
      {:else}
        <!-- Desktop Buttons -->
        <div class="buttons">
          <Button text="Website Docs" location="home" />
          <Button text="GitHub Repo" location="home" />
        </div>
      {/if}
      
      <!-- Theme Toggle (always visible) -->
      <ThemeToggle />
    </div>
  </div>
</nav>

<!-- Mobile Menu (conditionally rendered) -->
{#if isMenuOpen && isMobile}
  <NavMenu {titleLinkMap} />
{/if}

<style>
  nav {
    position: sticky;
    top: 0;
    width: 100%;
    z-index: 100;
  }

  .nav-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    width: 100%;
    max-width: 1024px;
    margin: 0 auto;
  }

  .nav-links {
    display: flex;
    gap: 2rem;
  }

  .right-side {
    display: flex;
    align-items: center;
    gap: 2rem;
    
    /* When in mobile view, make the right-side take full width */
    :global(.mobile) & {
      width: 100%;
      justify-content: space-between;
    }
  }

  .buttons {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .menu-button {
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  /* Responsive adjustments handled by JS logic instead of CSS media queries */
</style>