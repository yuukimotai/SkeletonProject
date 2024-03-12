import React, { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { gql, useMutation } from "urql";
import Modal from "react-modal";
import { UpdateReadingMemo } from "./gql/graphql.js";

const createReadingMemoMutation = gql`
mutation updateReadingMemo($input: UpdateReadingMemo!) {
  updateReadingMemo(input: $input) {
    memoId
    janCode
    title
    content
  }
}
`;
const EditReadingMemo = (data) => {
  const { register, handleSubmit } = useForm()
  const [updateMemoResult, updateMemoExecute] = useMutation(createReadingMemoMutation)
  const [title, setTitle] = useState("")
  const [content, setContent] = useState("")
  const [memoId, setMemoId] = useState(0)
  const [janCode, setJanCode] = useState("")
  const onSubmit = (data) => {
    console.log(data.memoId)
    let newMemo: UpdateReadingMemo = {
      memoId: data.memoId,
      janCode: data.janCode,
      title: data.title,
      content: data.content
    }
    let cookieValue = getCookie("obserbookstoken")
    updateMemoExecute({ input: newMemo }, {
      fetchOptions: {
        headers: {
          Authorization: `${cookieValue}`,
        },
      },
    })
    window.location.href = "/mybooklist";
  }
  const [updateMemoModalIsOpen, setIsOpen] = useState(false);
  const openModal = () => {
    setIsOpen(true);
  }
  const closeModal = () => {
    setIsOpen(false);
  }
  const customStyles = {
    content: {
      top: '50%',
      left: '50%',
      transform: 'translate(-50%, -50%)',
    },
    button: {
      backgroundcolor: 'black'
    }
  };

  useEffect(() => {
    setTitle(data["memo"].title)
    setContent(data["memo"].content)
    setMemoId(data["memo"].memoId)
    setJanCode(data["memo"].janCode)
  }, [data])
  useEffect(() => {
  }, [handleSubmit])

  return (
    <>
      <div>
        <button type="submit" onClick={openModal} className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800">
          更新
        </button>
        <Modal
          style={customStyles}
          isOpen={updateMemoModalIsOpen}
          onRequestClose={closeModal}
          contentLabel="EditMemoModal"
        >
          <h2 className="text-slate-950 mb-2">メモ更新</h2>
          <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col">
            <p className="flex flex-col mb-4">
              <input type="text" {...register("title")} value={title} onChange={(e) => { setTitle(e.target.value) }} className="mb-2 rounded" />
              <input type="textarea" {...register("content")} value={content} onChange={(e) => { setContent(e.target.value) }} className="h-52 rounded" />
              <input type="hidden" {...register("memoId")} value={memoId} />
              <input type="hidden" {...register("janCode")} value={janCode} />
            </p>
            <p className="flex flex-row justify-end">
              <button onClick={closeModal} className="text-red-700 hover:text-white border border-red-700 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-red-500 dark:text-red-500 dark:hover:text-white dark:hover:bg-red-600 dark:focus:ring-red-900">閉じる</button>
              <button type="submit" className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center mb-2 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800">更新</button>
            </p>
          </form>
        </Modal>
      </div>
    </>
  )
}

const getCookie = (name) => {
  const cookieValue = document.cookie.match('(^|;) ?' + name + '=([^;]*)(;|$)');
  return cookieValue ? cookieValue[2] : null;
}

export default EditReadingMemo;