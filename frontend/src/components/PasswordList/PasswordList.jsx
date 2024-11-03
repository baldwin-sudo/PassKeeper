import React, { useState } from "react";
import Password from "../Password/Password";
import "./passwordList.css";

function PasswordList({ passwordList }) {
  const [expandedIndex, setExpandedIndex] = useState(-1);

  const handleShow = (index) => {
    setExpandedIndex(expandedIndex === index ? null : index);
  };

  return (
    <div className="pwd-list-container">
      <div className="pwd-list-wrapper">
        {passwordList.map((passwordItem, index) => (
          <Password
            key={index}
            website={passwordItem.website}
            description={passwordItem.description}
            email={passwordItem.email}
            username={passwordItem.username}
            password={passwordItem.password}
            showInfo={expandedIndex === index}
            handleShow={() => handleShow(index)}
          />
        ))}
      </div>
    </div>
  );
}

export default PasswordList;
