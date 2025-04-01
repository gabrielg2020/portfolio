<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { ChevronLeft, ChevronRight } from '@lucide/svelte';
  
  // Props
  export let title: string = "Carousel";
  export let totalItems: number;
  export let itemsPerSlideDesktop: number = 3;
  export let itemsPerSlideTablet: number = 2;
  export let itemsPerSlideMobile: number = 1;
  
  // State
  let currentIndex: number = 0;
  let isTransitioning: boolean = false;
  let touchStart: number = 0;
  let touchEnd: number = 0;
  let itemsPerSlide: number = itemsPerSlideDesktop;
  let totalSlides: number = Math.ceil(totalItems / itemsPerSlide);
  let carouselRef: HTMLElement;
  
  // Calculate how many slides to show based on screen size
  const updateItemsPerSlide = () => {
    if (typeof window !== 'undefined') {
      if (window.innerWidth < 640) {
        itemsPerSlide = itemsPerSlideMobile;
      } else if (window.innerWidth < 1024) {
        itemsPerSlide = itemsPerSlideTablet;
      } else {
        itemsPerSlide = itemsPerSlideDesktop;
      }
      totalSlides = Math.ceil(totalItems / itemsPerSlide);
      
      // Make sure current index is still valid
      if (currentIndex >= totalSlides) {
        currentIndex = totalSlides - 1;
      }
    }
  };
  
  // Navigation functions
  const goToSlide = (index: number) => {
    if (isTransitioning) return;
    
    // Handle wrapping
    let newIndex = index;
    if (index < 0) newIndex = totalSlides - 1;
    if (index >= totalSlides) newIndex = 0;
    
    isTransitioning = true;
    currentIndex = newIndex;
    
    // Reset transition flag
    setTimeout(() => {
      isTransitioning = false;
    }, 500); // Match this with CSS transition time
  };
  
  const nextSlide = () => {
    goToSlide(currentIndex + 1);
  };
  
  const prevSlide = () => {
    goToSlide(currentIndex - 1);
  };
  
  // Touch event handlers
  const handleTouchStart = (e: TouchEvent) => {
    touchStart = e.touches[0].clientX;
  };
  
  const handleTouchMove = (e: TouchEvent) => {
    touchEnd = e.touches[0].clientX;
  };
  
  const handleTouchEnd = () => {
    if (touchStart - touchEnd > 75) {
      // Swipe left
      nextSlide();
    }
    
    if (touchStart - touchEnd < -75) {
      // Swipe right
      prevSlide();
    }
  };
  
  // Keyboard navigation
  const handleKeyDown = (e: KeyboardEvent) => {
    if (e.key === 'ArrowLeft' || e.key === 'h') {
      prevSlide();
    } else if (e.key === 'ArrowRight' || e.key === 'l') {
      nextSlide();
    }
  };
  
  // Lifecycle
  onMount(() => {
    updateItemsPerSlide();
    window.addEventListener('resize', updateItemsPerSlide);
    window.addEventListener('keydown', handleKeyDown);
    
    return () => {
      window.removeEventListener('resize', updateItemsPerSlide);
      window.removeEventListener('keydown', handleKeyDown);
    };
  });
</script>

<section class="py-16 bg-gray-50">
  <div class="container mx-auto px-6">
    <div class="flex justify-between items-center mb-8">
      <h2 class="text-2xl font-bold text-gray-800">{title}</h2>
      <div class="flex space-x-2">
        <button 
          on:click={prevSlide}
          disabled={isTransitioning}
          class="p-2 rounded-full bg-white border border-gray-200 text-gray-600 hover:bg-gray-100 hover:text-blue-600 transition-colors disabled:opacity-50"
          aria-label="Previous items"
        >
          <ChevronLeft size={20} />
        </button>
        <button 
          on:click={nextSlide}
          disabled={isTransitioning}
          class="p-2 rounded-full bg-white border border-gray-200 text-gray-600 hover:bg-gray-100 hover:text-blue-600 transition-colors disabled:opacity-50"
          aria-label="Next items"
        >
          <ChevronRight size={20} />
        </button>
      </div>
    </div>
    
    <!-- Carousel wrapper -->
    <div 
      class="relative overflow-hidden"
      bind:this={carouselRef}
      on:touchstart={handleTouchStart}
      on:touchmove={handleTouchMove}
      on:touchend={handleTouchEnd}
    >
      <div 
        class="flex transition-transform duration-500 ease-in-out"
        style="transform: translateX(-{currentIndex * 100}%)"
      >
        <!-- Generate slides -->
        {#each Array(totalSlides) as _, slideIndex}
          <div 
            class="w-full flex-shrink-0 flex gap-4"
            aria-hidden={currentIndex !== slideIndex}
          >
            <slot name="slide" {slideIndex} {itemsPerSlide} />
          </div>
        {/each}
      </div>
    </div>
    
    <!-- Dots navigation -->
    {#if totalSlides > 1}
      <div class="flex justify-center mt-6 space-x-2">
        {#each Array(totalSlides) as _, index}
          <!-- svelte-ignore element_invalid_self_closing_tag -->
          <button
            on:click={() => goToSlide(index)}
            class="h-2 rounded-full transition-all {
              index === currentIndex 
                ? 'bg-blue-600 w-4' 
                : 'bg-gray-300 hover:bg-gray-400 w-2'
            }"
            aria-label="Go to slide {index + 1}"
            aria-current={index === currentIndex ? 'true' : 'false'}
          />
        {/each}
      </div>
    {/if}
  </div>
</section>