<script>
  import Background from "./components/Home.svelte";
</script>

<main>
  <Background circleCount={100} minRadius={3} maxRadius={10} />
</main>

<style>
  main {
    width: 100%;
    height: 100vh;
  }
</style>
