import React, { useState, useEffect, useMemo } from "react";
import { gql, useQuery } from "urql";
import { render } from "react-dom";
import Modal from 'react-modal';
import CreateReadingMemo from "./CreateReadingMemo.js";
import EditReadingMemo from "./EditReadingMemo.js";
import DeleteReadingMemo from "./ConfirmationDeleteMemo.js";
import { GetAllMemo } from "./gql/graphql.js";


const getAllReadingMemo = gql`
query findAllReadingMemo($input: GetAllMemo!) {
  ReadingMemos(input: $input) {
    memoId
    janCode
    title
    content
  }
}
`;

const DEFAULT_PAGE_SIZE = 10;

function ReadingMemoList(props) {
  let cookieValue = getCookie("obserbookstoken")
  const getAllMemo: GetAllMemo = {
    janCode: props.book.janCode
  }
  const newContext = {
    fetchOptions: {
      headers: {
        Authorization: `${cookieValue}`,
      },
    },
  }
  const [getAllReadingMemoResult, excecuteGetAllReadingMemo] = useQuery({
    query: getAllReadingMemo,
    variables: { input: getAllMemo }
  })
  const { data, fetching, error } = getAllReadingMemoResult;
  const [memos, setMemos] = useState([])
  //const [memoList, setMemoList] = useState([])
  const [currentPage, setCurrentPage] = useState(1)
  const [posts, setPosts] = useState<any>([])
  const [totalPages, setTotalPages] = useState(0)
  const [count, setCount] = useState(0)
  const handlePageChange = (newPage: number) => {
    if (newPage >= 1 && newPage <= totalPages) {
      setCurrentPage(newPage);
      console.log("next")
    }
    if (newPage >= 1 && newPage === totalPages) {
      setCurrentPage(newPage);
      console.log("prev")
    }
  };
  const currentMemos = useMemo(() => {
    return memos
      .filter((memo) => memo["title"])
      .slice(
        (currentPage - 1) * DEFAULT_PAGE_SIZE,
        currentPage * DEFAULT_PAGE_SIZE
      );
  }, [currentPage, memos]);

  useEffect(() => {
    excecuteGetAllReadingMemo(newContext);
  }, [excecuteGetAllReadingMemo]);
  useEffect(() => {
    if (data && data.ReadingMemos !== undefined) {
      setMemos(data.ReadingMemos)
      setCount(data.ReadingMemos.length)
      setTotalPages(Math.ceil(count / DEFAULT_PAGE_SIZE))
    }
  }, [data])

  if (fetching) console.log("Loading...");
  if (error) console.error(error.message);

  return (
    <>
      <h3>読書メモリスト</h3>
      <div className="flex flex-row justify-end">
        <CreateReadingMemo book={props.book} />
      </div>
      <div className='overflow-y-scroll'>
        <div className='flex-none h-80'>
          {currentMemos.map((memo, index) => (
            <div key={index} className='memo-item border border-slate-700 dark:border-white rounded-lg p-2 mb-2'>
              <p className="flex justify-start text-slate-700 dark:text-white">{memo["title"]}</p>
              <p className="flex flex-row justify-end">
                <EditReadingMemo memo={memo} />
                <DeleteReadingMemo memo={memo} />
              </p>
            </div>
          ))}
        </div>
      </div>
    </>
  )
}

const getCookie = (name) => {
  const cookieValue = document.cookie.match(`(^|;) ?${name}=([^;]*)(;|$)`);
  return cookieValue ? cookieValue[2] : null;
};

export default ReadingMemoList;

