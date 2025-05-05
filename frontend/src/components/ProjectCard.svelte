<script lang="ts">
  import { Github, ExternalLink } from "@lucide/svelte";
  import TechStack from "./TechStack.svelte";

  // Props
  export let title: string = "Project Title";
  export let description: string =
    "A short description of the project that explains what it does and any notable features.";
  export let languages: string[] = ["TypeScript", "Go", "Python"];
  export let technologies: string[] = ["Docker", "Svelte", "SQLite"];
  export let githubUrl: string = "#";
  export let liveUrl: string | null = null;
</script>

<div
  class="p-5 bg-white dark:bg-gray-900 border-l-2 border-blue-600 transition-all duration-100 hover:border-l-4 h-full flex flex-col"
>
  <!-- Header with links -->
  <div class="flex justify-between items-start mb-3">
    <h3 class="text-lg font-semibold text-gray-800 dark:text-gray-100">
      {title}
    </h3>
    <div class="flex space-x-3">
      <a
        href={githubUrl}
        target="_blank"
        rel="noopener noreferrer"
        aria-label="GitHub Repository"
        class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
      >
        <Github size={18} />
      </a>

      {#if liveUrl}
        <a
          href={liveUrl}
          target="_blank"
          rel="noopener noreferrer"
          aria-label="Live Demo"
          class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
        >
          <ExternalLink size={18} />
        </a>
      {/if}
    </div>
  </div>

  <!-- Description -->
  <p class="text-sm text-gray-600 dark:text-gray-300 mb-4 flex-grow">
    {description}
  </p>

  <!-- Tech stack -->
  <div class="mt-auto">
    <TechStack {languages} {technologies} />
  </div>
</div>
