<script lang="ts">
  export let message: string = "Loading...";
</script>

<style>
  @keyframes bounce {
    0%, 100% { transform: translateY(0); }
    50% { transform: translateY(-10px); }
  }
  
  .dot {
    display: inline-block;
    width: 10px;
    height: 10px;
    margin: 0 4px;
    border-radius: 50%;
  }
  
  .dot-1 { animation: bounce 1s ease infinite; }
  .dot-2 { animation: bounce 1s ease infinite 0.2s; }
  .dot-3 { animation: bounce 1s ease infinite 0.4s; }
</style>

<div class="py-16 bg-gray-50 dark:bg-gray-800 transition-colors">
  <div class="container mx-auto px-6">
    <div class="flex flex-col items-center justify-center min-h-[300px]">
      <!-- Message -->
      <p class="text-gray-600 dark:text-gray-300 text-lg mb-4">{message}</p>
      
      <!-- Bouncing dots -->
      <div class="flex items-center">
        <div class="dot dot-1 bg-blue-600 dark:bg-blue-500"></div>
        <div class="dot dot-2 bg-blue-600 dark:bg-blue-500"></div>
        <div class="dot dot-3 bg-blue-600 dark:bg-blue-500"></div>
      </div>
    </div>
  </div>
</div>