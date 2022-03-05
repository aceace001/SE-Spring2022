import React, {SyntheticEvent, useState} from 'react';
import {Link, Navigate} from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.min.css';

const Home = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [navi, setNavi] = useState(false);
    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        setNavi(true);
    }

    if (navi) {
        return <Navigate to="/HomePost"/>;
    }

    return (
        <div className="home">
        <div className="container1 ">
            <div className={"row align-items-center g-lg-5 py-5"}>
                <div className={"col-lg-7 text-center text-lg-start"}>
                    <h1 className={"display-4 fw-bold lh-1 mb-3"}>
                        Welcome to the App!
                    </h1>
                    <p className={"col-lg-10 fs-4"}>
                        "Please register to use our app!"
                    </p>

                </div>
                <div className={"col-md-10 mx-auto col-lg-5"}>
                    <form className={"p-4 p-md-5 border rounded-3 bg-light"} onSubmit={submit}>
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
                        <button className="w-100 btn btn-lg btn-primary mb-2" type="submit">Login</button>
                        <small className="signup">New user? Click here to register.
                            <ul className="nav  ">
                                <li className="nav-item active">
                                    <Link to="/register" className="nav-link  text-blue">Register</Link>
                                </li>
                            </ul>
                        </small>
                    </form>
                </div>

            </div>
        </div>
        </div>
    );
};

export default Home;
