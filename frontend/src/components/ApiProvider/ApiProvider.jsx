// context/ApiContext.jsx
import React, { createContext, useContext } from "react";

// Create a Context with the api_url
const ApiContext = createContext();

export const ApiProvider = ({ children }) => {
  const api_url = import.meta.env.VITE_API_HOST;

  return <ApiContext.Provider value={api_url}>{children}</ApiContext.Provider>;
};

export const useApi = () => useContext(ApiContext);
