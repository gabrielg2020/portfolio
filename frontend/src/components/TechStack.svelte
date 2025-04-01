<script lang="ts">
  // Props
  export let languages: string[] = [];
  export let technologies: string[] = [];
</script>

<div class="space-y-2">
  <!-- Languages with dot separators -->
  {#if languages.length > 0}
    <div class="flex flex-wrap gap-2">
      {#each languages as lang, index}
        <span class="text-xs text-gray-600">
          {lang}{index < languages.length - 1 ? " •" : ""}
        </span>
      {/each}
    </div>
  {/if}
  
  <!-- Technologies as tags -->
  {#if technologies.length > 0}
    <div class="flex flex-wrap gap-2">
      {#each technologies as tech}
        <span class="inline-block text-xs px-2 py-1 bg-gray-100 text-gray-700 rounded">
          {tech}
        </span>
      {/each}
    </div>
  {/if}
</div>