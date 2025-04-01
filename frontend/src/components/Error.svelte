<script lang="ts">
  import { AlertTriangle } from "@lucide/svelte";
  
  export let message: string = "An error occurred.";
  export let retryFn: (() => void) | null = null;
</script>

<div class="py-16 bg-gray-50 dark:bg-gray-800 transition-colors">
  <div class="container mx-auto px-6">
    <div class="max-w-2xl mx-auto">
      <div class="bg-white dark:bg-gray-900 rounded-lg border border-red-200 dark:border-red-900 shadow-sm p-6">
        <div class="flex flex-col items-center sm:flex-row sm:items-start text-center sm:text-left">
          <div class="bg-red-100 dark:bg-red-900/20 p-3 rounded-full mb-4 sm:mb-0 sm:mr-4 flex-shrink-0">
            <AlertTriangle size={28} class="text-red-600 dark:text-red-500" />
          </div>
          
          <div>
            <h3 class="text-lg font-semibold text-gray-800 dark:text-gray-200 mb-2">Error</h3>
            <p class="text-gray-600 dark:text-gray-300 mb-4">{message}</p>
            
            {#if retryFn}
              <button 
                on:click={retryFn}
                class="inline-flex items-center px-4 py-2 bg-blue-600 dark:bg-blue-700 text-white rounded-md hover:bg-blue-700 dark:hover:bg-blue-600 transition-colors"
              >
                Try Again
              </button>
            {/if}
          </div>
        </div>
      </div>
    </div>
  </div>
</div>