import { defineStore } from 'pinia'
import type { User } from './auth'

export interface Match {
  id: string
  user: User
  matchScore: number
  isLiked: boolean
  isMatched: boolean
  lastMessage?: string
  lastMessageTime?: string
}

export const useMatchingStore = defineStore('matching', () => {
  const potentialMatches = ref<User[]>([])
  const currentMatch = ref<User | null>(null)
  const matches = ref<Match[]>([])
  const isLoading = ref(false)

  const fetchPotentialMatches = async () => {
    isLoading.value = true
    try {
      const response = await $fetch<User[]>('/api/matches/potential')
      potentialMatches.value = response
    } catch (error) {
      console.error('Error fetching potential matches:', error)
    } finally {
      isLoading.value = false
    }
  }

  const likeUser = async (userId: string) => {
    try {
      const response = await $fetch<{ isMatch: boolean }>('/api/matches/like', {
        method: 'POST',
        body: { userId }
      })
      
      if (response.isMatch) {
        // Handle new match
        await fetchMatches()
      }
      
      return response
    } catch (error) {
      console.error('Error liking user:', error)
      throw error
    }
  }

  const dislikeUser = async (userId: string) => {
    try {
      await $fetch('/api/matches/dislike', {
        method: 'POST',
        body: { userId }
      })
    } catch (error) {
      console.error('Error disliking user:', error)
      throw error
    }
  }

  const fetchMatches = async () => {
    isLoading.value = true
    try {
      const response = await $fetch<Match[]>('/api/matches')
      matches.value = response
    } catch (error) {
      console.error('Error fetching matches:', error)
    } finally {
      isLoading.value = false
    }
  }

  const getNextMatch = () => {
    if (potentialMatches.value.length > 0) {
      currentMatch.value = potentialMatches.value[0]
      potentialMatches.value = potentialMatches.value.slice(1)
    } else {
      currentMatch.value = null
    }
  }

  const getMatchScore = async (user1: User, user2: User) => {
    try {
      const response = await $fetch<{ score: number }>('/api/matches/score', {
        method: 'POST',
        body: { user1, user2 }
      })
      return response.score
    } catch (error) {
      console.error('Error getting match score:', error)
      return 0
    }
  }

  return {
    potentialMatches: readonly(potentialMatches),
    currentMatch: readonly(currentMatch),
    matches: readonly(matches),
    isLoading: readonly(isLoading),
    fetchPotentialMatches,
    likeUser,
    dislikeUser,
    fetchMatches,
    getNextMatch,
    getMatchScore
  }
}) 