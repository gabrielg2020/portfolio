<script lang="ts">
  import { Github, FileText } from "@lucide/svelte";
  
  // Props
  export let docsUrl: string = "/docs";
  export let repoUrl: string = "https://github.com/yourusername/your-repo";
  export let copyrightName: string = "Your Name";
  export let startYear: number = new Date().getFullYear();
  
  // Calculate current year for copyright
  const currentYear = new Date().getFullYear();
  
  // Build copyright text
  $: copyrightText = startYear === currentYear 
    ? `© ${currentYear} ${copyrightName}` 
    : `© ${startYear}–${currentYear} ${copyrightName}`;
</script>

<footer class="bg-gray-900 dark:bg-gray-950 text-gray-300 py-8 mt-auto transition-colors duration-200">
  <div class="container mx-auto px-6">
    <div class="flex flex-col md:flex-row justify-between items-center">
      <!-- Left side: Copyright -->
      <div class="mb-4 md:mb-0">
        <p class="text-sm">{copyrightText}</p>
      </div>
      
      <!-- Right side: Links -->
      <div class="flex space-x-6">
        <a 
          href={docsUrl}
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center text-gray-300 hover:text-white dark:hover:text-blue-400 transition-colors text-sm"
        >
          <FileText class="h-4 w-4 mr-2" />
          <span>Documentation</span>
        </a>
        
        <a 
          href={repoUrl}
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center text-gray-300 hover:text-white dark:hover:text-blue-400 transition-colors text-sm"
        >
          <Github class="h-4 w-4 mr-2" />
          <span>GitHub</span>
        </a>
      </div>
    </div>
  </div>
</footer>