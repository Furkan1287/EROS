import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import TinderCard from "react-tinder-card";

interface User {
  name: string;
  age?: number;
  bio?: string;
  photos?: string[];
  interests?: string[];
  preferences?: {
    ageRange: [number, number];
    distance: number;
    gender: string;
  };
}

interface DashboardCard {
  id: string;
  type: 'profile' | 'stats' | 'matches' | 'activity' | 'settings' | 'navigation';
  title: string;
  content: any;
  color: string;
  icon: string;
}

const Dashboard: React.FC = () => {
  const navigate = useNavigate();
  const [user, setUser] = useState<User | null>(null);
  const [currentCardIndex, setCurrentCardIndex] = useState(0);
  const [isEditing, setIsEditing] = useState(false);
  const [editForm, setEditForm] = useState({
    name: '',
    age: '',
    bio: '',
    interests: [] as string[],
    newInterest: ''
  });

  useEffect(() => {
    const userData = localStorage.getItem("user");
    if (userData) {
      const parsedUser = JSON.parse(userData);
      setUser(parsedUser);
      setEditForm({
        name: parsedUser.name || '',
        age: parsedUser.age?.toString() || '',
        bio: parsedUser.bio || '',
        interests: parsedUser.interests || [],
        newInterest: ''
      });
    }
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("user");
    navigate("/");
  };

  const handleSaveProfile = () => {
    const updatedUser = {
      ...user,
      name: editForm.name,
      age: parseInt(editForm.age) || 0,
      bio: editForm.bio,
      interests: editForm.interests
    };
    setUser(updatedUser);
    localStorage.setItem("user", JSON.stringify(updatedUser));
    setIsEditing(false);
  };

  const addInterest = () => {
    if (editForm.newInterest.trim() && !editForm.interests.includes(editForm.newInterest.trim())) {
      setEditForm(prev => ({
        ...prev,
        interests: [...prev.interests, editForm.newInterest.trim()],
        newInterest: ''
      }));
    }
  };

  const removeInterest = (interest: string) => {
    setEditForm(prev => ({
      ...prev,
      interests: prev.interests.filter(i => i !== interest)
    }));
  };

  const mockStats = {
    matches: 12,
    likes: 45,
    views: 128,
    conversations: 8
  };

  const mockRecentActivity = [
    { type: 'match', message: 'Yeni eşleşme: Ayşe ile tanıştın', time: '2 saat önce' },
    { type: 'like', message: 'Seni beğendi: Mehmet', time: '5 saat önce' },
    { type: 'message', message: 'Yeni mesaj: Zeynep', time: '1 gün önce' },
    { type: 'view', message: 'Profilini görüntüledi: Ali', time: '2 gün önce' }
  ];

  const mockMatches = [
    { id: 1, name: 'Ayşe', age: 24, image: 'https://images.unsplash.com/photo-1494790108755-2616b612b786?w=100&h=100&fit=crop', lastSeen: '2 saat önce' },
    { id: 2, name: 'Zeynep', age: 26, image: 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=100&h=100&fit=crop', lastSeen: '5 saat önce' },
    { id: 3, name: 'Elif', age: 23, image: 'https://images.unsplash.com/photo-1544005313-94ddf0286df2?w=100&h=100&fit=crop', lastSeen: '1 gün önce' }
  ];

  const dashboardCards: DashboardCard[] = [
    {
      id: 'profile',
      type: 'profile',
      title: 'Profil',
      color: '#ff3576',
      icon: 'fas fa-user',
      content: {
        user,
        isEditing,
        editForm,
        setIsEditing,
        handleSaveProfile,
        addInterest,
        removeInterest
      }
    },
    {
      id: 'stats',
      type: 'stats',
      title: 'İstatistikler',
      color: '#ffd700',
      icon: 'fas fa-chart-bar',
      content: { stats: mockStats }
    },
    {
      id: 'matches',
      type: 'matches',
      title: 'Eşleşmeler',
      color: '#00ff88',
      icon: 'fas fa-heart',
      content: { matches: mockMatches }
    },
    {
      id: 'activity',
      type: 'activity',
      title: 'Aktivite',
      color: '#00bfff',
      icon: 'fas fa-bell',
      content: { activities: mockRecentActivity }
    },
    {
      id: 'settings',
      type: 'settings',
      title: 'Ayarlar',
      color: '#ff6b35',
      icon: 'fas fa-cog',
      content: {}
    },
    {
      id: 'navigation',
      type: 'navigation',
      title: 'Navigasyon',
      color: '#9c27b0',
      icon: 'fas fa-compass',
      content: { handleLogout }
    }
  ];

  if (!user) {
    return (
      <div className="loading-container">
        <div className="loading-spinner"></div>
        <p>Yükleniyor...</p>
      </div>
    );
  }

  const currentCard = dashboardCards[currentCardIndex];

  const onSwipe = (direction: string) => {
    if (direction === 'left' && currentCardIndex < dashboardCards.length - 1) {
      setCurrentCardIndex(currentCardIndex + 1);
    } else if (direction === 'right' && currentCardIndex > 0) {
      setCurrentCardIndex(currentCardIndex - 1);
    }
  };

  const handleLike = () => onSwipe("right");
  const handleDislike = () => onSwipe("left");

  const renderCardContent = (card: DashboardCard) => {
    switch (card.type) {
      case 'profile':
        return (
          <div className="profile-card-content">
            <div className="profile-header">
              <div className="profile-avatar">
                <i className="fas fa-user-circle"></i>
              </div>
              <div className="profile-info">
                <h2>{user?.name}</h2>
                <p className="profile-age">{user?.age || 'Yaş belirtilmemiş'}</p>
                <p className="profile-bio">{user?.bio || 'Henüz bio eklenmemiş'}</p>
              </div>
              <button 
                onClick={() => setIsEditing(!isEditing)}
                className="edit-btn"
              >
                <i className="fas fa-edit"></i>
                {isEditing ? 'İptal' : 'Düzenle'}
              </button>
            </div>

            {isEditing ? (
              <div className="edit-form">
                <div className="form-group">
                  <label>İsim</label>
                  <input
                    type="text"
                    value={editForm.name}
                    onChange={(e) => setEditForm(prev => ({ ...prev, name: e.target.value }))}
                    placeholder="İsminizi girin"
                  />
                </div>
                <div className="form-group">
                  <label>Yaş</label>
                  <input
                    type="number"
                    value={editForm.age}
                    onChange={(e) => setEditForm(prev => ({ ...prev, age: e.target.value }))}
                    placeholder="Yaşınızı girin"
                  />
                </div>
                <div className="form-group">
                  <label>Bio</label>
                  <textarea
                    value={editForm.bio}
                    onChange={(e) => setEditForm(prev => ({ ...prev, bio: e.target.value }))}
                    placeholder="Kendinizden bahsedin..."
                    rows={3}
                  />
                </div>
                <div className="form-group">
                  <label>İlgi Alanları</label>
                  <div className="interests-input">
                    <input
                      type="text"
                      value={editForm.newInterest}
                      onChange={(e) => setEditForm(prev => ({ ...prev, newInterest: e.target.value }))}
                      placeholder="İlgi alanı ekle"
                      onKeyPress={(e) => e.key === 'Enter' && addInterest()}
                    />
                    <button onClick={addInterest} className="add-interest-btn">
                      <i className="fas fa-plus"></i>
                    </button>
                  </div>
                  <div className="interests-tags">
                    {editForm.interests.map((interest, index) => (
                      <span key={index} className="interest-tag">
                        {interest}
                        <button onClick={() => removeInterest(interest)} className="remove-interest">
                          <i className="fas fa-times"></i>
                        </button>
                      </span>
                    ))}
                  </div>
                </div>
                <div className="form-actions">
                  <button onClick={handleSaveProfile} className="save-btn">
                    <i className="fas fa-save"></i>
                    Kaydet
                  </button>
                </div>
              </div>
            ) : (
              <div className="profile-details">
                <div className="detail-section">
                  <h3>İlgi Alanları</h3>
                  <div className="interests-display">
                    {user.interests && user.interests.length > 0 ? (
                      user.interests.map((interest, index) => (
                        <span key={index} className="interest-badge">{interest}</span>
                      ))
                    ) : (
                      <p className="no-data">Henüz ilgi alanı eklenmemiş</p>
                    )}
                  </div>
                </div>
              </div>
            )}
          </div>
        );

      case 'stats':
        return (
          <div className="stats-card-content">
            <div className="stats-grid">
              <div className="stat-card">
                <i className="fas fa-heart"></i>
                <div className="stat-info">
                  <h3>{card.content.stats.matches}</h3>
                  <p>Eşleşme</p>
                </div>
              </div>
              <div className="stat-card">
                <i className="fas fa-thumbs-up"></i>
                <div className="stat-info">
                  <h3>{card.content.stats.likes}</h3>
                  <p>Beğeni</p>
                </div>
              </div>
              <div className="stat-card">
                <i className="fas fa-eye"></i>
                <div className="stat-info">
                  <h3>{card.content.stats.views}</h3>
                  <p>Görüntülenme</p>
                </div>
              </div>
              <div className="stat-card">
                <i className="fas fa-comments"></i>
                <div className="stat-info">
                  <h3>{card.content.stats.conversations}</h3>
                  <p>Sohbet</p>
                </div>
              </div>
            </div>
          </div>
        );

      case 'matches':
        return (
          <div className="matches-card-content">
            <div className="matches-list">
              {card.content.matches.map((match: any) => (
                <div key={match.id} className="match-card">
                  <div className="match-avatar">
                    <img src={match.image} alt={match.name} />
                  </div>
                  <div className="match-info">
                    <h4>{match.name}, {match.age}</h4>
                    <p>Son görülme: {match.lastSeen}</p>
                  </div>
                  <button className="message-btn">
                    <i className="fas fa-comment"></i>
                  </button>
                </div>
              ))}
            </div>
          </div>
        );

      case 'activity':
        return (
          <div className="activity-card-content">
            <div className="activity-list">
              {card.content.activities.map((activity: any, index: number) => (
                <div key={index} className="activity-item">
                  <div className={`activity-icon ${activity.type}`}>
                    <i className={`fas fa-${activity.type === 'match' ? 'heart' : activity.type === 'like' ? 'thumbs-up' : activity.type === 'message' ? 'comment' : 'eye'}`}></i>
                  </div>
                  <div className="activity-content">
                    <p>{activity.message}</p>
                    <span className="activity-time">{activity.time}</span>
                  </div>
                </div>
              ))}
            </div>
          </div>
        );

      case 'settings':
        return (
          <div className="settings-card-content">
            <div className="settings-list">
              <div className="setting-item">
                <div className="setting-info">
                  <h4>Bildirimler</h4>
                  <p>Yeni eşleşme ve mesaj bildirimleri</p>
                </div>
                <label className="toggle">
                  <input type="checkbox" defaultChecked />
                  <span className="slider"></span>
                </label>
              </div>
              <div className="setting-item">
                <div className="setting-info">
                  <h4>Konum Paylaşımı</h4>
                  <p>Yakındaki kullanıcıları görmek için</p>
                </div>
                <label className="toggle">
                  <input type="checkbox" />
                  <span className="slider"></span>
                </label>
              </div>
              <div className="setting-item">
                <div className="setting-info">
                  <h4>Gizlilik Modu</h4>
                  <p>Profilinizi gizli tutun</p>
                </div>
                <label className="toggle">
                  <input type="checkbox" />
                  <span className="slider"></span>
                </label>
              </div>
            </div>
            <button className="danger-btn">
              <i className="fas fa-trash"></i>
              Hesabı Sil
            </button>
          </div>
        );

      case 'navigation':
        return (
          <div className="navigation-card-content">
            <div className="nav-options">
              <button className="nav-option">
                <i className="fas fa-fire"></i>
                <span>Keşfet</span>
              </button>
              <button className="nav-option">
                <i className="fas fa-heart"></i>
                <span>Eşleşmeler</span>
              </button>
              <button className="nav-option">
                <i className="fas fa-comments"></i>
                <span>Mesajlar</span>
              </button>
              <button className="nav-option">
                <i className="fas fa-user"></i>
                <span>Profil</span>
              </button>
            </div>
            <button onClick={card.content.handleLogout} className="logout-btn">
              <i className="fas fa-sign-out-alt"></i>
              Çıkış Yap
            </button>
          </div>
        );

      default:
        return <div>Kart içeriği yükleniyor...</div>;
    }
  };

  return (
    <div className="dashboard-container">
      {/* Header */}
      <header className="dashboard-header">
        <div className="header-content">
          <h1>Hoş geldin, {user.name}!</h1>
          <div className="card-indicator">
            {currentCardIndex + 1} / {dashboardCards.length}
          </div>
        </div>
      </header>

      {/* Main Card Area */}
      <main className="dashboard-main">
        <div className="card-container">
          <TinderCard
            key={currentCard.id}
            onSwipe={onSwipe}
            preventSwipe={["down", "up"]}
          >
            <div 
              className="dashboard-card"
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
            disabled={currentCardIndex === dashboardCards.length - 1}
          >
            <i className="fas fa-arrow-right"></i>
          </button>
        </div>

        {/* Progress Indicator */}
        <div className="progress-indicator">
          <div className="progress-bar">
            <div 
              className="progress-fill" 
              style={{ width: `${((currentCardIndex + 1) / dashboardCards.length) * 100}%` }}
            ></div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default Dashboard; 