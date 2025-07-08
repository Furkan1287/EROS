// ai_models.go - AI modelleri ve kullanım alanları
package types

// AI Model Types
const (
	// Tüm AI servisleri için tek model
	ModelChatAnalysis    = "google/gemma-3-27b-it:free"
	ModelDateSuggestion  = "google/gemma-3-27b-it:free"
	ModelSecurityFilter  = "google/gemma-3n-e4b-it:free"
	ModelIceBreaker      = "google/gemma-3-27b-it:free"
	ModelProfileMatching = "google/gemma-3-27b-it:free"
)

// AI Task Types
const (
	TaskChatAnalysis    = "chat_analysis"
	TaskDateSuggestion  = "date_suggestion"
	TaskSecurityFilter  = "security_filter"
	TaskIceBreaker      = "ice_breaker"
	TaskProfileMatching = "profile_matching"
)

// Hobi Kategorileri
var HobbyCategories = map[string][]string{
	"Spor": {
		"Futbol", "Basketbol", "Voleybol", "Tenis", "Yüzme", "Koşu", "Fitness",
		"Yoga", "Pilates", "Dans", "Boks", "Kickbox", "Güreş", "Atletizm",
		"Bisiklet", "Kayak", "Snowboard", "Sörf", "Dalış", "Tırmanış",
	},
	"Sanat": {
		"Resim", "Çizim", "Fotoğrafçılık", "Müzik", "Gitar", "Piyano", "Şarkı söyleme",
		"Dans", "Tiyatro", "Sinema", "Yazı yazma", "Şiir", "El sanatları",
		"Takı yapımı", "Dikiş", "Örgü", "Seramik", "Heykel",
	},
	"Teknoloji": {
		"Programlama", "Web tasarımı", "Mobil uygulama", "Oyun geliştirme",
		"Yapay zeka", "Veri analizi", "Siber güvenlik", "Blockchain",
		"Robotik", "Drone", "3D yazıcı", "Elektronik",
	},
	"Doğa": {
		"Kamp", "Trekking", "Doğa yürüyüşü", "Dağcılık", "Balıkçılık",
		"Bahçıvanlık", "Bitki yetiştirme", "Kuş gözlemi", "Astronomi",
		"Fosil toplama", "Mineral koleksiyonu",
	},
	"Seyahat": {
		"Backpacking", "Kültür turları", "Şehir keşfi", "Müze gezileri",
		"Festival takibi", "Yolculuk", "Dil öğrenme", "Kültür değişimi",
	},
	"Yemek": {
		"Yemek yapma", "Pastacılık", "Kahve", "Şarap", "Bira yapımı",
		"Restoran keşfi", "Yemek fotoğrafçılığı", "Tarif geliştirme",
	},
	"Sosyal": {
		"Sohbet", "Kitap okuma", "Film izleme", "Dizi takibi", "Podcast",
		"Sosyal medya", "Blog yazma", "Vlog", "Topluluk aktiviteleri",
	},
	"Eğitim": {
		"Dil öğrenme", "Kurs alma", "Sertifika", "Online eğitim",
		"Kitap okuma", "Araştırma", "Öğretmenlik", "Mentorluk",
	},
	"İş": {
		"Girişimcilik", "Freelance", "Networking", "Konferans",
		"Workshop", "Mentorluk", "Danışmanlık",
	},
	"Diğer": {
		"Koleksiyon", "Bulmaca", "Satranç", "Poker", "Kumar",
		"Meditasyon", "Feng shui", "Astroloji", "Numeroloji",
	},
}

// Eğitim Seviyeleri
var EducationLevels = []string{
	"İlkokul",
	"Ortaokul",
	"Lise",
	"Üniversite (Önlisans)",
	"Üniversite (Lisans)",
	"Yüksek Lisans",
	"Doktora",
	"Diğer",
}

// Meslek Kategorileri
var JobCategories = map[string][]string{
	"Teknoloji": {
		"Yazılım Geliştirici", "Veri Bilimci", "Sistem Yöneticisi", "DevOps",
		"UI/UX Tasarımcı", "Proje Yöneticisi", "Test Uzmanı", "Güvenlik Uzmanı",
	},
	"Sağlık": {
		"Doktor", "Hemşire", "Eczacı", "Diş Hekimi", "Psikolog", "Fizyoterapist",
		"Beslenme Uzmanı", "Veteriner",
	},
	"Eğitim": {
		"Öğretmen", "Akademisyen", "Eğitmen", "Koç", "Mentor", "Danışman",
	},
	"Finans": {
		"Muhasebeci", "Finans Uzmanı", "Yatırım Danışmanı", "Bankacı",
		"Sigortacı", "Ekonomist", "Aktüer",
	},
	"Hukuk": {
		"Avukat", "Hakim", "Savcı", "Noter", "Hukuk Danışmanı",
	},
	"Medya": {
		"Gazeteci", "Editör", "Yazar", "Muhabir", "Yayıncı", "İçerik Üreticisi",
	},
	"Sanat": {
		"Müzisyen", "Aktör", "Yönetmen", "Sanatçı", "Tasarımcı", "Fotoğrafçı",
	},
	"Hizmet": {
		"Satış Temsilcisi", "Müşteri Hizmetleri", "Pazarlama Uzmanı",
		"İnsan Kaynakları", "İdari İşler", "Sekreter",
	},
	"Üretim": {
		"Mühendis", "Teknisyen", "Operatör", "Kalite Kontrol", "Üretim Sorumlusu",
	},
	"Diğer": {
		"Serbest Meslek", "Girişimci", "Memur", "İşçi", "Emekli", "Öğrenci",
	},
}
