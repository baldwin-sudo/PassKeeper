import React from "react";
import "./footer.css";

function Footer() {
  return (
    <div className="footer">
      <p>
        by{" "}
        <a
          href="https://github.com/baldwin-sudo"
          target="_blank"
          rel="noopener noreferrer"
          className="github-link"
        >
          baldwin-sudo
        </a>
      </p>
    </div>
  );
}

export default Footer;
