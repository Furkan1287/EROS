import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import TinderCard from "react-tinder-card";

// Onboarding (login/register) cards
const onboardingCards = [
  {
    id: "choose",
    type: "choose",
    title: "Üye misin?",
    description: "Devam etmek için lütfen bir seçim yap:",
    icon: "fas fa-question-circle",
    gradient: "linear-gradient(135deg, #18141a 0%, #2a1f2e 100%)"
  },
  {
    id: "login",
    type: "login",
    title: "Giriş Yap",
    icon: "fas fa-sign-in-alt",
    gradient: "linear-gradient(135deg, #18141a 0%, #00bfff 100%)"
  },
  {
    id: "register",
    type: "register",
    title: "Kayıt Ol",
    icon: "fas fa-user-plus",
    gradient: "linear-gradient(135deg, #18141a 0%, #ff3576 100%)"
  }
];

// Info/feature cards for main page
const infoCards = [
  {
    id: "ai",
    title: "Yapay Zeka Eşleştirme",
    description: "Gelişmiş yapay zeka algoritmaları ile sana en uygun eşleşmeleri sunuyoruz.",
    icon: "fas fa-robot",
    gradient: "linear-gradient(135deg, #8f5cff 0%, #00bfff 100%)"
  },
  {
    id: "blinddate",
    title: "Blind Date Sürprizi",
    description: "Blind Date özelliğiyle rastgele biriyle heyecan verici bir sohbete başla!",
    icon: "fas fa-user-secret",
    gradient: "linear-gradient(135deg, #ffb300 0%, #ff3576 100%)"
  },
  {
    id: "guvenli",
    title: "Gizlilik & Güvenlik",
    description: "Tüm verilerin güvende! EROS ile modern ve güvenli sosyal deneyim.",
    icon: "fas fa-shield-alt",
    gradient: "linear-gradient(135deg, #18141a 0%, #00bfff 100%)"
  },
  {
    id: "katil",
    title: "Hemen Katıl!",
    description: "Sen de EROS’a katıl, yeni nesil sosyal deneyimi keşfet!",
    icon: "fas fa-sign-in-alt",
    gradient: "linear-gradient(135deg, #00ff88 0%, #ff3576 100%)"
  }
];

const SwipePage: React.FC = () => {
  const navigate = useNavigate();
  // Onboarding modal state
  const [modalOpen, setModalOpen] = useState(false);
  const [onboardIndex, setOnboardIndex] = useState(0);
  const [formData, setFormData] = useState({ email: '', password: '', name: '' });
  const [isLoading, setIsLoading] = useState(false);
  // Info cards state
  const [infoIndex, setInfoIndex] = useState(0);

  // Open modal if not logged in
  useEffect(() => {
    const user = localStorage.getItem("user");
    if (!user) setModalOpen(true);
  }, []);
  // Listen for custom event from Navbar user icon
  useEffect(() => {
    const openModal = () => setModalOpen(true);
    window.addEventListener("openOnboardingModal", openModal);
    return () => window.removeEventListener("openOnboardingModal", openModal);
  }, []);

  // Onboarding modal swipe logic
  const onOnboardSwipe = (direction: string) => {
    if (onboardIndex === 0) {
      if (direction === "right") setOnboardIndex(2); // Register
      else if (direction === "left") setOnboardIndex(1); // Login
    } else if (direction === "left" && onboardingCards.length - 1) {
      setOnboardIndex(onboardIndex + 1);
    } else if (direction === "right" && onboardIndex > 0) {
      setOnboardIndex(onboardIndex - 1);
    }
  };
  const goToLogin = () => setOnboardIndex(1);
  const goToRegister = () => setOnboardIndex(2);

  // Info cards swipe logic
  const onInfoSwipe = (direction: string) => {
    if (direction === "left" && infoIndex < infoCards.length - 1) {
      setInfoIndex(infoIndex + 1);
    } else if (direction === "right" && infoIndex > 0) {
      setInfoIndex(infoIndex - 1);
    }
  };

  // Auth form logic
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData(prev => ({ ...prev, [e.target.name]: e.target.value }));
  };
  const handleLogin = async () => {
    setIsLoading(true);
    await new Promise(resolve => setTimeout(resolve, 1000));
    const userData = { name: formData.name || "Demo Kullanıcı", email: formData.email };
    localStorage.setItem("user", JSON.stringify(userData));
    setIsLoading(false);
    setModalOpen(false);
    navigate("/dashboard");
  };
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (onboardIndex === 2 && (!formData.name || !formData.email || !formData.password)) {
      alert("Lütfen tüm alanları doldurun!");
      return;
    }
    if (onboardIndex === 1 && (!formData.email || !formData.password)) {
      alert("Lütfen e-posta ve şifreyi girin!");
      return;
    }
    handleLogin();
  };

  // Render onboarding modal card content
  const renderOnboardCard = (card: typeof onboardingCards[0]) => {
    switch (card.type) {
      case 'choose':
        return (
          <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', gap: 24 }}>
            <p style={{ color: '#fff', fontSize: 18, marginBottom: 24 }}>{card.description}</p>
            <div style={{ display: 'flex', gap: 24 }}>
              <button className="final-btn primary" style={{ fontSize: 18, padding: '0.7rem 2.2rem', borderRadius: 12, background: 'linear-gradient(90deg,#ff3576,#ff6b9d)' }} onClick={goToRegister}>
                <i className="fas fa-user-plus"></i> Üye Değilim
              </button>
              <button className="final-btn primary" style={{ fontSize: 18, padding: '0.7rem 2.2rem', borderRadius: 12, background: 'linear-gradient(90deg,#00bfff,#00ff88)' }} onClick={goToLogin}>
                <i className="fas fa-sign-in-alt"></i> Üyeyim
              </button>
            </div>
          </div>
        );
      case 'login':
        return (
          <div className="form-card-content">
            <form onSubmit={handleSubmit} className="login-form">
              <div className="form-group">
                <label>E-posta</label>
                <input type="email" name="email" value={formData.email} onChange={handleInputChange} placeholder="E-posta adresinizi girin" required />
              </div>
              <div className="form-group">
                <label>Şifre</label>
                <input type="password" name="password" value={formData.password} onChange={handleInputChange} placeholder="Şifrenizi girin" required />
              </div>
              <button type="submit" className="submit-btn" disabled={isLoading}>
                {isLoading ? <div className="loading-spinner-small"></div> : (<><i className="fas fa-sign-in-alt"></i> Giriş Yap</>)}
              </button>
            </form>
          </div>
        );
      case 'register':
        return (
          <div className="form-card-content">
            <form onSubmit={handleSubmit} className="login-form">
              <div className="form-group">
                <label>İsim</label>
                <input type="text" name="name" value={formData.name} onChange={handleInputChange} placeholder="İsminizi girin" required />
              </div>
              <div className="form-group">
                <label>E-posta</label>
                <input type="email" name="email" value={formData.email} onChange={handleInputChange} placeholder="E-posta adresinizi girin" required />
              </div>
              <div className="form-group">
                <label>Şifre</label>
                <input type="password" name="password" value={formData.password} onChange={handleInputChange} placeholder="Şifrenizi girin" required />
              </div>
              <button type="submit" className="submit-btn" disabled={isLoading}>
                {isLoading ? <div className="loading-spinner-small"></div> : (<><i className="fas fa-user-plus"></i> Kayıt Ol</>)}
              </button>
            </form>
          </div>
        );
      default:
        return <div>Kart içeriği yükleniyor...</div>;
    }
  };

  // Render info card
  const renderInfoCard = (card: typeof infoCards[0]) => (
    <div className="info-card"
      style={{
        background: card.gradient,
        borderRadius: 24,
        boxShadow: "0 8px 32px rgba(0,0,0,0.25)",
        padding: "2.5rem 2rem 2rem 2rem",
        minWidth: 320,
        maxWidth: 400,
        minHeight: 340,
        color: "#fff",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        position: "relative"
      }}
    >
      <div style={{ fontSize: 40, marginBottom: 18 }}>
        <i className={card.icon}></i>
      </div>
      <h2 style={{ fontWeight: 700, fontSize: 26, marginBottom: 12, textAlign: "center" }}>{card.title}</h2>
      <p style={{ fontSize: 17, textAlign: "center", marginBottom: 18 }}>{card.description}</p>
      <div style={{ position: "absolute", bottom: 12, right: 18, fontSize: 13, opacity: 0.7 }}>
        Kart {infoIndex + 1} / {infoCards.length}
      </div>
    </div>
  );

  // Main render
  return (
    <div className="swipe-page" style={{ minHeight: "100vh", display: "flex", flexDirection: "column", justifyContent: "center", alignItems: "center" }}>
      {/* Onboarding Modal */}
      {modalOpen && (
        <div className="onboarding-modal-overlay">
          <div className="onboarding-modal-card">
            <TinderCard
              key={onboardingCards[onboardIndex].id}
              onSwipe={onOnboardSwipe}
              preventSwipe={["down", "up"]}
            >
              <div
                className="info-card onboarding-modal-content"
                style={{
                  background: onboardingCards[onboardIndex].gradient,
                  borderRadius: 32,
                  boxShadow: "0 16px 48px rgba(0,0,0,0.55)",
                  padding: "3.5rem 2.5rem 2.5rem 2.5rem",
                  minWidth: 380,
                  maxWidth: 480,
                  minHeight: 400,
                  color: "#fff",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                  justifyContent: "center",
                  position: "relative",
                  zIndex: 20
                }}
              >
                <div style={{ fontSize: 48, marginBottom: 18 }}>
                  <i className={onboardingCards[onboardIndex].icon}></i>
                </div>
                <h2 style={{ fontWeight: 700, fontSize: 28, marginBottom: 12, textAlign: "center" }}>{onboardingCards[onboardIndex].title}</h2>
                {onboardingCards[onboardIndex].description && <p style={{ fontSize: 18, textAlign: "center", marginBottom: 18 }}>{onboardingCards[onboardIndex].description}</p>}
                {renderOnboardCard(onboardingCards[onboardIndex])}
                <div style={{ position: "absolute", bottom: 12, right: 18, fontSize: 13, opacity: 0.7 }}>
                  Kart {onboardIndex + 1} / {onboardingCards.length}
                </div>
              </div>
            </TinderCard>
          </div>
          <div className="onboarding-modal-bg" onClick={() => setModalOpen(false)}></div>
        </div>
      )}
      {/* Info/feature cards (main page) */}
      {!modalOpen && (
        <main className="swipe-main" style={{ flex: 1, display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "center" }}>
          <div className="card-container">
            <TinderCard
              key={infoCards[infoIndex].id}
              onSwipe={onInfoSwipe}
              preventSwipe={["down", "up"]}
            >
              {renderInfoCard(infoCards[infoIndex])}
            </TinderCard>
          </div>
          {/* Action Buttons */}
          <div className="card-actions">
            <button
              onClick={() => onInfoSwipe("right")}
              className="action-btn dislike"
              title="Önceki Kart"
              disabled={infoIndex === 0}
            >
              <i className="fas fa-arrow-left"></i>
            </button>
            <button
              onClick={() => onInfoSwipe("left")}
              className="action-btn like"
              title="Sonraki Kart"
              disabled={infoIndex === infoCards.length - 1}
            >
              <i className="fas fa-arrow-right"></i>
            </button>
          </div>
          {/* Progress Indicator */}
          <div className="progress-indicator">
            <div className="progress-bar">
              <div
                className="progress-fill"
                style={{ width: `${((infoIndex + 1) / infoCards.length) * 100}%` }}
              ></div>
            </div>
          </div>
        </main>
      )}
    </div>
  );
};

export default SwipePage; 