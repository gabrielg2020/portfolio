<script lang="ts">
  export let text: string;

  let isHovered = false;
  let isClicked = false;
  let mouseX = 0;
  let mouseY = 0;
  let buttonRect: DOMRect;

  function handleMouseEnter(event: MouseEvent) {
    isHovered = true;
    updateMousePosition(event);
  }

  function handleMouseLeave() {
    if (!isClicked) {
      isHovered = false;
    }
  }

  function handleMouseMove(event: MouseEvent) {
    if (isHovered && !isClicked) {
      updateMousePosition(event);
    }
  }

  function handleClick(event: MouseEvent) {
    isClicked = true;
    updateMousePosition(event);

    // Reset after animation completes
    setTimeout(() => {
      isClicked = false;
      isHovered = false;
    }, 600); // Match with CSS animation duration
  }

  function updateMousePosition(event: MouseEvent) {
    if (!buttonRect) {
      buttonRect = (event.currentTarget as HTMLElement).getBoundingClientRect();
    }

    mouseX = event.clientX - buttonRect.left;
    mouseY = event.clientY - buttonRect.top;
  }
</script>

<button
  on:mouseenter={handleMouseEnter}
  on:mouseleave={handleMouseLeave}
  on:mousemove={handleMouseMove}
  on:click={handleClick}
  class="animated-button"
>
  {#if isHovered || isClicked}
    <span
      class="circle"
      class:expand={isClicked}
      style="left: {mouseX}px; top: {mouseY}px;"
    ></span>
  {/if}

  <span class="button-text">{text}</span>
</button>

<style>
  .animated-button {
    display: flex;
    justify-content: center;
    position: relative;
    padding: 10px 20px;
    background-color: transparent;
    color: white;
    border: 3px solid white;
    cursor: pointer;
    overflow: hidden;
    font-size: 1.5rem;
    width: 100%;
    transition: color 0.4s;
  }

  .button-text {
    font-family: 'Martian Mono', monospace;
    position: relative;
    transition: color 0.4s;
    z-index: 2;
  }

  .circle {
    position: absolute;
    width: 0;
    height: 0;
    background-color: white;
    border-radius: 50%;
    transform: translate(-50%, -50%);
    z-index: 1;
    opacity: 0.4;
    transition:
      width 0.2s ease-out,
      height 0.2s ease-out;
  }

  .animated-button:hover .circle:not(.expand) {
    width: 20px;
    height: 20px;
  }

  .circle.expand {
    width: 600px;
    height: 600px;
    opacity: 1;
    transition:
      width 0.6s ease-out,
      height 0.6s ease-out,
      opacity 0.6s;
  }

  .animated-button:hover {
    border-color: white;
  }

  /* Change text color when button is clicked and circle expands */
  .animated-button:active,
  .circle.expand ~ .button-text {
    color: black;
  }
</style>
