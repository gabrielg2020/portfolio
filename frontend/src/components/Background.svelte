<script lang="ts">
  import { onMount, onDestroy } from "svelte";

  // Props
  export let backgroundColor = "#121212";
  export let circleCount = 50;
  export let minRadius = 5;
  export let maxRadius = 20;

  // DOM Variables
  let canvas: HTMLCanvasElement;
  let ctx: CanvasRenderingContext2D;
  let width: number;
  let height: number;
  let animationId: number;

  // Animation Variables
  let circles = [];
  let mouse = { x: 0, y: 0, isActive: false };

  onMount(() => {
    // Initialise Canvas
    ctx = canvas.getContext("2d");
    resizeCanvas();

    // Create initial circles
    initCircles();

    // Start animation loop
    animationId = requestAnimationFrame(animate);

    // Add event listeners
    window.addEventListener("resize", resizeCanvas);
    window.addEventListener('mousemove', handleMouseMove);
    window.addEventListener('mousedown', () => { mouse.isActive = true; });
    window.addEventListener('mouseup', () => { mouse.isActive = false; });
    window.addEventListener('mouseleave', () => { mouse.isActive = false; });

    return () => {
      window.removeEventListener("resize", resizeCanvas);
      window.removeEventListener('mousemove', handleMouseMove);
      window.removeEventListener('mousedown', () => { mouse.isActive = true; });
      window.removeEventListener('mouseup', () => { mouse.isActive = false; });
      window.removeEventListener('mouseleave', () => { mouse.isActive = false; });
    };
  });

  onDestroy(() => {
    if (animationId) {
      cancelAnimationFrame(animationId);
    }
  });

  function resizeCanvas() {
    const container = canvas.parentElement;
    width = container.clientWidth;
    height = container.clientHeight;
    canvas.width = width;
    canvas.height = height;

    // Reinitialize circles when canvas is resized
    initCircles();
  }

  function initCircles() {
    circles = [];
    for (let i = 0; i < circleCount; i++) {
      circles.push({
        x: Math.random() * width,
        y: Math.random() * height,
        radius: minRadius + Math.random() * (maxRadius - minRadius),
        dx: (Math.random() - 0.5) * 2,
        dy: Math.random(),
        color: getRandomColor(0.5),
      originalColor: getRandomColor(0.5),
      });
    }
  }

  function getRandomColor(alpha = 1) {
    const lightness = Math.floor(Math.random() * 100);
    return `hsla(0, 0%, ${lightness}%, ${alpha})`;
  }

  function handleMouseMove(event) {
    const rect = canvas.getBoundingClientRect();
    mouse.x = event.clientX - rect.left;
    mouse.y = event.clientY - rect.top;
  }

  function animate() {
    ctx.clearRect(0, 0, width, height);
    ctx.fillStyle = backgroundColor;
    ctx.fillRect(0, 0, width, height);
    
    for (let i = 0; i < circles.length; i++) {
      const circle = circles[i];
      
      // Update position
      circle.x += circle.dx;
      circle.y += circle.dy;
      
      // Boundary checking (bounce off walls)
      if (circle.x - circle.radius < 0 || circle.x + circle.radius > width) {
        circle.dx = -circle.dx;
      }
      
      // Reset position if circle goes off bottom
      if (circle.y - circle.radius > height) {
        circle.y = -circle.radius;
        circle.x = Math.random() * width;
      }
      
      // Apply some gravity
      circle.dy += 0.01;
      
      // Add some friction
      circle.dx *= 0.99;
      circle.dy *= 0.99;
      
      // Draw the circle
      ctx.beginPath();
      ctx.arc(circle.x, circle.y, circle.radius, 0, Math.PI * 2);
      ctx.fillStyle = circle.color;
      ctx.fill();
    }
    
    animationId = requestAnimationFrame(animate);
  }
</script>

<div class="container">
  <canvas bind:this={canvas}></canvas>
  <slot />
</div>

<style>
    .container {
    position: relative;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }
  
  canvas {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 0;
  }
  
  :global(.container > *:not(canvas)) {
    position: relative;
    z-index: 1;
  }
</style>
