import React, { useState } from "react";
import "./passwordForm.css";

function PasswordForm({ setPasswordList }) {
  const [formData, setFormData] = useState({
    website: "",
    description: "",
    email: "",
    username: "",
    password: "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    setPasswordList((prevList) => [...prevList, formData]);
    // Reset the form
    setFormData({
      website: "",
      description: "",
      email: "",
      username: "",
      password: "",
    });
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Save new Password</h2>
      <div className="new-pwd-input-grp">
        <label htmlFor="website">Website</label>
        <input
          id="website"
          name="website"
          value={formData.website}
          onChange={handleChange}
          placeholder="Website"
          required
        />
      </div>
      <div className="new-pwd-input-grp">
        <label htmlFor="description">Description</label>
        <textarea
          id="description"
          name="description"
          value={formData.description}
          onChange={handleChange}
          placeholder="Description"
        />
      </div>
      <div className="new-pwd-input-grp">
        <label htmlFor="email">Email</label>
        <input
          id="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
          placeholder="Email"
        />
      </div>
      <div className="new-pwd-input-grp">
        <label htmlFor="username">Username</label>
        <input
          id="username"
          name="username"
          value={formData.username}
          onChange={handleChange}
          placeholder="Username"
        />
      </div>
      <div className="new-pwd-input-grp">
        <label htmlFor="password">Password</label>
        <input
          id="password"
          name="password"
          type="password"
          value={formData.password}
          onChange={handleChange}
          placeholder="Password"
          required
        />
      </div>
      <button id="add-pwd-btn" type="submit">
        Store Password
      </button>
    </form>
  );
}

export default PasswordForm;
