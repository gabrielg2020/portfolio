import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vite.dev/config/
export default defineConfig({
  plugins: [svelte()],
  build: {
    // Output to the backend's static directory when building
    outDir: '../backend/static',
    emptyOutDir: true,
    assetsDir: 'assets'
  },
  server: {
    host: '0.0.0.0',
    proxy: {
      // Proxy API requests to the backend
      '/api': {
        target: 'http://backend-dev:8080',
        changeOrigin: true
      }
    }
  }
})