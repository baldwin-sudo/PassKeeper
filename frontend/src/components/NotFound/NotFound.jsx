// components/NotFound/NotFound.js
import React from "react";
import "./NotFound.css"; // Optional styling
import { Link } from "react-router-dom";
function NotFound({ isLoggedIn }) {
  return (
    <div className="not-found-container">
      <h2>404 - Page Not Found</h2>
      <p>The page you are looking for does not exist.</p>
      <Link
        to="/"
        onClick={() => {
          console.log(isLoggedIn);
        }}
      >
        {isLoggedIn ? "click to go home ." : "click to login ."}
      </Link>
    </div>
  );
}

export default NotFound;
