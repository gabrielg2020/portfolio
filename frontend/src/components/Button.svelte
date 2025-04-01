<script lang="ts">
  import { ArrowRight } from "@lucide/svelte";

  // Props
  export let text: string = "Button";
  export let variant: "primary" | "outline" = "primary";
  export let size: "sm" | "md" | "lg" = "md";
  export let icon: typeof ArrowRight | null = null;
  export let href: string | null = null;

  // Element to be rendered (button or anchor)
  const element = href ? "a" : "button";

  // Internal state
  let isHovered = false;

  // Size styles mapping
  const sizeStyles = {
    sm: "text-xs py-2 px-3",
    md: "text-sm py-2.5 px-4",
    lg: "text-base py-3 px-5",
  };

  const variantStyles = {
    primary:
      "bg-blue-600 text-white border border-blue-600 hover:bg-blue-700 hover:border-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:bg-blue-700 dark:border-blue-700 dark:hover:bg-blue-600 dark:hover:border-blue-600 dark:focus:ring-blue-600 dark:focus:ring-offset-gray-900",
    outline:
      "bg-transparent text-gray-800 border border-gray-300 hover:bg-gray-50 focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 dark:text-gray-200 dark:border-gray-600 dark:hover:bg-gray-800 dark:focus:ring-gray-500 dark:focus:ring-offset-gray-900"
  };

  // Combined classes
  $: classes = `
    ${variantStyles[variant]}
    ${sizeStyles[size]}
    rounded-md font-medium transition-all duration-200
    inline-flex items-center justify-center gap-2
    focus:outline-none
  `;

  function handleHover(isHovering: boolean) {
    isHovered = isHovering;
  }

  // Determine if we should show the arrow icon
  $: showArrow = variant === "primary" && !icon;
</script>

<!-- svelte-ignore a11y-missing-attribute -->
<svelte:element
  this={element}
  {href}
  class={classes}
  on:click
  on:mouseenter={() => handleHover(true)}
  on:mouseleave={() => handleHover(false)}
  on:focus
  on:blur
  {...$$restProps}
>
  <!-- Button text -->
  <span>{text}</span>

  <!-- Right icon slot -->
  {#if icon}
    <span class="inline-flex">
      <svelte:component
        this={icon}
        size={size === "sm" ? 14 : size === "md" ? 16 : 18}
      />
    </span>
  {/if}

  <!-- Arrow icon for primary buttons without custom icon -->
  {#if showArrow}
    <span
      class="inline-flex {isHovered
        ? 'translate-x-0.5'
        : 'translate-x-0'} transition-transform duration-200"
    >
      <ArrowRight size={size === "sm" ? 14 : size === "md" ? 16 : 18} />
    </span>
  {/if}
</svelte:element>