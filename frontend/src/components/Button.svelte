<script lang="ts">
  import { ExternalLink, Download } from "@lucide/svelte";

  export let text: string = "Text";
  export let iconStr: string = "Link";
  export let location: string = "home";
  export let isInlineFlex: boolean = false;
  export let size: number = 15;
  export let strokeWidth: number = 2.2;
  $: icon = iconStr === "Link" ? ExternalLink : Download;
</script>

<a href="#{location}">
  <div class:inline-flex={isInlineFlex} class:flex={!isInlineFlex}>
    <span class="button">{text}</span>
    <span class="icon"
      ><svelte:component this={icon} {size} {strokeWidth} /></span
    >
  </div>
</a>

<style>
  div {
    align-items: center;
    justify-content: center;
    cursor: pointer;
    gap: 8px;
    border: 2px solid var(--text);
    padding: 4px;
  }

  .inline-flex {
    display: inline-flex !important;
  }

  .flex {
    display: flex !important;
  }
</style>
