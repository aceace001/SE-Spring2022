import React, {SyntheticEvent, useState} from 'react';
import {Navigate} from 'react-router-dom';

const Register = () => {
    const [first_name, setName] = useState('');
    const [last_name, setName1] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [navi, setNavi] = useState(false);

    console.log({
        first_name,
        last_name,
        email,
        password
    })

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        const response = await fetch('http://localhost:8081/api/register', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                first_name,
                last_name,
                email,
                password
            })
        })
        setNavi(true);


        const content = await response.json();

        console.log(content);
    }

    if (navi) {
        return <Navigate to="/login"/>;
    }

    return (
        <div className="register">
        <form className={"p-4 p-md-5 border rounded-3 bg-light"} onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Please register</h1>
            <div className={"form-floating mb-3"} id='first'>
                <input type="username" className="form-control" placeholder="First name" required
                       onChange={e => setName(e.target.value)}
                />
                <label>First name</label>
            </div>

            <div className={"form-floating mb-3"} id='last'>
                <input type="username" className="form-control" placeholder="Last Name" required
                       onChange={e => setName1(e.target.value)}
                />
                <label>Last name</label>
            </div>
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
            <button className="w-100 btn btn-lg btn-primary" type="submit">Sign Up</button>
        </form>
        </div>
    );
};

export default Register;
