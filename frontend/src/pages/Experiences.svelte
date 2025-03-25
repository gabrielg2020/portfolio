<script lang="ts">
  import Experience from "../components/Experience.svelte";
  import JumpButton from "../components/JumpButton.svelte";

  import { onMount } from "svelte";

  let experiences: {
    id: number;
    company: string;
    role: string;
    yearEnd: string;
    yearStart: string;
    points: string[];
    languages: string[];
    technologies: string[];
  }[] = [];
  let error: string | null = null;

  onMount(async () => {
    try {      
      const response = await fetch("/api/experiences");
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      experiences = data.experiences;
      console.log(data);
    } catch (err) {
      console.error("Failed to fetch experiences:", err);
      if (err instanceof Error) {
        error = err.message;
      } else {
        error = String(err);
      }
    }
  });
</script>

<div id="experiences">
  <div class="header">
    <h1>Experiences</h1>
  </div>
  {#if error}
    <div class="error">{error}</div>
  {:else}
    <div class="experiences-wrapper">
      {#each experiences as experience (experience.id)}
        <Experience
          company={experience.company}
          role={experience.role}
          yearStart={experience.yearStart}
          yearEnd={experience.yearEnd}
          points={experience.points}
          langauges={experience.languages}
          technologies={experience.technologies}
        />
      {/each}
    </div>
  {/if}
  <div class="footer">
    <JumpButton text="Contact me!" location="contact" />
  </div>
</div>

<style>
  #experiences {
    display: flex;
    flex-direction: column;
    padding-top: 5vh; /* Add some top padding */
  }

  .header {
    outline: 1px blue solid;
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 3vh; /* Space between header and experiences */
  }

  .experiences-wrapper {
    display: flex;
    gap: 1rem;
  }

  .footer {
    display: flex;
    justify-content: center;
    padding-top: 5vh;
  }

  h1 {
    font-family: "Martian Mono", monospace;
    margin: 0;
  }
</style>
