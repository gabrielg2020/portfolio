<script lang="ts">
  import ProjectCard from "../components/ProjectCard.svelte";
  import { onMount } from "svelte";
  import { GetProjects } from "../services/apiService";
  import type { Project } from "../services/apiService";

  let projects: Promise<Project[]>

  onMount(() => {
    projects = GetProjects()
  })
</script>

<div id="projects">
  <h1 class="h2">Projects</h1>
  <div class="projects">
    {#await projects}
      <p>Loading...</p>
    {:then loadedProjects}
      {#each loadedProjects as project}
      <ProjectCard 
        title={project.title} 
        smallDescription={project.smallDescription} 
        githubLink={project.githubLink} 
        languages={project.languages} 
        technologies={project.technologies} 
      />
      {/each}
    {:catch error}
      <p>Error loading projects: {error}</p>
    {/await}
  </div>
</div>

<style>
  .projects {
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 2rem;
    align-items: center;
    width: 100%;
  }
</style>
