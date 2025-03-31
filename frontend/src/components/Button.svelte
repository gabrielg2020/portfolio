<script lang="ts">
  import { ArrowRight } from "@lucide/svelte";

  export let text: string = "Button";
  export let variant: string = "primary";
  let isHovered = false;

  const getVariantStyles = () => {
    switch (variant) {
      case "primary":
        return "bg-blue-600 text-white hover:bg-blue-700";
      case "blank":
        return "bg-white text-gray-800 border border-gray-300 hover:bg-gray-50";
    }
  };

  const getTextStyles = () => {
    switch (variant) {
      case "primary":
        return "pl-4";
      case "blank":
        return "px-4";
    }
  };
</script>

<button
  class="rounded-md flex items-center justify-center gap-2 font-medium transition-all duration-200 {getVariantStyles()}"
  on:mouseenter={() => (isHovered = true)}
  on:mouseleave={() => (isHovered = false)}
>
  <span class="text-base font-medium py-3 {getTextStyles()}">{text}</span>
  {#if variant == "primary"}
    <span class="pr-4 {isHovered ? "translate-x-1" : "translate-x-0"} transition-transform duration-200"
      ><ArrowRight size="16" /></span
    >
  {/if}
</button>
