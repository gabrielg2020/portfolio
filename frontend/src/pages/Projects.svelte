<script lang="ts">
  import ProjectCard from "../components/ProjectCard.svelte";
</script>

<div id="projects">
  <h1 class="h2">Projects</h1>
  <div class="projects">
    <ProjectCard />
  </div>
</div>

<style>
  #projects {
    display:flex;
    justify-items: center;
    align-items: center;
    flex-direction: column;
  }

  .projects {
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 2rem;
    align-items: center;
    width: 100%;
  }
</style>
