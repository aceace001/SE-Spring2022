import React from 'react';
import {Link} from "react-router-dom";

const Nav = () => {

    return (
        <nav className="navb bg-dark text-white">
            <div className="container ">
                <div className={"d-flex flex-wrap mb-4 p-2"}>
                    <Link to="/" className="navbar-brand px-2 text-white">Home</Link>

                    <ul className="nav mb-2 mb-md-0">
                        <li className="nav-item active">
                            <Link to="/login" className="nav-link px-2 text-white">Login</Link>
                        </li>
                        <li className="nav-item active">
                            <Link to="/register" className="nav-link px-2 text-white">Register</Link>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    );
};

export default Nav;
