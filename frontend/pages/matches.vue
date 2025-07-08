<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-gray-800">Eşleşmeler</h1>
      <button 
        @click="fetchPotentialMatches"
        :disabled="isLoading"
        class="btn-secondary"
      >
        <svg v-if="isLoading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        Yeni Eşleşmeler
      </button>
    </div>

    <!-- Current Match Card -->
    <div v-if="currentMatch" class="bg-white rounded-2xl shadow-lg overflow-hidden">
      <div class="relative">
        <!-- Profile Image -->
        <div class="relative h-96 bg-gradient-to-br from-primary-100 to-secondary-100">
          <img 
            :src="currentMatch.photos?.[0] || '/default-avatar.png'"
            :alt="currentMatch.name"
            class="w-full h-full object-cover"
          />
          <div class="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent"></div>
          
          <!-- Match Info Overlay -->
          <div class="absolute bottom-0 left-0 right-0 p-6 text-white">
            <div class="flex justify-between items-end">
              <div>
                <h2 class="text-2xl font-bold mb-2">{{ currentMatch.name }}, {{ currentMatch.age }}</h2>
                <p class="text-white/90 mb-3">{{ currentMatch.bio }}</p>
                <div class="flex flex-wrap gap-2">
                  <span 
                    v-for="interest in currentMatch.interests?.slice(0, 3)" 
                    :key="interest"
                    class="bg-white/20 backdrop-blur-sm px-3 py-1 rounded-full text-sm"
                  >
                    {{ interest }}
                  </span>
                </div>
              </div>
              
              <!-- Match Score -->
              <div class="text-right">
                <div class="bg-white/20 backdrop-blur-sm rounded-full p-3">
                  <div class="text-2xl font-bold">{{ matchScore }}%</div>
                  <div class="text-sm opacity-90">Uyum</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="p-6">
        <div class="flex justify-center space-x-4">
          <button 
            @click="dislikeUser"
            class="w-16 h-16 bg-red-500 hover:bg-red-600 text-white rounded-full flex items-center justify-center transition-colors"
          >
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
          
          <button 
            @click="likeUser"
            class="w-16 h-16 bg-green-500 hover:bg-green-600 text-white rounded-full flex items-center justify-center transition-colors"
          >
            <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- No More Matches -->
    <div v-else class="text-center py-12">
      <div class="w-24 h-24 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <svg class="w-12 h-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
        </svg>
      </div>
      <h3 class="text-xl font-semibold text-gray-800 mb-2">Henüz Eşleşme Yok</h3>
      <p class="text-gray-600 mb-6">Yeni eşleşmeler için biraz bekleyin veya profil ayarlarınızı güncelleyin.</p>
      <button 
        @click="fetchPotentialMatches"
        class="btn-primary"
      >
        Yeni Eşleşmeler Ara
      </button>
    </div>

    <!-- Matches List -->
    <div v-if="matches.length > 0" class="bg-white rounded-2xl shadow-sm p-6">
      <h3 class="text-xl font-semibold mb-4 text-gray-800">Başarılı Eşleşmeler</h3>
      <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
        <div 
          v-for="match in matches" 
          :key="match.id"
          class="card-hover bg-gray-50 rounded-xl p-4 cursor-pointer"
          @click="openChat(match)"
        >
          <div class="flex items-center space-x-3">
            <img 
              :src="match.user.photos?.[0] || '/default-avatar.png'"
              :alt="match.user.name"
              class="w-12 h-12 rounded-full object-cover"
            />
            <div class="flex-1">
              <h4 class="font-semibold text-gray-800">{{ match.user.name }}</h4>
              <p class="text-sm text-gray-600">{{ match.user.age }} yaşında</p>
              <div class="flex items-center space-x-2 mt-1">
                <div class="text-xs bg-primary-100 text-primary-600 px-2 py-1 rounded-full">
                  {{ match.matchScore }}% uyum
                </div>
                <div v-if="match.isMatched" class="text-xs bg-green-100 text-green-600 px-2 py-1 rounded-full">
                  Eşleşti
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const matchingStore = useMatchingStore()
const { currentMatch, matches, isLoading } = storeToRefs(matchingStore)
const { fetchPotentialMatches, likeUser, dislikeUser, getMatchScore } = matchingStore

const matchScore = ref(0)

// Fetch initial data
onMounted(async () => {
  await fetchPotentialMatches()
  await matchingStore.fetchMatches()
})

// Calculate match score when current match changes
watch(currentMatch, async (newMatch) => {
  if (newMatch) {
    const authStore = useAuthStore()
    const currentUser = authStore.user
    if (currentUser) {
      matchScore.value = await getMatchScore(currentUser, newMatch)
    }
  }
})

const openChat = (match) => {
  // Navigate to chat with this match
  navigateTo(`/chat/${match.id}`)
}
</script> 