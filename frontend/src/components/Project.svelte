<script lang="ts">
  export let title: string = "Project Name";
  export let description: string =
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus.";
  export let langauges: string[] = ["go", "typescript"];
  export let technologies: string[] = ["docker", "svelte"];
</script>

<div class="project">
  <div class="header">
    <h3>{title}</h3>
  </div>
  <div class="content">
    <p>{description}</p>
  </div>
  <div class="footer">
    <div class="footer-wrapper">
      <h4>Languages</h4>
      <div class="image-wrapper">
        {#each langauges as language (language)}
          <img alt="logo of {language}" src="images/{language}.svg" />
        {/each}
      </div>
    </div>
    <div class="footer-wrapper">
      <h4>Technologies</h4>
      <div class="image-wrapper">
        {#each technologies as technology (technology)}
          <img alt="logo of {technology}" src="/images/{technology}.svg" />
        {/each}
      </div>
    </div>
  </div>
</div>

<style>
  .project {
    font-family: "Martian Mono", monospace;
    border: 1px red solid;
  }

  p {
    font-weight: 200;
  }

  h4 {
    font-weight: 400;
  }

  .image-wrapper {
    display: flex;
    justify-content: space-around;
    gap: 1rem;
    max-width: 48px;
  }

  .image-wrapper > img {
    width: 100%;
  }
</style>
