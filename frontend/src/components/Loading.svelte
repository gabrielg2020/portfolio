<script lang="ts">
  export let message: string = "Loading...";
</script>

<div class="py-16 bg-gray-50 dark:bg-gray-800 transition-colors">
  <div class="container mx-auto px-6">
    <div class="flex flex-col items-center justify-center min-h-[300px]">
      <!-- Spinner -->
      <div class="w-12 h-12 mb-4 relative">
        <div class="absolute top-0 left-0 right-0 bottom-0 rounded-full border-4 border-gray-200 dark:border-gray-700"></div>
        <div class="absolute top-0 left-0 right-0 bottom-0 rounded-full border-4 border-blue-600 dark:border-blue-500 border-t-transparent animate-spin"></div>
      </div>
      
      <!-- Message -->
      <p class="text-gray-600 dark:text-gray-300 text-lg">{message}</p>
    </div>
  </div>
</div>