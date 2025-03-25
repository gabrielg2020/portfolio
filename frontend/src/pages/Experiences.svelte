<script lang="ts">
  import Experinece from "../components/Experinece.svelte";
  import JumpButton from "../components/JumpButton.svelte";
</script>

<div id="experiences">
  <div class="header">
    <h1>Experiences</h1>
  </div>
  <div class="experiences-wrapper">
    <Experinece />
    <Experinece />
    <Experinece />
  </div>
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
    margin-bottom: 3vh; /* Space between header and projects */
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
