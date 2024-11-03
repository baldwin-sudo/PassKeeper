import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

function Signup({ updateLogginStatus }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate(); // Initialize navigate

  const handleSignup = (e) => {
    e.preventDefault();
    // Add your signup logic here
    console.log("User signed up:", { username, password });

    // Update login status to indicate the user is logged in
    updateLogginStatus(false);

    // Use navigate to go to /password-list without refreshing
    navigate("/password-list");
  };

  return (
    <div>
      <h2>Signup</h2>
      <form onSubmit={handleSignup}>
        <label>
          Username:
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </label>
        <br />
        <label>
          Password:
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </label>
        <br />
        <button type="submit">Sign Up</button>
      </form>
    </div>
  );
}

export default Signup;
