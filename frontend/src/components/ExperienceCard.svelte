<script lang="ts">
  import InfoTab from "./InfoTab.svelte";
  import { onMount } from "svelte";

  export let title: string = "Experience Title";
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
  export let company: string = "Company";
  export let yearStart: string = "2000";
  export let yearEnd: string = "2001";

  let isDesktop = false;
  let hasEnoughSpace = false;
  let subheaderRef: HTMLDivElement;

  function updateLayout() {
    isDesktop = window.innerWidth >= 1024;
    
    // Check if there's enough space in the subheader for horizontal layout
    if (subheaderRef) {
      const companyElement = subheaderRef.querySelector('.company');
      const tagsElement = subheaderRef.querySelector('.languages-technologies');
      
      if (companyElement && tagsElement) {
        const companyWidth = companyElement.getBoundingClientRect().width;
        const tagsWidth = tagsElement.getBoundingClientRect().width;
        const subheaderWidth = subheaderRef.getBoundingClientRect().width;
        
        // Add some buffer space (20px) for gap
        hasEnoughSpace = (companyWidth + tagsWidth + 20) <= subheaderWidth;
      }
    }
  }

  onMount(() => {
    updateLayout();
    window.addEventListener("resize", updateLayout);

    return () => {
      window.removeEventListener("resize", updateLayout);
    };
  });
</script>

<div class="experience-card">
  <div class="header">
    <h2 class={isDesktop ? "h5" : "h6"}>{title}</h2>
    <h6 class={isDesktop ? "sub-h1" : "sub-h2"}>{yearStart} - {yearEnd}</h6>
  </div>
  <div class="subheader" class:stacked={!hasEnoughSpace} bind:this={subheaderRef}>
    <h6 class={`company ${isDesktop ? "sub-h1" : "sub-h2"}`}>{company}</h6>
    <div class="languages-technologies">
      {#each languages as language, index (index)}
        <InfoTab text={language} />
      {/each}
      {#each technologies as technology, index (index)}
        <InfoTab text={technology} iconStr="box" />
      {/each}
    </div>
  </div>
  <div class="content">
    {#if !isDesktop}
      <p class="body-1">{smallDescription}</p>
    {:else}
      <ul class="description-list">
        {#each largeDescription as point, index (index)}
          <li class="body-1">{point}</li>
        {/each}
      </ul>
    {/if}
  </div>
</div>

<style>
  .experience-card {
    display: flex;
    flex-direction: column;
  }

  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 1rem;
    margin-bottom: 0.5rem;
  }

  .subheader {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 1rem;
  }
  
  .stacked {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }
  
  .company {
    margin: 0;
  }

  .languages-technologies {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    flex-wrap: wrap;
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