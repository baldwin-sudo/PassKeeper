import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { useState } from "react";
import PasswordList from "./components/PasswordList/PasswordList";
import PasswordForm from "./components/PasswordForm/PasswordForm"; // Import your PasswordForm component
import "./App.css";
import Search from "./components/Search/Search";
import Header from "./components/Header/Header";
import Abstract from "./Abstract/Abstract";
import Footer from "./components/Footer/Footer";
import { passwords } from "./data";

function App() {
  const [passwordList, setPasswordList] = useState(passwords);

  return (
    <Router>
      <Header />
      <Routes>
        <Route
          path="/"
          element={
            <>
              <Search />
              <PasswordList passwordList={passwordList} />
            </>
          }
        />
        <Route
          path="/add"
          element={<PasswordForm setPasswordList={setPasswordList} />}
        />{" "}
        <Route path="/abstract" element={<Abstract />} />
      </Routes>
      <Footer />
    </Router>
  );
}

export default App;
