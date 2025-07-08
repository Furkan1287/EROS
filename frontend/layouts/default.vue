<template>
  <div class="min-h-screen w-full bg-gradient-to-br from-pink-100 via-blue-100 to-yellow-100">
    <NuxtPage />
  </div>
</template>

<style>
html, body, #__nuxt {
  min-height: 100vh;
  width: 100vw;
  background: linear-gradient(135deg, #fde4ec 0%, #e0e7ff 50%, #fef9c3 100%) !important;
  box-sizing: border-box;
}
body {
  margin: 0;
  padding: 0;
}
</style>

<script setup>
const showLoginModal = ref(false)
const showRegisterModal = ref(false)
</script> 