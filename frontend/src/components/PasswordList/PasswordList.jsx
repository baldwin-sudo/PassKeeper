import React, { useState, useEffect } from "react";
import Password from "../Password/Password";
import "./passwordList.css";
import { useApi } from "../ApiProvider/ApiProvider";

function PasswordList() {
  const [expandedIndex, setExpandedIndex] = useState(-1);
  const [passwords, setPasswords] = useState([]);
  const api_url = useApi();
  const handleShow = (index) => {
    setExpandedIndex(expandedIndex === index ? null : index);
  };

  const fetchPasswords = async () => {
    try {
      const response = await fetch(`${api_url}/passwords`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      });

      const data = await response.json();
      console.log(data);
      setPasswords(data);
    } catch (error) {
      console.error("Error during login:", error);
      //TODO: Add error handling here
    }
  };

  // Adding useEffect to fetch passwords when the component mounts
  useEffect(() => {
    fetchPasswords();
  }, []); // Empty dependency array ensures it runs only once when the component mounts
  const handleDelete = async (id) => {
    console.log(id);
  };
  return (
    <div className="pwd-list-container">
      <div className="pwd-list-wrapper">
        {passwords.map((passwordItem, index) => (
          <Password
            key={index}
            website={passwordItem.website}
            description={passwordItem.description}
            email={passwordItem.email}
            username={passwordItem.username}
            password={passwordItem.plain_password}
            showInfo={expandedIndex === index}
            handleDelete={() => handleDelete(passwordItem.ID)}
            handleShow={() => handleShow(index)}
          />
        ))}
      </div>
    </div>
  );
}

export default PasswordList;
