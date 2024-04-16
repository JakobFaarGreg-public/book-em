import React, { useEffect } from "react";
import Head from "next/head";
import styles from "../styles/Home.module.css";
import SearchBar from "../components/SearchBar";

export default function Home() {
  useEffect(() => {
    console.log("Jakob");
  }, []);

  return (
    <div className={styles.container}>
      <Head>
        <title>Book-em</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <h1 className="title">Book-em</h1>
        <span>A library management system</span>
        <SearchBar />
      </main>

      <footer>
        <div>Created by Jakob Faarbæk Gregersen</div>
      </footer>
    </div>
  );
}
