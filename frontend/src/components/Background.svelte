<script lang="ts">
  import { onMount } from "svelte";

  // Props
  export let backgroundColor = "#121212";

  let canvas: HTMLCanvasElement;
  let ctx: CanvasRenderingContext2D;
  let width: number;
  let height: number;

  onMount(() => {
    // Initialise Canvas
    ctx = canvas.getContext("2d");
    resizeCanvas();

    // Add event listeners
    window.addEventListener("resize", resizeCanvas);

    return () => {
      window.removeEventListener("resize", resizeCanvas);
    };
  });

  function resizeCanvas() {
    const container = canvas.parentElement;
    width = container.clientWidth;
    height = container.clientHeight;
    canvas.width = width;
    canvas.height = height;
  }
</script>

<div class="container">
  <canvas bind:this={canvas}></canvas>
  <slot />
</div>

<style>
</style>
