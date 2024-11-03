import React from "react";
import "./abstract.css";

function Abstract() {
  return (
    <div className="abstract-wrapper">
      <ul className="abstract">
        <li>
          This <span className="highlight">password manager app</span> is
          designed to offer a <span className="highlight">secure</span> and
          centralized method for
          <span className="highlight"> storing sensitive information</span>,
          such as passwords. Access to all stored information is protected by a
          <span className="highlight"> master password</span>, which serves as
          the primary access key. Instead of storing the master password
          directly, the app securely saves a
          <span className="highlight"> hashed version</span> of it in the
          database.
        </li>

        <li>
          The hashing process ensures that the original password cannot be
          retrieved, as it uses a
          <span className="highlight"> hashing algorithm</span> to convert the
          password into a fixed-size string. When storing other passwords, we
          implement a
          <span className="highlight"> key derivation algorithm</span> such as
          <span className="highlight"> PBKDF2</span> (Password-Based Key
          Derivation Function 2), which creates a unique{" "}
          <span className="highlight"> encryption key</span> from the master
          password.
        </li>

        <li>
          This <span className="highlight"> derived key</span> is used to
          encrypt each password before storage, adding an extra layer of{" "}
          <span className="highlight"> security</span>. Each encryption includes
          unique salts and a high number of iterations, which strengthens
          resistance to brute-force attacks. To view stored passwords, the app
          requires the correct master password to regenerate the encryption key
          and securely decrypt the stored information.
        </li>

        <li>
          This setup delivers a <span className="highlight">reliable</span> and{" "}
          <span className="highlight">flexible </span>
          framework for securely managing encrypted passwords, making it both
          secure and user-friendly.
        </li>
      </ul>{" "}
    </div>
  );
}

export default Abstract;
