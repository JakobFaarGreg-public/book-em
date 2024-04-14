import React from "react";

interface BookDisplayInput {
  title: string;
  author: string;
}

const BookDisplay = ({ title, author }: BookDisplayInput) => {
  return (
    <div>
      <div>{title}</div>
      <div>{author}</div>
    </div>
  );
};
export default BookDisplay;
