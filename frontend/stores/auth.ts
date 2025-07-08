import { defineStore } from 'pinia'

export interface User {
  id: string
  name: string
  email: string
  bio?: string
  age?: number
  seriousness?: number
  height?: number
  weight?: number
  smokes?: boolean
  drinks?: boolean
  job?: string
  jobCategory?: string
  education?: string
  hobbies?: string[]
  hobbyCategories?: string[]
}

export interface RegisterData {
  name: string
  email: string
  password: string
  bio: string
  age: number
  seriousness: number
  height: number
  weight: number
  smokes: boolean
  drinks: boolean
  job: string
  jobCategory: string
  education: string
  hobbies: string[]
  hobbyCategories: string[]
}

export interface SimpleRegisterData {
  name: string
  email: string
  password: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const isAuthenticated = computed(() => !!user.value)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  // API base URL
  const API_BASE = 'http://localhost:8081/api'

  // Login with real API
  const login = async (email: string, password: string) => {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await $fetch<{token: string, user: User}>(`${API_BASE}/auth/login`, {
        method: 'POST',
        body: {
          email,
          password
        }
      })

      if (response.token) {
        user.value = response.user
        // Token'ı localStorage'a kaydet
        localStorage.setItem('auth_token', response.token)
        return { success: true }
      }
    } catch (err: any) {
      error.value = err.data?.message || 'Giriş başarısız'
      return { success: false, error: error.value }
    } finally {
      isLoading.value = false
    }
  }

  // Register with real API
  const register = async (data: RegisterData) => {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await $fetch<{user_id: string, message: string}>(`${API_BASE}/auth/register`, {
        method: 'POST',
        body: data
      })

      if (response.user_id) {
        // Kayıt başarılı, kullanıcıyı otomatik giriş yap
        await login(data.email, data.password)
        return { success: true }
      }
    } catch (err: any) {
      error.value = err.data?.message || 'Kayıt başarısız'
      return { success: false, error: error.value }
    } finally {
      isLoading.value = false
    }
  }

  // Simple register with real API
  const simpleRegister = async (data: SimpleRegisterData) => {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await $fetch<{user_id: string, message: string}>(`${API_BASE}/auth/simple-register`, {
        method: 'POST',
        body: data
      })

      if (response.user_id) {
        // Kayıt başarılı, kullanıcıyı otomatik giriş yap
        await login(data.email, data.password)
        return { success: true }
      }
    } catch (err: any) {
      error.value = err.data?.message || 'Kayıt başarısız'
      return { success: false, error: error.value }
    } finally {
      isLoading.value = false
    }
  }

  // Get form data (hobby categories, education levels, job categories)
  const getHobbyCategories = async () => {
    try {
      return await $fetch(`${API_BASE}/form/hobby-categories`)
    } catch (err) {
      console.error('Hobi kategorileri alınamadı:', err)
      return null
    }
  }

  const getEducationLevels = async () => {
    try {
      return await $fetch(`${API_BASE}/form/education-levels`)
    } catch (err) {
      console.error('Eğitim seviyeleri alınamadı:', err)
      return null
    }
  }

  const getJobCategories = async () => {
    try {
      return await $fetch(`${API_BASE}/form/job-categories`)
    } catch (err) {
      console.error('İş kategorileri alınamadı:', err)
      return null
    }
  }

  // Update user profile
  const updateProfile = async (userId: string, profileData: any) => {
    isLoading.value = true
    error.value = null
    
    try {
      const response = await $fetch(`${API_BASE}/users/${userId}`, {
        method: 'PUT',
        body: profileData
      })
      return { success: true }
    } catch (err: any) {
      error.value = err.data?.message || 'Profil güncellenemedi'
      return { success: false, error: error.value }
    } finally {
      isLoading.value = false
    }
  }

  const logout = () => {
    user.value = null
    localStorage.removeItem('auth_token')
  }

  // Check if user is already logged in (on app start)
  const checkAuth = () => {
    const token = localStorage.getItem('auth_token')
    if (token && !user.value) {
      // TODO: Token'ı validate et ve kullanıcı bilgilerini al
      // Şimdilik sadece token varsa authenticated sayıyoruz
    }
  }

  return {
    user,
    isAuthenticated,
    isLoading,
    error,
    login,
    register,
    simpleRegister,
    updateProfile,
    logout,
    getHobbyCategories,
    getEducationLevels,
    getJobCategories,
    checkAuth
  }
}) 