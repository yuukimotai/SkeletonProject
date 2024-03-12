import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App.js'
import Header from "./Header.js"
import { BrowserRouter, Route, Routes } from "react-router-dom"
import { Client, fetchExchange, Provider } from "urql"


const rakutenBookClient = new Client({
  exchanges: [fetchExchange],
  url: "http://localhost:8080/query",
})

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <Provider value={rakutenBookClient}>
        <Header />
        <App />
      </Provider>
    </BrowserRouter>
  </React.StrictMode>,
)