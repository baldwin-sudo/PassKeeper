import React from "react";
import "./search.css";
import search from "../../assets/search.png";

function Search() {
  const handleSearch = () => {
    console.log("search");
    // Implement your search functionality here
  };

  return (
    <div className="search-container">
      <div className="searchBar-container">
        <input
          type="text"
          id="search"
          placeholder="Search..."
          aria-label="Search"
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              handleSearch();
            }
          }}
        />
        <img
          src={search}
          alt="Search icon"
          className="search-icon"
          onClick={handleSearch}
        />
      </div>
    </div>
  );
}

export default Search;
