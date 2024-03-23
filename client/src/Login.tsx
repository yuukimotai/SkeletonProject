
import { useState } from "react";
import { useForm } from "react-hook-form"
import axios from "axios"
import Cookies from 'universal-cookie';


const baseUrl = "http://localhost:8080/authentication/login"
axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';//これの許可を後で調整する必要あり
interface LoginRequest {
  email: string;
  password: string;
}
export const Login = () => {
  const { register, handleSubmit } = useForm()
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const requestData: LoginRequest = {
    email: email, // emailはフォームから受け取った値
    password: password, // passwordはフォームから受け取った値
  };
  const onSubmit = (data) => {
    setEmail(data.email)
    setPassword(data.password)
    axios.post(baseUrl, requestData, {
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': 'http://localhost:5174',
      },
    }).then((response) => {
      var receivedJwt = response.data.jwt;
      // JWTをサーバーから受け取った後
      const cookies = new Cookies();
      cookies.set('obserbookstoken', receivedJwt, { path: '/' });
      if (response.data !== null) {
        alert(response.data.statustext)
        window.location.href = "/mybooklist";
      }
    });
  }

  return (
    <>
      <div className="border border-slate-700 dark:border-white rounded p-2">
        <h3>ログイン</h3>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="mb-2">
            <input type="email" placeholder="Email" {...register("email")} className="w-full" />
          </div>
          <div className="mb-4 p-p">
            <input type="password" placeholder="Password" {...register("password")} className="w-full" />
          </div>
          <div className="text-right">
            <input type="submit" value="ログイン" className="border border-slate-700 dark:border-white p-px rounded text-right" />
          </div>
        </form>
      </div>
    </>
  );
}

export default Login;