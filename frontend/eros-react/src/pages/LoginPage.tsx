import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import TinderCard from "react-tinder-card";

interface LoginCard {
  id: string;
  type: 'intro' | 'login' | 'register';
  title: string;
  content: any;
  color: string;
  icon: string;
}

const LoginPage: React.FC = () => {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);
  const [currentCardIndex, setCurrentCardIndex] = useState(0);
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    name: ''
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData(prev => ({
      ...prev,
      [e.target.name]: e.target.value
    }));
  };

  const handleLogin = async () => {
    setIsLoading(true);
    await new Promise(resolve => setTimeout(resolve, 1000));
    const userData = {
      name: formData.name || "Demo Kullanıcı",
      bio: "Merhaba! Ben yeni bir kullanıcıyım.",
      interests: ["Müzik", "Seyahat", "Spor"],
      email: formData.email
    };
    localStorage.setItem("user", JSON.stringify(userData));
    setIsLoading(false);
    navigate("/dashboard");
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (currentCardIndex === 2 && (!formData.name || !formData.email || !formData.password)) {
      alert("Lütfen tüm alanları doldurun!");
      return;
    }
    handleLogin();
  };

  const loginCards: LoginCard[] = [
    {
      id: 'intro',
      type: 'intro',
      title: 'EROS: Yeni Nesil Sosyal Deneyim',
      color: '#ff3576',
      icon: 'fas fa-bolt',
      content: {
        description: 'EROS, klasik arkadaşlık uygulamalarından farklı olarak tamamen kart tabanlı, sağa/sola kaydırmalı ve modern bir sosyal deneyim sunar. Her şey kartlarda! Üyelikle giriş yap, profilini oluştur, hobilerini ekle ve kartları kaydırarak yeni insanlarla tanış. EROS ile tanışmanın en yenilikçi yolu burada!'
      }
    },
    {
      id: 'login',
      type: 'login',
      title: 'Üye Girişi',
      color: '#00bfff',
      icon: 'fas fa-sign-in-alt',
      content: {
        isLogin: true,
        formData,
        handleInputChange,
        handleSubmit,
        isLoading,
        description: 'Zaten bir hesabın var mı? Hemen giriş yap ve kartlı sosyal deneyime katıl!'
      }
    },
    {
      id: 'register',
      type: 'register',
      title: 'Kayıt Ol',
      color: '#00ff88',
      icon: 'fas fa-user-plus',
      content: {
        isLogin: false,
        formData,
        handleInputChange,
        handleSubmit,
        isLoading,
        description: 'Hemen ücretsiz kaydol, profilini oluştur ve kartları kaydırarak yeni insanlarla tanış!'
      }
    }
  ];

  const currentCard = loginCards[currentCardIndex];

  const onSwipe = (direction: string) => {
    if (direction === 'left' && currentCardIndex < loginCards.length - 1) {
      setCurrentCardIndex(currentCardIndex + 1);
    } else if (direction === 'right' && currentCardIndex > 0) {
      setCurrentCardIndex(currentCardIndex - 1);
    }
  };

  const handleLike = () => onSwipe("right");
  const handleDislike = () => onSwipe("left");

  const renderCardContent = (card: LoginCard) => {
    switch (card.type) {
      case 'intro':
        return (
          <div className="welcome-card-content">
            <div className="welcome-logo">
              <i className={card.icon}></i>
              <span>EROS</span>
            </div>
            <h3 style={{color:'#ff3576', marginBottom: 16}}>{card.title}</h3>
            <p className="welcome-desc">{card.content.description}</p>
          </div>
        );
      case 'login':
      case 'register':
        return (
          <div className="form-card-content">
            <p style={{color:'#fff', marginBottom: 16}}>{card.content.description}</p>
            <form onSubmit={handleSubmit} className="login-form">
              {card.type === 'register' && (
                <div className="form-group">
                  <label>İsim</label>
                  <input
                    type="text"
                    name="name"
                    value={formData.name}
                    onChange={handleInputChange}
                    placeholder="İsminizi girin"
                    required
                  />
                </div>
              )}
              <div className="form-group">
                <label>E-posta</label>
                <input
                  type="email"
                  name="email"
                  value={formData.email}
                  onChange={handleInputChange}
                  placeholder="E-posta adresinizi girin"
                  required
                />
              </div>
              <div className="form-group">
                <label>Şifre</label>
                <input
                  type="password"
                  name="password"
                  value={formData.password}
                  onChange={handleInputChange}
                  placeholder="Şifrenizi girin"
                  required
                />
              </div>
              <button 
                type="submit" 
                className="submit-btn"
                disabled={isLoading}
              >
                {isLoading ? (
                  <div className="loading-spinner-small"></div>
                ) : (
                  <>
                    <i className={card.icon}></i>
                    {card.type === 'login' ? 'Giriş Yap' : 'Kayıt Ol'}
                  </>
                )}
              </button>
            </form>
          </div>
        );
      default:
        return <div>Kart içeriği yükleniyor...</div>;
    }
  };

  return (
    <div className="login-container">
      {/* Header */}
      <header className="login-header">
        <div className="header-content">
          <div className="logo">
            <i className="fas fa-heart"></i>
            <span>EROS</span>
          </div>
          <div className="card-indicator">
            {currentCardIndex + 1} / {loginCards.length}
          </div>
        </div>
      </header>
      {/* Main Card Area */}
      <main className="login-main">
        <div className="card-container">
          <TinderCard
            key={currentCard.id}
            onSwipe={onSwipe}
            preventSwipe={["down", "up"]}
          >
            <div 
              className="login-card"
              style={{ borderColor: currentCard.color }}
            >
              <div className="card-header" style={{ backgroundColor: currentCard.color }}>
                <i className={currentCard.icon}></i>
                <h2>{currentCard.title}</h2>
              </div>
              <div className="card-body">
                {renderCardContent(currentCard)}
              </div>
            </div>
          </TinderCard>
        </div>
        {/* Action Buttons */}
        <div className="card-actions">
          <button
            onClick={handleDislike}
            className="action-btn dislike"
            title="Önceki Kart"
            disabled={currentCardIndex === 0}
          >
            <i className="fas fa-arrow-left"></i>
          </button>
          <button
            onClick={handleLike}
            className="action-btn like"
            title="Sonraki Kart"
            disabled={currentCardIndex === loginCards.length - 1}
          >
            <i className="fas fa-arrow-right"></i>
          </button>
        </div>
        {/* Progress Indicator */}
        <div className="progress-indicator">
          <div className="progress-bar">
            <div 
              className="progress-fill" 
              style={{ width: `${((currentCardIndex + 1) / loginCards.length) * 100}%` }}
            ></div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default LoginPage; 