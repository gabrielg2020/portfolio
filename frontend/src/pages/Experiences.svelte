<script lang="ts">
  import ExperienceCard from "../components/ExperienceCard.svelte";
  import { onMount } from "svelte";
  import { GetExperiences } from "../services/apiService";
  import type { Experience } from "../services/apiService";

  let experiences: Promise<Experience[]>;

  onMount(() => {
    experiences = GetExperiences();
  });
</script>

<div id="experiences">
  <h1 class="h2">Experiences</h1>
  <div class="experiences">
    {#await experiences}
      <p>Loading...</p>
    {:then loadedExperiences}
      {#each loadedExperiences as experience}
        <ExperienceCard
          title={experience.title}
          smallDescription={experience.smallDescription}
          githubLink={experience.githubLink}
          languages={experience.languages}
          technologies={experience.technologies}
          yearStart={experience.yearStart}
          yearEnd={experience.yearEnd}
        />
      {/each}
    {:catch error}
      <p>Error loading experiences: {error}</p>
    {/await}
  </div>
</div>

<style>
  .experiences {
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 2rem;
    align-items: center;
    width: 100%;
  }
</style>
