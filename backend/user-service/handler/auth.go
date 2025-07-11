// auth.go - Kullanıcı kimlik doğrulama işlemleri
package handler

import (
	"encoding/json"
	"eros/user-service/model"
	"eros/user-service/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

// RegisterRequest - Kayıt isteği
type RegisterRequest struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	Bio             string   `json:"bio"`
	Age             int      `json:"age"`
	Seriousness     int      `json:"seriousness"`
	Height          int      `json:"height"`
	Weight          int      `json:"weight"`
	Smokes          bool     `json:"smokes"`
	Drinks          bool     `json:"drinks"`
	Job             string   `json:"job"`
	JobCategory     string   `json:"job_category"`
	Education       string   `json:"education"`
	Hobbies         []string `json:"hobbies"`
	HobbyCategories []string `json:"hobby_categories"`
}

// SimpleRegisterRequest - Basit kayıt isteği
type SimpleRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest - Giriş isteği
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register - Kullanıcı kaydı
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasyonlar
	if err := h.validateRegistration(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Şifre hash'leme
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Password hashing failed", http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Name:            req.Name,
		Email:           req.Email,
		Password:        string(hashedPassword),
		Bio:             req.Bio,
		Age:             req.Age,
		Seriousness:     req.Seriousness,
		Height:          req.Height,
		Weight:          req.Weight,
		Smokes:          req.Smokes,
		Drinks:          req.Drinks,
		Job:             req.Job,
		JobCategory:     req.JobCategory,
		Education:       req.Education,
		Hobbies:         req.Hobbies,
		HobbyCategories: req.HobbyCategories,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := h.userService.CreateUser(user); err != nil {
		log.Println("[CreateUser ERROR]", err)
		http.Error(w, "User creation failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}

// validateRegistration - Kayıt validasyonu
func (h *AuthHandler) validateRegistration(req RegisterRequest) error {
	// Temel alanlar
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return fmt.Errorf("name, email and password are required")
	}

	if len(req.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters")
	}

	// Yaş kontrolü
	if req.Age < 18 || req.Age > 100 {
		return fmt.Errorf("age must be between 18-100")
	}

	// Ciddiyet seviyesi kontrolü
	if req.Seriousness < 1 || req.Seriousness > 10 {
		return fmt.Errorf("seriousness level must be between 1-10")
	}

	// Boy kontrolü
	if req.Height < 140 || req.Height > 220 {
		return fmt.Errorf("height must be between 140-220 cm")
	}

	// Kilo kontrolü
	if req.Weight < 40 || req.Weight > 200 {
		return fmt.Errorf("weight must be between 40-200 kg")
	}

	// Hobi kontrolü
	if len(req.Hobbies) == 0 {
		return fmt.Errorf("at least one hobby must be selected")
	}

	if len(req.Hobbies) > 10 {
		return fmt.Errorf("maximum 10 hobbies can be selected")
	}

	// Hobi kategorileri kontrolü
	if len(req.HobbyCategories) == 0 {
		return fmt.Errorf("at least one hobby category must be selected")
	}

	return nil
}

// SimpleRegister - Basit kullanıcı kaydı (sadece temel bilgiler)
func (h *AuthHandler) SimpleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SimpleRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Invalid request body"})
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "name, email and password are required"})
		return
	}

	if len(req.Password) < 6 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "password must be at least 6 characters"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Password hashing failed"})
		return
	}

	user := &model.User{
		Name:            req.Name,
		Email:           req.Email,
		Password:        string(hashedPassword),
		Bio:             "",
		Age:             0,
		Seriousness:     5,
		Height:          0,
		Weight:          0,
		Smokes:          false,
		Drinks:          false,
		Job:             "",
		JobCategory:     "",
		Education:       "",
		Hobbies:         []string{},
		HobbyCategories: []string{},
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := h.userService.CreateUser(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "User creation failed"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"user_id": user.ID,
	})
}

// Login - Kullanıcı girişi
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Invalid request body"})
		return
	}

	user, err := h.userService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Invalid credentials"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"user":    user,
	})
}

// GetHobbyCategories - Hobi kategorilerini getir
func (h *AuthHandler) GetHobbyCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	hobbyCategories := map[string][]string{
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hobbyCategories)
}

// GetEducationLevels - Eğitim seviyelerini getir
func (h *AuthHandler) GetEducationLevels(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	educationLevels := []string{
		"İlkokul",
		"Ortaokul",
		"Lise",
		"Üniversite (Önlisans)",
		"Üniversite (Lisans)",
		"Yüksek Lisans",
		"Doktora",
		"Diğer",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(educationLevels)
}

// GetJobCategories - Meslek kategorilerini getir
func (h *AuthHandler) GetJobCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	jobCategories := map[string][]string{
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jobCategories)
}
