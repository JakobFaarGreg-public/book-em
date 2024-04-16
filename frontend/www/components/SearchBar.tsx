import React, { useEffect, useState } from "react";
import Link from "next/link";
import axios from "axios";
import { MagnifyingGlass } from "phosphor-react";
import { useRouter } from "next/router";

const SearchBar = () => {
  const [formData, setFormData] = useState<string>("");
  const router = useRouter()

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setFormData(value);
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    console.log(formData);
    router.push(`/search/results?book=${formData}`)
  };
  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={formData}
          name="searchBar"
          onChange={handleInputChange}
          placeholder="Seach for a book..."
        />
        <button type="submit">Submit</button>
      </form>

      <Link href={`/search?book=${formData}`}>
        <MagnifyingGlass size={24} />
      </Link>
    </div>
  );
};

export default SearchBar;
