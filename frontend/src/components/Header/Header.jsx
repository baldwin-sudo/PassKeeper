import React from "react";
import "./header.css";
import { Link, useLocation, useNavigate } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faList,
  faPlus,
  faFileAlt,
  faSignOutAlt,
} from "@fortawesome/free-solid-svg-icons";

function Header({ isLoggedIn, logoutUser }) {
  const location = useLocation(); // Get the current location
  const navigate = useNavigate();

  const handleLogout = () => {
    logoutUser();
    navigate("/login");
  };

  return (
    <div className="header-container">
      <h1>
        <strong>
          Pass <span className="keeper"> Keeper .</span>
        </strong>
      </h1>
      {isLoggedIn ? (
        <nav>
          <Link
            className={`link ${location.pathname === "/" ? "active" : ""}`}
            to="/"
          >
            <FontAwesomeIcon icon={faList} className="nav-icon" />
            <span className="nav-text">Password List</span>
          </Link>
          <Link
            className={`link ${location.pathname === "/add" ? "active" : ""}`}
            to="/add"
          >
            <FontAwesomeIcon icon={faPlus} className="nav-icon" />
            <span className="nav-text">Add Password</span>
          </Link>
          <Link
            className={`link ${
              location.pathname === "/abstract" ? "active" : ""
            }`}
            to="/abstract"
          >
            <FontAwesomeIcon icon={faFileAlt} className="nav-icon" />
            <span className="nav-text">Abstract</span>
          </Link>
          <Link className="link" onClick={handleLogout}>
            <FontAwesomeIcon icon={faSignOutAlt} className="nav-icon" />
            <span className="nav-text">Logout</span>
          </Link>
        </nav>
      ) : null}
    </div>
  );
}

export default Header;
