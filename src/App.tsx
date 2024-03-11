import { useState } from 'react'

import './App.css'
import { Routes, Route, Link, BrowserRouter } from "react-router-dom";
import Login from "./Login.js"
import Register from "./Register.js"
import RakutenBook from "./RakutenBookRequest.js"
import MyBookList from "./MyBookList.js"
import MyBookDetail from "./MyBookDetail.js"
import ReadingMemoList from './ReadingMemoList.js';

function App() {
  return (
    <>
      <div className="App">
        <Routes>
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/rakutenbook" element={<RakutenBook />} />
          <Route path="/mybooklist" element={<MyBookList />} />
          <Route path="/mybookdetail" element={<MyBookDetail />} />
          <Route path="/readingmemolist" element={<ReadingMemoList />} />
        </Routes >
      </div>
    </>
  )
}

export default App
