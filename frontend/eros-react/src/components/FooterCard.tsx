import React from "react";

const FooterCard: React.FC = () => {
  return (
    <footer className="footer-card footer-shake">
      <div className="footer-content">
        <div className="footer-logo-section">
          <div className="footer-logo">
            <i className="fas fa-heart"></i>
            <span>EROS</span>
          </div>
          <p className="footer-desc">
            EROS, yapay zeka destekli Blind Date, kişilik analizi ve ilişki beklentisi eşleştirmesiyle yeni nesil sosyal deneyim sunar. Güvenli, yenilikçi ve modern!
          </p>
        </div>
        <div className="footer-links-section">
          <a href="#" className="footer-link">Gizlilik</a>
          <a href="#" className="footer-link">Kullanım Şartları</a>
          <a href="#" className="footer-link">İletişim</a>
          <a href="#" className="footer-link">Çerez Ayarları</a>
        </div>
        <div className="footer-social-section">
          <a href="#" className="footer-social"><i className="fab fa-instagram"></i></a>
          <a href="#" className="footer-social"><i className="fab fa-twitter"></i></a>
          <a href="#" className="footer-social"><i className="fab fa-youtube"></i></a>
        </div>
        <div className="footer-apps-section">
          <button className="footer-app-btn"><i className="fab fa-apple"></i> App Store</button>
          <button className="footer-app-btn"><i className="fab fa-google-play"></i> Google Play</button>
        </div>
      </div>
      <div className="footer-bottom">
        <span>© {new Date().getFullYear()} EROS. Tüm hakları saklıdır.</span>
      </div>
    </footer>
  );
};

export default FooterCard; 