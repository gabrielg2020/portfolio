<script lang="ts">
  import Button from "../components/Button.svelte";
</script>

<section 
  id="home" 
  class="min-h-screen pt-28 px-6 flex flex-col items-start justify-center"
>
  <div class="container mx-auto max-w-4xl">
    <p class="font-medium mb-2">Hello 👋, I'm</p>
    
    <h1 class="text-5xl md:text-6xl font-bold mb-3 font-sans text-gray-900">
      Gabriel Guimaraes
    </h1>
    
    <p class="text-2xl md:text-3xl text-gray-700 mb-6">
      Software Engineer
    </p>
    
    <p class="text-lg text-gray-600 max-w-xl mb-10">
      Crafting creative, impactful solutions with clean code and clear purpose.
    </p>
    
    <div class="flex flex-wrap gap-4 mt-2">
      <Button text="View Projects" variant="primary" size="lg" href="#projects"/>
      <Button text="View Resume" variant="outline" size="lg"/>
    </div>
  </div>
</section>