<template>
  <div class="min-h-screen bg-gradient-to-br from-pink-900 via-purple-900 to-black text-white">
    <div class="container mx-auto px-4 py-8">
      <!-- Header -->
      <div class="text-center mb-8">
        <h1 class="text-4xl font-bold bg-gradient-to-r from-pink-500 to-fuchsia-600 bg-clip-text text-transparent mb-4">
          Profilini Tamamla
        </h1>
        <p class="text-pink-300">Daha iyi eşleşmeler için profil bilgilerini doldur</p>
      </div>

      <!-- Loading State -->
      <div v-if="isLoading" class="text-center py-8">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-pink-500 mx-auto"></div>
        <p class="text-pink-300 mt-2">Kaydediliyor...</p>
      </div>

      <!-- Error Message -->
      <div v-if="error" class="bg-red-900/50 border border-red-500 text-red-200 p-4 rounded-lg mb-6">
        {{ error }}
      </div>

      <!-- Profile Form -->
      <form v-if="!isLoading" @submit.prevent="handleSubmit" class="max-w-2xl mx-auto space-y-8">
        <!-- Basic Info -->
        <div class="bg-black/30 backdrop-blur-sm rounded-2xl p-6 border border-pink-500/30">
          <h2 class="text-2xl font-semibold text-pink-300 mb-4">Temel Bilgiler</h2>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-pink-200 mb-2">Kendiniz hakkında</label>
              <textarea 
                v-model="profile.bio" 
                placeholder="Kendiniz hakkında kısa bir yazı..." 
                class="w-full px-4 py-3 rounded-lg bg-black/50 border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500" 
                rows="3"
              ></textarea>
            </div>
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-pink-200 mb-2">Yaş</label>
                <input 
                  v-model.number="profile.age" 
                  type="number" 
                  min="18" 
                  max="100" 
                  class="w-full px-4 py-3 rounded-lg bg-black/50 border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-pink-200 mb-2">Ciddiyet (1-10)</label>
                <input 
                  v-model.number="profile.seriousness" 
                  type="number" 
                  min="1" 
                  max="10" 
                  class="w-full px-4 py-3 rounded-lg bg-black/50 border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500"
                />
              </div>
            </div>
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-pink-200 mb-2">Boy (cm)</label>
                <input 
                  v-model.number="profile.height" 
                  type="number" 
                  min="140" 
                  max="220" 
                  class="w-full px-4 py-3 rounded-lg bg-black/50 border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-pink-200 mb-2">Kilo (kg)</label>
                <input 
                  v-model.number="profile.weight" 
                  type="number" 
                  min="40" 
                  max="200" 
                  class="w-full px-4 py-3 rounded-lg bg-black/50 border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500"
                />
              </div>
            </div>
          </div>
        </div>

        <!-- Lifestyle -->
        <div class="bg-black/30 backdrop-blur-sm rounded-2xl p-6 border border-pink-500/30">
          <h2 class="text-2xl font-semibold text-pink-300 mb-4">Yaşam Tarzı</h2>
          <div class="space-y-4">
            <div class="flex space-x-6">
              <label class="flex items-center">
                <input v-model="profile.smokes" type="checkbox" class="mr-2 w-4 h-4 text-pink-500 bg-black border-pink-400 rounded focus:ring-pink-500" />
                <span class="text-white">Sigara içiyor</span>
              </label>
              <label class="flex items-center">
                <input v-model="profile.drinks" type="checkbox" class="mr-2 w-4 h-4 text-pink-500 bg-black border-pink-400 rounded focus:ring-pink-500" />
                <span class="text-white">Alkol kullanıyor</span>
              </label>
            </div>
          </div>
        </div>

        <!-- Education & Job -->
        <div class="bg-black/30 backdrop-blur-sm rounded-2xl p-6 border border-pink-500/30">
          <h2 class="text-2xl font-semibold text-pink-300 mb-4">Eğitim & İş</h2>
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-pink-200 mb-2">Eğitim Seviyesi</label>
              <select v-model="profile.education" class="w-full px-4 py-3 rounded-lg bg-black/50 border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500">
                <option value="">Eğitim seviyesi seçin</option>
                <option value="İlkokul">İlkokul</option>
                <option value="Ortaokul">Ortaokul</option>
                <option value="Lise">Lise</option>
                <option value="Üniversite">Üniversite</option>
                <option value="Yüksek Lisans">Yüksek Lisans</option>
                <option value="Doktora">Doktora</option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-pink-200 mb-2">Meslek</label>
              <input v-model="profile.job" type="text" placeholder="Mesleğiniz" class="w-full px-4 py-3 rounded-lg bg-black/50 border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500" />
            </div>
            <div>
              <label class="block text-sm font-medium text-pink-200 mb-2">İş Kategorisi</label>
              <select v-model="profile.jobCategory" class="w-full px-4 py-3 rounded-lg bg-black/50 border border-pink-400 text-white focus:outline-none focus:ring-2 focus:ring-pink-500">
                <option value="">İş kategorisi seçin</option>
                <option value="Teknoloji">Teknoloji</option>
                <option value="Sağlık">Sağlık</option>
                <option value="Eğitim">Eğitim</option>
                <option value="Finans">Finans</option>
                <option value="Sanat">Sanat</option>
                <option value="Spor">Spor</option>
                <option value="Diğer">Diğer</option>
              </select>
            </div>
          </div>
        </div>

        <!-- Hobbies -->
        <div class="bg-black/30 backdrop-blur-sm rounded-2xl p-6 border border-pink-500/30">
          <h2 class="text-2xl font-semibold text-pink-300 mb-4">Hobiler</h2>
          <p class="text-pink-200 text-sm mb-4">İlgi alanlarınızı seçin (en az 1)</p>
          <div class="grid grid-cols-2 gap-3">
            <label v-for="hobby in availableHobbies" :key="hobby" class="flex items-center p-3 rounded-lg bg-black/30 border border-pink-400/30 hover:border-pink-400 cursor-pointer">
              <input v-model="profile.hobbies" type="checkbox" :value="hobby" class="mr-3 w-4 h-4 text-pink-500 bg-black border-pink-400 rounded focus:ring-pink-500" />
              <span class="text-white">{{ hobby }}</span>
            </label>
          </div>
        </div>

        <!-- Submit Button -->
        <div class="text-center">
          <button type="submit" class="btn-primary px-8 py-3 text-lg">
            Profili Kaydet
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
const auth = useAuthStore()
const router = useRouter()

const isLoading = ref(false)
const error = ref('')

const profile = ref({
  bio: '',
  age: 25,
  seriousness: 5,
  height: 170,
  weight: 70,
  smokes: false,
  drinks: false,
  job: '',
  jobCategory: '',
  education: '',
  hobbies: []
})

const availableHobbies = [
  'Futbol', 'Basketbol', 'Müzik', 'Kitap okuma', 'Yemek yapma', 'Seyahat',
  'Fotoğrafçılık', 'Dans', 'Yoga', 'Fitness', 'Sinema', 'Tiyatro',
  'Resim', 'Çizim', 'Yazı yazma', 'Şiir', 'El sanatları', 'Bahçe işleri',
  'Programlama', 'Doğa yürüyüşü', 'Kamp', 'Fotoğrafçılık', 'Kahve'
]

const handleSubmit = async () => {
  isLoading.value = true
  error.value = ''

  try {
    // Hobi kategorilerini otomatik belirle
    const hobbyCategories = []
    if (profile.value.hobbies.some(h => ['Futbol', 'Basketbol', 'Fitness', 'Yoga'].includes(h))) {
      hobbyCategories.push('Spor')
    }
    if (profile.value.hobbies.some(h => ['Müzik', 'Resim', 'Çizim', 'Sinema', 'Tiyatro'].includes(h))) {
      hobbyCategories.push('Sanat')
    }
    if (profile.value.hobbies.some(h => ['Kitap okuma', 'Yazı yazma', 'Şiir'].includes(h))) {
      hobbyCategories.push('Edebiyat')
    }
    if (profile.value.hobbies.some(h => ['Programlama', 'Teknoloji'].includes(h))) {
      hobbyCategories.push('Teknoloji')
    }

    const updateData = {
      ...profile.value,
      hobbyCategories
    }

    // Backend'e profil güncelleme isteği gönder
    if (auth.user?.id) {
      const result = await auth.updateProfile(auth.user.id, updateData)
      if (result.success) {
        // Başarılı güncelleme sonrası ana sayfaya yönlendir
        await router.push('/')
      } else {
        error.value = result.error || 'Profil güncellenirken bir hata oluştu'
      }
    } else {
      error.value = 'Kullanıcı bilgisi bulunamadı'
    }
  } catch (err) {
    error.value = 'Profil güncellenirken bir hata oluştu'
  } finally {
    isLoading.value = false
  }
}

// Kullanıcı giriş yapmamışsa login sayfasına yönlendir
onMounted(() => {
  if (!auth.isAuthenticated) {
    router.push('/')
  }
})
</script>

<style scoped>
.btn-primary {
  @apply bg-gradient-to-r from-pink-500 to-fuchsia-600 hover:from-fuchsia-600 hover:to-pink-500 text-white font-bold rounded-lg shadow-lg transition-all;
}
</style> 