// components/Login/Login.js
import React, { useState } from "react";
import "./Login.css";
function Login({ onLogin }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const handleLogin = (e) => {
    e.preventDefault();
    // Simple example login validation (replace with real authentication in production)
    if (username === "user" && password === "password") {
      onLogin();
    } else {
      setError("* error");
    }
  };

  return (
    <div className="login-container">
      <h2>Login</h2>
      <form onSubmit={handleLogin}>
        <input
          type="text"
          value={username}
          placeholder="username ..."
          onChange={(e) => setUsername(e.target.value)}
        />
        <br />
        <input
          type="password"
          value={password}
          placeholder="password ..."
          onChange={(e) => setPassword(e.target.value)}
        />
        <br />
        <button type="submit">Login</button>{" "}
        <div className="error-message">{error}</div>
      </form>
    </div>
  );
}

export default Login;
