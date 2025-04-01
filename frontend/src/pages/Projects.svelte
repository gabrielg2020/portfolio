<script lang="ts">
  import Carousel from "../components/Carousel.svelte";
  import ProjectCard from "../components/ProjectCard.svelte";

  // AI Generated Projects - API Call in next PR
  const projects = [
    {
      title: "E-Commerce Platform",
      description:
        "A full-stack e-commerce application with cart functionality, user authentication, and payment processing.",
      languages: ["TypeScript", "CSS", "HTML"],
      technologies: ["React", "Node.js", "Express", "MongoDB"],
      githubUrl: "#",
      liveUrl: "#",
    },
    {
      title: "Task Management Dashboard",
      description:
        "A Kanban-style project management tool with drag and drop functionality and team collaboration features.",
      languages: ["JavaScript", "CSS", "HTML"],
      technologies: ["Vue.js", "Firebase", "Vuex"],
      githubUrl: "#",
      liveUrl: "#",
    },
    {
      title: "Data Visualization Tool",
      description:
        "An interactive dashboard showcasing various visualization techniques for complex datasets.",
      languages: ["TypeScript", "SCSS"],
      technologies: ["Angular", "D3.js", "PostgreSQL"],
      githubUrl: "#",
      liveUrl: null,
    },
    {
      title: "Content Management System",
      description:
        "A custom CMS designed for bloggers and content creators with markdown support and SEO optimization.",
      languages: ["JavaScript", "CSS", "HTML"],
      technologies: ["Next.js", "Supabase", "Tailwind CSS"],
      githubUrl: "#",
      liveUrl: "#",
    },
    {
      title: "Mobile Fitness App",
      description:
        "A cross-platform mobile application for tracking workouts, nutrition, and progress towards fitness goals.",
      languages: ["TypeScript", "CSS"],
      technologies: ["React Native", "Redux", "Firebase"],
      githubUrl: "#",
      liveUrl: null,
    },
    {
      title: "Real-time Chat Application",
      description:
        "A real-time messaging platform with private conversations, group chats, and media sharing capabilities.",
      languages: ["JavaScript", "CSS", "HTML"],
      technologies: ["Socket.io", "Express", "MongoDB"],
      githubUrl: "#",
      liveUrl: "#",
    },
  ];

  // Helper function to get projects for a specific slide
  function getProjectsForSlide(slideIndex: number, itemsPerSlide: number) {
    const startIdx = slideIndex * itemsPerSlide;
    return projects.slice(startIdx, startIdx + itemsPerSlide);
  }
</script>

<section id="projects">
  <Carousel
    title="Projects"
    totalItems={projects.length}
    itemsPerSlideDesktop={3}
    itemsPerSlideTablet={2}
    itemsPerSlideMobile={1}
  >
    <svelte:fragment slot="slide" let:slideIndex let:itemsPerSlide>
      {#each getProjectsForSlide(slideIndex, itemsPerSlide) as project (project.title)}
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
      {#each Array(itemsPerSlide - getProjectsForSlide(slideIndex, itemsPerSlide).length) as _, index (index)}
        <div class="flex-1 min-w-0"></div>
      {/each}
    </svelte:fragment>
  </Carousel>
</section>
