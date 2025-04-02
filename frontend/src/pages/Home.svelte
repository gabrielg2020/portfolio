<script lang="ts">
  import Button from "../components/Button.svelte";
</script>

<style>
  @keyframes wave {
    0% { transform: rotate(0deg); }
    10% { transform: rotate(14deg); }
    20% { transform: rotate(-8deg); }
    30% { transform: rotate(14deg); }
    40% { transform: rotate(-4deg); }
    50% { transform: rotate(10deg); }
    60% { transform: rotate(0deg); }
    100% { transform: rotate(0deg); }
  }
  
  .waving-hand {
    display: inline-block;
    transform-origin: 70% 70%;
    animation: wave 2.5s ease-in-out infinite;
  }
</style>

<section 
  id="home" 
  class="min-h-screen pt-28 px-6 flex flex-col items-start justify-center"
>
  <div class="container mx-auto max-w-4xl">
    <p class="text-blue-600 dark:text-blue-400 font-medium mb-2">Hello <span class="waving-hand">👋</span>, I'm</p>
    
    <h1 class="text-5xl md:text-6xl font-bold mb-3 font-sans text-gray-900 dark:text-gray-100">
      Gabriel Guimaraes
    </h1>
    
    <p class="text-2xl md:text-3xl text-gray-700 dark:text-gray-300 mb-6">
      Software Engineer
    </p>
    
    <p class="text-lg text-gray-600 dark:text-gray-400 max-w-xl mb-10">
      Crafting creative, impactful solutions with clean code and clear purpose.
    </p>
    
    <div class="flex flex-wrap gap-4 mt-2">
      <Button text="View Projects" variant="primary" size="lg" href="#projects"/>
      <Button text="View Resume" variant="outline" size="lg" href="documents/resume.pdf"/>
    </div>
  </div>
</section>