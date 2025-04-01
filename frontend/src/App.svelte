<script lang="ts">
  import NavBar from "./components/NavBar.svelte";
  import Experience from "./pages/Experience.svelte";
  import Home from "./pages/Home.svelte";
  import Projects from "./pages/Projects.svelte";
  import Contact from "./pages/Contact.svelte";
  import Footer from "./components/Footer.svelte";
  import { darkMode } from "./store/darkMode";
  import { onMount } from "svelte";

  let pageLinks: { label: string; href: string }[] = [
    { label: "Home", href: "#home" },
    { label: "Projects", href: "#projects" },
    { label: "Experience", href: "#experience" },
    { label: "Contact", href: "#contact" },
  ];

  let socials: { github: string; linkedin: string; resume: string } = {
    github: "https://github.com/gabrielg2020",
    linkedin: "https://linkedin.com/in/gabrielg4",
    resume: "/resume.pdf",
  };

  // Set up dark mode on initial page load
  onMount(() => {
    // Apply the dark mode class to the document if needed
    if ($darkMode) {
      document.documentElement.classList.add('dark');
    }
  });
</script>

<div class="min-h-screen flex flex-col dark:bg-gray-900 dark:text-gray-100 transition-colors duration-200">
  <NavBar name="Gabriel Guimaraes" links={pageLinks} {socials} />
  <main class="flex-grow">
    <Home />
    <Projects />
    <Experience />
    <Contact
      email="gabriel.mg04@outlook.com"
      location="East Midlands, England"
      {socials}
    />
  </main>
  <Footer
    docsUrl="#"
    repoUrl="https://github.com/gabrielg2020/portfolio"
    copyrightName="Gabriel Guimaraes"
    startYear={2025}
  />
</div>