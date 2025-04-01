<script lang="ts">
  import { Briefcase, GraduationCap, Calendar } from "@lucide/svelte";
  import TechStack from "./TechStack.svelte";
  
  // Props
  export let type: "work" | "education" = "work";
  export let organisation: string = "Company Name";
  export let role: string = "Job Title";
  export let startYear: string = "2020";
  export let endYear: string = "Present";
  export let description: string = "Description of responsibilities and achievements in this role.";
  export let languages: string[] = ["TypeScript", "Go", "Python"];
  export let technologies: string[] = ["Docker", "Svelte", "SQLite"];
  export let isLast: boolean = false;
</script>

<div class="relative pl-8 pb-12 group">
  <!-- Timeline line -->
  <div class="absolute left-0 top-0 h-full w-px bg-gray-200 dark:bg-gray-700 {isLast ? 'h-6' : ''}"></div>
  
  <!-- Timeline dot -->
  <div class="absolute left-0 top-0 -translate-x-1/2 w-6 h-6 rounded-full border-2 border-blue-600 dark:border-blue-500 bg-white dark:bg-gray-800 flex items-center justify-center">
    {#if type === "education"}
      <GraduationCap size={14} class="text-blue-600 dark:text-blue-400" />
    {:else}
      <Briefcase size={14} class="text-blue-600 dark:text-blue-400" />
    {/if}
  </div>
  
  <!-- Content -->
  <div>
    <!-- Header -->
    <div class="flex flex-col sm:flex-row sm:justify-between sm:items-start mb-2">
      <div>
        <h3 class="text-lg font-semibold text-gray-800 dark:text-gray-100">{role}</h3>
        <p class="text-sm text-gray-600 dark:text-gray-400">{organisation}</p>
      </div>
      <div class="flex items-center mt-1 sm:mt-0 text-sm text-gray-500 dark:text-gray-400">
        <Calendar size={14} class="mr-1" />
        <span>{startYear} — {endYear}</span>
      </div>
    </div>
    
    <!-- Description -->
    <p class="text-sm text-gray-600 dark:text-gray-300 mb-3">{description}</p>
    
    <!-- Skills -->
    <TechStack {languages} {technologies} />
  </div>
</div>