import React, { useEffect } from "react";
import Link from "next/link";
import { useRouter } from "next/router";
const SearchResultsPage = () => {
  const router = useRouter()
  const { book} = router.query
  useEffect(() => {
    console.log(book);

  },[])
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
