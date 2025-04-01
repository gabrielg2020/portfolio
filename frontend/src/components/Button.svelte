<script lang="ts">
  import { ExternalLink, Download } from "@lucide/svelte";

  export let text: string = "Text";
  export let iconStr: string = "Link";
  export let location: string = "home";
  export let isInlineFlex: boolean = true;
  export let size:number = 15;
  export let strokeWidth:number = 2.2;
  $: icon = iconStr === "Link" ? ExternalLink : Download;
</script>

<a href="#{location}">
  <div>
    <span class="button" class:inlineFlex={isInlineFlex}>{text}</span>
    <span class="icon"><svelte:component this={icon} size={size} strokeWidth={strokeWidth}/></span>
  </div>
</a>

<style>
  div {
    display: flex;
    align-items: center;
    cursor: pointer;
    gap: 8px;
    border: 2px solid var(--near-black);
    padding: 4px;
  }

  .inlineFlex {
    display: inline-flex !important;
  }
  
</style>
