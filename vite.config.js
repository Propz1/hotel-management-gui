import { defineConfig } from 'vite'
    import vue from '@vitejs/plugin-vue'
    import mkcert from 'vite-plugin-mkcert'
    import fs from 'fs';
    
    // https://vitejs.dev/config/
    export default defineConfig({
      plugins: [vue(),  mkcert() ],
      server: {
        origin: 'https://localhost:5173',
        https: {
          key: fs.readFileSync('./internal/certs/localhost+2-key.pem'),
          cert: fs.readFileSync('./internal/certs/localhost+2.pem'),
        } 
      }
      
    })
    