import React, { useState } from "react";
import "./password.css";
import show from "../../assets/show.png";
import hide from "../../assets/hide.png";
import update from "../../assets/update.png";
import save from "../../assets/save.png";
import delete_icon from "../../assets/delete.png";
function Password({
  showInfo,
  website,
  description,
  email,
  username,
  password,
  handleShow,
}) {
  const [showPassword, setShowPassword] = useState(false);
  const handleShowPassword = () => {
    setShowPassword(!showPassword);
  };
  return (
    <div className="password-item">
      {!showInfo ? (
        <div className="quick-pwd-viewer pwd" onClick={handleShow}>
          <p>{website}</p>
          <img
            src={show}
            className="show-pwd-btn pwd-btn"
            onClick={handleShow}
          />
        </div>
      ) : (
        <div className="long-pwd-viewer pwd ">
          <div className="creds-viewer">
            <div className="input-group">
              <label htmlFor="website">website</label>
              <input type="text" id="website" value={website} />
            </div>
            <div className="input-group">
              <label htmlFor="description">description</label>
              <textarea id="description" value={description} />
            </div>
            <div className="input-group">
              <label htmlFor="username">username</label>
              <input type="text" id="username" value={username} />
            </div>
            <div className="input-group">
              <label htmlFor="password" onClick={handleShowPassword}>
                password
              </label>
              <input
                type={showPassword ? "text" : "password"}
                id="password"
                value={password}
              />
            </div>
          </div>
          <div className="hide-btn-container">
            <img src={hide} className="hide-btn pwd-btn" onClick={handleShow} />
            <img src={update} className="hide-btn pwd-btn" alt="update" />
            <img src={delete_icon} className="hide-btn pwd-btn" alt="del" />
            <img src={save} className="hide-btn pwd-btn" alt="del" />
          </div>
        </div>
      )}
    </div>
  );
}

export default Password;
