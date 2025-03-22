<script lang="ts">
  import { ArrowDown } from "@lucide/svelte";

  export let text: string = "Button";
  export let location: string = "home";
  const icon = ArrowDown;
</script>

<a href="#{location}">
  <div>
    <span>{text}</span>
    <span class="icon"><svelte:component this={icon}/></span>
  </div>
</a>

<style>
  div {
    display: flex;
    justify-items: center;
    align-items: center;
    border-radius: 4px;
    border: 1px solid red;
    cursor: pointer;
    font-size: 1.1rem;
    font-family: "Martian Mono", monospace;
  }

  .icon {
    display: flex;
    align-items: center;
    justify-items: center;
  }
</style>
