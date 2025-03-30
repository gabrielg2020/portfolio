<script lang="ts">
  import Button from "./Button.svelte";
  import { Moon, Menu } from "@lucide/svelte";
  import NavMenu from "./NavMenu.svelte";
  import { slide } from "svelte/transition";

  let isMenuOpen = false;

  const toggleMenu = () => {
    isMenuOpen = !isMenuOpen;
  };

  const titleLinkMap = {
    "00 Home": "home",
    "01 Projects": "projects",
    "02 Experiences": "experiences",
    "03 Contact": "contact",
  };
</script>

<nav>
  <!-- Desktop View -->
  <div class="desktop-wrapper">
    <div class="nav-links">
      {#each Object.entries(titleLinkMap) as [title, link]}
        <a href="#{link}" class="sub-h2">{title}</a>
      {/each}
    </div>
    <div class="right-side">
      <div class="buttons">
        <Button text="Website Docs" location="home" />
        <Button text="GitHub Repo" location="home" />
      </div>
      <span class="icon action"><Moon /></span>
    </div>
  </div>

  <!-- Mobile View -->
  <div class="mobile-wrapper">
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div on:click={toggleMenu}>
      <span class="icon action"><Menu /></span>
    </div>
    <span class="icon action"><Moon /></span>
  </div>
</nav>

{#if isMenuOpen}
  <NavMenu {titleLinkMap} />
{/if}

<style>
  nav {
    position: sticky;
    top: 0;
  }

  .desktop-wrapper {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .nav-links {
    display: flex;
    gap: 2rem;
  }

  .right-side {
    display: flex;
    justify-content: right;
    gap: 2rem;
  }

  .buttons {
    display: flex;
    gap: 1rem;
  }

  .mobile-wrapper {
    display: none;
  }

  /* Mobile styles */
  @media (max-width: 1023px) {
    .desktop-wrapper {
      display: none;
    }

    .mobile-wrapper {
      display: flex;
      visibility: visible;
      justify-content: space-between;
      align-items: center;
    }
  }
</style>
