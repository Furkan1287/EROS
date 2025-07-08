import { defineStore } from 'pinia'

export interface Message {
  id: string
  senderId: string
  receiverId: string
  content: string
  timestamp: string
  isRead: boolean
  type: 'text' | 'image' | 'location'
}

export interface Chat {
  id: string
  participants: string[]
  lastMessage?: Message
  unreadCount: number
  isActive: boolean
}

export const useChatStore = defineStore('chat', () => {
  const chats = ref<Chat[]>([])
  const currentChat = ref<Chat | null>(null)
  const messages = ref<Message[]>([])
  const isLoading = ref(false)
  const isTyping = ref(false)

  const fetchChats = async () => {
    isLoading.value = true
    try {
      const response = await $fetch<Chat[]>('/api/chats')
      chats.value = response
    } catch (error) {
      console.error('Error fetching chats:', error)
    } finally {
      isLoading.value = false
    }
  }

  const fetchMessages = async (chatId: string) => {
    isLoading.value = true
    try {
      const response = await $fetch<Message[]>(`/api/chats/${chatId}/messages`)
      messages.value = response
    } catch (error) {
      console.error('Error fetching messages:', error)
    } finally {
      isLoading.value = false
    }
  }

  const sendMessage = async (chatId: string, content: string, type: 'text' | 'image' | 'location' = 'text') => {
    try {
      const response = await $fetch<Message>('/api/chats/messages', {
        method: 'POST',
        body: {
          chatId,
          content,
          type
        }
      })
      
      messages.value.push(response)
      
      // Update chat's last message
      const chat = chats.value.find(c => c.id === chatId)
      if (chat) {
        chat.lastMessage = response
        chat.unreadCount = 0
      }
      
      return response
    } catch (error) {
      console.error('Error sending message:', error)
      throw error
    }
  }

  const markAsRead = async (chatId: string) => {
    try {
      await $fetch(`/api/chats/${chatId}/read`, {
        method: 'POST'
      })
      
      // Update local state
      const chat = chats.value.find(c => c.id === chatId)
      if (chat) {
        chat.unreadCount = 0
      }
      
      messages.value.forEach(msg => {
        if (msg.senderId !== 'current-user-id') {
          msg.isRead = true
        }
      })
    } catch (error) {
      console.error('Error marking as read:', error)
    }
  }

  const startTyping = async (chatId: string) => {
    try {
      await $fetch(`/api/chats/${chatId}/typing`, {
        method: 'POST',
        body: { isTyping: true }
      })
      isTyping.value = true
    } catch (error) {
      console.error('Error starting typing:', error)
    }
  }

  const stopTyping = async (chatId: string) => {
    try {
      await $fetch(`/api/chats/${chatId}/typing`, {
        method: 'POST',
        body: { isTyping: false }
      })
      isTyping.value = false
    } catch (error) {
      console.error('Error stopping typing:', error)
    }
  }

  const getIceBreaker = async (userId: string) => {
    try {
      const response = await $fetch<{ message: string }>('/api/chat/ice-breaker', {
        method: 'POST',
        body: { userId }
      })
      return response.message
    } catch (error) {
      console.error('Error getting ice breaker:', error)
      return null
    }
  }

  const getDateSuggestion = async (userId: string) => {
    try {
      const response = await $fetch<{ suggestion: string }>('/api/chat/date-suggestion', {
        method: 'POST',
        body: { userId }
      })
      return response.suggestion
    } catch (error) {
      console.error('Error getting date suggestion:', error)
      return null
    }
  }

  return {
    chats: readonly(chats),
    currentChat: readonly(currentChat),
    messages: readonly(messages),
    isLoading: readonly(isLoading),
    isTyping: readonly(isTyping),
    fetchChats,
    fetchMessages,
    sendMessage,
    markAsRead,
    startTyping,
    stopTyping,
    getIceBreaker,
    getDateSuggestion
  }
}) 