<script lang="ts">
  import Carousel from "../components/Carousel.svelte";
  import ProjectCard from "../components/ProjectCard.svelte";
  import { GetProjects } from "../scripts/callers/getProjects";
  import type { Project } from "../scripts/callers/getProjects";
  const projectsPromise: Promise<Project[]> = GetProjects();

  // Helper function to get projects for a specific slide
  function getProjectsForSlide(projects: Project[], slideIndex: number, itemsPerSlide: number) {
    const startIdx = slideIndex * itemsPerSlide;
    return projects.slice(startIdx, startIdx + itemsPerSlide);
  }
</script>

<section id="projects">
  {#await projectsPromise}
    <div class="loading">Loading projects...</div>
  {:then projects}
    <Carousel
      title="Projects"
      totalItems={projects.length}
      itemsPerSlideDesktop={3}
      itemsPerSlideTablet={2}
      itemsPerSlideMobile={1}
    >
      <svelte:fragment slot="slide" let:slideIndex let:itemsPerSlide>
        {#each getProjectsForSlide(projects, slideIndex, itemsPerSlide) as project (project.title)}
          <div class="flex-1 min-w-0">
            <ProjectCard
              title={project.title}
              description={project.description}
              languages={project.languages}
              technologies={project.technologies}
              githubUrl={project.githubUrl}
              liveUrl={project.liveUrl}
            />
          </div>
        {/each}

        <!-- Add empty placeholders if needed to fill the slide -->
        {#each Array(itemsPerSlide - getProjectsForSlide(projects, slideIndex, itemsPerSlide).length) as _, index (index)}
          <div class="flex-1 min-w-0"></div>
        {/each}
      </svelte:fragment>
    </Carousel>
  {:catch error}
    <div class="error">
      <p>Error loading projects: {error.message}</p>
    </div>
  {/await}
</section>