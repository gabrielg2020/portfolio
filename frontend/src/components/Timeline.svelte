<script lang="ts">
  export let title: string = "Timeline";
</script>

<div class="bg-white p-6">
  <h2 class="text-2xl font-bold text-gray-800 mb-8">{title}</h2>
  <div class="relative">
    <slot></slot>
  </div>
</div>