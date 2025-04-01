<script lang="ts">
  import Timeline from "../components/Timeline.svelte";
  import ExperienceCard from "../components/ExperienceCard.svelte";
  import Loading from "../components/Loading.svelte";
  import Error from "../components/Error.svelte";
  import { GetExperiece } from "../scripts/callers/getExperience";
  import type { Experience } from "../scripts/callers/getExperience";
  
  // State management
  let experiencesPromise: Promise<Experience[]>;
  
  // Load experiences function (can be called for retries)
  function loadExperiences() {
    experiencesPromise = GetExperiece();
  }
  
  // Initial load
  loadExperiences();
</script>

<section id="experience" class="py-16">
  <div class="container mx-auto px-6">
    {#await experiencesPromise}
      <Loading message="Loading experience..." />
    {:then experiences}
      <Timeline title="Experience">
        {#each experiences as exp, index (index)}
          <ExperienceCard
            type={exp.type as "work" | "education"}
            organisation={exp.organisation}
            role={exp.role}
            startYear={exp.startYear}
            endYear={exp.endYear}
            description={exp.description}
            languages={exp.languages}
            technologies={exp.technologies}
            isLast={index === experiences.length - 1}
          />
        {/each}
      </Timeline>
    {:catch error}
      <Error 
        message={`Failed to load experience data: ${error.message}`} 
        retryFn={loadExperiences} 
      />
    {/await}
  </div>
</section>