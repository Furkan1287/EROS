import React from "react";
import { useNavigate } from "react-router-dom";

const Navbar: React.FC = () => {
  const navigate = useNavigate();
  const user = localStorage.getItem("user");
  const handleUserIconClick = () => {
    if (!user) {
      window.dispatchEvent(new Event("openOnboardingModal"));
    } else {
      navigate("/dashboard");
    }
  };
  return (
    <nav className="eros-navbar">
      <div className="navbar-content">
        <div className="navbar-left">
          <span className="navbar-logo" style={{cursor: 'pointer'}} onClick={() => navigate("/")}>
            <i className="fas fa-heart"></i> EROS
          </span>
        </div>
        <div className="navbar-right">
          <button className="navbar-login-btn" onClick={handleUserIconClick} title="Giriş Yap / Kayıt Ol">
            <i className="fas fa-user-circle"></i>
          </button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar; 