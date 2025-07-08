// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt',
    '@vueuse/nuxt'
  ],
  typescript: {
    strict: false
  },
  css: [
    '~/assets/css/main.css'
  ],
  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE || 'http://localhost:8080'
    }
  },
  app: {
    head: {
      title: 'EROS - Dating App',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Modern dating app with AI-powered matching' }
      ],
      link: [
        { rel: 'icon', type: 'image/png', href: '/eroslogo.png' }
      ]
    }
  }
})
