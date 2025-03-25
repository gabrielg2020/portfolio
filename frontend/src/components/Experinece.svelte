<script lang="ts">
  export let company: string = "Company Name";
  export let role: string = "Role Titile";
  export let yearStart: string = "2000";
  export let yearEnd: string = "2004";
  export let points: string[] = Array(3).fill(
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, mattis lectus sed, maximus dolor.",
  );
  export let langauges: string[] = ["go", "typescript"];
  export let technologies: string[] = ["docker", "svelte"];
</script>

<div class="experience">
  <div class="header">
    <div class="header-wrapper">
      <h3>{company}</h3>
      <h4>({yearStart} - {yearEnd})</h4>
    </div>
    <h4>{role}</h4>
  </div>
  <div class="content">
    <ul>
      {#each points as point, index (index)}
        <li>{point}</li>
      {/each}
    </ul>
    <div class="content-wrapper">
      <div class="languages">
        <h4>Languages</h4>
        <div class="image-wrapper">
          {#each langauges as langauge (langauge)}
            <img alt="{langauge} logo" src="/images/{langauge}.svg" />
          {/each}
        </div>
      </div>
      <div class="technologies">
        <h4>Technologies</h4>
        <div class="image-wrapper">
          {#each technologies as technology (technology)}
            <img alt="{technology} logo" src="/images/{technology}.svg" />
          {/each}
        </div>
      </div>
    </div>
  </div>
  <div class="footer"></div>
</div>

<style>
  .experience {
    font-family: "Martian Mono", monospace;
    border: 1px red solid;
    padding-right: 10px;
    padding-left: 10px;
  }

  .header-wrapper {
    display: flex;
    gap: 1rem;
    align-items: center;
  }

  p {
    font-weight: 200;
  }

  h4 {
    font-weight: 400;
  }

  .content-wrapper {
    display: flex;
    gap: 3rem;
  }

  .languages,
  .technologies {
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
