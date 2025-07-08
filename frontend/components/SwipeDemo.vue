<template>
  <div class="relative w-full max-w-xs mx-auto">
    <div v-if="current < demoUsers.length" class="bg-black border-2 border-pink-400 rounded-2xl shadow-xl p-6 flex flex-col items-center animate-fade-in">
      <img :src="demoUsers[current].photo" :alt="demoUsers[current].name" class="w-32 h-32 object-cover rounded-full border-4 border-pink-300 mb-4" />
      <h3 class="text-2xl font-bold text-pink-200 mb-1">{{ demoUsers[current].name }}, {{ demoUsers[current].age }}</h3>
      <p class="text-pink-100 mb-4 text-center">{{ demoUsers[current].bio }}</p>
      <div class="flex gap-6 mt-2">
        <button @click="swipe('left')" class="w-12 h-12 rounded-full bg-pink-900 hover:bg-pink-700 flex items-center justify-center text-white text-2xl shadow-lg transition-all">✖️</button>
        <button @click="swipe('right')" class="w-12 h-12 rounded-full bg-pink-500 hover:bg-pink-400 flex items-center justify-center text-white text-2xl shadow-lg transition-all">❤️</button>
      </div>
    </div>
    <div v-else class="text-center text-pink-200 py-12">
      <div class="text-4xl mb-4">🎉</div>
      <div class="text-lg">Kaydıracak kart kalmadı!<br>Giriş yaparak gerçek eşleşmeleri görebilirsin.</div>
    </div>
  </div>
</template>

<script setup>
const demoUsers = [
  {
    name: 'Zeynep',
    age: 27,
    bio: 'Kitap okumayı ve doğa yürüyüşlerini severim. Gerçek aşkı arıyorum.',
    photo: 'https://randomuser.me/api/portraits/women/44.jpg'
  },
  {
    name: 'Ahmet',
    age: 29,
    bio: 'Müzik, kahve ve uzun sohbetler. Ciddi ilişki arıyorum.',
    photo: 'https://randomuser.me/api/portraits/men/32.jpg'
  },
  {
    name: 'Elif',
    age: 25,
    bio: 'Sanat, tiyatro ve yeni yerler keşfetmek hobim.',
    photo: 'https://randomuser.me/api/portraits/women/68.jpg'
  }
]
const current = ref(0)

function swipe(dir) {
  current.value++
}
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.5s;
}
@keyframes fadeIn {
  from { opacity: 0; transform: scale(0.95); }
  to { opacity: 1; transform: scale(1); }
}
</style> 