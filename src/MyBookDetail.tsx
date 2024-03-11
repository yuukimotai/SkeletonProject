import React, { useState, useEffect } from "react";
import ReadingMemoList from "./ReadingMemoList.js";
import { Link, Routes, Route, useLocation } from "react-router-dom";

const MyBookDetail = () => {
  const location = useLocation();
  const locationState = location.state;
  const [showReadingMemoList, setShowReadingMemoList] = useState(false)

  console.log(locationState.book)

  return (
    <>
      <div className="border border-slate-700 dark:border-white rounded-lg p-2">
        <p className="m-1 line-clamp-1">{locationState["book"].title}</p>
        <img className="mx-auto" src={locationState["book"].mediumImageUrl} />
      </div>
      <div>
        <ReadingMemoList book={locationState.book} />
      </div>
    </>
  )
}

export default MyBookDetail;