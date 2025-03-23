<script lang="ts">
  import Project from "../components/Project.svelte";

  import { onMount } from "svelte";

  let projects: { id:number, description: string; githubLink: string; title: string }[] =
    [];
  let error: string | null = null;

  onMount(async () => {
    try {
      console.log("Requesting");
      const response = await fetch("/api/projects");
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      projects = data.projects;
      console.log(projects);
    } catch (err) {
      console.error("Failed to fetch projects:", err);
      if (err instanceof Error) {
        error = err.message;
      } else {
        error = String(err);
      }
    }
  });
</script>

<div id="projects">
  <div class="header">
    <h1>Project Showcase</h1>
    <h4>Open them up to learn more!</h4>
  </div>
  <div class="projects-wrapper">
    {#if error}
      <div class="error">{error}</div>
    {:else}
      <div class="projects-wrapper">
        {#each projects as project (project.id)}
          <Project title={project.title} description={project.description} />
        {/each}
      </div>
    {/if}
  </div>
</div>

<style>
  #projects {
    display: flex;
    flex-direction: column;
    padding-top: 5vh; /* Add some top padding */
  }

  .header {
    outline: 1px blue solid;
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 3vh; /* Space between header and projects */
  }

  .projects-wrapper {
    display: flex;

  }

  h1,
  h4 {
    font-family: "Martian Mono", monospace;
    margin: 0;
  }

  h1 {
    margin-bottom: 2vh;
  }

  h4 {
    font-weight: 400;
  }
</style>
