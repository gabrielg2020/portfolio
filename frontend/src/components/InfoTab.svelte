<script lang="ts">
  import { Box, Code } from "@lucide/svelte";
  export let text: string = "Text";
  export let iconStr: string = "code";
  $: icon = iconStr === "code" ? Code : Box;
</script>

<div>
  <span class="icon"><svelte:component this={icon} size=11/></span>
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
</style>