import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server:
  {
    hmr:true,
    port:8080,
    proxy: 
    {
      "/auth":
      {
        target:"http://8.137.100.0:8080",
        changeOrigin:true,
        cookiePathRewrite:
        {
          "^auth":"auth"
        }
      }
    }
  }
})
