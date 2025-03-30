<script lang="ts">
  import { Box, Code } from "@lucide/svelte";
  import { onMount } from "svelte";
  export let text: string = "Text";
  export let iconStr: string = "code";
  $: icon = iconStr === "code" ? Code : Box;

  let isDesktop = false;

  function updateLayout() {
    isDesktop = window.innerWidth >= 1024;
  }

  onMount(() => {
    updateLayout();
    window.addEventListener("resize", updateLayout);

    return () => {
      window.removeEventListener("resize", updateLayout);
    };
  });
</script>

<div>
  <span class="icon"><svelte:component this={icon} size={isDesktop ? 17 : 13}/></span>
  <span class="overline">{text}</span>
</div>

<style>
  div {
    display: inline-flex;
    align-items: center;
    gap: 4px;
    padding: 4px;
    color: var(--info-blue);
    background-color: var(--info-blue-transparent);
    box-shadow: inset 0 0 0 1px var(--info-blue);
    border-radius: 20px;
  }

  .overline {
    font-size: 0.875rem;
  }

  @media (max-width: 1024px) {
    .overline {
      font-size: 0.625rem;
    }
  }
</style>