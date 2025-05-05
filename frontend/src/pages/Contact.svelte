<script lang="ts">
  import { Mail, MapPin, Github, Linkedin, FileText } from "@lucide/svelte";

  // Props
  export let email: string = "your.email@example.com";
  export let location: string = "City, Country";
  export let socials: {
    github?: string;
    linkedin?: string;
    resume?: string;
  } = {
    github: "https://github.com/gabrielg2020",
    linkedin: "https://linkedin.com/in/gabrielg4",
    resume: "/resume.pdf",
  };
</script>

<section
  id="contact"
  class="py-8 bg-gray-50 dark:bg-gray-800 transition-colors duration-200"
>
  <div class="container mx-auto px-6">
    <h2 class="text-2xl font-bold text-gray-800 dark:text-gray-100 mb-8">
      Get In Touch
    </h2>

    <div class="max-w-2xl mx-auto">
      <div
        class="bg-white dark:bg-gray-900 p-8 rounded-lg border border-gray-100 dark:border-gray-700 shadow-sm"
      >
        <div class="space-y-8">
          <!-- Contact Information -->
          <div class="space-y-6">
            <div class="flex items-start">
              <Mail
                class="text-blue-600 dark:text-blue-400 mr-4 h-6 w-6 flex-shrink-0 mt-1"
              />
              <div>
                <p
                  class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
                >
                  Email
                </p>
                <a
                  href={`mailto:${email}`}
                  class="text-blue-600 dark:text-blue-400 hover:underline text-lg"
                >
                  {email}
                </a>
              </div>
            </div>

            <div class="flex items-start">
              <MapPin
                class="text-blue-600 dark:text-blue-400 mr-4 h-6 w-6 flex-shrink-0 mt-1"
              />
              <div>
                <p
                  class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
                >
                  Location
                </p>
                <p class="text-gray-800 dark:text-gray-200 text-lg">
                  {location}
                </p>
              </div>
            </div>
          </div>

          <!-- Divider -->
          <div class="border-t border-gray-200 dark:border-gray-700"></div>

          <!-- Social Links -->
          <div>
            <h3
              class="text-lg font-medium text-gray-800 dark:text-gray-200 mb-4"
            >
              Connect With Me
            </h3>
            <div class="flex flex-wrap gap-6">
              {#if socials.github}
                <a
                  href={socials.github}
                  target="_blank"
                  rel="noopener noreferrer"
                  class="flex items-center text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
                >
                  <Github class="h-5 w-5 mr-2" />
                  <span>GitHub</span>
                </a>
              {/if}

              {#if socials.linkedin}
                <a
                  href={socials.linkedin}
                  target="_blank"
                  rel="noopener noreferrer"
                  class="flex items-center text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
                >
                  <Linkedin class="h-5 w-5 mr-2" />
                  <span>LinkedIn</span>
                </a>
              {/if}

              {#if socials.resume}
                <a
                  href={socials.resume}
                  target="_blank"
                  rel="noopener noreferrer"
                  class="flex items-center text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 transition-colors"
                >
                  <FileText class="h-5 w-5 mr-2" />
                  <span>Resume</span>
                </a>
              {/if}
            </div>
          </div>

          <!-- Message -->
          <div
            class="bg-blue-50 dark:bg-blue-900/20 p-4 rounded-lg border-l-4 border-blue-600 dark:border-blue-500"
          >
            <p class="text-gray-700 dark:text-gray-300 italic">
              "I'm always open to discussing new projects, opportunities in
              software development, or tech topics. Feel free to reach out!"
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>

