<script lang="ts">
  import Timeline from "../components/Timeline.svelte";
  import ExperienceCard from "../components/ExperienceCard.svelte";
  
  // AI Generated Experiences - API Call in next PR
  const experiences = [
    {
      type: "work",
      organization: "Tech Innovations Inc.",
      role: "Senior Full Stack Developer",
      startYear: "2022",
      endYear: "Present",
      description: "Leading the development of the company's flagship product. Architected and implemented new features, mentored junior developers, and improved overall system performance by 40%.",
      languages: ["TypeScript", "HTML", "CSS", "SQL"],
      technologies: ["React", "Node.js", "Express", "PostgreSQL", "Docker", "AWS"]
    },
    {
      type: "work",
      organization: "Digital Solutions Ltd.",
      role: "Full Stack Developer",
      startYear: "2020",
      endYear: "2022",
      description: "Developed and maintained multiple client-facing web applications. Implemented responsive designs and optimized database queries resulting in a 25% improvement in page load times.",
      languages: ["JavaScript", "HTML", "CSS", "Python"],
      technologies: ["Vue.js", "Django", "MongoDB", "Webpack"]
    },
    {
      type: "education",
      organization: "University of Technology",
      role: "BSc Computer Science",
      startYear: "2016",
      endYear: "2020",
      description: "Graduated with honors. Focused on software engineering and data structures. Completed a capstone project building a real-time collaborative coding platform.",
      languages: ["Java", "C++", "JavaScript"],
      technologies: ["Spring", "ReactJS", "Socket.io", "MySQL"]
    }
  ];
</script>

<section id="experience" class="py-16">
  <div class="container mx-auto px-6">
    <Timeline title="Experience">
      {#each experiences as exp, index (index)}
        <ExperienceCard
          type={exp.type as "work" | "education"}
          organization={exp.organization}
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
  </div>
</section>