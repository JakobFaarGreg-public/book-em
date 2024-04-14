import React from "react";
import Link from "next/link";
const SearchResultsPage = () => {
  return (
    <>
      <div id="Sidebar">
        <Link href={"/"}>Back to home</Link>
        <div>Search Results</div>
      </div>
      <div id="ScrollableBookList"></div>
    </>
  );
};
export default SearchResultsPage;
