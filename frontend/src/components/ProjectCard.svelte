<script lang="ts">
  import { ExternalLink } from "@lucide/svelte";
  import InfoTab from "./InfoTab.svelte";
  import { onMount } from "svelte";

  export let title: string = "Project Title";
  export let smallDescription: string =
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, Vivamus sed velit mattis, Vivamus sed velit.";
  export let largeDescription: string[] = [
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, Vivamus sed velit mattis, Vivamus sed velit.",
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, Vivamus sed velit mattis, Vivamus sed velit.",
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut ultricies placerat interdum. Vivamus sed velit mattis, Vivamus sed velit mattis, Vivamus sed velit.",
  ];
  export let githubLink: string = "www.github.com/gabrielg2020";
  export let languages: string[] = ["GoLang", "TypeScript"];
  export let technologies: string[] = ["Docker", "Svelte"];

  let isDesktop = false;

  function updateLayout() {
    isDesktop = window.innerWidth >= 1024;
  }

  onMount(() => {
    updateLayout();
    window.addEventListener("resize", updateLayout);

    return () => {
      window.removeEventListener("resize", updateLayout);
    };
  });
</script>

<div class="project-card">
  <div class="header">
    <h2 class={isDesktop ? "h5" : "h6"}>{title}</h2>
    <span class="icon action"><ExternalLink size=27/></span>
  </div>
  {#if isDesktop}
    <div class="languages-technologies">
      {#each languages as language}
        <InfoTab text={language} />
      {/each}
      {#each technologies as technology}
        <InfoTab text={technology} iconStr="box"/>
      {/each}
    </div>
  {/if}
  <div class="content">
    {#if !isDesktop}
      <p class="body-1">{smallDescription}</p>
    {:else}
      <ul class="description-list">
        {#each largeDescription as point}
          <li class="body-1">{point}</li>
        {/each}
      </ul>
    {/if}
  </div>
  {#if !isDesktop}
    <div class="languages-technologies">
      {#each languages as language}
        <InfoTab text={language} />
      {/each}
      {#each technologies as technology}
      <InfoTab text={technology} iconStr="box"/>
      {/each}
    </div>
  {/if}
</div>


<style>
  .project-card {
    display: flex;
    flex-direction: column;
  }

  .header {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 0.5rem;
  }

  .languages-technologies {
    display: flex;
    gap: 0.5rem;
    margin-bottom: 1rem;
  }
  
  .content {
    margin-bottom: 0.5rem;
  }

  p {
    max-width: 60ch;
    margin: 0;
  }
  
  .description-list {
    list-style-type: disc;
    padding-left: 1.5rem;
    margin: 0;
  }
  
  .description-list li {
    max-width: 60ch;
    margin-bottom: 1rem;
  }
  
  .description-list li:last-child {
    margin-bottom: 0;
  }
</style>