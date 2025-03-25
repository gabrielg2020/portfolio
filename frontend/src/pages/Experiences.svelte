<script lang="ts">
  import JumpButton from "../components/JumpButton.svelte";
</script>

<div id="experiences">
  <div class="header">
    <h1>Work Experience</h1>
    <h4>Open them up to learn more!</h4>
  </div>
  <div class="experiences-wrapper"></div>
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

  .footer {
    display: flex;
    justify-content: center;
    padding-top: 5vh;
  }

  h1,
  h4 {
    font-family: "Martian Mono", monospace;
    margin: 0;
  }

  h1 {
    margin-bottom: 2vh;
  }

  h4 {
    font-weight: 400;
  }
</style>
