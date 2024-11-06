// components/Login/Login.js
import React, { useState } from "react";
import "./Login.css";
import { useApi } from "../ApiProvider/ApiProvider";

function Login({ onLogin }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const api_url = useApi();
  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`${api_url}/login/${username}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ master_password: password }),
      });

      if (response.ok) {
        const data = await response.json();
        onLogin(data.username); // Call onLogin with the username from the response
      } else {
        setError("* Invalid username or password");
      }
    } catch (error) {
      console.error("Error during login:", error);
      setError("* Network or server error");
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
        <button type="submit">Login</button>
        <div className="error-message">{error}</div>
      </form>
    </div>
  );
}

export default Login;
