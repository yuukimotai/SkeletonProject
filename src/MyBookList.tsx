import React, { useEffect, useState } from "react";
import { gql, useQuery } from "urql";
import { Link, Outlet, useNavigate } from "react-router-dom";

const findMyBooksQuery = gql`
  query findAllMyBook {
    findAllMyBook {
      title
      author
      janCode
      publisherName
      itemUrl
      largeImageUrl
      mediumImageUrl
      myBook
    }
  }
`;

const MyBookList = () => {
  let cookieValue = getCookie("obserbookstoken");
  const [myBooksResult, executeFindMyBooksQuery] = useQuery({ query: findMyBooksQuery });
  const { data, fetching, error } = myBooksResult;
  const [books, setBooks] = useState([]);
  const [list, setList] = useState([])
  const newContext = {
    fetchOptions: {
      headers: {
        Authorization: `${cookieValue}`,
      },
    },
  }

  useEffect(() => {
    executeFindMyBooksQuery(newContext);
  }, [executeFindMyBooksQuery]);
  useEffect(() => {
    if (data && data.findAllMyBook !== undefined) {
      console.log(data);
      setBooks(data.findAllMyBook);
    }
  }, [data])

  if (fetching) console.log("Loading...");
  if (error) console.error(error.message);

  useEffect(() => {
    setList(
      books.map((book, index) => (
        <section className="m-2 w-52 border border-slate-700 dark:border-white rounded p-4">
          <p className="m-1 line-clamp-1" key={index}>{book.title}</p>
          <img className="mx-auto mb-2" src={book.mediumImageUrl} />
          <p className="text-center border border-white rounded">
            <Link to="../mybookdetail" state={{ book: book }}>詳細</Link>
          </p>
        </section >
      ))
    );
  }, [books])

  return (
    <>
      <div className='flex flex-row flex-wrap	max-w-screen-md'>
        {list}
      </div>
    </>
  );
}

const getCookie = (name) => {
  const cookieValue = document.cookie.match(`(^|;) ?${name}=([^;]*)(;|$)`);
  return cookieValue ? cookieValue[2] : null;
};

export default MyBookList;
