<script lang="ts">
  import Button from "../components/JumpButton.svelte";
</script>

<div id="home">
  <div class="wrapper">
    <div class="content">
      <h1>Gabriel Guimaraes</h1>
      <h2>Software Engineer</h2>
    </div>
    <Button text="Take a look at my projects!" location="projects"/>
  </div>
</div>

<style>
  /* div {
    outline: 1px red solid;
  } */

  .wrapper {
    outline: 1px blue solid;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    position: absolute;
    left: 50%;
    top: 40%;
    -webkit-transform: translate(-50%, -50%);
    transform: translate(-50%, -50%);
    width: 90vw;
  }

  h1, h2 {
    font-family: 'Martian Mono', monospace;
    margin: 0;
  }

  h1 {
    font-size: 4rem;
  }

  h2 {
    font-size: 2rem;
    font-weight: 400;
    margin-top: 0.5rem;
    margin-bottom: 2vh;
  }
</style>
