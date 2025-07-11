import React from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import SwipePage from "./pages/SwipePage";
import FooterCard from "./components/FooterCard";
import Navbar from "./components/Navbar";

function RequireAuth({ children }: React.PropsWithChildren) {
  const user = localStorage.getItem("user");
  return user ? <>{children}</> : <Navigate to="/" replace />;
}

function App() {
  return (
    <>
      <Router>
        <Navbar />
        <Routes>
          <Route path="/" element={<SwipePage />} />
          <Route
            path="/dashboard"
            element={
              <RequireAuth>
                <Dashboard />
              </RequireAuth>
            }
          />
        </Routes>
      </Router>
      <FooterCard />
    </>
  );
}

export default App;
