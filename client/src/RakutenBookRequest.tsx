import React, { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import axios from "axios";
import { gql, useMutation } from "urql";
import { CreateBookMutation, NewBook } from "./gql/graphql.js";

const RakutenBook = () => {
  const { register, handleSubmit } = useForm()
  const [books, setBooks] = useState([])
  const [list, setList] = useState([])
  const [currentPage, setCurrentPage] = useState(0)
  const [pageCount, setPageCount] = useState(0)
  const [newUrl, setNewUrl] = useState("")
  const [paginateFlag, setPaginateFlag] = useState(false)

  let url = `https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404?format=json&applicationId=1073592281532444208`
  const onSubmit = (data) => {
    if (data.title != "") {
      let encodedTitle = encodeURIComponent(data.title)
      url += "&title=" + encodedTitle
    }
    if (data.authorText != "") {
      let encodedAuthor = encodeURIComponent(data.authorText)
      url += "&author=" + encodedAuthor
    }
    if (data.ISBNText != "") {
      let encodedISBN = encodeURIComponent(data.ISBNText)
      url += "&isbn=" + encodedISBN
    }
    if (data.publisherNameText != "") {
      let encodedPublishName = encodeURIComponent(data.publisherNameText)
      url += "&publisherName=" + encodedPublishName
    }
    if (data.booksGenreIdText != "") {
      let encodedBooksGenreId = encodeURIComponent(data.booksGenreId)
      url += "&booksGenreId" + encodedBooksGenreId
    }
    setNewUrl(url)
    axios.get(url)
      .then(res => {
        setCurrentPage(res.data.page)
        setPageCount(res.data.pageCount)
        setBooks(res.data.Items)
        setPaginateFlag(true)
      })
  }

  useEffect(() => {
    console.log(`current:`, currentPage)
    console.log(`pages:`, pageCount)
  }, [pageCount])

  const createBookMutation = gql`
  mutation createBook($input: NewBook!) {
    createBook(input: $input) {
      title
      author
      janCode
      publisherName
      itemUrl
      largeImageUrl
      mediumImageUrl
      myBook
      attentionBook
    }
  }
  `;
  const [createBookResult, createBook] = useMutation(createBookMutation);

  useEffect(() => {
    // ステートが更新されたらlistを再生成
    const AddMyBook = async (newBookData) => {
      let newBook: NewBook = {
        title: newBookData["Item"].title,
        author: newBookData["Item"].author,
        janCode: newBookData["Item"].isbn,
        publisherName: newBookData["Item"].publisherName,
        itemUrl: newBookData["Item"].itemUrl,
        largeImageUrl: newBookData["Item"].largeImageUrl,
        mediumImageUrl: newBookData["Item"].mediumImageUrl,
        myBook: true,
        attentionBook: false
      }
      let cookieValue = getCookie("obserbookstoken")
      createBook({ input: newBook }, {
        fetchOptions: {
          headers: {
            Authorization: `${cookieValue}`,
          },
        },
      });
    };
    function AddAttentionBook(newBookData) {
      let newBook: NewBook = {
        title: newBookData["Item"].title,
        author: newBookData["Item"].author,
        publisherName: newBookData["Item"].publisherName,
        janCode: newBookData["Item"].isbn,
        itemUrl: newBookData["Item"].itemUrl,
        largeImageUrl: newBookData["Item"].largeImageUrl,
        mediumImageUrl: newBookData["Item"].mediumImageUrl,
        attentionBook: true
      }
      let cookieValue = getCookie("obserbookstoken")
      createBook({ input: newBook }, {
        fetchOptions: {
          headers: {
            Authorization: `${cookieValue}`,
          },
        },
      });
      console.log("気になる本に追加しました")
    };
    setList(
      books.map((book, index) => (
        <section className="w-1/4 border dark:border-white p-2 mx-8 mb-4 rounded">
          <p className="m-1 line-clamp-1" key={index}>{book["Item"].title}</p>
          <img className="mx-auto mb-2" src={book["Item"].largeImageUrl} />
          <p className="text-center mb-4">
            <button className="text-xs me-2 border border-slate-700 dark:border-white p-2" onClick={() => AddMyBook(book)}>マイブック</button>
            <button className="text-xs border border-slate-700 dark:border-white p-2" onClick={() => AddAttentionBook(book)}>気になる！</button>
          </p>
        </section>
      ))
    );
  }, [books]); // booksの変更を監視

  const handlePageChange = async (newPage: number) => {
    setCurrentPage(newPage)
    console.log(newUrl + `&page=${newPage}`)
    axios.get(newUrl + `&page=${newPage}`)
      .then(res => {
        setCurrentPage(res.data.page)
        setPageCount(res.data.pageCount)
        setBooks(res.data.Items)
      })
  };

  return (
    <>
      <div className="border rounded p-2 mb-8">
        <h3>RakutenBookで検索</h3>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="grid grid-flow-row auto-rows-max">
            <input placeholder="タイトル" {...register("title")} />
            <input type="hidden" id="authorText" placeholder="著者" {...register("authorText")} />
            <input type="hidden" id="ISBNText" placeholder="ISBN" {...register("ISBNText")} />
            <input type="hidden" id="publisherNameText" placeholder="出版社名" {...register("publisherNameText")} />
            <input type="hidden" id="booksGenreIdText" placeholder="楽天ブックスジャンルID" {...register("booksGenreIdText")} />
          </div>
          <div>
            <input type="checkbox" id="author" onChange={InputFieldDisplay} />
            <label htmlFor="author">著者</label>
            <input type="checkbox" id="ISBN" onChange={InputFieldDisplay} />
            <label htmlFor="ISBN">ISBN</label>
            <input type="checkbox" id="publisherName" onChange={InputFieldDisplay} />
            <label htmlFor="publisherName">出版社名</label>
            <input type="checkbox" id="booksGenreId" onChange={InputFieldDisplay} />
            <label htmlFor="booksGenreId">楽天ブックスジャンルID</label>
          </div>
          <div>
            <input type="submit" value="検索" className="border border-slate-700 dark:border-white p-px rounded" />
          </div>
        </form>
      </div>
      <div className='flex flex-row flex-wrap max-w-screen-md mb-4'>
        {list}
      </div>
      {paginateFlag ?
        <div className="text-center">
          <button id='pagination-prev' className="me-2 border border-slate-700 dark:border-white p-2" onClick={() => handlePageChange(currentPage - 1)} disabled={currentPage <= 1}>
            前へ
          </button>
          <button id='pagination-next' className="border border-slate-700 dark:border-white p-2" onClick={() => handlePageChange(currentPage + 1)} disabled={currentPage === pageCount}>
            次へ
          </button>
        </div>
        : <div></div>}
    </>
  )
}

const getCookie = (name) => {
  const cookieValue = document.cookie.match('(^|;) ?' + name + '=([^;]*)(;|$)');
  return cookieValue ? cookieValue[2] : null;
}

function InputFieldDisplay() {
  let authorCheckBox = document.getElementById("author") as HTMLInputElement
  let authorInput = document.getElementById("authorText") as HTMLInputElement
  let isbnCheckBox = document.getElementById("ISBN") as HTMLInputElement
  let isbnInput = document.getElementById("ISBNText") as HTMLInputElement
  let publisherNameCheckBox = document.getElementById("publisherName") as HTMLInputElement
  let publisherInput = document.getElementById("publisherNameText") as HTMLInputElement
  let booksGenreIdCheckBox = document.getElementById("booksGenreId") as HTMLInputElement
  let booksGenreIdInput = document.getElementById("booksGenreIdText") as HTMLInputElement

  if (authorCheckBox.checked === true) {
    authorInput.type = "text"
  } else {
    authorInput.type = "hidden"
  }
  if (isbnCheckBox.checked === true) {
    isbnInput.type = "text"
  } else {
    isbnInput.type = "hidden"
  }
  if (publisherNameCheckBox.checked === true) {
    publisherInput.type = "text"
  } else {
    publisherInput.type = "hidden"
  }
  if (booksGenreIdCheckBox.checked === true) {
    booksGenreIdInput.type = "text"
  } else {
    booksGenreIdInput.type = "hidden"
  }
}

export default RakutenBook;