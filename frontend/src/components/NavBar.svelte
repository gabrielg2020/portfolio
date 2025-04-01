<script lang="ts">
  import { onMount } from 'svelte';
  import { Github, Linkedin, FileText, Menu, X } from '@lucide/svelte';
  import { darkMode } from '../store/darkMode';
  import ThemeToggle from './ThemeToggle.svelte';

  // Props
  export let name: string = "Your Name";
  export let links: { label: string; href: string }[] = [
    { label: "Home", href: "#home" },
    { label: "Projects", href: "#projects" },
    { label: "Experience", href: "#experience" },
    { label: "Contact", href: "#contact" }
  ];
  export let socials: {
    github?: string;
    linkedin?: string;
    resume?: string;
  } = {
    github: "https://github.com/yourusername",
    linkedin: "https://linkedin.com/in/yourusername",
    resume: "/resume.pdf"
  };

  // Internal state
  let isMenuOpen: boolean = false;
  let isScrolled: boolean = false;
  let isVisible: boolean = true;
  let lastScrollY: number = 0;

  onMount(() => {
    const handleScroll = () => {
      const currentScrollY = window.scrollY;
      
      // Check if scrolled past threshold
      isScrolled = currentScrollY > 20;
      
      // Hide when scrolling down, show when scrolling up
      if (currentScrollY > lastScrollY) {
        isVisible = false; // Scrolling down
      } else {
        isVisible = true; // Scrolling up
      }
      
      // Update last scroll position
      lastScrollY = currentScrollY;
    };
    
    window.addEventListener('scroll', handleScroll);
    
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  });

  function toggleMenu(): void {
    isMenuOpen = !isMenuOpen;
  }

  // Close menu when clicking a link
  function closeMenu(): void {
    isMenuOpen = false;
  }
</script>

<header class="{isScrolled ? 'py-3 bg-white dark:bg-gray-900 shadow-sm' : 'py-5 bg-white dark:bg-gray-900'} {isVisible ? 'translate-y-0' : '-translate-y-full'} fixed top-0 left-0 right-0 z-50 transition-all duration-300">
  <nav class="container mx-auto px-6 flex justify-between items-center">
    <!-- Logo/Name -->
    <a href="#home" class="text-lg font-medium text-gray-800 dark:text-gray-200 hover:text-blue-600 dark:hover:text-blue-400 transition-colors">
      {name}
    </a>

    <!-- Desktop Navigation -->
    <div class="hidden md:flex items-center space-x-8">
      <!-- Main Links -->
      <ul class="flex space-x-6">
        {#each links as link, index (index)}
          <li>
            <a 
              href={link.href} 
              class="text-sm text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
            >
              {link.label}
            </a>
          </li>
        {/each}
      </ul>

      <!-- Social Links -->
      <div class="flex items-center space-x-4">
        {#if socials.github}
          <a 
            href={socials.github}
            target="_blank"
            rel="noopener noreferrer"
            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
            aria-label="GitHub"
          >
            <Github size={18} />
          </a>
        {/if}
        
        {#if socials.linkedin}
          <a 
            href={socials.linkedin}
            target="_blank"
            rel="noopener noreferrer"
            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
            aria-label="LinkedIn"
          >
            <Linkedin size={18} />
          </a>
        {/if}
        
        {#if socials.resume}
          <a 
            href={socials.resume}
            target="_blank"
            rel="noopener noreferrer"
            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
            aria-label="Resume"
          >
            <FileText size={18} />
          </a>
        {/if}
      </div>

      <!-- Dark Mode Toggle -->
      <ThemeToggle size={18} />
    </div>

    <!-- Mobile Menu Button -->
    <button 
      class="md:hidden text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
      on:click={toggleMenu}
      aria-label="Toggle Menu"
    >
      {#if isMenuOpen}
        <X size={24} />
      {:else}
        <Menu size={24} />
      {/if}
    </button>
  </nav>

  <!-- Mobile Menu -->
  <div class="md:hidden bg-white dark:bg-gray-900 absolute left-0 right-0 shadow-md transition-all duration-300 overflow-hidden {isMenuOpen ? 'max-h-96 py-4' : 'max-h-0'}">
    <div class="container mx-auto px-6">
      <ul class="space-y-4 mb-6">
        {#each links as link, index (index)}
          <li>
            <a 
              href={link.href} 
              class="block text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
              on:click={closeMenu}
            >
              {link.label}
            </a>
          </li>
        {/each}
      </ul>

      <div class="flex items-center space-x-4 pb-4">
        {#if socials.github}
          <a 
            href={socials.github}
            target="_blank"
            rel="noopener noreferrer"
            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
            aria-label="GitHub"
          >
            <Github size={20} />
          </a>
        {/if}
        
        {#if socials.linkedin}
          <a 
            href={socials.linkedin}
            target="_blank"
            rel="noopener noreferrer"
            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
            aria-label="LinkedIn"
          >
            <Linkedin size={20} />
          </a>
        {/if}
        
        {#if socials.resume}
          <a 
            href={socials.resume}
            target="_blank"
            rel="noopener noreferrer"
            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
            aria-label="Resume"
          >
            <FileText size={20} />
          </a>
        {/if}

        <div class="ml-auto">
          <ThemeToggle size={20} />
        </div>
      </div>
    </div>
  </div>
</header>