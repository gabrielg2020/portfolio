<script lang="ts">
  import Carousel from "../components/Carousel.svelte";
  import ProjectCard from "../components/ProjectCard.svelte";
  import Loading from "../components/Loading.svelte";
  import Error from "../components/Error.svelte";
  import { GetProjects } from "../scripts/callers/getProjects";
  import type { Project } from "../scripts/callers/getProjects";

  // State management
  let projectsPromise: Promise<Project[]>;

  // Load projects function (can be called for retries)
  function loadProjects() {
    projectsPromise = GetProjects();
  }

  // Initial load
  loadProjects();

  // Helper function to get projects for a specific slide
  function getProjectsForSlide(
    projects: Project[],
    slideIndex: number,
    itemsPerSlide: number,
  ) {
    const startIdx = slideIndex * itemsPerSlide;
    return projects.slice(startIdx, startIdx + itemsPerSlide);
  }
</script>

<section id="projects">
  {#await projectsPromise}
    <Loading message="Loading projects..." />
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
              githubUrl={project.githubURL}
              liveUrl={project.liveURL}
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
    <Error
      message={`Failed to load projects: ${error.message}`}
      retryFn={loadProjects}
    />
  {/await}
</section>

