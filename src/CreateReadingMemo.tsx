import React, { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { gql, useMutation } from "urql";
import Modal from "react-modal";
import { NewReadingMemo } from "./gql/graphql.js";

const createReadingMemoMutation = gql`
mutation createReadingMemo($input: NewReadingMemo!) {
  createReadingMemo(input:$input) {
    janCode
    title
    content
  }
}
`;
const CreateReadingMemo = (props) => {
  const { register, handleSubmit, setValue, watch } = useForm()
  const [createMemoResult, createMemoExecute] = useMutation(createReadingMemoMutation)
  const [title, setTitle] = useState("")
  const [content, setContent] = useState("")
  const onSubmit = (data) => {
    let newMemo: NewReadingMemo = {
      memoId: "",
      janCode: props["book"].janCode,
      title: data.title,
      content: data.content
    }
    let cookieValue = getCookie("obserbookstoken")
    createMemoExecute({ input: newMemo }, {
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
  }
  const closeModal = () => {
    setIsOpen(false);
  }
  const customStyles = {
    content: {
      width: '600px',
      height: '400px',
      top: '50%',
      left: '50%',
      transform: 'translate(-50%, -50%)',
      borderradius: '8px',
      bodercolor: '#333333',
      backgroundColor: '#f9f9f9'
    },
  };
  useEffect(() => {
  }, [watch]);

  return (
    <>
      <div>
        <button type="submit" onClick={openModal} className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center mb-2 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800">メモ作成</button>
        <Modal
          style={customStyles}
          isOpen={modalIsOpen}
          onRequestClose={closeModal}
          contentLabel="CreateMemoModal"
        >
          <h2 className="text-slate-950 mb-2">メモ作成</h2>
          <form onSubmit={handleSubmit(onSubmit)}>
            <p className="flex flex-col mb-4">
              <input type="text" placeholder="タイトル" {...register("title")} className="mb-2 rounded" />
              <input type="textarea" placeholder="内容" {...register("content")} className="h-52 rounded" />
            </p>
            <p className="flex flex-row justify-end">
              <button onClick={closeModal} className="text-red-700 hover:text-white border border-red-700 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 dark:border-red-500 dark:text-red-500 dark:hover:text-white dark:hover:bg-red-600 dark:focus:ring-red-900">閉じる</button>
              <button type="submit" className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800">作成</button>
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

export default CreateReadingMemo;