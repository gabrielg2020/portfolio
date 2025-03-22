<script lang="ts">
  import Project from "../components/Project.svelte";
</script>

<div id="projects">
  <div class="header">
    <h1>Project Showcase</h1>
    <h4>Open them up to learn more!</h4>
  </div>
  <div class="projects-wrapper">
    <Project />
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
    max-width: 16rem;
  }

  h1, h4 {
    font-family: 'Martian Mono', monospace;
    margin: 0;
  }

  h1 {
    margin-bottom: 2vh;
  }
</style>
