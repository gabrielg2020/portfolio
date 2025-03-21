<script lang="ts">
  import InteractiveBackground from "./Background.svelte";
  import AnimatedButton from "./AnimatedButton.svelte";
</script>

<InteractiveBackground circleCount={500} minRadius={3} maxRadius={10}>
  <div class="home-container">
    <h1>Gabriel Guimaraes</h1>
    <h2>Software Engineer</h2>
    <div class="button-container">
      <AnimatedButton text="Projects" />
      <AnimatedButton text="Experience" />
      <AnimatedButton text="Contact" />
    </div>
  </div>
</InteractiveBackground>

<style>
  .home-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    position: absolute;
    left: 50%;
    top: 50%;
    -webkit-transform: translate(-50%, -50%);
    transform: translate(-50%, -50%);

    min-height: 100vh;
    color: white;
  }

  h1 {
    font-size: 4rem;
    margin-bottom: 0.5rem;
  }

  h2 {
    font-size: 2.5rem;
    margin-bottom: 2rem;
    opacity: 0.8;
  }

  .button-container {
    width: 100%;
    display: flex;
    gap: 1rem;
  }
</style>
