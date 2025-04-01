<script lang="ts">
  export let title: string = "Timeline";
</script>

<div class="bg-white dark:bg-gray-900 p-6 transition-colors duration-200">
  <h2 class="text-2xl font-bold text-gray-800 dark:text-gray-100 mb-8">{title}</h2>
  <div class="relative">
    <slot></slot>
  </div>
</div>