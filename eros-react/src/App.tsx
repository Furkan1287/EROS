import React, { useState, useEffect } from "react";

const API_BASE = "http://localhost:8081/api";

function App() {
  const [mode, setMode] = useState<'login' | 'register'>("login");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [user, setUser] = useState<any>(null);
  const [terminalOutput, setTerminalOutput] = useState<string[]>([]);

  useEffect(() => {
    // Happy hacking ekranı
    const messages = [
      "EROS - Dating App v1.0",
      "Initializing system...",
      "Loading user interface...",
      "Connecting to backend...",
      "System ready!",
      "",
      "Type 'help' for commands or use the interface below."
    ];
    
    let index = 0;
    const interval = setInterval(() => {
      if (index < messages.length) {
        setTerminalOutput(prev => [...prev, messages[index]]);
        index++;
      } else {
        clearInterval(interval);
      }
    }, 500);

    return () => clearInterval(interval);
  }, []);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError("");
    setTerminalOutput(prev => [...prev, "> Attempting login..."]);
    
    try {
      const res = await fetch(`${API_BASE}/auth/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });
      const data = await res.json();
      if (data.success && data.user) {
        setUser(data.user);
        setTerminalOutput(prev => [...prev, "> Login successful!", `> Welcome ${data.user.name}!`]);
      } else {
        setError(data.error || "Giriş başarısız");
        setTerminalOutput(prev => [...prev, `> Error: ${data.error || "Login failed"}`]);
      }
    } catch (err) {
      setError("Sunucu hatası");
      setTerminalOutput(prev => [...prev, "> Server error"]);
    } finally {
      setLoading(false);
    }
  };

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError("");
    setTerminalOutput(prev => [...prev, "> Attempting registration..."]);
    
    try {
      const res = await fetch(`${API_BASE}/auth/simple-register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, email, password }),
      });
      const data = await res.json();
      if (data.success) {
        setTerminalOutput(prev => [...prev, "> Registration successful!", "> Auto-login in progress..."]);
        // Kayıt başarılı, otomatik giriş yap
        await handleLogin(e);
      } else {
        setError(data.error || "Kayıt başarısız");
        setTerminalOutput(prev => [...prev, `> Error: ${data.error || "Registration failed"}`]);
      }
    } catch (err) {
      setError("Sunucu hatası");
      setTerminalOutput(prev => [...prev, "> Server error"]);
    } finally {
      setLoading(false);
    }
  };

  if (user) {
    return (
      <div className="min-h-screen bg-black text-green-400 p-4 font-mono">
        <div className="max-w-4xl mx-auto">
          <div className="mb-4">
            <div className="text-green-500">$ eros --user {user.name}</div>
            <div className="text-green-400">Welcome to EROS Dashboard!</div>
          </div>
          <div className="bg-gray-900 p-4 rounded border border-green-500">
            <div className="text-yellow-400 mb-2">User Info:</div>
            <div>ID: {user.id}</div>
            <div>Name: {user.name}</div>
            <div>Email: {user.email}</div>
            <div>Created: {new Date(user.created_at).toLocaleString()}</div>
          </div>
          <div className="mt-4">
            <button 
              className="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700"
              onClick={() => {
                setUser(null);
                setTerminalOutput(prev => [...prev, "> Logged out"]);
              }}
            >
              Logout
            </button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-black text-green-400 p-4 font-mono">
      <div className="max-w-4xl mx-auto">
        {/* Terminal Header */}
        <div className="bg-gray-900 p-2 rounded-t border border-green-500">
          <div className="flex items-center space-x-2">
            <div className="w-3 h-3 bg-red-500 rounded-full"></div>
            <div className="w-3 h-3 bg-yellow-500 rounded-full"></div>
            <div className="w-3 h-3 bg-green-500 rounded-full"></div>
            <span className="ml-2 text-sm">eros-terminal</span>
          </div>
        </div>
        
        {/* Terminal Output */}
        <div className="bg-gray-900 p-4 border-l border-r border-green-500 min-h-96">
          {terminalOutput.map((line, index) => (
            <div key={index} className="mb-1">
              {line}
            </div>
          ))}
          {loading && <div className="text-yellow-400">{'>'} Processing...</div>}
        </div>

        {/* Terminal Input */}
        <div className="bg-gray-900 p-4 rounded-b border border-green-500">
          <div className="flex items-center">
            <span className="text-green-500 mr-2">$</span>
            <div className="flex-1">
              {error && <div className="text-red-400 mb-2">{error}</div>}
              
              <div className="flex space-x-4 mb-4">
                <button
                  className={`px-4 py-2 rounded ${mode === "login" ? "bg-green-600 text-white" : "bg-gray-700 text-green-400"}`}
                  onClick={() => setMode("login")}
                >
                  login
                </button>
                <button
                  className={`px-4 py-2 rounded ${mode === "register" ? "bg-green-600 text-white" : "bg-gray-700 text-green-400"}`}
                  onClick={() => setMode("register")}
                >
                  register
                </button>
              </div>

              {mode === "login" ? (
                <form onSubmit={handleLogin} className="space-y-2">
                  <div>
                    <span className="text-green-500">email:</span>
                    <input
                      type="email"
                      className="ml-2 bg-black border border-green-500 px-2 py-1 text-green-400 focus:outline-none"
                      value={email}
                      onChange={e => setEmail(e.target.value)}
                      placeholder="user@example.com"
                      required
                    />
                  </div>
                  <div>
                    <span className="text-green-500">password:</span>
                    <input
                      type="password"
                      className="ml-2 bg-black border border-green-500 px-2 py-1 text-green-400 focus:outline-none"
                      value={password}
                      onChange={e => setPassword(e.target.value)}
                      placeholder="********"
                      required
                    />
                  </div>
                  <button type="submit" className="bg-green-600 text-white px-4 py-1 rounded hover:bg-green-700">
                    execute login
                  </button>
                </form>
              ) : (
                <form onSubmit={handleRegister} className="space-y-2">
                  <div>
                    <span className="text-green-500">name:</span>
                    <input
                      type="text"
                      className="ml-2 bg-black border border-green-500 px-2 py-1 text-green-400 focus:outline-none"
                      value={name}
                      onChange={e => setName(e.target.value)}
                      placeholder="John Doe"
                      required
                    />
                  </div>
                  <div>
                    <span className="text-green-500">email:</span>
                    <input
                      type="email"
                      className="ml-2 bg-black border border-green-500 px-2 py-1 text-green-400 focus:outline-none"
                      value={email}
                      onChange={e => setEmail(e.target.value)}
                      placeholder="user@example.com"
                      required
                    />
                  </div>
                  <div>
                    <span className="text-green-500">password:</span>
                    <input
                      type="password"
                      className="ml-2 bg-black border border-green-500 px-2 py-1 text-green-400 focus:outline-none"
                      value={password}
                      onChange={e => setPassword(e.target.value)}
                      placeholder="********"
                      required
                    />
                  </div>
                  <button type="submit" className="bg-green-600 text-white px-4 py-1 rounded hover:bg-green-700">
                    execute register
                  </button>
                </form>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App; 