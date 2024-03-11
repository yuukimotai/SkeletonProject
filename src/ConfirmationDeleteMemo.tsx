import React, { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { gql, useMutation } from "urql";
import Modal from "react-modal";
import { DeleteReadingMemo } from "./gql/graphql.js";

const deleteReadingMemoMutation = gql`
mutation deleteReadingMemo($input: DeleteReadingMemo!) {
  deleteReadingMemo(input:$input)
}
`;
const ConfirmationDeleteReadingMemo = (props) => {
  const [janCode, setJanCode] = useState("")
  const [title, setTitle] = useState("")
  const [content, setContent] = useState("")
  const { register, handleSubmit } = useForm()
  const [deleteMemoResult, deleteMemoExecute] = useMutation(deleteReadingMemoMutation)
  const onSubmit = (data) => {
    let memo: DeleteReadingMemo = {
      janCode: data.janCode,
      memoId: props["memo"].memoId
    }
    console.log(memo)
    let cookieValue = getCookie("obserbookstoken")
    deleteMemoExecute({ input: memo }, {
      fetchOptions: {
        headers: {
          Authorization: `${cookieValue}`,
        },
      },
    })
    window.location.href = "/mybooklist";
  }
  const [modalIsOpen, setIsOpen] = useState(false);
  const openModal = () => {
    setIsOpen(true);
    setTitle(props["memo"].title)
    setContent(props["memo"].content)
    setJanCode(props["memo"].janCode)
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
  }, [handleSubmit])

  return (
    <>
      <div>
        <button type="submit" onClick={openModal} className="text-red-700 hover:text-white border border-red-700 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:border-red-500 dark:text-red-500 dark:hover:text-white dark:hover:bg-red-600 dark:focus:ring-red-900">削除</button>
        <Modal
          style={customStyles}
          isOpen={modalIsOpen}
          onRequestClose={closeModal}
          contentLabel="ConfirmationDeleteMemoModal"
        >
          <h2 className="text-slate-950 mb-2">削除確認</h2>
          <p className="text-slate-950 mb-4">以下の内容のメモを削除します。よろしいでしょうか？</p>
          <div className="border border-gray-800 rounded p-2 mb-4">
            <p className="text-slate-950 mb-2">
              {title}
            </p>
            <p className="text-slate-950">
              {content}
            </p>
          </div>
          <form onSubmit={handleSubmit(onSubmit)}>
            <p className="flex flex-row justify-end">
              <input type="hidden" {...register("janCode")} value={janCode} />
              <button onClick={closeModal} className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800">閉じる</button>
              <button type="submit" className="text-red-700 hover:text-white border border-red-700 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center mb-2 dark:border-red-500 dark:text-red-500 dark:hover:text-white dark:hover:bg-red-600 dark:focus:ring-red-900">削除</button>
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

export default ConfirmationDeleteReadingMemo;