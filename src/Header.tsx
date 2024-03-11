import React from "react";
import { Link } from "react-router-dom"

const Header = () => {
  return (
    <>
      <header>
        <ul className="flex">
          <li className="p-2"><Link to="/Login">ログイン</Link></li>
          <li className="p-2"><Link to="/Register">新規登録</Link></li>
          <li className="p-2"><Link to="/RakutenBook">楽天ブック検索</Link></li>
          <li className="p-2"><Link to="/mybooklist">マイブック</Link></li>
        </ul>
      </header>
    </>
  )
}

export default Header;