<script lang="ts">
  import Experiences from "./pages/Experiences.svelte";
  import Home from "./pages/Home.svelte";
  import Projects from "./pages/Projects.svelte";
</script>

<Home />
<Projects />
<Experiences />

<style>
  :global(html, body) {
    height: 100%;
    margin: 0;
    padding: 0;
    scroll-behavior: smooth; /* For smooth scrolling */
  }

  :global(body) {
    overflow-y: scroll;
    scroll-snap-type: y mandatory; /* Makes sections snap when scrolling */
  }
</style>
