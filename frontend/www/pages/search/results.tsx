import React, { useEffect, useState } from "react";
import Link from "next/link";
import { useRouter } from "next/router";
import axios from "axios";
import SearchBar from "../../components/SearchBar";
const SearchResultsPage = () => {
  const [data, setData] = useState();
  const router = useRouter();
  const { book } = router.query;
  useEffect(() => {
    console.log(book);
    getBooks(book);
  }, []);

  const getBooks = async (book) => {
    axios
      .get(`http://127.0.0.1:3333/v1/books?book=${book}`)
      .then((response) => {
        setData(response.data);
      })
      .catch((error) => console.error(error));
  };

  return (
    <>
      <div id="Sidebar">
        <Link href={"/"}>Back to home</Link>
        <SearchBar />
      </div>
    </>
  );
};
export default SearchResultsPage;
