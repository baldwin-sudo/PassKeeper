import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import { useEffect, useState } from "react";
import PasswordList from "./components/PasswordList/PasswordList";
import PasswordForm from "./components/PasswordForm/PasswordForm";
import "./App.css";
import Search from "./components/Search/Search";
import Header from "./components/Header/Header";
import Abstract from "./Abstract/Abstract";
import Footer from "./components/Footer/Footer";
import { passwords } from "./data";
import Login from "./components/Login/Login";
import NotFound from "./components/NotFound/NotFound";

function App() {
  const [passwordList, setPasswordList] = useState(passwords);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    // Check local storage for authentication status
    const loggedInStatus = localStorage.getItem("isLoggedIn") === "true";
    setIsLoggedIn(loggedInStatus);
  }, []);

  const logoutUser = () => {
    setIsLoggedIn(false);
    localStorage.setItem("isLoggedIn", "false"); // Update local storage
  };

  const loginUser = () => {
    setIsLoggedIn(true);
    localStorage.setItem("isLoggedIn", "true"); // Update local storage
  };

  return (
    <Router>
      <Header logoutUser={logoutUser} isLoggedIn={isLoggedIn} />

      <Routes>
        <Route
          path="/"
          element={
            isLoggedIn ? (
              <>
                <Search />
                <PasswordList passwordList={passwordList} />
              </>
            ) : (
              <Navigate to="/login" replace />
            )
          }
        />
        <Route
          path="/add"
          element={
            isLoggedIn ? (
              <PasswordForm setPasswordList={setPasswordList} />
            ) : (
              <Navigate to="/login" replace />
            )
          }
        />
        <Route
          path="/abstract"
          element={isLoggedIn ? <Abstract /> : <Navigate to="/login" replace />}
        />
        <Route
          path="/login"
          element={
            isLoggedIn ? (
              <Navigate to="/" replace />
            ) : (
              <Login onLogin={loginUser} />
            )
          }
        />
        <Route path="*" element={<NotFound isLoggedIn={isLoggedIn} />} />
      </Routes>
      <Footer />
    </Router>
  );
}

export default App;
