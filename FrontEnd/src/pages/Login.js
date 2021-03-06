import React, {SyntheticEvent, useState} from 'react';
import { Navigate } from "react-router-dom";


const Login = (props: { setName: (name: string) => void }) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [navi, setNavi] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        await fetch('http://localhost:8081/api/login',{
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        })
        setNavi(true);
    }

    if (navi) {
        return <Navigate to="/Homepost"/>;
    }

    return (
        <div className="login">
        <form className={"p-4 p-md-5 border rounded-3 bg-light"} onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please login</h1>
            <div className={"form-floating mb-3"} id='email'>
                <input type="email" className="form-control" placeholder="Email address" required
                       onChange={e => setEmail(e.target.value)}
                />
                <label>Email Address</label>
            </div>
            <div className={"form-floating mb-3"} id='password'>
                <input type="password" className="form-control" placeholder="Password" required
                       onChange={e => setPassword(e.target.value)}
                />
                <label>Password</label>
            </div>
            <button className="w-100 btn btn-lg btn-primary" type="submit">Login</button>
        </form>
        </div>
    );
};

export default Login;
