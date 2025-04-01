<script lang="ts">
  import ExperienceCard from "./components/ExperienceCard.svelte";
import NavBar from "./components/NavBar.svelte";
  import Home from "./pages/Home.svelte";
  import Projects from "./pages/Projects.svelte";

  let pageLinks: {label: string; href:string}[] = [
    { label: "Home", href: "#home" },
    { label: "Projects", href: "#projects" },
    { label: "Experience", href: "#experience" },
    { label: "Contact", href: "#contact" }
  ]

  let socials: {github: string; linkedin: string; resume: string} = {
    github: "https://github.com/gabrielg2020",
    linkedin: "https://linkedin.com/in/gabrielg4",
    resume: "/resume.pdf"
  }
</script>

<div class="min-h-screen">
  <NavBar name="Gabriel Guimaraes" links={pageLinks} socials={socials}/>
  <Home />
  <Projects />
  <ExperienceCard />
</div>

