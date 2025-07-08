<template>
  <div v-if="show" class="fixed inset-0 z-50 flex items-center justify-center bg-black/70 overflow-y-auto p-4">
    <div class="bg-black rounded-2xl shadow-2xl p-8 w-full max-w-sm relative border border-pink-500 my-8">
      <button @click="$emit('close')" class="absolute top-3 right-3 text-pink-300 hover:text-white text-2xl">&times;</button>
      <h2 class="text-2xl font-bold text-center mb-6 bg-gradient-to-r from-pink-500 to-fuchsia-600 bg-clip-text text-transparent">{{ mode === 'login' ? 'Giriş Yap' : 'Kayıt Ol' }}</h2>
      
      <!-- Loading State -->
      <div v-if="auth.isLoading" class="text-center py-8">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-pink-500 mx-auto"></div>
        <p class="text-pink-300 mt-2">İşleniyor...</p>
      </div>

      <!-- Error Message -->
      <div v-if="auth.error" class="bg-red-900/50 border border-red-500 text-red-200 p-3 rounded-lg mb-4">
        {{ auth.error }}
      </div>

      <!-- Login Form -->
      <form v-if="mode === 'login' && !auth.isLoading" @submit.prevent="handleLogin" class="space-y-4">
        <input v-model="loginForm.email" type="email" placeholder="E-posta" class="w-full px-4 py-2 rounded-lg bg-black border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500" required />
        <input v-model="loginForm.password" type="password" placeholder="Şifre" class="w-full px-4 py-2 rounded-lg bg-black border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500" required />
        <button type="submit" class="w-full btn-primary py-2 text-lg">Giriş Yap</button>
      </form>

      <!-- Simple Register Form -->
      <form v-if="mode === 'register' && !auth.isLoading" @submit.prevent="handleRegister" class="space-y-4">
        <input v-model="registerForm.name" type="text" placeholder="Adınız" class="w-full px-4 py-2 rounded-lg bg-black border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500" required />
        <input v-model="registerForm.email" type="email" placeholder="E-posta" class="w-full px-4 py-2 rounded-lg bg-black border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500" required />
        <input v-model="registerForm.password" type="password" placeholder="Şifre (en az 6 karakter)" class="w-full px-4 py-2 rounded-lg bg-black border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500" required />
        
        <div class="text-center text-sm text-pink-300 mt-4">
          <p>Kayıt olduktan sonra profil bilgilerinizi tamamlayabilirsiniz.</p>
        </div>

        <button type="submit" class="w-full btn-primary py-2 text-lg">Kayıt Ol</button>
      </form>

      <div class="text-center mt-4">
        <button v-if="mode === 'login'" @click="mode = 'register'" class="text-pink-300 hover:underline">Hesabın yok mu? Kayıt Ol</button>
        <button v-else @click="mode = 'login'" class="text-pink-300 hover:underline">Zaten hesabın var mı? Giriş Yap</button>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  show: Boolean
})
const emit = defineEmits(['close'])

const mode = ref('login')
const auth = useAuthStore()

// Login form
const loginForm = ref({
  email: '',
  password: ''
})

// Simple register form
const registerForm = ref({
  name: '',
  email: '',
  password: ''
})

const handleLogin = async () => {
  const result = await auth.login(loginForm.value.email, loginForm.value.password)
  if (result.success) {
    emit('close')
    loginForm.value = { email: '', password: '' }
  }
}

const handleRegister = async () => {
  const result = await auth.simpleRegister(registerForm.value)
  if (result.success) {
    emit('close')
    registerForm.value = { name: '', email: '', password: '' }
    // Profil tamamlama sayfasına yönlendir
    await navigateTo('/profile-setup')
  }
}

// Reset form when modal opens
watch(() => props.show, (newVal) => {
  if (newVal) {
    auth.error = null
    loginForm.value = { email: '', password: '' }
    registerForm.value = { name: '', email: '', password: '' }
  }
})
</script>

<style scoped>
.btn-primary {
  @apply bg-gradient-to-r from-pink-500 to-fuchsia-600 hover:from-fuchsia-600 hover:to-pink-500 text-white font-bold rounded-lg shadow-lg transition-all;
}
</style> 