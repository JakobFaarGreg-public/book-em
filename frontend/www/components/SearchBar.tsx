import React from "react";
import Link from "next/link";
import { MagnifyingGlass } from "phosphor-react";
const SearchBar = () => {
  return (
    <div>
      <input id="searchBar" placeholder="Seach for a book..." />
      <Link href={"/search/results"}>
        <MagnifyingGlass size={24} />
      </Link>
    </div>
  );
};

export default SearchBar;
