# EROS - Modern Go Microservices Dating App

## Özellikler
- Go ile yazılmış mikroservis mimarisi (user, match, chat, api-gateway)
- SQLite veritabanı
- Nuxt 3 + Tailwind CSS frontend
- AI destekli eşleşme, analiz ve güvenlik (OpenRouter API)
- JWT ile authentication (hazır altyapı)
- Modern, mobil uyumlu arayüz

---

## Kurulum

### 1. Depoyu Klonla
```sh
git clone https://github.com/kullaniciadi/EROS.git
cd EROS
```

### 2. Ortam Değişkenleri (.env)
Kök dizinde `.env` dosyası oluştur:
```env
# Database
DB_PATH=./eros.db

# OpenRouter API (anahtarını buraya yaz)
OPENROUTER_API_KEY=sk-or-***************

# Servis Portları
USER_SERVICE_PORT=8081
MATCH_SERVICE_PORT=8082
CHAT_SERVICE_PORT=8083

# JWT Secret
JWT_SECRET=your_jwt_secret_here

# Log Level
LOG_LEVEL=info
```
> **Dikkat:** OpenRouter API anahtarını asla Github'a yükleme!  
> `.env` dosyanı `.gitignore`'da tut.

---

### 3. Bağımlılıkları Kur

#### Backend (Go)
```sh
cd backend/user-service && go mod tidy
cd ../match-service && go mod tidy
cd ../chat-service && go mod tidy
cd ../../..
```

#### Frontend (Nuxt)
```sh
cd frontend
npm install
cd ..
```

---

### 4. Tüm Servisleri Başlat

Kök dizinde:
```sh
chmod +x start-backend.sh
./start-backend.sh
```
Her servis kendi log dosyasına (`user-service.log` vs.) log yazar.

---

### 5. Frontend'i Başlat

```sh
cd frontend
npm run dev
```
Uygulama: [http://localhost:3000](http://localhost:3000)

---

## Kullanım

- **Kayıt Ol:** Sadece ad, e-posta ve şifre ile hızlı kayıt.
- **Profil Tamamlama:** Giriş sonrası profilini detaylı doldur.
- **Eşleşme & Chat:** Swipe ve chat özellikleri.
- **AI Destekli:** Sohbet analizi, güvenlik filtresi, eşleşme önerileri.

---

## Geliştirici Notları

- Her mikroservisin kendi `main.go` dosyası vardır.
- Ortak tipler ve yardımcılar `backend/shared` dizinindedir.
- OpenRouter API anahtarı sadece backend'de kullanılır, frontend'e asla koyma!
- Veritabanı şeması otomatik oluşur, ilk çalıştırmada `eros.db` dosyası oluşur.

---

## Katkı ve Lisans

- PR ve issue açabilirsin.
- MIT Lisansı.

---

**Ekstra:**
- Hata ayıklama için backend loglarını (`user-service.log`) kontrol et.
- .env dosyanı asla Github'a yükleme! 