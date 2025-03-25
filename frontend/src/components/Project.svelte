<script lang="ts">
  export let title: string = "Project Name";
  export let description: string =
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor. Quisque pulvinar tincidunt lacus, quis cursus nibh pretium eget. Proin porta mi at pulvinar maximus.";
  export let langauges: string[] = ["go", "typescript"];
  export let technologies: string[] = ["docker", "svelte"];
  export let githubLink: string;
</script>

<div class="project">
  <div class="header">
    <a href="/api/project/CHANGEME">
      <h3>{title}</h3>
    </a>
    <a href={githubLink}>
      <img alt="github logo" src="/images/github.svg" />
    </a>
  </div>
  <div class="content">
    <p>{description}</p>
  </div>
  <div class="footer">
    <div class="footer-wrapper">
      <div class="languages">
        <h4>Languages</h4>
        <div class="image-wrapper">
          {#each langauges as language (language)}
            <img alt="{language} language logo" src="images/{language}.svg" />
          {/each}
        </div>
      </div>
    </div>
    {#if technologies && technologies.length > 0}
      <div class="footer-wrapper">
        <div class="technologies">
          <h4>Technologies</h4>
          <div class="image-wrapper">
            {#each technologies as technology (technology)}
              <img alt="{technology} logo" src="/images/{technology}.svg" />
            {/each}
          </div>
        </div>
      </div>
    {/if}
  </div>
</div>

<style>
  .project {
    font-family: "Martian Mono", monospace;
    border: 1px red solid;
    padding-right: 10px;
    padding-left: 10px;
  }

  p {
    font-weight: 200;
  }

  h4 {
    font-weight: 400;
  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .header > a > img {
    width: 36px;
    height: 36px;
  }

  .footer {
    display: flex;
    gap: 15%;
    flex-wrap: wrap;
  }

  .languages, .technologies {
    display: flex;
    flex-direction: column;
  }

  .image-wrapper {
    display: flex;
    justify-content: space-around;
    gap: 1rem;
    max-width: 48px;
    max-height: 48px;
  }

  .image-wrapper > img {
    width: 100%;
  }
</style>
