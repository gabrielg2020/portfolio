<script lang="ts">
  import NavBar from "./components/NavBar.svelte";
  import Experience from "./pages/Experience.svelte";
  import Home from "./pages/Home.svelte";
  import Projects from "./pages/Projects.svelte";
  import Contact from "./pages/Contact.svelte";
  import Footer from "./components/Footer.svelte";

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
</script>

<div class="min-h-screen">
  <NavBar name="Gabriel Guimaraes" links={pageLinks} {socials} />
  <Home />
  <Projects />
  <Experience />
  <Contact
    email="gabriel.mg04@outlook.com"
    location="East Midlands, England"
    {socials}
  />
  <Footer
    docsUrl="#"
    repoUrl="https://github.com/gabrielg2020/portfolio"
    copyrightName="Gabriel Guimaraes"
    startYear={2025}
  />
</div>
