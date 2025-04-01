<script lang="ts">
  import Timeline from "../components/Timeline.svelte";
  import ExperienceCard from "../components/ExperienceCard.svelte";
  import { GetExperiece } from "../scripts/callers/getExperience";
  import type { Experience } from "../scripts/callers/getExperience";
  const experiencesPromise: Promise<Experience[]> = GetExperiece();
</script>

<section id="experience" class="py-16">
  <div class="container mx-auto px-6">
    <Timeline title="Experience">
      {#await experiencesPromise}
        <p>Waiting</p>
      {:then experiences}
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
      {/await}
    </Timeline>
  </div>
</section>
