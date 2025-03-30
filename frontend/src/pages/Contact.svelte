<script lang="ts">
 import Button from "../components/Button.svelte";
</script>

<div id="contact">
  <h1 class="h2">Contact</h1>
  <div class="content">
    <Button text="GitHub"/>
    <Button text="Linkedin"/>
    <Button text="Resume" iconStr="download"/>
  </div>
</div>

<style>
  .content {
    display: flex;
    gap: 2rem;
  }
</style>

