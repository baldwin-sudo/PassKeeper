import React, { useState } from "react";
import "./passwordForm.css";
import { useApi } from "../ApiProvider/ApiProvider";

function PasswordForm({ setPasswordList }) {
  const [formData, setFormData] = useState({
    website: "",
    description: "",
    email: "",
    username: "",
    password: "",
  });
  const api_url = useApi();
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      // Correctly reference formData fields
      const { website, description, email, username, password } = formData;

      const response = await fetch(`${api_url}/passwords/create`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          website: website,
          description: description,
          email: email,
          username: username, // Correctly reference the state values
          plain_password: password, // Correctly reference the state values
        }),
      });

      const data = await response.json();
      console.table(data);
    } catch (error) {
      console.error("Error during login:", error);
      //TODO ADD ERROR IN FORM
      // setError("* Network or server error");
    }

    setPasswordList((prevList) => [...prevList, formData]);
    // Optionally reset the form after successful submission
    // setFormData({
    //   website: "",
    //   description: "",
    //   email: "",
    //   username: "",
    //   password: "",
    // });
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
