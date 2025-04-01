<script lang="ts">
  // Props
  export let languages: string[] = [];
  export let technologies: string[] = [];
</script>

<div class="space-y-2">
  <!-- Languages with dot separators -->
  {#if Array.isArray(languages) && languages.length > 0}
    <div class="flex flex-wrap gap-1">
      {#each languages as lang, index (index)}
        <span class="text-xs text-gray-600 dark:text-gray-400">
          {lang}{index < languages.length - 1 ? " •" : ""}
        </span>
      {/each}
    </div>
  {/if}
  
  <!-- Technologies as tags -->
  {#if  Array.isArray(technologies) && technologies.length > 0}
    <div class="flex flex-wrap gap-2">
      {#each technologies as tech, index (index)}
        <span class="inline-block text-xs px-2 py-1 bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300 rounded">
          {tech}
        </span>
      {/each}
    </div>
  {/if}
</div>