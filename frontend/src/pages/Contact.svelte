<script lang="ts">
  import JumpButton from "../components/JumpButton.svelte";
</script>

<div id="contact">
  <div class="header">
    <h1>Contact Me</h1>
  </div>
  <div class="content">
    <div class="content-wrapper">
      <a href="http://github.com/gabrielg2020">
        <img alt="github logo" src="/images/github.svg" />
      </a>
      <a href="http://github.com/gabrielg2020">
        <h4>gabrielg2020</h4>
      </a>
    </div>
    <div class="content-wrapper">
      <a href="http://linkedin.com/in/gabrielg4">
        <img alt="linkedin logo" src="/images/linkedin.svg" />
      </a>
      <a href="http://github.com/gabrielg2020">
        <h4>gabrielg4</h4>
      </a>
    </div>
  </div>
  <div class="footer">
    <JumpButton text="Go to top!" location="home" arrowDirection="up" />
  </div>
</div>

<style>
  #contact {
    display: flex;
    flex-direction: column;
    padding-top: 5vh; /* Add some top padding */
  }

  .header {
    outline: 1px blue solid;
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 10vh; /* Space between header and projects */
  }

  .content {
    display: flex;
    justify-content: center;
    gap: 4rem;
  }

  .content-wrapper > a {
    display: flex;
    justify-content: center;
  }
  .content-wrapper > a > img {
    min-width: 7rem;
  }

  .footer {
    display: flex;
    justify-content: center;
    padding-top: 10vh;
  }

  h1,
  h4 {
    font-family: "Martian Mono", monospace;
    margin: 0;
  }

  h4 {
    margin-top: 1rem;
  }
</style>
