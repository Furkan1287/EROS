<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-gray-800">Mesajlar</h1>
      <button 
        @click="fetchChats"
        :disabled="isLoading"
        class="btn-secondary"
      >
        <svg v-if="isLoading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        Yenile
      </button>
    </div>

    <!-- Chat List -->
    <div v-if="chats.length > 0" class="bg-white rounded-2xl shadow-sm">
      <div 
        v-for="chat in chats" 
        :key="chat.id"
        class="p-4 border-b border-gray-100 last:border-b-0 hover:bg-gray-50 cursor-pointer transition-colors"
        @click="openChat(chat)"
      >
        <div class="flex items-center space-x-3">
          <!-- Chat Avatar -->
          <div class="relative">
            <img 
              :src="getChatAvatar(chat)" 
              :alt="getChatName(chat)"
              class="w-12 h-12 rounded-full object-cover"
            />
            <div 
              v-if="chat.isActive"
              class="absolute -bottom-1 -right-1 w-4 h-4 bg-green-500 border-2 border-white rounded-full"
            ></div>
          </div>
          
          <!-- Chat Info -->
          <div class="flex-1 min-w-0">
            <div class="flex justify-between items-start">
              <div>
                <h3 class="font-semibold text-gray-800 truncate">{{ getChatName(chat) }}</h3>
                <p v-if="chat.lastMessage" class="text-sm text-gray-600 truncate mt-1">
                  {{ chat.lastMessage.content }}
                </p>
              </div>
              <div class="text-right">
                <div v-if="chat.unreadCount > 0" class="bg-primary-500 text-white text-xs rounded-full px-2 py-1 mb-1">
                  {{ chat.unreadCount }}
                </div>
                <div v-if="chat.lastMessage" class="text-xs text-gray-500">
                  {{ formatTime(chat.lastMessage.timestamp) }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="text-center py-12">
      <div class="w-24 h-24 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <svg class="w-12 h-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
        </svg>
      </div>
      <h3 class="text-xl font-semibold text-gray-800 mb-2">Henüz Mesaj Yok</h3>
      <p class="text-gray-600 mb-6">Eşleştiğiniz kişilerle sohbet başlatın ve gerçek bağlantılar kurun.</p>
      <NuxtLink to="/matches" class="btn-primary">
        Eşleşmeleri Gör
      </NuxtLink>
    </div>
  </div>
</template>

<script setup>
const chatStore = useChatStore()
const { chats, isLoading } = storeToRefs(chatStore)
const { fetchChats } = chatStore

// Fetch chats on mount
onMounted(() => {
  fetchChats()
})

const getChatAvatar = (chat) => {
  // This would be implemented based on your chat structure
  return '/default-avatar.png'
}

const getChatName = (chat) => {
  // This would be implemented based on your chat structure
  return 'Chat User'
}

const formatTime = (timestamp) => {
  const date = new Date(timestamp)
  const now = new Date()
  const diffInHours = (now.getTime() - date.getTime()) / (1000 * 60 * 60)
  
  if (diffInHours < 1) {
    return 'Az önce'
  } else if (diffInHours < 24) {
    return `${Math.floor(diffInHours)}s önce`
  } else {
    return date.toLocaleDateString('tr-TR')
  }
}

const openChat = (chat) => {
  navigateTo(`/chat/${chat.id}`)
}
</script> 