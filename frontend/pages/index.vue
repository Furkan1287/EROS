<template>
  <div class="relative min-h-screen overflow-x-hidden bg-gradient-to-br from-pink-100 via-blue-100 to-yellow-100">
    <!-- Animated amorphous background shapes -->
    <div class="pointer-events-none absolute inset-0 -z-10">
      <svg class="absolute left-[-10vw] top-[-10vw] w-[60vw] h-[60vw] opacity-60 animate-float-slow" viewBox="0 0 600 600">
        <defs>
          <radialGradient id="g1" cx="50%" cy="50%" r="50%">
            <stop offset="0%" stop-color="#a5b4fc"/>
            <stop offset="100%" stop-color="#fbc2eb"/>
          </radialGradient>
        </defs>
        <ellipse cx="300" cy="300" rx="300" ry="250" fill="url(#g1)"/>
      </svg>
      <svg class="absolute right-[-15vw] top-[20vh] w-[50vw] h-[50vw] opacity-50 animate-float" viewBox="0 0 600 600">
        <defs>
          <radialGradient id="g2" cx="50%" cy="50%" r="50%">
            <stop offset="0%" stop-color="#fbc2eb"/>
            <stop offset="100%" stop-color="#fcd34d"/>
          </radialGradient>
        </defs>
        <ellipse cx="300" cy="300" rx="300" ry="200" fill="url(#g2)"/>
      </svg>
      <svg class="absolute left-[10vw] bottom-[-10vw] w-[60vw] h-[60vw] opacity-40 animate-float-reverse" viewBox="0 0 600 600">
        <defs>
          <radialGradient id="g3" cx="50%" cy="50%" r="50%">
            <stop offset="0%" stop-color="#a7ffeb"/>
            <stop offset="100%" stop-color="#fbc2eb"/>
          </radialGradient>
        </defs>
        <ellipse cx="300" cy="300" rx="300" ry="250" fill="url(#g3)"/>
      </svg>
    </div>

    <!-- Navbar -->
    <nav class="w-full flex items-center justify-between px-10 py-6 bg-gradient-to-r from-pink-100/80 via-yellow-100/80 to-blue-100/80 backdrop-blur-md shadow-none z-20">
      <span class="text-4xl font-black bg-gradient-to-r from-pink-400 via-fuchsia-400 to-yellow-400 bg-clip-text text-transparent tracking-tight select-none">EROS</span>
      <div class="flex gap-4">
        <button class="px-7 py-3 rounded-full bg-gradient-to-r from-yellow-400 via-pink-400 to-fuchsia-400 text-white text-lg font-bold shadow hover:scale-105 transition-all duration-200" @click="showLogin = true">Giriş Yap</button>
        <button class="px-7 py-3 rounded-full border-2 border-fuchsia-300 text-fuchsia-700 text-lg font-bold shadow hover:bg-fuchsia-100/60 hover:text-fuchsia-900 transition-all duration-200" @click="showRegister = true">Kayıt Ol</button>
      </div>
    </nav>

    <!-- Hero Section -->
    <section class="pt-24 pb-8 flex flex-col items-center w-full relative z-10">
      <h1 class="text-[8vw] md:text-7xl font-black bg-gradient-to-r from-pink-400 via-fuchsia-400 to-yellow-400 bg-clip-text text-transparent drop-shadow-2xl tracking-tight text-center select-none w-full">
        EROS
      </h1>
      <p class="text-3xl md:text-4xl font-light text-fuchsia-700 mt-4 mb-2 text-center w-full select-none drop-shadow">
        Gerçek aşkı bul, sağa kaydır!
      </p>
      <p class="text-xl text-fuchsia-500 mb-8 text-center w-full select-none drop-shadow">
        Modern, güvenli ve eğlenceli eşleşme deneyimi.
      </p>
      <div class="flex gap-8 justify-center mt-6 flex-wrap w-full">
        <button class="px-12 py-5 rounded-full bg-gradient-to-r from-yellow-400 via-pink-400 to-fuchsia-400 text-white text-2xl font-bold shadow-xl hover:scale-110 hover:from-fuchsia-400 hover:to-yellow-400 transition-all duration-300" @click="showLogin = true">
          Giriş Yap
        </button>
        <button class="px-12 py-5 rounded-full border-2 border-fuchsia-300 text-fuchsia-700 text-2xl font-bold shadow-xl hover:bg-fuchsia-100/60 hover:text-fuchsia-900 transition-all duration-300" @click="showRegister = true">
          Kayıt Ol
        </button>
      </div>
    </section>

    <!-- Tinder Swipe Demo -->
    <section v-if="!isAuthenticated" class="w-full flex flex-col items-center justify-center py-16 relative z-10">
      <div class="flex flex-col items-center w-full">
        <div class="flex flex-row items-center justify-center gap-8 w-full">
          <!-- Sol (X) butonu -->
          <button class="rounded-full bg-white/80 hover:bg-pink-200 text-4xl shadow-lg w-20 h-20 flex items-center justify-center transition-all duration-200">
            <span class="text-pink-500">✖️</span>
          </button>
          <!-- SwipeDemo kartı -->
          <div class="w-[340px] h-[480px] md:w-[400px] md:h-[560px] bg-white/90 rounded-3xl shadow-2xl flex items-center justify-center relative overflow-hidden">
            <SwipeDemo />
          </div>
          <!-- Sağ (Kalp) butonu -->
          <button class="rounded-full bg-white/80 hover:bg-fuchsia-200 text-4xl shadow-lg w-20 h-20 flex items-center justify-center transition-all duration-200">
            <span class="text-fuchsia-500">❤️</span>
          </button>
        </div>
        <div class="mt-8 text-fuchsia-400 text-lg font-semibold tracking-wide">Sağa kaydır, eşleş!</div>
      </div>
    </section>

    <!-- Final CTA -->
    <footer class="w-full flex justify-center py-16 relative z-10 bg-gradient-to-r from-pink-100/80 via-yellow-100/80 to-blue-100/80 backdrop-blur-md shadow-none">
      <button class="px-20 py-8 rounded-full bg-gradient-to-r from-yellow-400 via-pink-400 to-fuchsia-400 text-white text-3xl font-bold shadow-2xl hover:scale-110 hover:from-fuchsia-400 hover:to-yellow-400 transition-all duration-300" @click="showRegister = true">
        Hemen Başla
      </button>
    </footer>

    <!-- Auth Modal -->
    <AuthModal :show="showLogin || showRegister" @close="showLogin = false; showRegister = false" :mode="showRegister ? 'register' : 'login'" />
  </div>
</template>

<script setup>
import { useAuthStore } from '~/stores/auth'
import AuthModal from '~/components/AuthModal.vue'
import SwipeDemo from '~/components/SwipeDemo.vue'
import { ref, computed } from 'vue'

const auth = useAuthStore()
const isAuthenticated = computed(() => auth.isAuthenticated)
const showLogin = ref(false)
const showRegister = ref(false)
</script>

<style scoped>
@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-18px); }
}
.animate-float {
  animation: float 5s ease-in-out infinite;
}
.animate-float-slow {
  animation: float 12s ease-in-out infinite;
}
.animate-float-reverse {
  animation: float 7s ease-in-out infinite reverse;
}
</style> 