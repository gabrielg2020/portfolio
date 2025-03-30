<script lang="ts">
  import { MoveDown } from "@lucide/svelte";
  import { onMount } from "svelte";
  
  // Define size variants
  const sizes = {
    desktop: {
      titleClass: "h2",
      subtitleClass: "h3",
      ctaClass: "h4",
      iconSize: 64
    },
    tablet: {
      titleClass: "h3",
      subtitleClass: "h4",
      ctaClass: "h5",
      iconSize: 32
    },
    mobile: {
      titleClass: "h5",
      subtitleClass: "h6",
      ctaClass: "sub-h1",
      iconSize: 32
    }
  };
  
  // Default to desktop - will update on mount and resize
  let activeSize = sizes.desktop;
  
  function updateSize() {
    const width = window.innerWidth;
    if (width <= 767) {
      activeSize = sizes.mobile;
    } else if (width <= 1023) {
      activeSize = sizes.tablet;
    } else {
      activeSize = sizes.desktop;
    }
  }
  
  onMount(() => {
    updateSize();
    window.addEventListener('resize', updateSize);
    
    return () => {
      window.removeEventListener('resize', updateSize);
    };
  });
</script>

<div class="hero">
  <div class="wrapper">
    <div class="header">
      <h1 class={activeSize.titleClass}>Gabriel Guimaraes</h1>
      <h2 class={activeSize.subtitleClass}>Software Engineer</h2>
    </div>
    <div class="footer">
      <h4 class={activeSize.ctaClass}>Scroll to see my projects!</h4>
      <div class="arrow-container">
        <MoveDown size={activeSize.iconSize} />
      </div>
    </div>
  </div>
</div>

<style>
  .hero {
    display: flex;
    justify-content: center;
    position: absolute;
    left: 50%;
    top: 40%;
    -webkit-transform: translate(-50%, -50%);
    transform: translate(-50%, -50%);
    width: 100%;
  }

  .wrapper {
    display: flex;
    flex-direction: column;
    margin-top: 2rem
  }

  .header {
    display: flex;
    flex-direction: column;
    align-items: baseline;
    padding-bottom: 10vh;
  }

  .header > h1 {
    padding-bottom: 1.5rem;
  }

  .footer {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 2rem;
  }
  
  /* Arrow animation styles */
  .arrow-container {
    animation: float 1.5s ease-in-out infinite;
  }
  
  @keyframes float {
    0% {
      transform: translateY(0px);
    }
    50% {
      transform: translateY(-10px);
    }
    100% {
      transform: translateY(0px);
    }
  }
</style>