import { defineConfig } from 'vite'
    import vue from '@vitejs/plugin-vue'
    import mkcert from 'vite-plugin-mkcert'
    import fs from 'fs';
    
    // https://vitejs.dev/config/
    export default defineConfig({
      plugins: [vue(),  mkcert() ],
      server: {
        origin: 'https://ormand.localhost:5173',
   
        https: {
          key: fs.readFileSync('./internal/certs/local-key.pem'),
          cert: fs.readFileSync('./internal/certs/local-cert.pem'),
        }, 

        // // add the next lines if you're using windows and hot reload doesn't work
        watch: {
          usePolling: true
        },
      },
      
    })
    